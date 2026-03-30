---
title: "Development Guide"
title_dir: "Development"
description: "Guide for contributing to DocsGo development"
keywords: "development, contributing, build, source code"
order: 90
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/develop"
---

# Development Guide

Welcome to DocsGo development! This guide will help you set up your development environment and contribute to the project.

## Prerequisites

- Go 1.22 or higher
- Git
- Make (optional)

## Setup

### 1. Fork and Clone

```bash
# Fork on GitHub, then clone
git clone https://github.com/YOUR_USERNAME/docs-go.git
cd docs-go
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Run Development Server

```bash
go run main.go
```

Visit http://localhost:8080

## Project Structure

```
docs-go/
├── app/                 # Application layer
│   ├── app.go          # Route setup
│   └── docs/           # Document handlers
├── pkg/                # Core packages
│   ├── auth/          # Authentication
│   ├── config/        # Configuration
│   ├── doc/           # Document processing
│   ├── httpd/         # HTTP server
│   ├── resp/          # Response utilities
│   ├── search/        # Search engine
│   ├── watcher/       # File watching
│   └── webfs/         # Embedded assets
├── web/               # Frontend assets
│   ├── static/        # CSS/JS files
│   └── views/         # HTML templates
├── docs/              # Documentation
├── data/              # Data directory
└── main.go            # Entry point
```

## Development Workflow

### Making Changes

1. Create a new branch:
   ```bash
   git checkout -b feature/my-feature
   ```

2. Make your changes
3. Test locally
4. Commit with clear messages
5. Push and create pull request

### Code Style

We follow standard Go conventions:

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...
```

### Commit Messages

Use conventional commits:

```
feat: add new search feature
fix: resolve memory leak
docs: update README
refactor: simplify auth logic
test: add unit tests
```

### Testing

```bash
# Run all tests
go test ./...

# Run specific package
go test ./pkg/doc

# With coverage
go test -cover ./...
```

## Key Components

### Document Processing (`pkg/doc/`)

- `document.go` - Markdown/HTML parsing
- `doctree.go` - Document tree structure
- `watcher.go` - File system monitoring

### Search (`pkg/search/`)

- `indexer.go` - Search index management
- `searcher.go` - Query execution
- `storage.go` - SQLite operations

### HTTP Server (`pkg/httpd/`)

- `server.go` - Server initialization

### Authentication (`pkg/auth/`)

- `cookie.go` - Session management

## Building

### Development Build

```bash
go build -o docs-go
```

### Production Build

```bash
# Cross-platform build
./build.sh

# Or manual build with optimizations
go build -ldflags "-s -w" -o docs-go
```

### Embedding Assets

Assets in `web/` are embedded using Go 1.16+ embed feature:

```go
//go:embed web
var webFS embed.FS
```

Set `EMBED_WEB=true` to use embedded assets.

## Debugging

### Enable Debug Logs

```go
// In code
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

### Common Issues

**Changes not reflecting:**
- Check `EMBED_WEB` setting
- Restart server after template changes

**Build errors:**
- Ensure Go 1.22+
- Run `go mod tidy`

**Test failures:**
- Check test data exists
- Verify file permissions

## Contributing

### Before Submitting

- [ ] Code follows style guidelines
- [ ] Tests pass: `go test ./...`
- [ ] Documentation updated
- [ ] Commit messages are clear
- [ ] PR description explains changes

### Pull Request Process

1. Update README.md if needed
2. Update CHANGELOG.md
3. Link related issues
4. Request review from maintainers

### Code Review

- Be respectful and constructive
- Focus on code, not person
- Explain reasoning for suggestions
- Approve when ready

## Release Process

1. Update version in code
2. Update CHANGELOG.md
3. Create git tag: `git tag -a v0.x.x -m "Version 0.x.x"`
4. Push tag: `git push origin v0.x.x`
5. GitHub Actions builds releases automatically

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [Goldmark Markdown](https://github.com/yuin/goldmark)
- [Project Issues](https://github.com/jinsi-ai/docs-go/issues)

## Getting Help

- GitHub Issues: Bug reports and features
- Discussions: Questions and ideas
- Email: 309328809@qq.com

Thank you for contributing to DocsGo! 🎉

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
