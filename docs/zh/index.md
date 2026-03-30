---
title: "DocsGo"
title_dir: "首页"
description: "基于Go语言的轻量级文档网站生成器，支持Markdown和HTML，无需编译，即写即看"
keywords: "Go文档生成器,Markdown服务器,文档管理系统,知识库工具,技术文档平台"
order: 1
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/"
---

# DocsGo - 轻量级文档网站生成器

**DocsGo** 是一款基于 **Go语言** 开发的开源 **Markdown文档服务器**，专为技术团队、开发者和企业设计，提供实时文档渲染、全文搜索和权限管理功能。

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

### 📝 多格式支持
- **Markdown**：完整支持 GitHub Flavored Markdown
- **HTML**：直接放置 `.html` 文件，自动渲染
- **智能识别**：自动识别 `index.html` / `index.md` / `README.md` 优先级
- **混合使用**：同一站点可同时使用 Markdown 和 HTML

## 文档导航

- [快速入门](./quickstart) - 5分钟快速上手
- [配置指南](./config) - 完整配置说明
- [使用指南](./usage) - 日常使用教程
- [部署指南](./deployment) - 生产环境部署
- [API参考](./api) - API接口文档
- [最佳实践](./best-practices) - 推荐实践方案
- [常见问题](./faq) - 常见问题解答
- [更新日志](./changelog) - 版本更新记录
- [开发指南](./develop) - 开发参与指南

## 快速开始

```bash
# 下载并运行
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64

# 访问 http://localhost:8080
```

## 开源协议

[MIT License](https://github.com/jinsi-ai/docs-go/blob/main/LICENSE) © 2025 [近思AI](./brand)

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> 由 <strong><a href="./brand">近思AI</a></strong> 精心打造 ❤️
  </p>
  <p style="font-size: 0.9em; color: #666;">
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
  <p style="font-size: 0.8em; color: #999;">
    <a href="https://github.com/jinsi-ai/docs-go">GitHub</a> • 
    <a href="https://github.com/jinsi-ai/docs-go/issues">问题反馈</a> • 
    <a href="../en/">English</a>
  </p>
</div>
