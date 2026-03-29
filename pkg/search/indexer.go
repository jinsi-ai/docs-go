package search

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Indexer 索引器实现
type Indexer struct {
	storage *Storage
	docsDir string
}

// NewIndexer 创建新的索引器
func NewIndexer(storage *Storage, docsDir string) *Indexer {
	return &Indexer{
		storage: storage,
		docsDir: docsDir,
	}
}

// BuildFullIndex 构建全量索引
func (i *Indexer) BuildFullIndex() error {
	log.Printf("开始构建全量索引...")
	startTime := time.Now()

	// 清除现有索引
	if err := i.storage.ClearAll(); err != nil {
		return fmt.Errorf("清除现有索引失败: %v", err)
	}

	// 遍历文档目录
	var indexedCount int
	err := filepath.Walk(i.docsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理.md文件
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".md") {
			if err := i.indexDocument(path); err != nil {
				log.Printf("索引文档失败 %s: %v", path, err)
				// 继续处理其他文件
				return nil
			}
			indexedCount++
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("遍历文档目录失败: %v", err)
	}

	duration := time.Since(startTime)
	log.Printf("全量索引构建完成，共索引 %d 个文档，耗时 %v", indexedCount, duration)

	return nil
}

// UpdateIncrementalIndex 更新增量索引
func (i *Indexer) UpdateIncrementalIndex(paths []string) error {
	log.Printf("开始更新增量索引，处理 %d 个文件", len(paths))
	startTime := time.Now()

	var processedCount int
	for _, path := range paths {
		// 检查文件是否存在
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// 文件不存在，从索引中删除
			docID := generateDocumentID(path)
			if err := i.storage.DeleteDocument(docID); err != nil {
				log.Printf("删除文档失败 %s: %v", path, err)
			}
			processedCount++
			continue
		}

		// 只处理.md文件
		if strings.HasSuffix(strings.ToLower(path), ".md") {
			if err := i.indexDocument(path); err != nil {
				log.Printf("索引文档失败 %s: %v", path, err)
				continue
			}
			processedCount++
		}
	}

	duration := time.Since(startTime)
	log.Printf("增量索引更新完成，处理 %d 个文件，耗时 %v", processedCount, duration)

	return nil
}

// indexDocument 索引单个文档
func (i *Indexer) indexDocument(filePath string) error {
	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	// 解析文档标题和内容
	title, contentText := parseMarkdownContent(string(content))

	// 生成文档ID
	docID := generateDocumentID(filePath)

	// 将文件路径转换为相对路径（去除/docs前缀）
	relPath, err := filepath.Rel(i.docsDir, filePath)
	if err != nil {
		// 如果转换失败，使用原始路径
		relPath = filePath
	}
	// 转换为HTTP路径格式
	httpPath := "/" + strings.ReplaceAll(relPath, "\\", "/")

	// 创建文档索引
	doc := &DocumentIndex{
		ID:        docID,
		Title:     title,
		Content:   contentText,
		Path:      httpPath,
		UpdatedAt: time.Now(),
	}

	// 保存文档到存储
	if err := i.storage.SaveDocument(doc); err != nil {
		return fmt.Errorf("保存文档失败: %v", err)
	}

	// 构建倒排索引
	if err := i.buildInvertedIndex(doc); err != nil {
		return fmt.Errorf("构建倒排索引失败: %v", err)
	}

	return nil
}

// buildInvertedIndex 构建倒排索引
func (i *Indexer) buildInvertedIndex(doc *DocumentIndex) error {
	// 分词处理
	terms := i.tokenizeText(doc.Content)
	titleTerms := i.tokenizeText(doc.Title)

	// 合并所有词汇（包括标题和内容）
	allTerms := make(map[string][]int)

	// 处理标题词汇
	for _, term := range titleTerms {
		if i.isStopWord(term) {
			continue
		}
		// 标题词汇权重更高
		allTerms[term] = []int{-1} // 使用-1表示标题位置
	}

	// 处理内容词汇
	for pos, term := range terms {
		if i.isStopWord(term) {
			continue
		}
		if positions, exists := allTerms[term]; exists {
			allTerms[term] = append(positions, pos)
		} else {
			allTerms[term] = []int{pos}
		}
	}

	// 保存倒排索引
	for term, positions := range allTerms {
		if err := i.storage.SaveInvertedIndex(term, doc.ID, positions); err != nil {
			return err
		}
	}

	return nil
}

// GetDocumentCount 获取文档数量
func (i *Indexer) GetDocumentCount() (int, error) {
	stats, err := i.storage.GetStats()
	if err != nil {
		return 0, err
	}
	return stats.TotalDocuments, nil
}

// GetStats 获取索引统计信息
func (i *Indexer) GetStats() (*IndexStats, error) {
	return i.storage.GetStats()
}

// removeDocument 从索引中删除文档
func (i *Indexer) removeDocument(filePath string) error {
	docID := generateDocumentID(filePath)
	return i.storage.DeleteDocument(docID)
}

// isDocumentFile 判断是否为文档文件
func (i *Indexer) isDocumentFile(filePath string) bool {
	return strings.HasSuffix(strings.ToLower(filePath), ".md")
}

// Close 关闭索引器
func (i *Indexer) Close() {
	// 不需要特殊处理，索引器关闭时自动清理
}

