---
title: "配置说明"
title_dir: "配置"
order: 10
---

# 配置说明

## 应用配置 (data/.env)

```bash
# 服务端口
PORT=8080

# 文档目录
DOCS_DIR=./docs

# 站点密码（可选）
PASSWORD_SITE=yourpassword

# 网站标题
SITE_TITLE=文档中心

# 嵌入静态资源（默认true）
EMBED_WEB=true

# 自动索引
AUTO_INDEX=false
```

## 命令行参数

```bash
./docs-go -port 8080 -docs ./docs -password-site mypass
```

**优先级**: 命令行 > .env > 默认值

## Frontmatter 配置

```yaml
---
title: "页面标题"
title_dir: "目录显示名"
order: 1
password: "访问密码"
---
```

| 参数 | 说明 |
|------|------|
| `title` | 页面标题 |
| `title_dir` | 目录导航显示名 |
| `order` | 排序权重（越小越靠前）|
| `password` | 文档访问密码 |
