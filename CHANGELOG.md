# Changelog

## v0.2 (2025-03-29)

### 改进
- **默认嵌入静态资源** - `EMBED_WEB` 默认改为 `true`，单文件即可部署
- **自动创建默认文档** - docs 目录不存在时自动生成 README.md 引导页
- **禁用页面弹性滚动** - 添加 `overscroll-behavior: none` 防止页面弹动
- **移除 bin 目录跟踪** - 避免仓库包含编译产物

### 文档
- **优化 SEO** - docs/index.md 添加关键词和详细使用场景
- **丰富 README** - 添加多平台下载说明和对比表格

**完整变更**: https://github.com/jinsi-ai/docs-go/compare/v0.1...v0.2

## v0.1 (2025-03-29)

### 功能
- 实时 Markdown 渲染（支持 GFM）
- 自动目录树生成
- 全文搜索（SQLite FTS）
- 文档级密码保护
- 站点级密码保护
- 静态资源嵌入（可选）
- 响应式 UI
- 代码高亮

### 技术栈
- Go + Gin
- Tailwind CSS
- goldmark
- SQLite

**完整变更**: https://github.com/jinsi-ai/docs-go/commits/v0.1
