---
title: "开发指南"
title_dir: "开发"
order: 30
---

# 开发指南

## 项目结构

```
docs-go/
├── app/              # 应用层
│   ├── app.go        # 路由设置
│   └── docs/         # 文档处理器
├── pkg/              # 核心包
│   ├── config/       # 配置管理
│   ├── doc/          # 文档处理
│   ├── httpd/        # HTTP服务
│   └── webfs/        # 嵌入资源
├── web/              # 前端资源
│   ├── static/       # CSS/JS
│   └── views/        # HTML模板
└── main.go           # 入口
```

## 常用命令

```bash
# 开发运行
go run main.go

# 编译
go build -o docs-go

# 格式化代码
go fmt ./...

# 检查
go vet ./...
```

## 嵌入资源开发

默认 `EMBED_WEB=true`，所有资源已打包在二进制中。

开发时设置 `EMBED_WEB=false`，从文件系统加载便于实时调试：

```bash
# 启用嵌入后重新编译
go build -o docs-go
```

此时无需部署 `web/` 目录，单个二进制文件即可运行。

## 自定义模板

修改 `web/views/` 下的 HTML 文件：

- `doc.html` - 文档页面
- `password.html` - 密码验证页
- `error.html` - 错误页面

模板变量：`{{.title}}`, `{{.content}}`, `{{.docTree}}`, `{{.siteTitle}}`
