---
title: "Usage Guide"
title_dir: "Usage"
description: "Learn how to use DocsGo effectively"
keywords: "usage, guide, howto, workflow"
order: 30
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/usage"
---

# Usage Guide

Learn the daily workflows for using DocsGo effectively.

## Local Development Workflow

### 1. Start the Server

```bash
./docs-go
```

The server will:
- Watch for file changes in `docs/` directory
- Auto-rebuild document tree and search index
- Serve at http://localhost:8080

### 2. Edit Documents

Create or edit Markdown files in the `docs/` directory:

```bash
# Create a new document
echo "# My New Doc" > docs/my-doc.md

# Edit existing document
vim docs/guide.md
```

Changes are reflected instantly in the browser.

### 3. Organize Your Documentation

Recommended structure:

```
docs/
├── index.md              # Home page
├── guide/
│   ├── index.md         # Guide overview
│   ├── getting-started.md
│   └── advanced.md
├── api/
│   ├── index.md
│   └── reference.md
└── README.md            # GitHub compatibility
```

## Writing Documents

### Markdown Support

DocsGo supports full Markdown syntax plus extensions:

```markdown
# Heading 1
## Heading 2
### Heading 3

**Bold text** and *italic text*

- Bullet list
- Another item
  - Nested item

1. Numbered list
2. Second item

[Link text](url)

![Image alt](image-url)

| Table | Column |
|-------|--------|
| Data  | Value  |

```code block
func main() {
    println("Hello")
}
```
```

### Code Highlighting

Supported languages: Go, Python, JavaScript, Java, C++, Rust, and 100+ more.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, DocsGo!")
}
```

### Frontmatter

Add metadata at the top of documents:

```yaml
---
title: "My Document"
order: 1
password: "secret"
---
```

## Search Functionality

### Automatic Indexing

- All Markdown content is automatically indexed
- Index updates when files change
- Full-text search with Chinese support

### Rebuilding Index

Delete `data/search.db` and restart to rebuild:

```bash
rm data/search.db
./docs-go
```

## Password Protection

### Site-wide Password

Set in `data/.env`:

```bash
PASSWORD_SITE=YourPassword
```

### Document-level Password

In document frontmatter:

```yaml
---
password: "DocumentSecret"
---
```

## Best Practices

1. **Use meaningful file names**: `installation.md` not `doc1.md`
2. **Set order explicitly**: Use `order` in frontmatter
3. **Write good titles**: Clear and descriptive
4. **Organize with folders**: Group related documents
5. **Use frontmatter**: Add metadata for better SEO

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
