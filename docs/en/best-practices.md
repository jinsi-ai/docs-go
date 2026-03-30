---
title: "Best Practices"
title_dir: "Best Practices"
description: "Recommended practices for using DocsGo effectively"
keywords: "best practices, recommendations, tips, guidelines"
order: 60
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/best-practices"
---

# Best Practices

Follow these recommendations to get the most out of DocsGo.

## Document Organization

### Structure Your Content

```
docs/
├── index.md              # Landing page with overview
├── getting-started/      # New user content
│   ├── index.md
│   ├── installation.md
│   └── quickstart.md
├── guides/              # Detailed guides
│   ├── index.md
│   ├── configuration.md
│   └── deployment.md
├── reference/           # API/reference docs
│   ├── index.md
│   └── api.md
└── README.md           # GitHub compatibility
```

### Naming Conventions

- Use lowercase with hyphens: `getting-started.md` not `GettingStarted.md`
- Be descriptive: `configuration-guide.md` not `guide1.md`
- Use index.md for directory overviews

## Writing Guidelines

### Frontmatter Best Practices

Always include these fields:

```yaml
---
title: "Clear, Descriptive Title"
description: "Brief description for SEO"
keywords: "relevant, keywords, here"
order: 10
---
```

### Content Structure

```markdown
# Page Title

Brief introduction (2-3 sentences)

## Section 1

Content...

### Subsection

More details...

## Section 2

Content...

## See Also

- [Related Doc](./related)
- [External Link](https://example.com)
```

### Writing Tips

1. **Start with context** - Explain why this matters
2. **Use examples** - Show, don't just tell
3. **Keep it scannable** - Use headings, lists, tables
4. **Link generously** - Connect related content
5. **Update regularly** - Keep docs current

## Configuration Best Practices

### Development

```bash
# data/.env
PORT=3000
DOCS_DIR=./docs-dev
EMBED_WEB=false              # Fast reload
```

### Production

```bash
# data/.env
PORT=80
DOCS_DIR=./docs
EMBED_WEB=true               # Single binary
PASSWORD_SITE=SecurePass     # Protect content
AUTO_INDEX=true              # Fresh index on start
```

## Security Practices

### Password Protection

- Use strong passwords (12+ characters)
- Change passwords periodically
- Don't commit passwords to git
- Use environment variables for sensitive data

### Access Control

```bash
# Site-wide protection
PASSWORD_SITE=StrongPass123!

# Document-level protection
# In frontmatter:
# password: "DocSpecificPass"
```

## Performance Optimization

### Search Performance

- Keep documents under 10,000 words each
- Use descriptive titles
- Include keywords naturally
- Delete and rebuild index if slow

### Large Documentation Sites

For sites with 1000+ documents:

```bash
# Split into categories
docs/
├── user-guide/         # User documentation
├── admin-guide/        # Admin documentation
├── api-reference/      # API docs
└── internal/           # Private docs (password protected)
```

## SEO Optimization

### Meta Information

Every document should have:

```yaml
---
title: "Descriptive Title - DocsGo"
description: "Clear, keyword-rich description of content"
keywords: "docs-go, documentation, your-topic"
---
```

### Content SEO

- Use descriptive H1 headings
- Include keywords naturally
- Add alt text to images
- Link to related content
- Keep URLs readable

## Maintenance

### Regular Tasks

- [ ] Review and update outdated content
- [ ] Check for broken links
- [ ] Verify search functionality
- [ ] Update dependencies
- [ ] Backup documentation

### Version Control

```bash
# Keep docs in git
git add docs/
git commit -m "docs: update installation guide"

# Tag releases
git tag -a v1.0.0 -m "Documentation v1.0"
```

## Troubleshooting

### Common Issues

**Search not working:**
```bash
rm data/search.db
./docs-go  # Will rebuild index
```

**Changes not appearing:**
- Check file permissions
- Verify file is in DOCS_DIR
- Restart server if needed

**Slow performance:**
- Reduce document size
- Rebuild search index
- Check disk space

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
