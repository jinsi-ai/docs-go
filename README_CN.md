<div align="center">
  <img src="web/static/logo.svg" alt="DocsGo Logo" width="120" height="120">
  <h1>DocsGo - Go语言实时文档服务器</h1>
  <p>
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/go-%3E%3D1.22-blue" alt="Go Version"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-green" alt="License"></a>
    <a href="https://github.com/jinsi-ai/docs-go/releases"><img src="https://img.shields.io/github/v/release/jinsi-ai/docs-go" alt="Release"></a>
  </p>
</div>

**DocsGo** 基于 **Go语言** 的高性能 **Markdown文档服务器**，支持实时文档渲染、全文搜索和权限管理。作为 Hugo、Jekyll 等静态网站生成器的轻量级替代方案。


<div align="center">中文文档 | <a href="README.md">English Documentation</a></div>
<div align="center"><em> 近思AI（JinSi-AI）开源项目  | 近思切问，AI务实 | NearThink AI, Pragmatic Tech</em></div>


## 为什么选择 DocsGo?

在现代软件开发中，技术文档的管理和分享至关重要：

- ❌ **静态生成器**（Hugo/Jekyll）：每次修改需重新编译，无法实时预览
- ❌ **云端文档**（Notion/语雀）：数据托管在第三方，隐私和定制受限
- ❌ **传统Wiki**（Confluence）：臃肿复杂，学习成本高

✅ **DocsGo**：单文件部署、实时渲染、数据自主、开箱即用

## 核心功能

- **实时渲染** - 修改 Markdown 文件后自动刷新，无需重新编译
- **全文搜索** - 基于 SQLite FTS 的实时搜索，自动索引，支持中文
- **密码保护** - 站点级和文档级双重加密保护
- **嵌入部署** - 静态资源打包到二进制，单文件部署
- **响应式 UI** - 完美适配桌面和移动端
- **代码高亮** - 支持 100+ 编程语言语法高亮

## 适用场景

- **技术文档中心** - 开发团队 API 文档、架构文档
- **产品知识库** - 用户手册、FAQ、更新日志
- **个人知识管理** - 学习笔记、技术博客
- **企业内网文档** - 内部知识库、培训资料

## 快速开始

### 方式一：下载预编译版本（推荐）

```bash
# Linux
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64

# Windows
curl -LO https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-windows-amd64.exe
./docs-go-windows-amd64.exe

# macOS
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-darwin-amd64
chmod +x docs-go-darwin-amd64
./docs-go-darwin-amd64
```

### 方式二：源码编译

```bash
# 克隆项目
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# 编译（需要 Go 1.22+）
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
EMBED_WEB=true          # 嵌入静态资源（默认true）
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

| 参数 | 说明 |
|------|------|
| `title` | 页面标题 |
| `order` | 排序权重（越小越靠前）|
| `password` | 文档访问密码 |

## 与 Hugo 对比

| 特性 | DocsGo | Hugo |
|------|--------|------|
| 架构模式 | 实时渲染服务端 | 静态编译生成 |
| 开发体验 | 实时预览，无需编译 | 修改后需重新生成 |
| 部署方式 | 单文件零依赖 | 需部署静态文件 |
| 搜索功能 | 内置全文搜索 | 需额外配置 |
| 权限控制 | 内置密码保护 | 需外部系统 |
| 适用场景 | 动态文档、知识库 | 静态博客、官网 |

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

- **后端**: Go + Gin Web Framework
- **前端**: Tailwind CSS + 原生 JavaScript
- **Markdown**: goldmark（支持 GFM、代码高亮）
- **搜索**: SQLite FTS5 全文索引
- **文件监控**: fsnotify 跨平台监控

## 开发

```bash
# 运行
go run main.go

# 格式化代码
go fmt ./...

# 检查
go vet ./...

# 交叉编译
./build.sh  # 或 build.bat（Windows）
```

## 许可证

[MIT](LICENSE) © 近思AI <309328809@qq.com>

- **GitHub**: https://github.com/jinsi-ai/docs-go
- **问题反馈**: https://github.com/jinsi-ai/docs-go/issues

## 关键词

Go语言文档生成器、Markdown服务器、文档管理系统、知识库工具、技术文档平台、实时预览、静态网站生成器替代品、开源文档工具、自托管Wiki、企业内部文档、团队协作工具

---

## 关于近思AI

> **近思切问，AI务实**  
> **NearThink AI, Pragmatic Tech**

**近思AI**（JinSi AI）名字出自儒家经典《论语·子张》：**"博学而笃志，切问而近思，仁在其中矣。"**

### 我们的理念

在AI浪潮中，有人在追逐风口，有人在贩卖焦虑，有人在讲高深概念。

**近思AI，选择另一条路。**

"近思"意味着从身边的问题出发，思考那些真正重要的事。我们不空谈未来，只解决当下。不追逐热点，只深挖本质。

**务实、落地、循序渐进——从近处走向远方。**

### 我们的坚持

- **务实** | Pragmatic — 解决真问题，不追风口
- **深度** | Deep Thinking — 不止于表面，深挖底层逻辑  
- **渐进** | Progressive — 由浅入深，陪伴成长
- **真诚** | Genuine — 做比说重要，落地比概念重要

### 近思AI出品

本项目秉承近思AI的理念：**为真问题提供务实工具**。我们相信好的文档工具应该简单、快速、开箱即用——让你专注于重要的事：创作优质内容。

---


**近思AI，解决真问题**  
**JinSi AI—Solving Real Problems**
