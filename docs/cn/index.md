---
title: "DocsGo"
description: "基于Go语言的实时文档渲染服务，支持Markdown和HTML，无需编译，即写即看"
keywords: "Go文档生成器,Markdown服务器,文档管理系统,知识库工具,技术文档平台,实时预览,静态网站生成器替代品"
order: 1
---

# DocsGo - 轻量级文档网站生成器

**DocsGo** 是一款基于 **Go语言** 开发的开源 **Markdown文档服务器**，专为技术团队、开发者和企业设计，提供实时文档渲染、全文搜索和权限管理功能。作为传统 **静态网站生成器**（如Hugo、Jekyll）的轻量级替代品，DocsGo采用服务端实时渲染架构，让您无需编译即可即时预览文档更新。

> 近思切问，AI务实 | NearThink AI, Pragmatic Tech

## 为什么选择 DocsGo?

在现代软件开发中，技术文档的管理和分享至关重要。DocsGo 解决了传统文档工具的痛点：

- ❌ **静态生成器**（Hugo/Jekyll）：每次修改需重新编译，无法实时预览
- ❌ **云端文档**（Notion/语雀）：数据托管在第三方，隐私和定制受限  
- ❌ **传统Wiki**（Confluence）：臃肿复杂，学习成本高

✅ **DocsGo**：单文件部署、实时渲染、数据自主、开箱即用

## 核心功能特性

### 🚀 实时渲染与热重载
- **即时预览**：修改 Markdown 文件后自动刷新，无需手动编译
- **文件监控**：基于 fsnotify 实时监控文档目录变化
- **自动索引**：文档变更后自动更新搜索索引和目录树

### 🔍 强大的全文搜索
- **SQLite FTS**：基于 SQLite 全文搜索引擎，支持中文分词
- **毫秒响应**：大规模文档库也能快速检索
- **自动索引**：无需手动维护，后台自动完成

### 🔒 企业级权限管理
- **站点级密码**：保护整个文档站点，适合内部文档
- **文档级加密**：单个文档独立密码，灵活控制访问
- **Cookie认证**：安全的会话管理，支持24小时免登录

### 📦 极简部署方案
- **单文件部署**：单个二进制文件，无需依赖（支持嵌入式资源）
- **跨平台支持**：Windows、Linux、macOS（amd64/arm64）
- **Docker就绪**：可轻松容器化部署

### 🎨 现代化界面
- **响应式设计**：完美适配桌面、平板和手机
- **代码高亮**：支持 100+ 编程语言的语法高亮
- **目录导航**：自动生成文档树，支持自定义排序

## 适用场景与使用案例

### 1. 技术团队文档中心
适合软件开发团队构建内部技术文档库：
- API 接口文档
- 系统架构文档  
- 开发规范指南
- 部署运维手册

### 2. 产品知识库
为产品团队提供用户友好的文档中心：
- 产品使用手册
- FAQ 常见问题
- 版本更新日志
- 新功能介绍

### 3. 个人知识管理
开发者和作家的个人笔记系统：
- 学习笔记整理
- 技术博客写作
- 项目文档托管
- 电子书发布

### 4. 企业内网文档
替代 Confluence 的轻量级方案：
- 企业内部知识库
- 部门文档共享
- 培训资料管理
- 流程规范文档

## 快速开始指南

### 方式一：下载预编译版本（推荐）

```bash
# 下载最新版本（Linux示例）
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64

# 访问 http://localhost:8080
```

### 方式二：源码编译安装

```bash
# 克隆仓库
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# 编译（需要Go 1.22+）
go build -o docs-go

# 运行
./docs-go
```

### 创建你的第一篇文档

```bash
# 在 docs/ 目录创建 Markdown 文件
cat > docs/getting-started.md << 'EOF'
---
title: "快速入门"
order: 1
---

# 欢迎使用 DocsGo

这是你的第一篇文档！

## 功能特性

- 支持 **Markdown** 语法
- 实时预览，无需刷新
- 代码高亮显示

```go
package main

func main() {
    println("Hello, DocsGo!")
}
```
EOF
```

## 配置说明

创建 `data/.env` 文件：

```bash
# 服务配置
PORT=8080
DOCS_DIR=./docs
SITE_TITLE="我的文档中心"

# 安全配置
PASSWORD_SITE=yourpassword    # 站点访问密码（可选）

# 部署配置
EMBED_WEB=true               # 嵌入静态资源（生产环境推荐）
```

## 文档 Frontmatter 配置

每篇文档支持 YAML Frontmatter 元数据：

```yaml
---
title: "文档标题"           # 页面标题
order: 1                   # 排序权重（越小越靠前）
password: "访问密码"        # 文档级密码（可选）
---
```

## 技术架构

- **后端框架**: Go 1.22 + Gin Web Framework
- **前端技术**: HTML5 + Tailwind CSS + 原生 JavaScript
- **Markdown引擎**: goldmark（支持 GFM、代码高亮）
- **搜索引擎**: SQLite FTS5 全文索引
- **文件监控**: fsnotify 跨平台文件系统监控

## 与 Hugo 对比

| 特性 | DocsGo | Hugo |
|------|--------|------|
| 架构模式 | 实时渲染服务端 | 静态编译生成 |
| 开发体验 | 实时预览，无需编译 | 修改后需重新生成 |
| 部署复杂度 | 单文件零依赖 | 需要部署静态文件 |
| 搜索功能 | 内置全文搜索 | 需额外配置 |
| 权限控制 | 内置密码保护 | 需外部系统支持 |
| 适用场景 | 动态文档、知识库 | 静态博客、官网 |

## 开源协议

DocsGo 采用 **MIT 开源协议**，您可以自由使用、修改和分发。

- **GitHub 仓库**: https://github.com/jinsi-ai/docs-go
- **问题反馈**: https://github.com/jinsi-ai/docs-go/issues
- **作者**: 近思AI <309328809@qq.com>

## 相关资源

- **Go语言官网**: https://golang.org/
- **Markdown语法**: https://www.markdownguide.org/
- **SQLite FTS**: https://www.sqlite.org/fts5.html

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

**近思AI，解决真问题**  
**JinSi AI—Solving Real Problems**

---

**开始构建你的文档站点吧！** 🚀
