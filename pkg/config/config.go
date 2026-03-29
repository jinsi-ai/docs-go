package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config 应用程序配置结构体
type Config struct {
	DocsDir      string // 文档目录
	Port         string // 服务器端口
	PasswordSite string // 站点密码
	SiteTitle    string // 网站标题
	EmbedWeb     bool   // 是否嵌入web静态资源
}

// LoadEnv 从data目录下的.env文件加载配置
func (c *Config) LoadEnv() error {
	envPath := filepath.Join("data", ".env")

	// 检查.env文件是否存在
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Printf("提示: .env配置文件不存在，跳过环境变量加载: %s", envPath)
		return nil
	}

	// 加载.env文件
	err := godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("加载.env文件失败: %v", err)
	}

	log.Printf("成功加载.env配置文件: %s", envPath)

	// 从环境变量读取配置
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
	if docsDir := os.Getenv("DOCS_DIR"); docsDir != "" {
		c.DocsDir = docsDir
	}
	if passwordSite := os.Getenv("PASSWORD_SITE"); passwordSite != "" {
		c.PasswordSite = passwordSite
	}
	if siteTitle := os.Getenv("SITE_TITLE"); siteTitle != "" {
		c.SiteTitle = siteTitle
	}
	if embedWeb := os.Getenv("EMBED_WEB"); embedWeb == "true" {
		c.EmbedWeb = true
	}

	return nil
}

// ParseFlags 从命令行参数解析配置
func (c *Config) ParseFlags() {
	portFlag := flag.String("port", c.Port, "服务器端口")
	docsDirFlag := flag.String("docs", c.DocsDir, "文档目录路径")
	passwordSiteFlag := flag.String("password-site", c.PasswordSite, "站点密码，为空则不验证")
	flag.Parse()

	if *portFlag != "" {
		c.Port = *portFlag
	}
	if *docsDirFlag != "" {
		c.DocsDir = *docsDirFlag
	}
	if *passwordSiteFlag != "" {
		c.PasswordSite = *passwordSiteFlag
	}

	c.SaveEnv()
}

// Validate 验证配置是否有效
func (c *Config) Validate() error {
	// 检查 docs 目录是否存在，如果不存在则创建
	if _, err := os.Stat(c.DocsDir); os.IsNotExist(err) {
		log.Printf("警告: 文档目录 %s 不存在，创建空目录", c.DocsDir)
		if err := os.MkdirAll(c.DocsDir, 0755); err != nil {
			return fmt.Errorf("创建文档目录失败: %v", err)
		}
	}

	return nil
}

// SaveEnv 将当前配置保存到data/.env文件
func (c *Config) SaveEnv() {
	envPath := filepath.Join("data", ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Printf("提示: .env配置文件不存在，创建空文件: %s", envPath)
		if err := os.MkdirAll(filepath.Dir(envPath), 0755); err != nil {
			log.Printf("警告: 创建data目录失败: %v", err)
		}
		if _, err := os.Create(envPath); err != nil {
			log.Printf("警告: 创建.env文件失败: %v", err)
		}
		//写入配置到.env文件
		embedWebStr := "false"
		if c.EmbedWeb {
			embedWebStr = "true"
		}
		envContent := fmt.Sprintf(
			"PORT=%s\nDOCS_DIR=%s\nPASSWORD_SITE=%s\nSITE_TITLE=%s\nEMBED_WEB=%s\n",
			c.Port, c.DocsDir, c.PasswordSite, c.SiteTitle, embedWebStr,
		)
		if err := os.WriteFile(envPath, []byte(envContent), 0644); err != nil {
			log.Printf("警告: 写入.env文件失败: %v", err)
		}
	}
}
