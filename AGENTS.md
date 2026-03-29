# AGENTS.md - AI Coding Agents Guide

## 项目概述
这是一个基于 Go 语言和 Gin 框架的轻量级文档网站生成器，支持 Markdown 和 HTML 文档的实时渲染与展示。

## 常用命令

### 构建
```bash
# 编译项目
go build -o docs-go.exe

# 运行项目
go run main.go

# 带参数运行
go run main.go -docs ./my-docs -port 8081
```

### 测试
```bash
# 运行所有测试
go test ./...

# 运行单个包的测试
go test ./pkg/doc

# 运行单个测试函数
go test -run TestFunctionName ./pkg/doc

# 详细输出
go test -v ./...

# 带覆盖率测试
go test -cover ./...
```

### 代码检查
```bash
# 格式化代码
go fmt ./...

# 代码审查
go vet ./...

# 整理依赖
go mod tidy

# 下载依赖
go mod download
```

## 代码风格指南

### 导入顺序
导入必须按以下顺序分组，每组之间空一行：
1. 标准库
2. 第三方库
3. 项目内部包

示例：
```go
import (
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"

    "docs-go/pkg/config"
    "docs-go/pkg/doc"
)
```

### 命名规范
- **包名**: 小写，简短，避免下划线（如 `doc`, `auth`, `config`）
- **文件名**: 小写，使用下划线分隔（如 `doctree.go`, `file_watcher.go`）
- **结构体**: 大写驼峰（如 `DocTree`, `CookieManager`）
- **接口**: 大写驼峰，通常以 `er` 结尾（如 `Reader`, `Handler`）
- **函数**: 导出函数大写驼峰，私有函数小写驼峰
- **变量**: 驼峰命名，布尔变量用 `is`/`has` 前缀
- **常量**: 大写下划线（如 `CookieExpire`, `DefaultPort`）

### 注释规范
- 所有导出元素必须有注释
- 注释以被描述对象的名称开头
- 使用中文注释
- 函数注释说明功能和参数

示例：
```go
// DocTree 文档树管理结构体
type DocTree struct {
    Root   *DocNode       // 根节点
    Mutex  sync.RWMutex   // 用于保护并发访问
    Config *config.Config // 配置引用
}

// NewDocTree 创建新的文档树实例
func NewDocTree(config *config.Config) *DocTree {
    return &DocTree{
        Config: config,
    }
}
```

### 错误处理
- 错误尽早返回，避免深层嵌套
- 使用 `fmt.Errorf` 包装错误，添加上下文
- 日志记录使用 `log.Printf`
- 关键错误使用 `log.Fatalf` 终止程序

示例：
```go
func (dt *DocTree) Init() error {
    // 验证文档目录存在
    if err := dt.Config.Validate(); err != nil {
        return err
    }
    // ...
}
```

### 结构体标签
使用 json 标签进行序列化控制，小写字段名：
```go
type DocNode struct {
    Name     string     `json:"name"`
    Path     string     `json:"path"`
    IsDir    bool       `json:"isDir"`
    Children []*DocNode `json:"children"`
}
```

### 并发安全
- 共享状态使用 `sync.RWMutex` 保护
- 读操作使用 `RLock`，写操作使用 `Lock`
- 使用 `defer` 确保解锁

示例：
```go
func (dt *DocTree) GetActiveTree(activePath string) *DocNode {
    dt.Mutex.RLock()
    defer dt.Mutex.RUnlock()
    // ...
}
```

## 项目结构
```
docs-go/
├── app/              # 应用层代码
│   ├── app.go        # 路由设置和初始化
│   └── docs/         # 文档处理模块
├── pkg/              # 核心包模块
│   ├── auth/         # 认证相关
│   ├── config/       # 配置管理
│   ├── doc/          # 文档处理核心
│   ├── httpd/        # HTTP服务器
│   ├── resp/         # 响应处理
│   └── watcher/      # 文件监控
├── web/              # 前端资源
│   ├── static/       # 静态文件
│   └── views/        # HTML模板
├── main.go           # 程序入口
└── go.mod            # Go模块定义
```

## 配置说明
- 配置文件位于 `data/.env`
- 支持命令行参数覆盖环境变量
- 参数优先级: 命令行 > 环境变量 > 默认值

## 注意事项
1. 使用 `gofmt` 格式化所有代码
2. 确保 `go mod tidy` 后提交
3. 文档目录默认是 `docs`，可自定义
4. 文件监控使用 `fsnotify`，修改文档后自动刷新
5. 支持密码保护，通过 frontmatter 配置
