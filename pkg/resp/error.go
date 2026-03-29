package resp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Err404 处理404错误
func Err404(c *gin.Context) {
	renderError(c, http.StatusNotFound,
		"404 页面未找到",
		fmt.Sprintf("文档文件 %s 不存在", c.Request.URL.Path),
	)
}

// Err500 处理500错误
func Err500(c *gin.Context, err error) {
	renderError(c, http.StatusInternalServerError,
		"500 内部服务器错误",
		fmt.Sprintf("服务器内部错误: %v", err),
	)
}

// renderError 渲染错误页面
func renderError(c *gin.Context, statusCode int, errorTitle, errorMessage string) {
	// 从上下文中获取 siteTitle，如果没有则使用默认值
	siteTitle, exists := c.Get("siteTitle")
	if !exists || siteTitle == "" {
		siteTitle = "文档中心"
	}
	c.HTML(statusCode, "error.html", gin.H{
		"error":     errorTitle,
		"message":   errorMessage,
		"siteTitle": siteTitle,
	})
}
