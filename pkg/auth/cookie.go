package auth

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
)

// CookieManager Cookie管理器
type CookieManager struct {
	PasswordSalt string
}

// CookieInfo Cookie信息
type CookieInfo struct {
	Name     string
	Value    string
	Expire   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

// 常量
const (
	CookiePrefix = "doc_auth_"
	CookieExpire = 24
	CookieSalt   = "doc-cookie-salt-2025"
)

// NewCookieManager 创建Cookie管理器
func NewCookieManager() *CookieManager {
	return &CookieManager{
		PasswordSalt: CookieSalt,
	}
}

// 加密密码
func (cm *CookieManager) Encrypt(password string) string {
	hash := md5.Sum([]byte(password + cm.PasswordSalt))
	return hex.EncodeToString(hash[:])
}

// 获取Cookie名称
func (cm *CookieManager) GetCookieName(pagePath string) string {
	hash := md5.Sum([]byte(pagePath))
	return CookiePrefix + hex.EncodeToString(hash[:])
}

// 生成认证值
func (cm *CookieManager) GenerateAuthValue() string {
	return cm.Encrypt("authenticated")
}

// 创建Cookie对象
func (cm *CookieManager) CreateCookie(pagePath string) *http.Cookie {
	name := cm.GetCookieName(pagePath)
	value := cm.GenerateAuthValue()

	return &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   CookieExpire * 3600,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

// 验证Cookie值
func (cm *CookieManager) ValidateCookieValue(cookieValue string) bool {
	return cookieValue == cm.GenerateAuthValue()
}

// 获取Cookie信息
func (cm *CookieManager) GetCookieInfo(pagePath string) CookieInfo {
	return CookieInfo{
		Name:     cm.GetCookieName(pagePath),
		Value:    cm.GenerateAuthValue(),
		Expire:   CookieExpire,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}
}

// 检查Cookie是否过期
func (cm *CookieManager) IsCookieExpired(cookie *http.Cookie) bool {
	if cookie.Expires.IsZero() {
		return false
	}
	return cookie.Expires.Before(time.Now())
}
