package search

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// SearchWatcher 搜索文件监听器
type SearchWatcher struct {
	indexer   *Indexer
	watcher   *fsnotify.Watcher
	isRunning bool
	mutex     sync.Mutex
	debounce  map[string]time.Time
}

// NewSearchWatcher 创建新的搜索文件监听器
func NewSearchWatcher(indexer *Indexer) *SearchWatcher {
	return &SearchWatcher{
		indexer:  indexer,
		debounce: make(map[string]time.Time),
	}
}

// Start 启动搜索文件监听器
func (sw *SearchWatcher) Start(docsDir string) error {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	if sw.isRunning {
		return nil
	}

	var err error
	sw.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	// 递归监听文档目录
	if err := sw.watchRecursive(docsDir); err != nil {
		sw.watcher.Close()
		return err
	}

	sw.isRunning = true

	// 启动监听goroutine
	go sw.run()

	log.Printf("搜索文件监听器已启动，监控目录: %s", docsDir)
	log.Println("文件变化时将自动更新搜索索引")

	return nil
}

// run 运行监听循环
func (sw *SearchWatcher) run() {
	defer sw.stop()

	for {
		select {
		case event, ok := <-sw.watcher.Events:
			if !ok {
				return
			}

			// 忽略临时文件
			if sw.isTempFile(event.Name) {
				continue
			}

			// 使用防抖机制处理文件变化
			sw.handleFileEvent(event)

		case err, ok := <-sw.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("搜索文件监听错误: %v", err)
		}
	}
}

// handleFileEvent 处理文件变化事件
func (sw *SearchWatcher) handleFileEvent(event fsnotify.Event) {
	// 防抖处理：避免短时间内重复处理同一文件
	now := time.Now()
	if lastTime, exists := sw.debounce[event.Name]; exists {
		if now.Sub(lastTime) < 2*time.Second {
			return // 2秒内重复事件，忽略
		}
	}
	sw.debounce[event.Name] = now

	// 清理过期的防抖记录
	go sw.cleanupDebounce()

	log.Printf("搜索索引检测到文件变化: %s, 操作: %s", event.Name, event.Op.String())

	// 根据事件类型处理
	if event.Op&fsnotify.Remove == fsnotify.Remove || 
	   event.Op&fsnotify.Rename == fsnotify.Rename {
		// 文件删除或重命名
		if err := sw.indexer.removeDocument(event.Name); err != nil {
			log.Printf("删除文档索引失败 %s: %v", event.Name, err)
		} else {
			log.Printf("已从搜索索引中移除: %s", event.Name)
		}
	} else if event.Op&fsnotify.Write == fsnotify.Write || 
	          event.Op&fsnotify.Create == fsnotify.Create {
		// 文件创建或修改
		if !sw.indexer.isDocumentFile(event.Name) {
			return
		}

		if err := sw.indexer.indexDocument(event.Name); err != nil {
			log.Printf("更新文档索引失败 %s: %v", event.Name, err)
		} else {
			log.Printf("已更新搜索索引: %s", event.Name)
		}
	}
}

// cleanupDebounce 清理过期的防抖记录
func (sw *SearchWatcher) cleanupDebounce() {
	now := time.Now()
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	for path, timestamp := range sw.debounce {
		if now.Sub(timestamp) > 10*time.Second {
			delete(sw.debounce, path)
		}
	}
}

// watchRecursive 递归添加目录到监听器
func (sw *SearchWatcher) watchRecursive(dirPath string) error {
	// 添加当前目录到监听器
	if err := sw.watcher.Add(dirPath); err != nil {
		return err
	}

	// 递归添加子目录
	return filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
			return sw.watcher.Add(path)
		}
		return nil
	})
}

// isTempFile 判断是否为临时文件
func (sw *SearchWatcher) isTempFile(filePath string) bool {
	baseName := filepath.Base(filePath)
	return strings.Contains(filePath, ".swp") ||
		strings.Contains(filePath, "~$") ||
		strings.HasPrefix(baseName, ".") ||
		strings.HasSuffix(baseName, ".tmp")
}

// stop 停止监听器
func (sw *SearchWatcher) stop() {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	if sw.watcher != nil {
		sw.watcher.Close()
	}
	sw.isRunning = false
	log.Println("搜索文件监听器已停止")
}

// Stop 手动停止监听器
func (sw *SearchWatcher) Stop() {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	if sw.isRunning && sw.watcher != nil {
		sw.watcher.Close()
		sw.isRunning = false
	}
}

// IsRunning 检查监听器是否在运行
func (sw *SearchWatcher) IsRunning() bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return sw.isRunning
}