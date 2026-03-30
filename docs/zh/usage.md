---
title: "使用指南"
title_dir: "使用"
description: "学习如何高效使用DocsGo"
keywords: "使用, 指南, 教程, 工作流"
order: 30
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/usage"
---

# 使用指南

学习DocsGo的日常使用工作流。

## 本地开发工作流

### 1. 启动服务

```bash
./docs-go
```

服务将：
- 监视 `docs/` 目录的文件变化
- 自动重建文档树和搜索索引
- 在 http://localhost:8080 提供服务

### 2. 编辑文档

在 `docs/` 目录创建或编辑Markdown文件：

```bash
# 创建新文档
echo "# 我的文档" > docs/my-doc.md

# 编辑现有文档
vim docs/guide.md
```

更改会立即反映在浏览器中。

### 3. 组织文档

推荐结构：

```
docs/
├── index.md              # 首页
├── guide/
│   ├── index.md         # 指南概览
│   ├── getting-started.md
│   └── advanced.md
├── api/
│   ├── index.md
│   └── reference.md
└── README.md            # GitHub兼容
```

## 编写文档

### Markdown支持

DocsGo支持完整Markdown语法及扩展：

```markdown
# 标题1
## 标题2
### 标题3

**粗体** 和 *斜体*

- 列表项
- 另一项
  - 嵌套项

1. 编号列表
2. 第二项

[链接文字](url)

![图片alt](图片url)

| 表格 | 列 |
|------|-----|
| 数据 | 值 |

```代码块
func main() {
    println("Hello")
}
```
```

### 代码高亮

支持语言：Go、Python、JavaScript、Java、C++、Rust等100+种。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, DocsGo!")
}
```

### Frontmatter

在文档顶部添加元数据：

```yaml
---
title: "我的文档"
order: 1
password: "secret"
---
```

## 搜索功能

### 自动索引

- 所有Markdown内容自动索引
- 文件更改时索引自动更新
- 支持中文全文搜索

### 重建索引

删除 `data/search.db` 并重启：

```bash
rm data/search.db
./docs-go
```

## 密码保护

### 站点密码

在 `data/.env` 中设置：

```bash
PASSWORD_SITE=YourPassword
```

### 文档密码

在文档frontmatter中：

```yaml
---
password: "DocumentSecret"
---

## HTML 文件支持

除了 Markdown，DocsGo 也支持直接使用 HTML 文件：

```bash
# 创建 HTML 文件
cat > docs/custom-page.html << 'EOF'
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>自定义页面</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        h1 { color: #1e3a5f; }
    </style>
</head>
<body>
    <h1>这是一个自定义 HTML 页面</h1>
    <p>你可以使用任何 HTML、CSS 和 JavaScript。</p>
</body>
</html>
EOF
```

### 使用场景

- 复杂的交互式文档
- 需要自定义样式的页面
- 嵌入第三方组件
- 纯静态页面

### 注意事项

- HTML 文件不会自动生成目录导航
- HTML 内容不参与全文搜索
- 建议与 Markdown 混合使用

### 文件优先级

同一目录下多个入口文件时，访问优先级：

1. `index.html` - 最高优先级
2. `README.md` - 次之
3. `index.md` - 再次之

示例：
```
docs/
├── guide/
│   ├── index.html    ← 访问 /guide/ 时显示这个
│   ├── README.md     ← 会被忽略
│   └── index.md      ← 会被忽略
```

## 最佳实践

1. **使用有意义的文件名**：`installation.md` 而不是 `doc1.md`
2. **显式设置排序**：使用 `order`
3. **写好标题**：清晰且有描述性
4. **用文件夹组织**：将相关文档分组
5. **使用frontmatter**：添加元数据优化SEO

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
