package docs

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"docs-go/pkg/auth"
	"docs-go/pkg/config"
	"docs-go/pkg/doc"

	"github.com/gin-gonic/gin"
)

// DocHandler 文档处理器结构体
type DocHandler struct {
	Config        *config.Config
	DocTree       *doc.DocTree
	CookieManager *auth.CookieManager
}

// NewDocHandler 创建新的文档处理器
func NewDocHandler(config *config.Config, docTree *doc.DocTree) *DocHandler {
	return &DocHandler{
		Config:        config,
		DocTree:       docTree,
		CookieManager: auth.NewCookieManager(),
	}
}

// buildDocPath 构建文档完整路径
func (dh *DocHandler) buildDocPath(path string) string {
	return filepath.Join(dh.Config.DocsDir, path)
}

// HandleDocPage 处理文档页面请求
func (dh *DocHandler) HandleDocPage(c *gin.Context) {
	// 设置 siteTitle 到上下文中，供错误页面使用
	siteTitle := dh.Config.SiteTitle
	if siteTitle == "" {
		siteTitle = "文档中心"
	}
	c.Set("siteTitle", siteTitle)

	pathUrl := c.Request.URL.Path
	ext := strings.ToLower(filepath.Ext(pathUrl))
	log.Printf("Request path: %s, File extension: %s", pathUrl, ext)

	// 检查是否为搜索页面
	if pathUrl == "/search" {
		dh.SearchPage(c)
		return
	}

	// 检查站点密码验证
	if !dh.checkSiteAccess(c) {
		return
	}

	// 根据文件类型调用相应的处理函数
	switch {
	case ext == "" || ext == ".md":
		// Markdown文件处理
		dh.FileMarkdown(c, pathUrl)
	case ext == ".html" || ext == ".htm":
		// HTML文件处理
		dh.FileHtml(c, pathUrl)
	default:
		// 非文档类型文件，使用静态文件处理
		dh.FileStatic(c, pathUrl)
	}
}

// checkSiteAccess 检查站点访问权限
func (dh *DocHandler) checkSiteAccess(c *gin.Context) bool {
	// 如果没有设置站点密码，直接通过
	if dh.Config.PasswordSite == "" {
		return true
	}

	// 检查站点Cookie
	if dh.checkSiteCookie(c) {
		return true
	}

	// 需要站点密码验证，渲染站点密码页面
	siteTitle := dh.Config.SiteTitle
	if siteTitle == "" {
		siteTitle = "文档中心"
	}
	c.HTML(http.StatusOK, "password.html", gin.H{
		"pagePath":       "站点密码",
		"docTree":        dh.DocTree.GetActiveTree("/"),
		"isSitePassword": true,
		"siteTitle":      siteTitle,
	})
	return false
}

// SearchPage 渲染搜索页面
func (dh *DocHandler) SearchPage(c *gin.Context) {
	query := c.Query("q")

	siteTitle := dh.Config.SiteTitle
	if siteTitle == "" {
		siteTitle = "文档中心"
	}
	c.HTML(http.StatusOK, "search.html", gin.H{
		"title":     "搜索文档",
		"query":     query,
		"docTree":   dh.DocTree.GetActiveTree("/"),
		"results":   nil,
		"total":     0,
		"siteTitle": siteTitle,
	})
}
