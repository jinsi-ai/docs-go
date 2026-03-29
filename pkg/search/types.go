package search

import (
	"time"
)

// DocumentIndex 文档索引结构
type DocumentIndex struct {
	ID        string    `json:"id"`        // 文档路径作为ID
	Title     string    `json:"title"`     // 文档标题
	Content   string    `json:"content"`   // 文档内容（不含密码）
	Path      string    `json:"path"`      // 文档路径
	UpdatedAt time.Time `json:"updated_at"`// 更新时间
}

// SearchResult 搜索结果结构
type SearchResult struct {
	Document  *DocumentIndex `json:"document"`
	Excerpt   string         `json:"excerpt"`   // 包含关键词的摘要
	Positions []int          `json:"positions"` // 关键词位置（用于锚点定位）
	Score     float64        `json:"score"`     // 相关性分数
}

// SearchRequest 搜索请求
type SearchRequest struct {
	Query     string `json:"query"`     // 搜索关键词
	Limit     int    `json:"limit"`     // 结果数量限制
	Offset    int    `json:"offset"`    // 分页偏移量
}

// SearchResponse 搜索响应
type SearchResponse struct {
	Results    []*SearchResult `json:"results"`
	TotalCount int             `json:"total_count"`
	QueryTime  time.Duration   `json:"query_time"`
}

// IndexStats 索引统计信息
type IndexStats struct {
	TotalDocuments int       `json:"total_documents"`
	LastUpdated    time.Time `json:"last_updated"`
	IndexSize      int64     `json:"index_size"`
}

// IndexUpdateType 索引更新类型
const (
	UpdateTypeFull    = "full"    // 全量更新
	UpdateTypeIncremental = "incremental" // 增量更新
)

// IndexUpdateEvent 索引更新事件
type IndexUpdateEvent struct {
	Type      string    `json:"type"`      // 更新类型
	Paths     []string  `json:"paths"`     // 更新的文件路径
	Timestamp time.Time `json:"timestamp"` // 更新时间
}