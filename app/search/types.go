package search

import (
	"docs-go/pkg/config"
	"docs-go/pkg/search"
)

// SearchHandler 搜索处理器
type SearchHandler struct {
	indexer   *search.Indexer
	searcher  *search.Searcher
	watcher   *search.SearchWatcher
	config    *config.Config
}

// SearchRequest 搜索请求
type SearchRequest struct {
	Query string `form:"query" json:"query" binding:"required"`
	Limit int    `form:"limit" json:"limit" default:"10"`
}

// SearchResponse 搜索响应
type SearchResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Data    *SearchResultData `json:"data,omitempty"`
}

// SearchResultData 搜索结果数据
type SearchResultData struct {
	Query   string                `json:"query"`
	Results []*search.SearchResult `json:"results"`
	Total   int                   `json:"total"`
}

// IndexUpdateRequest 索引更新请求
type IndexUpdateRequest struct {
	ForceRebuild bool `form:"force" json:"force"`
}

// IndexUpdateResponse 索引更新响应
type IndexUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Count   int    `json:"count,omitempty"`
}