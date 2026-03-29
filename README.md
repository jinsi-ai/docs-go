# DocsGo

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Release](https://img.shields.io/github/v/release/jinsi-ai/docs-go)](https://github.com/jinsi-ai/docs-go/releases)

基于 Go 语言的实时文档渲染服务，支持 Markdown 和 HTML，无需编译，即写即看。

## 特性

- **实时渲染** - 修改文档后自动刷新，无需重新编译
- **全文搜索** - 基于 SQLite 的实时搜索，自动索引
- **密码保护** - 支持站点级和文档级加密
- **嵌入部署** - 可选将静态资源打包到二进制文件
- **响应式 UI** - 支持桌面和移动端

## 快速开始

### 方式一：下载预编译版本

```bash
# 下载最新版本（以 Linux 为例）
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64
```

### 方式二：源码编译

```bash
# 克隆项目
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# 编译
go build -o docs-go

# 运行
./docs-go

# 访问 http://localhost:8080
```

将 Markdown 文件放入 `docs/` 目录即可自动渲染。

## 配置

创建 `data/.env`：

```bash
PORT=8080
DOCS_DIR=./docs
SITE_TITLE=文档中心
PASSWORD_SITE=          # 站点密码（可选）
EMBED_WEB=true          # 嵌入静态资源（默认true，开发时可设为false）
```

或使用命令行参数（优先级更高）：

```bash
./docs-go -port 8080 -docs ./docs -password-site mypass
```

## Frontmatter

在文档开头添加 YAML 配置：

```yaml
---
title: "页面标题"
order: 1
password: "访问密码"    # 可选
---
```

**支持的配置项：**

| 参数 | 说明 |
|------|------|
| `title` | 页面标题 |
| `order` | 排序权重（越小越靠前）|
| `password` | 文档访问密码 |

## 目录结构

```
docs/                           # 文档目录
├── index.md                    # 首页
├── guide/
│   ├── index.md               # 目录索引
│   └── quick-start.md         # 文档
└── README.md                   # 兼容 GitHub
```

**文件优先级：**
1. `index.html`
2. `README.md` / `readme.md`
3. `index.md`

## 嵌入部署

默认：`EMBED_WEB=true`，所有资源已打包在二进制中

开发时：设置 `EMBED_WEB=false`，从文件系统加载便于实时调试：

```bash
# 启用嵌入并编译
go build -o docs-go

# 仅需复制二进制文件和文档目录
./docs-go
```

## 项目结构

```
docs-go/
├── docs/           # 文档目录
├── data/           # 数据和配置
├── web/            # 前端资源
├── app/            # 应用层
└── pkg/            # 核心包
```

## 技术栈

- **后端**: Go + Gin
- **前端**: Tailwind CSS
- **Markdown**: goldmark (GFM)
- **搜索**: SQLite FTS
- **文件监控**: fsnotify

## 开发

```bash
# 运行
go run main.go

# 格式化代码
go fmt ./...

# 检查
go vet ./...
```

## 许可证

[MIT](LICENSE) © 近思AI <309328809@qq.com>

**GitHub**: https://github.com/jinsi-ai/docs-go