// tokenizeText 文本分词
func (i *Indexer) tokenizeText(text string) []string {
	// 使用简单的空格分词
	words := strings.Fields(text)

	// 过滤和处理词条
	terms := make([]string, 0, len(words))
	seen := make(map[string]bool)

	for _, word := range words {
		// 清理词条
		word = strings.TrimSpace(strings.ToLower(word))

		// 过滤停用词，支持单个字母的关键词（如编程语言名称：Go, C, R等）
		if len(word) < 1 || i.isStopWord(word) {
			continue
		}

		// 去重
		if !seen[word] {
			terms = append(terms, word)
			seen[word] = true
		}
	}

	return terms
}

// isStopWord 判断是否为停用词
func (i *Indexer) isStopWord(word string) bool {
	// 简单的停用词列表，可以根据需要扩展
	stopWords := map[string]bool{
		"的": true, "了": true, "在": true, "是": true, "我": true,
		"有": true, "和": true, "就": true, "不": true, "人": true,
		"都": true, "一": true, "一个": true, "上": true, "也": true,
		"很": true, "到": true, "说": true, "要": true, "去": true,
		"你": true, "会": true, "着": true, "没有": true, "看": true,
		"好": true, "自己": true, "这": true, "那": true, "这个": true,
		"那个": true, "但是": true, "因为": true, "所以": true, "如果": true,
		"the": true, "and": true, "or": true, "but": true, "is": true,
		"are": true, "was": true, "were": true, "be": true, "been": true,
		"have": true, "has": true, "had": true, "do": true, "does": true,
		"did": true, "will": true, "would": true, "could": true, "should": true,
		"can": true, "may": true, "might": true, "must": true,
	}

	return stopWords[word]
}

// parseMarkdownContent 解析Markdown内容，提取标题和纯文本，过滤掉frontmatter中的敏感信息
func parseMarkdownContent(content string) (string, string) {
	lines := strings.Split(content, "\n")

	var title string
	var contentLines []string
	contentStart := 0

	// 检查是否有frontmatter
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

	// 如果存在frontmatter，跳过整个frontmatter部分
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
		}
	}

	// 从内容开始位置处理文档
	for i := contentStart; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// 提取标题（以#开头的行）
		if strings.HasPrefix(line, "#") && title == "" {
			title = strings.TrimSpace(strings.TrimLeft(line, "#"))
		}

		// 跳过代码块标记和空行
		if line == "" || strings.HasPrefix(line, "```") {
			continue
		}

		// 只保留文本内容，移除所有Markdown标记
		cleanLine := extractPlainText(line)
		if cleanLine != "" {
			contentLines = append(contentLines, cleanLine)
		}
	}

	// 如果没有找到标题，使用文件名
	if title == "" {
		title = "未命名文档"
	}

	contentText := strings.Join(contentLines, " ")
	return title, contentText
}

// extractPlainText 提取纯文本内容，移除所有Markdown和HTML标记
func extractPlainText(line string) string {
	// 移除标题标记
	line = strings.TrimLeft(line, "#")
	line = strings.TrimSpace(line)

	// 移除所有Markdown格式标记
	line = removeAllMarkdownMarkup(line)

	// 移除HTML标签（如果有）
	line = removeHTMLTags(line)

	return strings.TrimSpace(line)
}

// removeAllMarkdownMarkup 移除所有Markdown标记
func removeAllMarkdownMarkup(line string) string {
	// 移除粗体、斜体标记
	line = strings.ReplaceAll(line, "**", "")
	line = strings.ReplaceAll(line, "__", "")
	line = strings.ReplaceAll(line, "*", "")
	line = strings.ReplaceAll(line, "_", "")

	// 移除删除线标记
	line = strings.ReplaceAll(line, "~~", "")

	// 移除代码标记
	line = strings.ReplaceAll(line, "`", "")

	// 移除引用标记
	line = strings.TrimLeft(line, ">")

	// 移除链接标记 [text](url) -> text
	if strings.Contains(line, "[") && strings.Contains(line, "]") && strings.Contains(line, "(") && strings.Contains(line, ")") {
		// 提取链接文本
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")
		if start >= 0 && end > start {
			linkText := line[start+1 : end]
			// 移除整个链接标记，只保留文本
			line = line[:start] + linkText
		}
	}

	// 移除图片标记 ![alt](url) -> alt
	if strings.Contains(line, "![") && strings.Contains(line, "]") && strings.Contains(line, "(") && strings.Contains(line, ")") {
		// 提取alt文本
		start := strings.Index(line, "![")
		if start >= 0 {
			altStart := start + 2
			altEnd := strings.Index(line[altStart:], "]")
			if altEnd >= 0 {
				altText := line[altStart : altStart+altEnd]
				// 移除整个图片标记，只保留alt文本
				line = line[:start] + altText
			}
		}
	}

	return line
}

// removeHTMLTags 移除HTML标签
func removeHTMLTags(line string) string {
	// 简单的HTML标签移除
	var result strings.Builder
	inTag := false

	for _, char := range line {
		if char == '<' {
			inTag = true
			continue
		}
		if char == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteRune(char)
		}
	}

	return result.String()
}

// generateDocumentID 生成文档ID
func generateDocumentID(filePath string) string {
	// 使用文件路径作为ID，确保唯一性
	return filePath
}
