package doc

import (
	"docs-go/pkg/config"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// FileWatcher 文件监控器结构体
type FileWatcher struct {
	DocTree   *DocTree // 文档树引用
	Watcher   *fsnotify.Watcher
	IsRunning bool
}

// NewFileWatcher 创建新的文件监控器
func NewFileWatcher(docTree *DocTree, config *config.Config) *FileWatcher {
	return &FileWatcher{
		DocTree: docTree,
	}
}

// Start 启动文件监控
func (fw *FileWatcher) Start() error {
	var err error

	// 创建新的监控器
	fw.Watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	// 递归监控文档目录
	if err = fw.watchRecursive(fw.DocTree.Config.DocsDir); err != nil {
		fw.Watcher.Close()
		return err
	}

	// 设置运行状态
	fw.IsRunning = true

	// 启动监控goroutine
	go fw.run()

	log.Println("文件监控已启动，当文档文件或目录变化时将自动更新目录树")
	log.Println("提示: 请刷新页面查看最新的目录结构")

	return nil
}

// run 运行监控循环
func (fw *FileWatcher) run() {
	defer fw.stop()

	for {
		select {
		case event, ok := <-fw.Watcher.Events:
			if !ok {
				return
			}

			// 忽略临时文件事件
			if fw.isTempFile(event.Name) {
				continue
			}

			// 记录文件变化
			log.Printf("检测到文件变化: %s, 操作: %s", event.Name, event.Op.String())

			// 重建文档树
			if err := fw.DocTree.Rebuild(); err != nil {
				log.Printf("重建文档树失败: %v", err)
			}

		case err, ok := <-fw.Watcher.Errors:
			if !ok {
				return
			}
			log.Printf("文件监控错误: %v", err)
		}
	}
}

// watchRecursive 递归添加目录到监控器
func (fw *FileWatcher) watchRecursive(dirPath string) error {
	// 添加当前目录到监控器
	if err := fw.Watcher.Add(dirPath); err != nil {
		return err
	}

	// 递归添加子目录
	return filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
			return fw.Watcher.Add(path)
		}
		return nil
	})
}

// isTempFile 判断是否为临时文件
func (fw *FileWatcher) isTempFile(filePath string) bool {
	baseName := filepath.Base(filePath)
	return strings.Contains(filePath, ".swp") ||
		strings.Contains(filePath, "~$") ||
		strings.HasPrefix(baseName, ".")
}

// stop 停止监控
func (fw *FileWatcher) stop() {
	if fw.Watcher != nil {
		fw.Watcher.Close()
	}
	fw.IsRunning = false
	log.Println("文件监控已停止")
}

// Stop 手动停止监控
func (fw *FileWatcher) Stop() {
	if fw.IsRunning && fw.Watcher != nil {
		// 向监控通道发送关闭信号
		close(fw.Watcher.Events)
		close(fw.Watcher.Errors)
	}
}
