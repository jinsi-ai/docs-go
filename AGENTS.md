# AGENTS.md - AI Coding Agents Guide

## Project Overview
基于 Go 语言和 Gin 框架的轻量级文档网站生成器，支持 Markdown 和 HTML 文档的实时渲染与展示，内置全文搜索功能。

## Build Commands

```bash
# Compile
 go build -o docs-go.exe

# Run
 go run main.go

# Run with args
 go run main.go -docs ./my-docs -port 8081

# Cross-platform build (see build.sh)
 bash build.sh
```

## Test Commands

```bash
# Run all tests
 go test ./...

# Run single package tests
 go test ./pkg/doc

# Run single test function
 go test -run TestFunctionName ./pkg/doc

# Verbose output
 go test -v ./...

# With coverage
 go test -cover ./...
```

## Lint/Format Commands

```bash
# Format code
 go fmt ./...

# Vet code
 go vet ./...

# Tidy dependencies
 go mod tidy

# Download dependencies
 go mod download
```

## Code Style Guidelines

### Import Order
1. Standard library
2. Project internal packages
3. Third-party packages

```go
import (
    "html/template"
    "log"

    "docs-go/pkg/config"
    "docs-go/pkg/doc"

    "github.com/gin-gonic/gin"
)
```

### Naming Conventions
- **Packages**: lowercase, short (e.g., `doc`, `auth`, `config`)
- **Files**: lowercase with underscores (e.g., `doctree.go`)
- **Types/Structs**: PascalCase (e.g., `DocTree`, `DocHandler`)
- **Interfaces**: PascalCase with `er` suffix (e.g., `Handler`, `Reader`)
- **Functions**: PascalCase for exported, camelCase for private
- **Variables**: camelCase, booleans use `is`/`has` prefix
- **Constants**: UPPER_SNAKE_CASE

### Comments
- All exported items must have comments
- Comments start with the item name
- Use Chinese for comments

```go
// DocTree 文档树管理结构体
type DocTree struct {
    Root   *DocNode       // 根节点
    Mutex  sync.RWMutex   // 用于保护并发访问
    Config *config.Config // 配置引用
}

// NewDocTree 创建新的文档树实例
func NewDocTree(config *config.Config) *DocTree {
    return &DocTree{Config: config}
}
```

### Error Handling
- Return errors early, avoid deep nesting
- Use `log.Printf` for warnings, `log.Fatalf` for critical errors
- Wrap errors with context using `fmt.Errorf`

```go
func (dt *DocTree) Init() error {
    if err := dt.Config.Validate(); err != nil {
        return err
    }
    // ...
}
```

### Struct Tags
Use `json` tags with lowercase field names:

```go
type DocNode struct {
    Name     string     `json:"name"`
    Path     string     `json:"path"`
    IsDir    bool       `json:"isDir"`
    Children []*DocNode `json:"children"`
}
```

### Concurrency
- Use `sync.RWMutex` for shared state
- Use `RLock` for reads, `Lock` for writes
- Always use `defer` to unlock

```go
func (dt *DocTree) GetActiveTree(activePath string) *DocNode {
    dt.Mutex.RLock()
    defer dt.Mutex.RUnlock()
    // ...
}
```

## Project Structure

```
docs-go/
├── app/              # Application layer
│   ├── app.go        # Route setup
│   ├── docs/         # Document handlers
│   └── search/       # Search handlers
├── pkg/              # Core packages
│   ├── auth/         # Authentication
│   ├── config/       # Configuration
│   ├── doc/          # Document processing
│   ├── httpd/        # HTTP server
│   ├── resp/         # Response utilities
│   ├── search/       # Search engine
│   ├── watcher/      # File watching
│   └── webfs/        # Embedded web assets
├── web/              # Frontend assets
│   ├── static/       # Static files
│   └── views/        # HTML templates
├── data/             # Data files (.env, search.db)
├── docs/             # Default docs directory
├── main.go           # Entry point
└── go.mod            # Go module
```

## Configuration
- Config file: `data/.env`
- Priority: CLI args > env vars > defaults
- Key settings: `DOCS_DIR`, `PORT`, `PASSWORD_SITE`, `SITE_TITLE`

## Key Dependencies
- `github.com/gin-gonic/gin` - Web framework
- `github.com/fsnotify/fsnotify` - File watching
- `github.com/yuin/goldmark` - Markdown parsing
- `github.com/mattn/go-sqlite3` - SQLite for search

## Notes
1. Run `go fmt ./...` before committing
2. Run `go mod tidy` after adding/removing imports
3. Default docs dir is `docs/`, customizable via flag
4. File watcher auto-refreshes on document changes
5. Password protection via frontmatter (`password:` field)
