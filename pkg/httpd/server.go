package httpd

import (
	"docs-go/app"
	"log"

	"docs-go/pkg/config"
)

func StartServer(config *config.Config) {
	// 设置路由
	router := app.SetupApp(config)

	// 启动服务器
	log.Printf("文档服务器启动在 http://localhost:%s", config.Port)
	log.Printf("正在服务文档目录: %s", config.DocsDir)
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
