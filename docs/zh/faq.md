---
title: "常见问题"
title_dir: "FAQ"
description: "DocsGo常见问题解答"
keywords: "FAQ, 常见问题, 帮助, 故障排除"
order: 70
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/faq"
---

# 常见问题

## 一般问题

### DocsGo是什么？

DocsGo是一个轻量级的实时文档服务器，使用Go构建。与静态站点生成器不同，它能即时渲染Markdown文件，无需编译。

### 谁应该使用DocsGo？

- 需要内部文档的技术团队
- 想要实时预览文档的开发者
- 需要自托管方案的组织
- 想要简单快速文档的任何人

### DocsGo与Hugo有何不同？

| 特性 | DocsGo | Hugo |
|------|--------|------|
| 构建时间 | 实时 | 需要编译 |
| 部署 | 单文件 | 静态文件 |
| 搜索 | 内置 | 需要插件 |
| 密码保护 | 内置 | 需要外部系统 |

## 安装

### 系统要求是什么？

- **二进制**：无，可在任何OS上运行
- **源码**：Go 1.22或更高版本

### 支持哪些平台？

- Linux (AMD64, ARM64)
- macOS (Intel, Apple Silicon)
- Windows (AMD64)

### 可以在Docker中运行吗？

可以！请参阅[部署指南](./deployment)获取Docker说明。

## 使用

### 如何添加新文档？

在 `docs/` 目录创建 `.md` 文件：

```bash
echo "# 我的文档" > docs/my-doc.md
```

它会立即出现在 http://localhost:8080/my-doc

### 如何组织文档？

使用文件夹和frontmatter：

```
docs/
├── guide/
│   ├── index.md      (order: 1)
│   └── advanced.md   (order: 2)
└── index.md          (order: 1)
```

### 可以在Markdown中使用HTML吗？

可以，支持原始HTML用于高级格式化。

### 如何添加图片？

将图片放在 `docs/` 或 `web/static/`：

```markdown
![Alt text](./images/screenshot.png)
```

## 配置

### 在哪里设置端口？

在 `data/.env` 中：

```bash
PORT=3000
```

或通过命令行：

```bash
./docs-go -port 3000
```

### 如何用密码保护文档？

站点级密码：

```bash
# data/.env
PASSWORD_SITE=YourPassword
```

文档级密码：

```yaml
---
password: "Secret123"
---
```

### 可以禁用搜索吗？

搜索始终启用，但可以删除 `data/search.db`，除非你访问搜索，否则不会重建。

## 故障排除

### 服务器无法启动

**检查：**
1. 端口未被占用：`lsof -i :8080`
2. 文档目录存在：`ls docs/`
3. 二进制文件有执行权限：`chmod +x docs-go`

### 更改不显示

1. 检查文件在正确的目录
2. 确认文件扩展名是 `.md`
3. 刷新浏览器（强制刷新：Ctrl+F5）
4. 检查服务器日志

### 搜索不工作

重建搜索索引：

```bash
rm data/search.db
./docs-go
```

### 性能缓慢

1. 减小文档大小
2. 重建搜索索引
3. 检查可用磁盘空间
4. 生产环境使用 `EMBED_WEB=true`

## 高级

### 可以自定义主题吗？

可以！编辑 `web/views/` 中的模板：

- `doc.html` - 文档页面
- `password.html` - 密码页面
- `error.html` - 错误页面

### 如何添加自定义CSS？

添加CSS到 `web/static/css/custom.css` 并在模板中包含。

### 可以使用DocsGo写博客吗？

可以，虽然它针对文档优化。对于博客，考虑：
- 按日期组织文章
- 在frontmatter中使用标签
- 创建索引页面

### 有API吗？

有！请参阅[API参考](./api)了解端点。

## 贡献

### 如何贡献？

1. Fork仓库
2. 进行更改
3. 提交pull request

详情请参阅[开发指南](./develop)。

### 在哪里报告bug？

GitHub Issues: https://github.com/jinsi-ai/docs-go/issues

### 有路线图吗？

查看GitHub Issues了解计划功能和里程碑。

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
