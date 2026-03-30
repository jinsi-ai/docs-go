---
title: "配置指南"
title_dir: "配置"
description: "DocsGo完整配置指南"
keywords: "配置, 设置, 环境变量, 配置说明"
order: 20
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/config"
---

# 配置指南

DocsGo可通过环境变量、`.env`文件或命令行参数进行配置。

## 配置方法

**优先级顺序**（从高到低）：
1. 命令行参数
2. 环境变量
3. `data/.env` 文件
4. 默认值

## 环境变量

创建 `data/.env` 文件：

```bash
# 服务配置
PORT=8080                    # 服务端口（默认：8080）
DOCS_DIR=./docs             # 文档目录（默认：./docs）
SITE_TITLE="我的文档"        # 网站标题（默认：文档中心）

# 安全配置
PASSWORD_SITE=             # 站点密码（可选）

# 部署配置
EMBED_WEB=true            # 嵌入静态资源（默认：true）
AUTO_INDEX=false          # 启动时自动索引（默认：false）
```

## 命令行参数

```bash
./docs-go -port 8080 -docs ./docs -password-site mypassword
```

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `-port` | 服务端口 | 8080 |
| `-docs` | 文档目录 | ./docs |
| `-password-site` | 站点密码 | "" |
| `-embed-web` | 嵌入静态资源 | true |

## 文档 Frontmatter

每篇文档支持YAML frontmatter：

```yaml
---
title: "页面标题"           # 页面标题
title_dir: "显示名称"       # 导航显示名称
order: 1                   # 排序权重（越小越靠前）
password: "secret"        # 文档密码（可选）
---

你的Markdown内容...
```

### Frontmatter 选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| `title` | 页面标题 | 文件名 |
| `title_dir` | 目录导航名称 | 文件名 |
| `order` | 排序权重 | -1（自然排序） |
| `password` | 访问密码 | "" |
| `description` | 页面描述 | "" |
| `keywords` | 页面关键词 | "" |

## 配置示例

### 开发环境

```bash
# data/.env
PORT=3000
DOCS_DIR=./docs-dev
EMBED_WEB=false              # 从文件系统加载，便于开发
```

### 生产环境

```bash
# data/.env
PORT=80
DOCS_DIR=./docs
SITE_TITLE="公司文档"
PASSWORD_SITE=SecurePass123
EMBED_WEB=true              # 嵌入所有资源
AUTO_INDEX=true             # 启动时重建索引
```

### 私密文档

```bash
# data/.env
PORT=8080
PASSWORD_SITE=YourSecretPassword
SITE_TITLE="内部文档"
```

## 高级配置

### 自定义模板

在 `web/views/` 目录放置自定义HTML模板：

- `doc.html` - 文档页面模板
- `password.html` - 密码保护页面
- `error.html` - 错误页面

模板变量：
- `{{.title}}` - 页面标题
- `{{.content}}` - 文档内容（HTML）
- `{{.docTree}}` - 文档树结构
- `{{.siteTitle}}` - 站点标题
- `{{.query}}` - 搜索关键词

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
