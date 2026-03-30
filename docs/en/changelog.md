---
title: "Changelog"
title_dir: "Changelog"
description: "Version history and release notes for DocsGo"
keywords: "changelog, release notes, version history, updates"
order: 80
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/changelog"
---

# Changelog

All notable changes to DocsGo are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

### Added
- Full internationalization (i18n) support
- Bilingual documentation (English/Chinese)
- Brand story and philosophy documentation
- Comprehensive API reference

## [0.2.0] - 2025-03-30

### Added
- Multi-language documentation support
- New logo and branding
- Brand story page
- Complete English documentation
- Enhanced frontmatter support with i18n fields

### Changed
- Improved README with logo
- Better organized docs structure
- Enhanced footer with brand information

### Fixed
- Documentation links consistency
- Frontmatter parsing improvements

## [0.1.0] - 2025-03-20

### Added
- Initial release
- Real-time Markdown rendering
- Full-text search with SQLite FTS
- Password protection (site and document level)
- Responsive web interface
- Syntax highlighting for 100+ languages
- File watching with auto-reload
- Cross-platform support (Windows, Linux, macOS)
- Embedded static assets for single-binary deployment
- Document tree navigation
- Frontmatter support (title, order, password)
- Cookie-based authentication

### Features
- **Real-time Preview**: Edit Markdown and see changes instantly
- **Full-text Search**: Fast search across all documents
- **Password Protection**: Secure sensitive documentation
- **Single Binary**: Easy deployment with embedded resources
- **Mobile Responsive**: Works on all devices
- **Code Highlighting**: Beautiful syntax highlighting

## Versioning

DocsGo follows [Semantic Versioning](https://semver.org/):

- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

## Release Schedule

- **Patch releases**: As needed for bug fixes
- **Minor releases**: Monthly with new features
- **Major releases**: Annually or for breaking changes

## Upgrade Guide

### From 0.1.x to 0.2.x

1. Backup your `docs/` directory
2. Download new binary or rebuild from source
3. Update `data/.env` if needed
4. Restart server

No breaking changes in configuration.

## Roadmap

### Planned for 0.3.0
- [ ] Plugin system
- [ ] Theme customization
- [ ] Advanced search filters
- [ ] Export to PDF
- [ ] Multi-user support

### Planned for 1.0.0
- [ ] Stable API
- [ ] Plugin marketplace
- [ ] Advanced analytics
- [ ] Enterprise features

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
