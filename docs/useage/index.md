---
title: "使用指南"
title_dir: "使用"
order: 20
---

# 使用指南

## 本地开发

适合个人或团队内部文档编写：

1. 启动服务：`./docs-go`
2. 编辑 `docs/` 下的 Markdown 文件
3. 浏览器实时预览更新

## 线上部署

生产环境部署建议：

```bash
# 1. 配置环境
cp data/.env.example data/.env
# 编辑 .env 设置 PORT 和 PASSWORD_SITE

# 2. 后台运行
nohup ./docs-go > app.log 2>&1 &

# 3. 使用 Nginx 反向代理
```

## 目录组织建议

```
docs/
├── index.md          # 首页
├── guide/
│   ├── index.md      # 目录说明
│   └── quick-start.md
├── api/
│   └── reference.md
└── README.md         # 兼容 GitHub
```

**文件优先级**: `index.html` > `README.md` > `index.md`

## 搜索功能

- 全文搜索自动索引所有 Markdown 内容
- 修改文档后索引自动更新
- 支持中文分词

删除 `data/search.db` 可重建索引。
