---
title: "项目简介"
order: 1
---

# DocsGo - 轻量级文档网站生成器

基于 Go 语言的实时文档渲染服务，支持 Markdown 和 HTML，无需编译，即写即看。

## 核心特性

- **实时渲染** - 修改文档后自动刷新，无需重新编译
- **文件监控** - 自动检测文档变化，更新目录和搜索索引
- **全文搜索** - 基于 SQLite 的实时搜索
- **密码保护** - 支持站点级和文档级加密
- **嵌入部署** - 可选将静态资源打包到二进制文件

## 快速开始

```bash
# 下载并运行
./docs-go

# 访问文档
open http://localhost:8080
```

将 Markdown 文件放入 `docs/` 目录即可自动渲染。

## 技术栈

- **后端**: Go + Gin
- **前端**: Tailwind CSS + 原生 JS
- **Markdown**: goldmark (GFM 支持)
- **搜索**: SQLite 全文索引

## 适用场景

- 团队技术文档中心
- 项目文档托管
- 个人知识库
- API 文档站点

## 项目地址

GitHub: https://github.com/jinsi-ai/docs-go

**License**: MIT | **Author**: 近思AI <309328809@qq.com>
