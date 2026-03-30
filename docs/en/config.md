---
title: "Configuration Guide"
title_dir: "Configuration"
description: "Complete configuration guide for DocsGo"
keywords: "configuration, config, setup, environment variables"
order: 20
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/config"
---

# Configuration Guide

DocsGo can be configured through environment variables, `.env` files, or command-line arguments.

## Configuration Methods

**Priority Order** (highest to lowest):
1. Command-line arguments
2. Environment variables
3. `data/.env` file
4. Default values

## Environment Variables

Create `data/.env` file:

```bash
# Server Configuration
PORT=8080                    # Server port (default: 8080)
DOCS_DIR=./docs             # Documentation directory (default: ./docs)
SITE_TITLE="My Docs"        # Site title (default: Documentation Center)

# Security Configuration
PASSWORD_SITE=             # Site-wide password (optional)

# Deployment Configuration
EMBED_WEB=true            # Embed static resources (default: true)
AUTO_INDEX=false          # Auto-index on startup (default: false)
```

## Command-Line Arguments

```bash
./docs-go -port 8080 -docs ./docs -password-site mypassword
```

| Argument | Description | Default |
|----------|-------------|---------|
| `-port` | Server port | 8080 |
| `-docs` | Documentation directory | ./docs |
| `-password-site` | Site password | "" |
| `-embed-web` | Embed static resources | true |

## Document Frontmatter

Each document supports YAML frontmatter:

```yaml
---
title: "Page Title"           # Page title
title_dir: "Display Name"     # Navigation display name
order: 1                      # Sort order (smaller = first)
password: "secret"           # Document password (optional)
---

Your markdown content here...
```

### Frontmatter Options

| Option | Description | Default |
|--------|-------------|---------|
| `title` | Page title | Filename |
| `title_dir` | Directory navigation name | Filename |
| `order` | Sort weight | -1 (natural sort) |
| `password` | Access password | "" |
| `description` | Meta description | "" |
| `keywords` | Meta keywords | "" |

## Configuration Examples

### Development Environment

```bash
# data/.env
PORT=3000
DOCS_DIR=./docs-dev
EMBED_WEB=false              # Load from filesystem for development
```

### Production Environment

```bash
# data/.env
PORT=80
DOCS_DIR=./docs
SITE_TITLE="Company Docs"
PASSWORD_SITE=SecurePass123
EMBED_WEB=true              # Embed all assets
AUTO_INDEX=true             # Rebuild index on startup
```

### Private Documentation

```bash
# data/.env
PORT=8080
PASSWORD_SITE=YourSecretPassword
SITE_TITLE="Internal Docs"
```

## Advanced Configuration

### Custom Templates

Place custom HTML templates in `web/views/`:

- `doc.html` - Document page template
- `password.html` - Password protection page
- `error.html` - Error page

Template variables:
- `{{.title}}` - Page title
- `{{.content}}` - Document content (HTML)
- `{{.docTree}}` - Document tree structure
- `{{.siteTitle}}` - Site title
- `{{.query}}` - Search query

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
