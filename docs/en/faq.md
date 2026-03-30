---
title: "FAQ"
title_dir: "FAQ"
description: "Frequently asked questions about DocsGo"
keywords: "FAQ, frequently asked questions, help, troubleshooting"
order: 70
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/faq"
---

# Frequently Asked Questions

## General Questions

### What is DocsGo?

DocsGo is a lightweight, real-time documentation server built with Go. It renders Markdown files instantly without compilation, unlike static site generators.

### Who should use DocsGo?

- Technical teams needing internal documentation
- Developers wanting live-preview docs
- Organizations requiring self-hosted solutions
- Anyone who wants simple, fast documentation

### How is DocsGo different from Hugo?

| Feature | DocsGo | Hugo |
|---------|--------|------|
| Build time | Real-time | Compilation required |
| Deployment | Single binary | Static files |
| Search | Built-in | Plugin required |
| Password protection | Built-in | External system |

## Installation

### What are the system requirements?

- **Binary**: None, runs on any OS
- **Source**: Go 1.22 or higher

### Which platforms are supported?

- Linux (AMD64, ARM64)
- macOS (Intel, Apple Silicon)
- Windows (AMD64)

### Can I run it in Docker?

Yes! See the [Deployment Guide](./deployment) for Docker instructions.

## Usage

### How do I add a new document?

Create a `.md` file in the `docs/` directory:

```bash
echo "# My Document" > docs/my-doc.md
```

It appears instantly at http://localhost:8080/my-doc

### How do I organize documents?

Use folders and frontmatter:

```
docs/
├── guide/
│   ├── index.md      (order: 1)
│   └── advanced.md   (order: 2)
└── index.md          (order: 1)
```

### Can I use HTML in Markdown?

Yes, raw HTML is supported for advanced formatting.

### How do I add images?

Place images in `docs/` or `web/static/`:

```markdown
![Alt text](./images/screenshot.png)
```

## Configuration

### Where do I set the port?

In `data/.env`:

```bash
PORT=3000
```

Or via command line:

```bash
./docs-go -port 3000
```

### How do I protect my docs with a password?

Site-wide password:

```bash
# data/.env
PASSWORD_SITE=YourPassword
```

Document-level password:

```yaml
---
password: "Secret123"
---
```

### Can I disable search?

Search is always enabled, but you can remove `data/search.db` and it won't be rebuilt unless you access search.

## Troubleshooting

### Server won't start

**Check:**
1. Port is not in use: `lsof -i :8080`
2. Docs directory exists: `ls docs/`
3. Binary has execute permission: `chmod +x docs-go`

### Changes not showing

1. Check file is in correct directory
2. Verify file extension is `.md`
3. Refresh browser (hard refresh: Ctrl+F5)
4. Check server logs

### Search not working

Rebuild the search index:

```bash
rm data/search.db
./docs-go
```

### Slow performance

1. Reduce document size
2. Rebuild search index
3. Check available disk space
4. Use `EMBED_WEB=true` for production

## Advanced

### Can I customize the theme?

Yes! Edit templates in `web/views/`:

- `doc.html` - Document page
- `password.html` - Password page
- `error.html` - Error page

### How do I add custom CSS?

Add CSS to `web/static/css/custom.css` and include in templates.

### Can I use DocsGo for a blog?

Yes, though it's optimized for documentation. For blogs, consider:
- Organizing posts by date
- Using tags in frontmatter
- Creating an index page

### Is there an API?

Yes! See the [API Reference](./api) for endpoints.

## Contributing

### How can I contribute?

1. Fork the repository
2. Make your changes
3. Submit a pull request

See [Development Guide](./develop) for details.

### Where can I report bugs?

GitHub Issues: https://github.com/jinsi-ai/docs-go/issues

### Is there a roadmap?

Check GitHub Issues for planned features and milestones.

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
