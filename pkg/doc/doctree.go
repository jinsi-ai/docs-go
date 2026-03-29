package doc

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"docs-go/pkg/config"
)

// DocNode 文档节点结构体
type DocNode struct {
	Name     string     `json:"name"`     // 节点名称
	Path     string     `json:"path"`     // 节点路径
	IsDir    bool       `json:"isDir"`    // 是否为目录
	Children []*DocNode `json:"children"` // 子节点
	Active   bool       `json:"active"`   // 是否激活状态
	Sort     int        `json:"-"`        // 排序权重（负数表示使用自然排序）
}

// DocTree 文档树管理结构体
type DocTree struct {
	Root   *DocNode       // 根节点
	Mutex  sync.RWMutex   // 用于保护并发访问
	Config *config.Config // 配置引用
}

// NewDocTree 创建新的文档树实例
func NewDocTree(config *config.Config) *DocTree {
	return &DocTree{
		Config: config,
	}
}

// Init 初始化文档树
func (dt *DocTree) Init() error {
	// 验证文档目录存在
	if err := dt.Config.Validate(); err != nil {
		return err
	}

	// 创建根节点
	dt.Root = &DocNode{
		Name:     "文档根目录",
		Path:     "/",
		Children: []*DocNode{},
		IsDir:    true,
	}

	// 构建文档树
	return dt.Build()
}

// Build 构建文档树
func (dt *DocTree) Build() error {
	return dt.buildRecursive(dt.Config.DocsDir, dt.Root)
}

// Rebuild 重建文档树
func (dt *DocTree) Rebuild() error {
	newRoot := &DocNode{
		Name:     "文档根目录",
		Path:     "/",
		Children: []*DocNode{},
		IsDir:    true,
	}

	if err := dt.buildRecursive(dt.Config.DocsDir, newRoot); err != nil {
		return err
	}

	dt.Mutex.Lock()
	dt.Root = newRoot
	dt.Mutex.Unlock()

	return nil
}

// GetActiveTree 获取标记了激活状态的文档树副本
func (dt *DocTree) GetActiveTree(activePath string) *DocNode {
	dt.Mutex.RLock()
	defer dt.Mutex.RUnlock()

	// 深拷贝并标记激活节点
	return dt.cloneAndMarkActive(dt.Root, activePath)
}

