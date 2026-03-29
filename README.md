# DocsGo

基于 Go 语言的实时文档渲染服务，支持 Markdown 和 HTML，无需编译，即写即看。

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

## 特性

- **实时渲染** - 修改文档后自动刷新，无需重新编译
- **全文搜索** - 基于 SQLite 的实时搜索，自动索引
- **密码保护** - 支持站点级和文档级加密
- **嵌入部署** - 可选将静态资源打包到二进制文件
- **响应式 UI** - 支持桌面和移动端

## 快速开始

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
EMBED_WEB=false         # 嵌入静态资源
```

或使用命令行参数：

```bash
./docs-go -port 8080 -docs ./docs -password-site mypass
```

## Frontmatter

```yaml
---
title: "页面标题"
order: 1
password: "访问密码"    # 可选
---
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

## 许可证

[MIT](LICENSE) © 近思AI <bitepeng@qq.com>

**GitHub**: https://github.com/jinsi-ai/docs-go
