<div align="center">
  <img src="web/static/logo.svg" alt="DocsGo Logo" width="120" height="120">
  <h1>DocsGo - Real-time Documentation Server in Go</h1>
  <p>
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/go-%3E%3D1.22-blue" alt="Go Version"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-green" alt="License"></a>
    <a href="https://github.com/jinsi-ai/docs-go/releases"><img src="https://img.shields.io/github/v/release/jinsi-ai/docs-go" alt="Release"></a>
  </p>
</div>

**DocsGo** is a high-performance **Markdown documentation server** built with **Go**, featuring real-time rendering, full-text search, and access control. A lightweight alternative to static site generators like Hugo and Jekyll.

> 近思切问，AI务实 | NearThink AI, Pragmatic Tech

## Why DocsGo?

In modern software development, managing and sharing technical documentation is crucial:

- ❌ **Static Generators** (Hugo/Jekyll): Require recompilation on every change, no real-time preview
- ❌ **Cloud Docs** (Notion/Yuque): Data hosted by third parties, limited privacy and customization
- ❌ **Traditional Wiki** (Confluence): Bloated and complex, high learning curve

✅ **DocsGo**: Single binary deployment, real-time rendering, data ownership, ready to use

## Key Features

- **Real-time Rendering** - Auto-refresh on file changes, no recompilation needed
- **Full-text Search** - SQLite FTS-based search with automatic indexing and Chinese support
- **Password Protection** - Site-level and document-level dual encryption
- **Embedded Deployment** - Static assets bundled into binary for single-file deployment
- **Responsive UI** - Perfectly adapted for desktop and mobile devices
- **Syntax Highlighting** - Support for 100+ programming languages

## Use Cases

- **Technical Documentation Hub** - API docs, architecture docs for dev teams
- **Product Knowledge Base** - User manuals, FAQs, changelogs
- **Personal Knowledge Management** - Learning notes, tech blogs
- **Enterprise Internal Docs** - Internal knowledge base, training materials

## Quick Start

### Option 1: Download Pre-built Binaries (Recommended)

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

### Option 2: Build from Source

```bash
# Clone repository
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# Build (requires Go 1.22+)
go build -o docs-go

# Run
./docs-go

# Visit http://localhost:8080
```

Place Markdown files in the `docs/` directory for automatic rendering.

## Configuration

Create `data/.env`:

```bash
PORT=8080
DOCS_DIR=./docs
SITE_TITLE=Documentation Center
PASSWORD_SITE=          # Site password (optional)
EMBED_WEB=true          # Embed static assets (default: true)
```

Or use command line arguments (higher priority):

```bash
./docs-go -port 8080 -docs ./docs -password-site mypass
```

## Frontmatter

Add YAML configuration at the beginning of documents:

```yaml
---
title: "Page Title"
order: 1
password: "access password"    # Optional
---
```

| Parameter | Description |
|-----------|-------------|
| `title` | Page title |
| `order` | Sort weight (smaller = higher) |
| `password` | Document access password |

## DocsGo vs Hugo

| Feature | DocsGo | Hugo |
|---------|--------|------|
| Architecture | Real-time server | Static compilation |
| Dev Experience | Live preview, no compile | Rebuild after changes |
| Deployment | Single binary, zero deps | Static files deployment |
| Search | Built-in full-text | Requires additional setup |
| Access Control | Built-in password | Requires external system |
| Best For | Dynamic docs, knowledge base | Static blogs, websites |

## Project Structure

```
docs-go/
├── docs/           # Documentation directory
├── data/           # Data and configuration
├── web/            # Frontend assets
├── app/            # Application layer
└── pkg/            # Core packages
```

## Tech Stack

- **Backend**: Go + Gin Web Framework
- **Frontend**: Tailwind CSS + Vanilla JavaScript
- **Markdown**: goldmark (GFM support, syntax highlighting)
- **Search**: SQLite FTS5 full-text indexing
- **File Watching**: fsnotify cross-platform monitoring

## Development

```bash
# Run
go run main.go

# Format code
go fmt ./...

# Vet
go vet ./...

# Cross-compile
./build.sh  # or build.bat (Windows)
```

## License

[MIT](LICENSE) © JinSi AI <309328809@qq.com>

- **GitHub**: https://github.com/jinsi-ai/docs-go
- **Issues**: https://github.com/jinsi-ai/docs-go/issues

## Keywords

Go documentation generator, Markdown server, document management system, knowledge base tool, technical documentation platform, live preview, static site generator alternative, open source documentation tool, self-hosted wiki, enterprise documentation, team collaboration tool

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

[中文文档](README_CN.md) | English Documentation
