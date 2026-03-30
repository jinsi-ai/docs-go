---
title: "DocsGo"
title_dir: "Documentation"
description: "Lightweight documentation site generator based on Go, supporting Markdown and HTML, no compilation needed"
keywords: "Go documentation generator, Markdown server, documentation management system, knowledge base tool"
order: 1
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/"
---

# DocsGo - Lightweight Documentation Site Generator

**DocsGo** is an open-source **Markdown documentation server** built with **Go**, designed for technical teams, developers, and enterprises. It provides real-time document rendering, full-text search, and access control features.

> 近思切问，AI务实 | NearThink AI, Pragmatic Tech

## Why Choose DocsGo?

In modern software development, managing and sharing technical documentation is crucial. DocsGo addresses the pain points of traditional documentation tools:

- ❌ **Static Generators** (Hugo/Jekyll): Require recompilation on every change, no real-time preview
- ❌ **Cloud Documentation** (Notion/Notion): Data hosted by third parties, limited privacy and customization
- ❌ **Traditional Wiki** (Confluence): Bloated and complex, high learning curve

✅ **DocsGo**: Single binary deployment, real-time rendering, data ownership, ready to use

## Core Features

### 🚀 Real-time Rendering & Hot Reload
- **Instant Preview**: Auto-refresh after modifying Markdown files, no manual compilation required
- **File Monitoring**: Real-time monitoring of document directory changes based on fsnotify
- **Auto Indexing**: Automatically updates search index and directory tree after document changes

### 🔍 Powerful Full-text Search
- **SQLite FTS**: Full-text search engine based on SQLite, supporting Chinese word segmentation
- **Millisecond Response**: Fast retrieval even with large-scale document libraries
- **Auto Indexing**: No manual maintenance needed, completed automatically in the background

### 🔒 Enterprise-grade Access Control
- **Site-level Password**: Protect the entire documentation site, suitable for internal documents
- **Document-level Encryption**: Individual document passwords for flexible access control
- **Cookie Authentication**: Secure session management, supporting 24-hour login-free access

### 📦 Minimal Deployment Solution
- **Single Binary Deployment**: Single binary file, no dependencies (supports embedded resources)
- **Cross-platform Support**: Windows, Linux, macOS (amd64/arm64)
- **Docker Ready**: Can be easily containerized

### 🎨 Modern Interface
- **Responsive Design**: Perfectly adapted for desktop, tablet, and mobile devices
- **Code Highlighting**: Syntax highlighting support for 100+ programming languages
- **Directory Navigation**: Automatic document tree generation with custom sorting support

### 📝 Multi-format Support
- **Markdown**: Full support for GitHub Flavored Markdown
- **HTML**: Place `.html` files directly, automatically rendered
- **Smart Recognition**: Automatically recognize priority of `index.html` / `index.md` / `README.md`
- **Mixed Usage**: Can use Markdown and HTML simultaneously on the same site

## Documentation

- [Quick Start](./quickstart) - Get started in 5 minutes
- [Configuration](./config) - Complete configuration guide
- [Usage Guide](./usage) - Daily usage instructions
- [Deployment](./deployment) - Production deployment guide
- [API Reference](./api) - API documentation
- [Best Practices](./best-practices) - Recommended practices
- [FAQ](./faq) - Frequently asked questions
- [Changelog](./changelog) - Version history
- [Development](./develop) - Development guide

## Quick Start

```bash
# Download and run
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64

# Visit http://localhost:8080
```

## License

[MIT License](https://github.com/jinsi-ai/docs-go/blob/main/LICENSE) © 2025 [JinSi AI](./brand)

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> is crafted with ❤️ by <strong><a href="./brand">JinSi AI</a></strong>
  </p>
  <p style="font-size: 0.9em; color: #666;">
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
  <p style="font-size: 0.8em; color: #999;">
    <a href="https://github.com/jinsi-ai/docs-go">GitHub</a> • 
    <a href="https://github.com/jinsi-ai/docs-go/issues">Issues</a> • 
    <a href="../zh/">中文</a>
  </p>
</div>
