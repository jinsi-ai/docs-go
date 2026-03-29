package search

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// Searcher 搜索器实现
type Searcher struct {
	storage *Storage
}

// NewSearcher 创建新的搜索器
func NewSearcher(storage *Storage) *Searcher {
	return &Searcher{
		storage: storage,
	}
}

// Search 执行搜索
func (s *Searcher) Search(req *SearchRequest) (*SearchResponse, error) {
	startTime := time.Now()

	// 清理查询字符串
	query := strings.TrimSpace(req.Query)
	if query == "" {
		return &SearchResponse{
			Results:    []*SearchResult{},
			TotalCount: 0,
			QueryTime:  time.Since(startTime),
		}, nil
	}

	// 设置默认限制
	limit := req.Limit
	if limit <= 0 {
		limit = 20
	}

	// 使用SQLite FTS5进行全文搜索
	results, err := s.storage.FullTextSearch(query, limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("全文搜索失败: %v", err)
	}

	// 获取总结果数
	totalCount, err := s.storage.GetTotalSearchCount(query)
	if err != nil {
		totalCount = len(results) // 如果获取总数失败，使用当前结果数
	}

	// 如果没有找到结果，尝试使用倒排索引进行更精确的搜索
	if len(results) == 0 {
		results, err = s.searchWithInvertedIndex(query, limit, req.Offset)
		if err != nil {
			return nil, fmt.Errorf("倒排索引搜索失败: %v", err)
		}

		// 重新计算总数
		totalCount = len(results)
	}

	// 为结果生成摘要（如果FTS没有提供）
	for _, result := range results {
		if result.Excerpt == "" {
			result.Excerpt = s.generateExcerpt(result.Document.Content, query)
		}

		// 计算关键词位置
		result.Positions = s.findKeywordPositions(result.Document.Content, query)
	}

	return &SearchResponse{
		Results:    results,
		TotalCount: totalCount,
		QueryTime:  time.Since(startTime),
	}, nil
}

// searchWithInvertedIndex 使用倒排索引进行搜索
func (s *Searcher) searchWithInvertedIndex(query string, limit int, offset int) ([]*SearchResult, error) {
	// 分词处理
	terms := s.tokenizeText(query)

	if len(terms) == 0 {
		return []*SearchResult{}, nil
	}

	// 获取所有文档
	allDocs, err := s.storage.GetAllDocuments()
	if err != nil {
		return nil, err
	}

	var results []*SearchResult

	// 对每个文档进行评分
	for _, doc := range allDocs {
		score := s.calculateRelevanceScore(doc, terms)

		if score > 0 {
			results = append(results, &SearchResult{
				Document: doc,
				Excerpt:  s.generateExcerpt(doc.Content, query),
				Score:    score,
			})
		}
	}

	// 按分数排序
	sortResultsByScore(results)

	// 应用分页
	start := offset
	if start > len(results) {
		start = len(results)
	}
	end := start + limit
	if end > len(results) {
		end = len(results)
	}

	if start >= len(results) {
		return []*SearchResult{}, nil
	}

	return results[start:end], nil
}

// calculateRelevanceScore 计算文档相关性分数
func (s *Searcher) calculateRelevanceScore(doc *DocumentIndex, terms []string) float64 {
	var score float64

	content := strings.ToLower(doc.Content)
	title := strings.ToLower(doc.Title)

	for _, term := range terms {
		if s.isStopWord(term) {
			continue
		}

		term = strings.ToLower(term)

		// 标题匹配权重更高
		if strings.Contains(title, term) {
			score += 10.0
		}

		// 内容匹配
		if strings.Contains(content, term) {
			score += 1.0

			// 计算词频（简单实现）
			frequency := strings.Count(content, term)
			score += float64(frequency) * 0.1
		}
	}

	return score
}

// safeTruncate 安全截断字符串，避免截断UTF-8字符
func safeTruncate(s string, maxBytes int) string {
	if len(s) <= maxBytes {
		return s
	}

	// 确保不会截断UTF-8字符
	for maxBytes > 0 {
		if utf8.RuneStart(s[maxBytes-1]) {
			break
		}
		maxBytes--
	}

	return s[:maxBytes]
}

// generateExcerpt 生成包含关键词的摘要
func (s *Searcher) generateExcerpt(content string, query string) string {
	content = strings.ToLower(content)
	query = strings.ToLower(query)

	// 查找关键词位置
	index := strings.Index(content, query)
	if index == -1 {
		// 如果没有找到完整匹配，尝试分词匹配
		terms := s.tokenizeText(query)
		for _, term := range terms {
			if !s.isStopWord(term) {
				index = strings.Index(content, strings.ToLower(term))
				if index != -1 {
					break
				}
			}
		}
	}

	if index == -1 {
		// 返回内容的前400个字符（约200字），使用安全截断
		if len(content) > 400 {
			return safeTruncate(content, 400) + "..."
		}
		return content
	}

	// 提取关键词周围的文本 - 扩大范围到前后各100个字符
	start := index - 100
	if start < 0 {
		start = 0
	}

	end := index + len(query) + 100
	if end > len(content) {
		end = len(content)
	}

	// 确保截断位置不会破坏UTF-8字符
	start = s.findSafeStart(content, start)
	end = s.findSafeEnd(content, end)

	excerpt := content[start:end]

	// 添加省略号
	if start > 0 {
		excerpt = "..." + excerpt
	}
	if end < len(content) {
		excerpt = excerpt + "..."
	}

	// 高亮关键词
	excerpt = strings.ReplaceAll(excerpt, query, "<mark>"+query+"</mark>")

	return excerpt
}

// findSafeStart 找到安全的起始位置，避免截断UTF-8字符
func (s *Searcher) findSafeStart(content string, start int) int {
	for start < len(content) {
		if utf8.RuneStart(content[start]) {
			return start
		}
		start++
	}
	return len(content)
}

// findSafeEnd 找到安全的结束位置，避免截断UTF-8字符
func (s *Searcher) findSafeEnd(content string, end int) int {
	for end > 0 && end < len(content) {
		if utf8.RuneStart(content[end]) {
			return end
		}
		end--
	}
	return end
}

// findKeywordPositions 查找关键词位置
func (s *Searcher) findKeywordPositions(content string, query string) []int {
	var positions []int

	content = strings.ToLower(content)
	query = strings.ToLower(query)

	// 查找所有出现位置
	start := 0
	for {
		index := strings.Index(content[start:], query)
		if index == -1 {
			break
		}

		positions = append(positions, start+index)
		start = start + index + len(query)
	}

	return positions
}

// sortResultsByScore 按分数排序结果
func sortResultsByScore(results []*SearchResult) {
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

// tokenizeText 文本分词
func (s *Searcher) tokenizeText(text string) []string {
	// 使用简单的空格分词
	words := strings.Fields(text)

	// 过滤和处理词条
	terms := make([]string, 0, len(words))
	seen := make(map[string]bool)

	for _, word := range words {
		// 清理词条
		word = strings.TrimSpace(strings.ToLower(word))

		// 过滤停用词，支持单个字母的关键词（如编程语言名称：Go, C, R等）
		if len(word) < 1 || s.isStopWord(word) {
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
func (s *Searcher) isStopWord(word string) bool {
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

// Close 关闭搜索器
func (s *Searcher) Close() {
	// 不需要特殊处理，搜索器关闭时自动清理
}
