package docs

import (
	"docs-go/pkg/doc"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 站点密码验证请求
type SitePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

// 密码验证请求
type PasswordRequest struct {
	PagePath string `json:"pagePath" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CheckSitePassword 处理站点密码验证请求
func (dh *DocHandler) CheckSitePassword(c *gin.Context) {
	// 检查是否设置了站点密码
	if dh.Config.PasswordSite == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "站点密码验证成功",
		})
		return
	}

	var req SitePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
		})
		return
	}

	// 验证站点密码
	if dh.verifySitePassword(req.Password) {
		// 设置站点Cookie
		dh.setSiteCookie(c)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "站点密码验证成功",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "站点密码错误",
		})
	}
}

// CheckPassword 处理密码验证请求
func (dh *DocHandler) CheckDocPassword(c *gin.Context) {
	var req PasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
		})
		return
	}

	if req.PagePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "文档路径不能为空",
		})
		return
	}

	// 验证页面路径
	path := doc.NormalizePath(dh.Config.DocsDir, req.PagePath)
	filePath := dh.buildDocPath(path)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "文档不存在",
		})
		return
	}

	// 加载文档
	document, err := doc.NewDocument(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "文档加载失败",
		})
		return
	}

	// 验证密码
	if dh.verifyPassword(document, req.Password) {
		// 设置Cookie
		dh.setCookie(c, req.PagePath)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "密码验证成功",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "密码错误",
		})
	}
}

// 检查文档密码
func (dh *DocHandler) checkDocPassword(c *gin.Context, filePath string, pagePath string) bool {
	// 加载文档
	document, err := doc.NewDocument(filePath)
	if err != nil {
		return true // 加载失败时默认允许访问
	}

	// 检查是否有密码设置
	fm := document.GetFrontmatter()
	if fm == nil || fm.Password == "" {
		return true // 没有设置密码，直接通过
	}

	// 检查Cookie
	if dh.checkCookie(c, pagePath) {
		return true // Cookie验证通过
	}

	// 需要密码验证，渲染密码页面
	c.HTML(http.StatusOK, "password.html", gin.H{
		"pagePath": pagePath,
		"docTree":  dh.DocTree.GetActiveTree(pagePath),
	})
	return false
}

// 验证站点密码
func (dh *DocHandler) verifySitePassword(password string) bool {
	if dh.Config.PasswordSite == "" {
		return true // 没有设置站点密码，直接通过
	}

	// 检查是否匹配加密后的密码或明文密码
	encrypted := dh.CookieManager.Encrypt(password)
	return encrypted == dh.Config.PasswordSite || password == dh.Config.PasswordSite
}

// 验证密码
func (dh *DocHandler) verifyPassword(document *doc.Document, password string) bool {
	fm := document.GetFrontmatter()
	if fm == nil || fm.Password == "" {
		return true // 没有设置密码，直接通过
	}

	// 检查是否匹配加密后的密码或明文密码
	encrypted := dh.CookieManager.Encrypt(password)
	return encrypted == fm.Password || password == fm.Password
}

// 检查Cookie
func (dh *DocHandler) checkCookie(c *gin.Context, pagePath string) bool {
	name := dh.CookieManager.GetCookieName(pagePath)
	cookie, err := c.Cookie(name)
	if err != nil {
		return false
	}
	return dh.CookieManager.ValidateCookieValue(cookie)
}

// 检查站点Cookie
func (dh *DocHandler) checkSiteCookie(c *gin.Context) bool {
	name := dh.CookieManager.GetCookieName("site")
	cookie, err := c.Cookie(name)
	if err != nil {
		return false
	}
	return dh.CookieManager.ValidateCookieValue(cookie)
}

// 设置站点Cookie
func (dh *DocHandler) setSiteCookie(c *gin.Context) {
	name := dh.CookieManager.GetCookieName("site")
	value := dh.CookieManager.GenerateAuthValue()

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   24 * 3600, // 24小时
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, cookie)
}

// 设置Cookie
func (dh *DocHandler) setCookie(c *gin.Context, pagePath string) {
	cookie := dh.CookieManager.CreateCookie(pagePath)
	http.SetCookie(c.Writer, cookie)
}
