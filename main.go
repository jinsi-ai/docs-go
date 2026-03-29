package main

import (
	"embed"
	"log"

	"docs-go/pkg/config"
	"docs-go/pkg/httpd"
	"docs-go/pkg/webfs"
)

//go:embed web
var webFS embed.FS

func init() {
	// 初始化嵌入的web资源
	webfs.SetFS(webFS)
}

func main() {
	// 初始化配置
	config := &config.Config{
		DocsDir: "docs",
		Port:    "8080",
	}

	// 先加载.env文件配置
	if err := config.LoadEnv(); err != nil {
		log.Printf("警告: 加载.env文件失败: %v", err)
	}

	// 再解析命令行参数（命令行参数优先级最高）
	config.ParseFlags()

	// 启动HTTP服务器
	httpd.StartServer(config)
}
