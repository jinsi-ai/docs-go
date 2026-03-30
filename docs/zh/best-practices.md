---
title: "最佳实践"
title_dir: "最佳实践"
description: "高效使用DocsGo的推荐实践"
keywords: "最佳实践, 推荐, 技巧, 指南"
order: 60
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/best-practices"
---

# 最佳实践

遵循这些建议，充分发挥DocsGo的潜力。

## 文档组织

### 结构化内容

```
docs/
├── index.md              # 首页概览
├── getting-started/      # 新用户内容
│   ├── index.md
│   ├── installation.md
│   └── quickstart.md
├── guides/              # 详细指南
│   ├── index.md
│   ├── configuration.md
│   └── deployment.md
├── reference/           # API/参考文档
│   ├── index.md
│   └── api.md
└── README.md           # GitHub兼容
```

### 命名规范

- 使用小写和连字符：`getting-started.md` 而不是 `GettingStarted.md`
- 有描述性：`configuration-guide.md` 而不是 `guide1.md`
- 目录概览使用index.md

## 编写指南

### Frontmatter最佳实践

始终包含这些字段：

```yaml
---
title: "清晰、描述性的标题"
description: "SEO友好的内容简介"
keywords: "相关, 关键词, 此处"
order: 10
---
```

### 内容结构

```markdown
# 页面标题

简介（2-3句话）

## 章节1

内容...

### 子章节

更多详情...

## 章节2

内容...

## 另请参阅

- [相关文档](./related)
- [外部链接](https://example.com)
```

### 编写技巧

1. **从上下文开始** - 解释为什么这很重要
2. **使用示例** - 展示，不要只说
3. **保持可扫描性** - 使用标题、列表、表格
4. **慷慨链接** - 连接相关内容
5. **定期更新** - 保持文档最新

## 配置最佳实践

### 开发环境

```bash
# data/.env
PORT=3000
DOCS_DIR=./docs-dev
EMBED_WEB=false              # 快速重载
```

### 生产环境

```bash
# data/.env
PORT=80
DOCS_DIR=./docs
EMBED_WEB=true               # 单文件
PASSWORD_SITE=SecurePass     # 保护内容
AUTO_INDEX=true              # 启动时重建索引
```

## 安全实践

### 密码保护

- 使用强密码（12+字符）
- 定期更换密码
- 不要将密码提交到git
- 敏感数据使用环境变量

### 访问控制

```bash
# 站点级保护
PASSWORD_SITE=StrongPass123!

# 文档级保护
# 在frontmatter中：
# password: "DocSpecificPass"
```

## 性能优化

### 搜索性能

- 每篇文档控制在10000字以内
- 使用描述性标题
- 自然包含关键词
- 如果变慢，删除并重建索引

### 大型文档站点

对于1000+文档的站点：

```bash
docs/
├── user-guide/         # 用户文档
├── admin-guide/        # 管理员文档
├── api-reference/      # API文档
└── internal/           # 私密文档（密码保护）
```

## SEO优化

### 元信息

每篇文档应包含：

```yaml
---
title: "描述性标题 - DocsGo"
description: "清晰、富含关键词的内容描述"
keywords: "docs-go, 文档, 你的主题"
---
```

### 内容SEO

- 使用描述性H1标题
- 自然包含关键词
- 为图片添加alt文本
- 链接到相关内容
- 保持URL可读

## 维护

### 定期任务

- [ ] 审查并更新过时内容
- [ ] 检查失效链接
- [ ] 验证搜索功能
- [ ] 更新依赖
- [ ] 备份文档

### 版本控制

```bash
# 将文档保存在git中
git add docs/
git commit -m "docs: 更新安装指南"

# 标记发布
git tag -a v1.0.0 -m "文档v1.0"
```

## 故障排除

### 常见问题

**搜索不工作：**
```bash
rm data/search.db
./docs-go  # 将重建索引
```

**更改不显示：**
- 检查文件权限
- 确认文件在DOCS_DIR中
- 重启服务器

**性能缓慢：**
- 减小文档大小
- 重建搜索索引
- 检查磁盘空间

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