// buildRecursive 递归构建文档树
func (dt *DocTree) buildRecursive(dirPath string, parent *DocNode) error {
	// 读取目录内容
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// 跳过隐藏文件和目录
		if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "." {
			continue
		}

		// 构建完整路径
		path := filepath.Join(dirPath, entry.Name())
		relPath, err := filepath.Rel(dt.Config.DocsDir, path)
		if err != nil {
			continue
		}

		// 构建URL路径
		httpPath := "/" + strings.ReplaceAll(relPath, "\\", "/")
		if entry.IsDir() {
			httpPath += "/"
		}

		// 提取节点名称
		nodeName := entry.Name()
		if !entry.IsDir() {
			// 支持.md和.html文件，以及README.md
			ext := filepath.Ext(path)
			lowerName := strings.ToLower(nodeName)
			if ext == ".md" {
				nodeName = strings.TrimSuffix(nodeName, ".md")
				// 处理README.md作为目录索引
				if lowerName == "readme.md" {
					nodeName = "首页"
				}
			} else if ext == ".html" {
				nodeName = strings.TrimSuffix(nodeName, ".html")
			}
		}

		// 创建节点
		isDir := entry.IsDir()
		newNode := &DocNode{
			Name:     nodeName,
			Path:     httpPath,
			Children: []*DocNode{},
			IsDir:    isDir,
			Sort:     -1, // 默认使用自然排序
		}

		// 如果是Markdown或HTML文件，尝试解析frontmatter获取标题和排序信息
		if !isDir && (filepath.Ext(path) == ".md" || filepath.Ext(path) == ".html") {
			document, err := NewDocument(path)
			if err == nil {
				frontmatter := document.GetFrontmatter()
				if frontmatter != nil {
					// 优先使用frontmatter中的title作为节点名称
					if frontmatter.Title != "" {
						newNode.Name = frontmatter.Title
					}
					// 获取排序信息
					if frontmatter.Sort > 0 {
						newNode.Sort = frontmatter.Sort
					}
				}
			}
		} else if isDir {
			// 如果是目录，检查是否存在index.md、README.md或index.html文件并解析其frontmatter
			// 优先级: index.md > README.md > index.html
			var document *Document
			var err error

			// 先尝试读取index.md
			indexPath := filepath.Join(path, "index.md")
			if document, err = NewDocument(indexPath); err != nil {
				// 如果index.md不存在，尝试读取README.md
				readmePath := filepath.Join(path, "README.md")
				if document, err = NewDocument(readmePath); err != nil {
					// 如果README.md也不存在，尝试读取index.html
					indexPath = filepath.Join(path, "index.html")
					document, err = NewDocument(indexPath)
				}
			}

			// 如果任一文件读取成功，解析frontmatter
			if err == nil {
				frontmatter := document.GetFrontmatter()
				if frontmatter != nil {
					if frontmatter.TitleDir != "" {
						newNode.Name = frontmatter.TitleDir
					}
					// 获取目录的排序信息
					if frontmatter.Sort > 0 {
						newNode.Sort = frontmatter.Sort
					}
				}
			}
		}

		// 如果是目录，递归构建子目录
		if isDir {
			if err := dt.buildRecursive(path, newNode); err != nil {
				log.Printf("构建子目录树失败 %s: %v", path, err)
			}
		} else if filepath.Ext(path) != ".md" && filepath.Ext(path) != ".html" {
			// 只添加.md和.html文件
			continue
		}

		parent.Children = append(parent.Children, newNode)
	}

	// 对子节点进行排序
	dt.sortChildren(parent)

	return nil
}

// sortChildren 对子节点进行排序
func (dt *DocTree) sortChildren(parent *DocNode) {
	if len(parent.Children) == 0 {
		return
	}

	// 使用稳定排序，保持相同排序值的自然顺序
	sort.SliceStable(parent.Children, func(i, j int) bool {
		nodeI := parent.Children[i]
		nodeJ := parent.Children[j]

		// 特殊处理：根目录的index.md文件应该显示在最顶部（最高优先级）
		if parent.Path == "/" && (nodeI.Path == "/index.md" || nodeI.Path == "/index.html") && nodeI.Sort == 1 {
			return true
		}
		if parent.Path == "/" && (nodeJ.Path == "/index.md" || nodeJ.Path == "/index.html") && nodeJ.Sort == 1 {
			return false
		}

		// 比较排序值（次高优先级）
		if nodeI.Sort >= 0 && nodeJ.Sort >= 0 {
			// 两者都有排序值，按排序值排序
			return nodeI.Sort < nodeJ.Sort
		} else if nodeI.Sort >= 0 {
			// 只有i有排序值，i排在前面
			return true
		} else if nodeJ.Sort >= 0 {
			// 只有j有排序值，j排在后面
			return false
		}

		// 目录优先于文件（最低优先级）
		if nodeI.IsDir && !nodeJ.IsDir {
			return true
		}
		if !nodeI.IsDir && nodeJ.IsDir {
			return false
		}

		// 两者都没有排序值，按自然排序（文件名）
		return nodeI.Name < nodeJ.Name
	})
}

// cloneAndMarkActive 克隆节点并标记激活状态
func (dt *DocTree) cloneAndMarkActive(node *DocNode, activePath string) *DocNode {
	clone := &DocNode{
		Name:     node.Name,
		Path:     node.Path,
		Children: make([]*DocNode, len(node.Children)),
		IsDir:    node.IsDir,
		Active:   node.Path == activePath || (node.IsDir && strings.HasPrefix(activePath, node.Path)),
	}

	for i, child := range node.Children {
		clone.Children[i] = dt.cloneAndMarkActive(child, activePath)
	}

	return clone
}
