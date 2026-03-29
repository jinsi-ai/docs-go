# docs-go

## 项目简介
docs-go 是一个功能完整的轻量级文档网站生成器，基于 Go 语言开发，支持 Markdown 和 HTML 文档的实时渲染与展示。项目提供了完整的文档管理、密码保护、文件监控和响应式UI界面。

## 核心功能
- **多格式文档支持**: 同时支持 Markdown (.md) 和 HTML (.html) 文档
- **智能导航菜单**: 自动生成树状导航菜单，支持目录层级和文件排序
- **Frontmatter 支持**: 通过 YAML frontmatter 设置文档标题、排序和密码保护
- **实时文件监控**: 使用 fsnotify 监控文档变化，自动更新目录结构
- **密码保护机制**: 支持文档级别的密码保护，基于 Cookie 的认证管理
- **响应式设计**: 适配桌面和移动设备的现代化 UI 界面
- **代码高亮**: 集成 goldmark-highlighting 提供语法高亮支持
- **面包屑导航**: 自动生成当前页面的面包屑导航路径

## 技术栈
- **后端框架**: Go + Gin Web Framework
- **前端技术**: HTML5 + Tailwind CSS + 原生 JavaScript
- **Markdown 引擎**: goldmark (支持 GFM 扩展)
- **文件监控**: fsnotify (实时监控文件变化)
- **代码高亮**: goldmark-highlighting (基于 Chroma)
- **认证管理**: 自定义 Cookie 认证系统

## 项目结构
```
docs-go/
├── app/                    # 应用层代码
│   ├── app.go             # 路由设置和初始化
│   └── docs/              # 文档处理模块
│       ├── auth.go        # 认证和密码验证
│       ├── doc.go         # 文档处理器定义
│       └── file.go        # 文件处理逻辑
├── pkg/                    # 核心包模块
│   ├── auth/              # 认证相关
│   │   └── cookie.go      # Cookie管理器
│   ├── config/            # 配置管理
│   │   └── config.go      # 配置结构体和解析
│   ├── doc/               # 文档处理核心
│   │   ├── doctree.go     # 文档树构建和管理
│   │   └── document.go    # 文档渲染和frontmatter解析
│   ├── httpd/             # HTTP服务器
│   │   └── server.go      # 服务器启动
│   ├── json/              # JSON响应处理
│   │   ├── error.go       # 错误响应
│   │   └── msg.go         # 消息响应
│   └── watcher/           # 文件监控
│       └── filewatcher.go # 文件变化监控器
├── web/                    # 前端资源
│   ├── static/            # 静态文件
│   │   └── css/           # 样式文件
│   └── views/             # HTML模板
│       ├── doc.html       # 文档页面模板
│       ├── error.html     # 错误页面模板
│       └── password.html  # 密码验证页面模板
├── main.go                # 程序入口
├── go.mod                 # Go模块依赖
├── go.sum                 # 依赖校验文件
└── readme.md              # 项目说明文档
```

## 快速开始

### 环境要求
- Go 1.22 或更高版本

### 安装和启动
1. 克隆或下载项目到本地
2. 进入项目目录：`cd docs-go-min`
3. 安装依赖：`go mod tidy`
4. 启动服务器：`go run main.go`
5. 访问 `http://localhost:8080` 查看文档

### 默认配置
- **服务器端口**: 8080
- **文档目录**: `../docs` (相对于可执行文件位置)
- **支持文件类型**: `.md`, `.html`, `.htm`

### 自定义配置
可以通过命令行参数自定义配置：
```bash
# 自定义文档目录
go run main.go -docs ./my-docs

# 自定义端口
go run main.go -port 8081

# 同时指定多个参数
go run main.go -docs ./my-docs -port 8081
```

## 文档组织

### 文件命名规则
- **Markdown文件**: 使用 `.md` 扩展名
- **HTML文件**: 使用 `.html` 或 `.htm` 扩展名
- **索引文件**: 目录中的 `index.md` 或 `index.html` 作为该目录的默认页面
- **URL路径**: 文件路径映射为URL路径，例如 `docs/api/guide.md` → `/api/guide`

### 文件优先级
当同一路径存在多个文件时，按以下优先级显示：
1. `.html` 文件
2. `.md` 文件
3. 目录索引文件 (`index.html` > `index.md`)

## Frontmatter 配置

### 基本语法
在文档文件开头添加 YAML frontmatter 块：
```yaml
---
title: 自定义标题     # 文档显示标题（覆盖文件名）
title_dir: 目录名     # 目录显示名称（仅对目录有效）
sort: 1              # 排序权重（数值越小越靠前）
password: mypass     # 文档密码（可选）
---
```

### 文档标题配置
```yaml
---
title: API使用指南    # 覆盖文件名作为显示标题
sort: 10             # 设置排序权重
---
```

### 目录标题配置
在目录的 `index.md` 或 `index.html` 中设置：
```yaml
---
title_dir: API文档    # 设置目录在导航菜单中的显示名称
title: 首页           # 目录页面的标题
sort: 5              # 目录排序权重
---
```

### 密码保护配置
```yaml
---
title: 机密文档
password: secret123   # 设置访问密码
---
```

## HTML文档支持

### 基本用法
- HTML文件可以直接放置在文档目录中
- 支持与Markdown文件混合使用
- 系统会原样输出HTML内容，不进行额外处理

### Frontmatter支持
HTML文件同样支持frontmatter配置：
```html
---
title: 自定义HTML页面
title_dir: HTML示例
sort: 90
---
<!DOCTYPE html>
<html>
<head>
    <title>我的HTML页面</title>
</head>
<body>
    <h1>这是HTML内容</h1>
    <p>直接显示，无需转换</p>
</body>
</html>
```

