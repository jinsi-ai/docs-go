// Package webfs 提供web资源的嵌入支持
package webfs

import (
	"io/fs"
)

// FS 存储嵌入的文件系统
var (
	WebFS    fs.FS
	ViewsFS  fs.FS
	StaticFS fs.FS
)

// SetFS 设置嵌入的文件系统
// 在 main.go 中调用，传入 embed.FS
func SetFS(web fs.FS) {
	WebFS = web
	ViewsFS, _ = fs.Sub(web, "web/views")
	StaticFS, _ = fs.Sub(web, "web/static")
}

// IsAvailable 检查嵌入资源是否可用
func IsAvailable() bool {
	return WebFS != nil
}
