package doc

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Document 文档结构体
type Document struct {
	Content     []byte
	Frontmatter *Frontmatter
	FilePath    string
	FileInfo    os.FileInfo
	IsMarkdown  bool
	engine      goldmark.Markdown
}

// Frontmatter 文档前置元数据
type Frontmatter struct {
	Title    string `yaml:"title"`     // 文档标题
	TitleDir string `yaml:"title_dir"` // 目录标题
	Sort     int    `yaml:"sort"`      // 排序值
	Password string `yaml:"password"`  // 密码
}

// NewDocument 创建新文档
func NewDocument(filePath string) (*Document, error) {
	doc := &Document{
		FilePath:   filePath,
		IsMarkdown: strings.HasSuffix(strings.ToLower(filePath), ".md"),
	}

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	doc.Content = content

	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	doc.FileInfo = fileInfo

	// 解析frontmatter
	doc.parseFrontmatter()

	return doc, nil
}

// parseFrontmatter 解析文档的前置元数据
func (doc *Document) parseFrontmatter() {
	frontmatter := &Frontmatter{}

	// 简单的frontmatter解析逻辑
	lines := strings.Split(string(doc.Content), "\n")
	contentStart := 0

	// 检查第一行或第二行（兼容html标记前加注释）是否是frontmatter的开始分隔符
	hasFrontmatter := false
	frontmatterStart := 0

	for i := 0; i < 2 && i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "---" {
			hasFrontmatter = true
			frontmatterStart = i
			break
		}
		// 兼容HTML注释包围的frontmatter
		if line == "<!--" && i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "---" {
			hasFrontmatter = true
			frontmatterStart = i + 1
			break
		}
	}

	if hasFrontmatter {
		// 从frontmatter开始位置的下一个行开始解析
		for i := frontmatterStart + 1; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])

			// 如果遇到结束分隔符，标记内容开始位置并退出循环
			if line == "---" {
				contentStart = i + 1
				// 如果是HTML注释包围的frontmatter，需要跳过-->注释行
				if frontmatterStart > 0 && strings.TrimSpace(lines[frontmatterStart-1]) == "<!--" {
					contentStart = i + 2 // 跳过---和-->两行
				}
				break
			}

			// 解析frontmatter键值对
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					value = strings.Trim(value, `"'`)

					switch key {
					case "title":
						frontmatter.Title = value
					case "title_dir":
						frontmatter.TitleDir = value
					case "sort", "order":  // 同时支持sort和order字段
						if sortVal, err := strconv.Atoi(value); err == nil {
							frontmatter.Sort = sortVal
						}
					case "password":
						frontmatter.Password = value
					}
				}
		}
	}

	doc.Frontmatter = frontmatter

	// 如果有frontmatter，移除frontmatter部分
	if contentStart > 0 {
		doc.Content = []byte(strings.Join(lines[contentStart:], "\n"))
	}
}

// Render 渲染文档内容
func (doc *Document) Render() (template.HTML, error) {
	if doc.IsMarkdown {
		return doc.renderMarkdown()
	}
	return template.HTML(doc.Content), nil // HTML直接返回
}

// renderMarkdown 渲染Markdown文档
func (doc *Document) renderMarkdown() (template.HTML, error) {
	if doc.engine == nil {
		doc.engine = goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				highlighting.Highlighting,
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithUnsafe(),
			),
		)
	}

	var buf bytes.Buffer
	if err := doc.engine.Convert(doc.Content, &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

// GetTitle 获取文档标题
func (doc *Document) GetTitle() string {
	// 优先使用frontmatter中的title
	if doc.Frontmatter != nil && doc.Frontmatter.Title != "" {
		return doc.Frontmatter.Title
	}

	// 从路径中提取标题
	title := filepath.Base(doc.FilePath)
	if doc.IsMarkdown && strings.HasSuffix(title, ".md") {
		title = strings.TrimSuffix(title, ".md")
	} else if strings.HasSuffix(title, ".html") || strings.HasSuffix(title, ".htm") {
		title = strings.TrimSuffix(title, ".html")
		title = strings.TrimSuffix(title, ".htm")
	}

	if title == "index" {
		dirName := filepath.Base(filepath.Dir(doc.FilePath))
		if dirName != "." {
			title = dirName
		} else {
			title = "首页"
		}
	}

	return title
}

// GetFilePath 获取文档文件路径
func (doc *Document) GetFilePath() string {
	return doc.FilePath
}

// GetFileInfo 获取文件信息
func (doc *Document) GetFileInfo() os.FileInfo {
	return doc.FileInfo
}

// GetFrontmatter 获取前置元数据
func (doc *Document) GetFrontmatter() *Frontmatter {
	return doc.Frontmatter
}