### 路径映射示例
- `docs/api/index.html` → 通过 `/api` 访问
- `docs/guide/tutorial.html` → 通过 `/guide/tutorial` 访问

## 高级功能

### 实时文件监控
项目集成了文件监控功能，当文档目录中的文件发生变化时，系统会自动：
- 重新构建文档树结构
- 更新导航菜单
- 无需重启服务器即可看到最新内容

### 密码保护系统
#### 密码设置
在文档的frontmatter中添加 `password` 字段：
```yaml
---
title: 机密文档
password: mypassword123
---
```

#### 认证机制
- 基于Cookie的认证系统
- 密码验证成功后24小时内免登录
- 支持明文密码和MD5加密密码
- 每个文档可以设置独立的密码

#### 认证流程
1. 用户访问受密码保护的文档
2. 系统检查Cookie认证状态
3. 如果未认证，显示密码输入界面
4. 验证成功后设置认证Cookie
5. 跳转到目标文档页面

### 响应式UI设计
- **桌面端**: 左侧固定导航栏，右侧内容区域
- **移动端**: 可折叠的汉堡菜单，适配小屏幕
- **导航树**: 支持目录展开/折叠，自动高亮当前页面
- **面包屑**: 显示当前页面的层级路径

## 开发指南

### 项目架构
项目采用分层架构设计：
- **main.go**: 程序入口，配置初始化
- **app层**: 路由设置和业务逻辑组装
- **pkg层**: 核心功能模块（认证、配置、文档处理等）
- **web层**: 前端模板和静态资源

### 核心模块说明

#### 文档树构建 (pkg/doc/doctree.go)
- 递归扫描文档目录结构
- 解析frontmatter获取标题和排序信息
- 构建树状导航数据结构
- 支持目录和文件的智能排序

#### 文档渲染 (pkg/doc/document.go)
- 支持Markdown和HTML双格式
- 集成goldmark引擎，支持GFM语法
- 代码高亮基于Chroma引擎
- Frontmatter解析和内容分离

#### 文件监控 (pkg/watcher/filewatcher.go)
- 基于fsnotify的实时文件监控
- 自动忽略临时文件和隐藏文件
- 文件变化时自动重建文档树
- 线程安全的并发处理

#### 认证管理 (pkg/auth/cookie.go)
- 基于MD5的密码加密
- Cookie生命周期管理
- 路径相关的认证机制
- 安全的Cookie设置

### 扩展开发
如需扩展功能，可以修改以下模块：
- 添加新的文档格式支持：修改 `pkg/doc/document.go`
- 自定义UI样式：修改 `web/static/css/` 中的样式文件
- 添加新的API接口：在 `app/docs/` 中添加处理器
- 修改认证逻辑：调整 `pkg/auth/` 相关代码

## 故障排除

### 常见问题

#### 文档目录不存在
```bash
# 系统会自动创建文档目录
# 或手动创建目录
mkdir ../docs
```

#### 端口被占用
```bash
# 使用其他端口
 go run main.go -port 8081
```

#### 文件变化未检测到
- 确保文件监控服务正常运行
- 检查文件权限和路径正确性
- 重启服务器以重新初始化监控

#### 密码认证失败
- 检查frontmatter中的password字段格式
- 清除浏览器Cookie重新认证
- 确认密码输入正确

### 日志调试
启动时添加详细日志输出：
```bash
go run main.go 2>&1 | tee server.log
```

## 性能优化

### 文档树缓存
- 文档树结构在内存中缓存
- 文件变化时自动更新缓存
- 减少重复的文件系统扫描

### 静态资源优化
- 使用Tailwind CSS压缩版本
- 内联关键CSS和JavaScript
- 浏览器缓存优化

### 并发处理
- 使用读写锁保护文档树访问
- 文件监控使用单独的goroutine
- HTTP请求处理并发安全

## 许可证

本项目采用 MIT 许可证，详见项目根目录的 LICENSE 文件。

## 贡献指南

欢迎提交 Issue 和 Pull Request 来改进项目。

### 开发环境设置
1. Fork 项目到个人仓库
2. 克隆项目：`git clone https://github.com/your-username/docs-go.git`
3. 创建功能分支：`git checkout -b feature/new-feature`
4. 提交更改：`git commit -am 'Add new feature'`
5. 推送到分支：`git push origin feature/new-feature`
6. 创建 Pull Request

### 代码规范
- 遵循 Go 语言官方代码规范
- 使用 `gofmt` 格式化代码
- 添加必要的注释和文档
- 确保测试通过

## 更新日志

### v1.0.0 (当前版本)
- ✅ 支持 Markdown 和 HTML 文档渲染
- ✅ 自动生成树状导航菜单
- ✅ Frontmatter 配置支持
- ✅ 实时文件监控
- ✅ 文档级密码保护
- ✅ 响应式 UI 设计
- ✅ 代码高亮显示
- ✅ 面包屑导航

## 相关项目

- [goldmark](https://github.com/yuin/goldmark) - Go 语言的 Markdown 解析器
- [gin](https://github.com/gin-gonic/gin) - Go 语言的 Web 框架
- [fsnotify](https://github.com/fsnotify/fsnotify) - 文件系统监控库

## 联系方式

如有问题或建议，请通过以下方式联系：
- 提交 GitHub Issue
- 发送邮件至项目维护者

---

**docs-go** - 让文档管理更简单、更高效！