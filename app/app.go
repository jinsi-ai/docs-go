package app

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"docs-go/app/docs"
	"docs-go/app/search"
	"docs-go/pkg/config"
	"docs-go/pkg/doc"
	"docs-go/pkg/webfs"

	"github.com/gin-gonic/gin"
)

// SetupApp 设置所有路由
func SetupApp(config *config.Config) *gin.Engine {
	// 初始化文档树
	docTree := doc.NewDocTree(config)
	if err := docTree.Init(); err != nil {
		panic("初始化文档树失败: " + err.Error())
	}

	// 初始化并启动文件监控器
	fileWatcher := doc.NewFileWatcher(docTree, config)
	if err := fileWatcher.Start(); err != nil {
		panic("启动文件监控器失败: " + err.Error())
	}

	// 初始化文档处理器
	docHandler := docs.NewDocHandler(config, docTree)

	// 初始化搜索处理器 - 使用应用程序目录存储索引数据
	searchHandler, err := search.NewSearchHandler(config, ".")
	if err != nil {
		log.Printf("警告: 初始化搜索处理器失败: %v", err)
	}
	// 如果data/search.db不存在，则全量更新索引
	indexPath := filepath.Join("data", "search.db")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		log.Printf("提示: 搜索索引文件 %s 不存在，进行全量更新", indexPath)
		if searchHandler != nil {
			if err := searchHandler.UpdateIndexOnStartup(); err != nil {
				log.Printf("警告: 全量更新搜索索引失败: %v", err)
			} else {
				log.Println("搜索索引全量更新完成")
			}
		}
	}

	router := gin.Default()

	// 根据配置决定使用嵌入资源还是文件系统
	if config.EmbedWeb && webfs.IsAvailable() {
		// 使用嵌入的HTML模板
		tmpl, err := template.ParseFS(webfs.ViewsFS, "*.html")
		if err != nil {
			panic("加载嵌入模板失败: " + err.Error())
		}
		router.SetHTMLTemplate(tmpl)
		// 使用嵌入的静态文件
		router.StaticFS("/static", http.FS(webfs.StaticFS))
		log.Printf("使用嵌入的web资源")
	} else {
		// 从文件系统加载
		router.LoadHTMLGlob(filepath.Join("./web/views", "*.html"))
		router.Static("/static", "./web/static")
		log.Printf("使用文件系统的web资源")
	}

	// API路由
	router.POST("/api/auth/doc-password", docHandler.CheckDocPassword)
	router.POST("/api/auth/site-password", docHandler.CheckSitePassword)

	// 注册搜索路由
	if searchHandler != nil {
		searchHandler.RegisterRoutes(router)
	}

	// 页面路由
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/index")
	})

	// 处理所有其他文档页面请求
	router.NoRoute(docHandler.HandleDocPage)

	return router
}
