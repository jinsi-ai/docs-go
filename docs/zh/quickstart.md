---
title: "快速入门"
title_dir: "快速入门"
description: "5分钟快速上手DocsGo"
keywords: "快速入门, 教程, 指南, docs-go"
order: 10
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/quickstart"
---

# 快速入门

5分钟让DocsGo运行起来。

## 前提条件

- Go 1.22+（从源码构建）
- 或直接下载预编译二进制文件

## 安装

### 方式一：下载预编译版本（推荐）

根据平台选择：

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
# 使用浏览器或curl下载
curl -LO https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-windows-amd64.exe
```

### 方式二：源码编译

```bash
# 克隆仓库
git clone https://github.com/jinsi-ai/docs-go.git
cd docs-go

# 编译（需要Go 1.22+）
go build -o docs-go

# 二进制文件已就绪
```

## 首次运行

### 1. 创建文档目录

```bash
mkdir -p docs
```

### 2. 创建第一篇文档

**方式一：使用 Markdown（推荐）**

创建 `docs/index.md`：

```markdown
---
title: "欢迎使用"
order: 1
---

# 欢迎来到我的文档

这是你的第一篇文档！

## 功能特性

- 实时预览
- 全文搜索
- 密码保护
```

**方式二：使用 HTML**

DocsGo 也支持直接放置 HTML 文件：

```bash
cat > docs/index.html << 'EOF'
<!DOCTYPE html>
<html>
<head>
    <title>我的文档</title>
</head>
<body>
    <h1>欢迎使用 DocsGo</h1>
    <p>这是一个 HTML 文档示例。</p>
</body>
</html>
EOF
```

**文件优先级**：当存在同名文件时，按以下顺序优先：
1. `index.html`
2. `README.md`
3. `index.md`

### 3. 启动服务

```bash
./docs-go
```

### 4. 打开浏览器

访问：http://localhost:8080

🎉 完成！你的文档站点已运行。

## 接下来

- [配置指南](./config) - 自定义你的设置
- [使用指南](./usage) - 学习日常工作流
- [部署指南](./deployment) - 部署到生产环境

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
