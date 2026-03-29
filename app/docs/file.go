package docs

import (
	"docs-go/pkg/doc"
	"docs-go/pkg/resp"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// FileStatic 处理静态文件请求
func (dh *DocHandler) FileStatic(c *gin.Context, pathUrl string) {
	staticFile := dh.buildDocPath(pathUrl)
	log.Printf("Static file: %s", pathUrl)
	c.File(staticFile)
}

// FileHtml 处理HTML文件请求
func (dh *DocHandler) FileHtml(c *gin.Context, pathUrl string) {
	normalizedPath := doc.NormalizePath(dh.Config.DocsDir, pathUrl)
	filePath := dh.buildDocPath(normalizedPath)
	log.Printf("HTML file: %s -> %s", pathUrl, normalizedPath)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resp.Err404(c)
		return
	}

	// 检查文档密码
	if !dh.checkDocPassword(c, filePath, pathUrl) {
		return // 密码验证未通过，已渲染密码页面
	}

	// 加载和渲染HTML文档
	document, err := doc.NewDocument(filePath)
	if err != nil {
		resp.Err500(c, err)
		return
	}

	htmlContent, err := document.Render()
	if err != nil {
		resp.Err500(c, err)
		return
	}

	// 直接返回HTML内容
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, string(htmlContent))
}

// FileMarkdown 处理Markdown文件请求
func (dh *DocHandler) FileMarkdown(c *gin.Context, pathUrl string) {
	normalizedPath := doc.NormalizePath(dh.Config.DocsDir, pathUrl)
	filePath := filepath.Join(dh.Config.DocsDir, normalizedPath)
	log.Printf("Markdown file: %s -> %s", pathUrl, normalizedPath)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resp.Err404(c)
		return
	}

	// 检查文档密码
	if !dh.checkDocPassword(c, filePath, pathUrl) {
		return // 密码验证未通过，已渲染密码页面
	}

	// 加载和渲染Markdown文档
	document, err := doc.NewDocument(filePath)
	if err != nil {
		resp.Err500(c, err)
		return
	}

	htmlContent, err := document.Render()
	if err != nil {
		resp.Err500(c, err)
		return
	}

	// 使用模板渲染Markdown内容
	siteTitle := dh.Config.SiteTitle
	if siteTitle == "" {
		siteTitle = "文档中心"
	}
	c.HTML(http.StatusOK, "doc.html", gin.H{
		"title":      document.GetTitle(),
		"content":    htmlContent,
		"docTree":    dh.DocTree.GetActiveTree(pathUrl),
		"basePath":   "/",
		"breadcrumb": doc.GenBreadcrumb(normalizedPath),
		"siteTitle":  siteTitle,
	})
}
