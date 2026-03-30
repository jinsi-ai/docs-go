---
title: "开发指南"
title_dir: "开发"
description: "参与DocsGo开发的指南"
keywords: "开发, 贡献, 构建, 源码"
order: 90
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/develop"
---

# 开发指南

欢迎参与DocsGo开发！本指南将帮助你搭建开发环境并为项目做贡献。

## 前提条件

- Go 1.22或更高版本
- Git
- Make（可选）

## 环境搭建

### 1. Fork和克隆

```bash
# 在GitHub上Fork，然后克隆
git clone https://github.com/YOUR_USERNAME/docs-go.git
cd docs-go
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 运行开发服务器

```bash
go run main.go
```

访问 http://localhost:8080

## 项目结构

```
docs-go/
├── app/                 # 应用层
│   ├── app.go          # 路由设置
│   └── docs/           # 文档处理器
├── pkg/                # 核心包
│   ├── auth/          # 认证
│   ├── config/        # 配置
│   ├── doc/           # 文档处理
│   ├── httpd/         # HTTP服务器
│   ├── resp/          # 响应工具
│   ├── search/        # 搜索引擎
│   ├── watcher/       # 文件监控
│   └── webfs/         # 嵌入资源
├── web/               # 前端资源
│   ├── static/        # CSS/JS文件
│   └── views/         # HTML模板
├── docs/              # 文档
├── data/              # 数据目录
└── main.go            # 入口点
```

## 开发工作流

### 进行更改

1. 创建新分支：
   ```bash
   git checkout -b feature/my-feature
   ```

2. 进行更改
3. 本地测试
4. 提交清晰的提交信息
5. 推送并创建pull request

### 代码风格

我们遵循标准Go约定：

```bash
# 格式化代码
go fmt ./...

# 运行检查
go vet ./...
```

### 提交信息

使用约定式提交：

```
feat: 添加新搜索功能
fix: 解决内存泄漏
docs: 更新README
refactor: 简化认证逻辑
test: 添加单元测试
```

### 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包
go test ./pkg/doc

# 带覆盖率
go test -cover ./...
```

## 核心组件

### 文档处理 (`pkg/doc/`)

- `document.go` - Markdown/HTML解析
- `doctree.go` - 文档树结构
- `watcher.go` - 文件系统监控

### 搜索 (`pkg/search/`)

- `indexer.go` - 搜索索引管理
- `searcher.go` - 查询执行
- `storage.go` - SQLite操作

### HTTP服务器 (`pkg/httpd/`)

- `server.go` - 服务器初始化

### 认证 (`pkg/auth/`)

- `cookie.go` - 会话管理

## 构建

### 开发构建

```bash
go build -o docs-go
```

### 生产构建

```bash
# 跨平台构建
./build.sh

# 或手动优化构建
go build -ldflags "-s -w" -o docs-go
```

### 嵌入资源

`web/` 中的资源使用Go 1.16+的embed功能嵌入：

```go
//go:embed web
var webFS embed.FS
```

设置 `EMBED_WEB=true` 使用嵌入资源。

## 调试

### 启用调试日志

```go
// 在代码中
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

### 常见问题

**更改不反映：**
- 检查 `EMBED_WEB` 设置
- 模板更改后重启服务器

**构建错误：**
- 确保Go 1.22+
- 运行 `go mod tidy`

**测试失败：**
- 检查测试数据存在
- 验证文件权限

## 贡献

### 提交前

- [ ] 代码遵循风格指南
- [ ] 测试通过：`go test ./...`
- [ ] 文档已更新
- [ ] 提交信息清晰
- [ ] PR描述解释更改

### Pull Request流程

1. 如需更新README.md
2. 更新CHANGELOG.md
3. 链接相关问题
4. 请求维护者审查

### 代码审查

- 尊重和建设性
- 关注代码而非个人
- 解释建议的原因
- 准备好时批准

## 发布流程

1. 在代码中更新版本
2. 更新CHANGELOG.md
3. 创建git标签：`git tag -a v0.x.x -m "版本 0.x.x"`
4. 推送标签：`git push origin v0.x.x`
5. GitHub Actions自动构建发布

## 资源

- [Go文档](https://golang.org/doc/)
- [Gin框架](https://gin-gonic.com/)
- [Goldmark Markdown](https://github.com/yuin/goldmark)
- [项目Issues](https://github.com/jinsi-ai/docs-go/issues)

## 获取帮助

- GitHub Issues：Bug报告和功能
- Discussions：问题和想法
- 邮箱：309328809@qq.com

感谢为DocsGo做贡献！🎉

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
