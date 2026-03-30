---
title: "Quick Start"
title_dir: "Quick Start"
description: "Get started with DocsGo in 5 minutes"
keywords: "quickstart, getting started, tutorial, docs-go"
order: 10
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/quickstart"
---

# Quick Start

Get DocsGo up and running in just 5 minutes.

## Prerequisites

- Go 1.22+ (for building from source)
- Or download pre-built binaries

## Installation

### Option 1: Download Pre-built Binaries (Recommended)

Choose your platform:

**Linux (AMD64)**
```bash
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
mv docs-go-linux-amd64 docs-go
```

**Linux (ARM64)**
```bash
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-arm64
chmod +x docs-go-linux-arm64
mv docs-go-linux-arm64 docs-go
```

**macOS (Intel)**
```bash
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-darwin-amd64
chmod +x docs-go-darwin-amd64
mv docs-go-darwin-amd64 docs-go
```

**macOS (Apple Silicon)**
```bash
wget https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-darwin-arm64
chmod +x docs-go-darwin-arm64
mv docs-go-darwin-arm64 docs-go
```

**Windows**
```powershell
# Download using browser or curl
curl -LO https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-windows-amd64.exe
```

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# Build (requires Go 1.22+)
go build -o docs-go

# The binary is now ready to use
```

## First Run

### 1. Create Your Docs Directory

```bash
mkdir -p docs
```

### 2. Create Your First Document

Create `docs/index.md`:

```markdown
---
title: "Welcome"
order: 1
---

# Welcome to My Docs

This is your first document!

## Features

- Real-time preview
- Full-text search
- Password protection
```

### 3. Start the Server

```bash
./docs-go
```

### 4. Open Your Browser

Visit: http://localhost:8080

🎉 That's it! Your documentation site is running.

## What's Next?

- [Configuration Guide](./config) - Customize your setup
- [Usage Guide](./usage) - Learn daily workflows
- [Deployment Guide](./deployment) - Go to production

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
