---
title: "DocsGo"
description: "Lightweight documentation site generator based on Go, supporting Markdown and HTML, no compilation needed, write and view instantly"
keywords: "Go documentation generator, Markdown server, documentation management system, knowledge base tool, technical documentation platform, real-time preview, static site generator alternative"
order: 1
---

# DocsGo - Lightweight Documentation Site Generator

**DocsGo** is an open-source **Markdown documentation server** built with **Go**, designed for technical teams, developers, and enterprises. It provides real-time document rendering, full-text search, and access control features. As a lightweight alternative to traditional **static site generators** (such as Hugo, Jekyll), DocsGo adopts a server-side real-time rendering architecture, allowing you to preview document updates instantly without compilation.

> 近思切问，AI务实 | NearThink AI, Pragmatic Tech

## Why Choose DocsGo?

In modern software development, managing and sharing technical documentation is crucial. DocsGo addresses the pain points of traditional documentation tools:

- ❌ **Static Generators** (Hugo/Jekyll): Require recompilation on every change, no real-time preview
- ❌ **Cloud Documentation** (Notion/Yuque): Data hosted by third parties, limited privacy and customization
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

## Use Cases

### 1. Technical Team Documentation Center
Suitable for software development teams to build internal technical documentation libraries:
- API interface documentation
- System architecture documents
- Development standards and guidelines
- Deployment and operation manuals

### 2. Product Knowledge Base
Provides a user-friendly documentation center for product teams:
- Product user manuals
- FAQ (Frequently Asked Questions)
- Version changelog
- New feature introductions

### 3. Personal Knowledge Management
Personal note-taking system for developers and writers:
- Learning notes organization
- Technical blog writing
- Project document hosting
- E-book publishing

### 4. Enterprise Internal Documentation
A lightweight alternative to Confluence:
- Enterprise internal knowledge base
- Department document sharing
- Training materials management
- Process and standard documentation

## Quick Start Guide

### Option 1: Download Pre-built Binaries (Recommended)

```bash
# Download latest version (Linux example)
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64

# Visit http://localhost:8080
```

### Option 2: Build from Source

```bash
# Clone repository
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# Build (requires Go 1.22+)
go build -o docs-go

# Run
./docs-go
```

### Create Your First Document

```bash
# Create Markdown file in docs/ directory
cat > docs/getting-started.md << 'EOF'
---
title: "Getting Started"
order: 1
---

# Welcome to DocsGo

This is your first document!

## Features

- Support for **Markdown** syntax
- Real-time preview, no refresh needed
- Code highlighting display

```go
package main

func main() {
    println("Hello, DocsGo!")
}
```
EOF
```

## Configuration

Create `data/.env` file:

```bash
# Service Configuration
PORT=8080
DOCS_DIR=./docs
SITE_TITLE="My Documentation Center"

# Security Configuration
PASSWORD_SITE=yourpassword    # Site access password (optional)

# Deployment Configuration
EMBED_WEB=true               # Embed static resources (recommended for production)
```

## Document Frontmatter Configuration

Each document supports YAML Frontmatter metadata:

```yaml
---
title: "Document Title"           # Page title
order: 1                          # Sort weight (smaller = higher)
password: "access password"       # Document-level password (optional)
---
```

## Technical Architecture

- **Backend Framework**: Go 1.22 + Gin Web Framework
- **Frontend Technology**: HTML5 + Tailwind CSS + Vanilla JavaScript
- **Markdown Engine**: goldmark (supports GFM, code highlighting)
- **Search Engine**: SQLite FTS5 full-text indexing
- **File Monitoring**: fsnotify cross-platform file system monitoring

## DocsGo vs Hugo

| Feature | DocsGo | Hugo |
|---------|--------|------|
| Architecture | Real-time rendering server | Static compilation |
| Development Experience | Real-time preview, no compile | Rebuild after changes |
| Deployment Complexity | Single binary, zero deps | Static files deployment required |
| Search | Built-in full-text | Requires additional setup |
| Access Control | Built-in password protection | Requires external system support |
| Best For | Dynamic docs, knowledge base | Static blogs, websites |

## Open Source License

DocsGo adopts the **MIT Open Source License**, you are free to use, modify, and distribute.

- **GitHub Repository**: https://github.com/jinsi-ai/docs-go
- **Issue Feedback**: https://github.com/jinsi-ai/docs-go/issues
- **Author**: JinSi AI <309328809@qq.com>

## Related Resources

- **Go Official Website**: https://golang.org/
- **Markdown Syntax**: https://www.markdownguide.org/
- **SQLite FTS**: https://www.sqlite.org/fts5.html

---

## About JinSi AI

> **近思切问，AI务实**  
> **NearThink AI, Pragmatic Tech**

**JinSi AI** (近思AI) derives its name from the Confucian classic *The Analects*: **"博学而笃志，切问而近思，仁在其中矣"** ("Extensive study with steadfast purpose, earnest inquiry with reflection on things near at hand—benevolence lies therein").

### Our Philosophy

In the wave of AI, some chase trends, some sell anxiety, and some speak in abstract concepts.

**JinSi AI chooses a different path.**

"近思" (*NearThink*) means starting from problems at hand, thinking about what truly matters. We don't talk about the distant future—we solve present problems. We don't chase hot topics—we dig into fundamentals.

**Pragmatic, grounded, step by step—from near to far.**

### What We Stand For

- **Pragmatic** | 务实 — Solve real problems, not chase hype
- **Deep Thinking** | 深度 — Go beyond surface, dig into fundamentals  
- **Progressive** | 渐进 — From shallow to deep, accompany your growth
- **Genuine** | 真诚 — Authentic content, no empty talk

### Created By JinSi AI

This project is built with the JinSi AI philosophy: **pragmatic tools for real problems**. We believe good documentation tools should be simple, fast, and just work—so you can focus on what matters: creating great content.

**近思AI，解决真问题**  
**JinSi AI—Solving Real Problems**

---

**Start building your documentation site!** 🚀
