package search

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"docs-go/pkg/config"
	"docs-go/pkg/resp"
	"docs-go/pkg/search"

	"github.com/gin-gonic/gin"
)

// NewSearchHandler 创建新的搜索处理器
func NewSearchHandler(config *config.Config, appDir string) (*SearchHandler, error) {
	dataDir := filepath.Join(appDir, "data")
	storage, err := search.NewStorage(dataDir)
	if err != nil {
		return nil, fmt.Errorf("初始化SQLite存储失败: %v", err)
	}

	// 初始化SQLite索引器
	indexer := search.NewIndexer(storage, config.DocsDir)

	// 初始化SQLite搜索器
	searcher := search.NewSearcher(storage)

	// 初始化并启动文件监听器
	watcher := search.NewSearchWatcher(indexer)
	if err := watcher.Start(config.DocsDir); err != nil {
		log.Printf("警告: 启动搜索文件监听器失败: %v", err)
	}

	handler := &SearchHandler{
		indexer:  indexer,
		searcher: searcher,
		watcher:  watcher,
		config:   config,
	}

	return handler, nil
}

// Search 处理搜索请求
func (h *SearchHandler) Search(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.BadRequest(c, "无效的请求参数: "+err.Error())
		return
	}

	// 设置默认限制
	if req.Limit <= 0 {
		req.Limit = 10
	}

	// 执行搜索
	searchResp, err := h.searcher.Search(&search.SearchRequest{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: 0,
	})
	if err != nil {
		log.Printf("搜索失败: %v", err)
		resp.InternalServerError(c, "搜索失败: "+err.Error())
		return
	}

	resp.Success(c, "搜索成功", &SearchResultData{
		Query:   req.Query,
		Results: searchResp.Results,
		Total:   searchResp.TotalCount,
	})
}

// UpdateIndex 更新搜索索引
func (h *SearchHandler) UpdateIndex(c *gin.Context) {
	var req IndexUpdateRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.BadRequest(c, "无效的请求参数: "+err.Error())
		return
	}

	var err error

	if req.ForceRebuild {
		// 强制重建索引
		err = h.indexer.BuildFullIndex()
	} else {
		// 增量更新索引 - 暂时只支持全量更新
		err = h.indexer.BuildFullIndex()
	}

	if err != nil {
		log.Printf("索引更新失败: %v", err)
		resp.InternalServerError(c, "索引更新失败: "+err.Error())
		return
	}

	// 获取更新后的文档数量
	count, err := h.indexer.GetDocumentCount()
	if err != nil {
		log.Printf("获取文档数量失败: %v", err)
	}

	resp.Success(c, "索引更新成功", gin.H{
		"count": count,
	})
}

// GetIndexStatus 获取索引状态
func (h *SearchHandler) GetIndexStatus(c *gin.Context) {
	count, err := h.indexer.GetDocumentCount()
	if err != nil {
		resp.InternalServerError(c, "获取索引状态失败: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"document_count": count,
			"index_ready":    count > 0,
		},
	})
}

// RegisterRoutes 注册搜索路由
func (h *SearchHandler) RegisterRoutes(router *gin.Engine) {
	searchGroup := router.Group("/api/search")
	{
		searchGroup.GET("/", h.Search)
		searchGroup.POST("/update", h.UpdateIndex)
		searchGroup.GET("/status", h.GetIndexStatus)
	}
}

// UpdateIndexOnStartup 启动时全量更新搜索索引
func (h *SearchHandler) UpdateIndexOnStartup() error {
	log.Println("开始全量构建搜索索引...")

	// 执行全量索引构建
	if err := h.indexer.BuildFullIndex(); err != nil {
		return fmt.Errorf("全量索引构建失败: %v", err)
	}

	// 获取索引文档数量
	count, err := h.indexer.GetDocumentCount()
	if err != nil {
		log.Printf("获取索引文档数量失败: %v", err)
	} else {
		log.Printf("搜索索引构建完成，共索引 %d 个文档", count)
	}

	return nil
}

// Close 关闭搜索处理器资源
func (h *SearchHandler) Close() {
	if h.watcher != nil {
		h.watcher.Stop()
	}
	if h.indexer != nil {
		h.indexer.Close()
	}
	if h.searcher != nil {
		h.searcher.Close()
	}
}
