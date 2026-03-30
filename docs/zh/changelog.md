---
title: "更新日志"
title_dir: "更新日志"
description: "DocsGo版本历史和发布说明"
keywords: "更新日志, 发布说明, 版本历史, 更新"
order: 80
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/changelog"
---

# 更新日志

DocsGo的所有重要更改都记录在此文件中。

格式基于[Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)。

## [未发布]

### 新增
- 完整国际化（i18n）支持
- 双语文档（英文/中文）
- 品牌故事和理念文档
- 完整API参考

## [0.2.0] - 2025-03-30

### 新增
- 多语言文档支持
- 新Logo和品牌
- 品牌故事页面
- 完整英文文档
- 增强的frontmatter支持，带i18n字段

### 变更
- 改进带Logo的README
- 更好的文档结构组织
- 增强的页脚品牌信息

### 修复
- 文档链接一致性
- Frontmatter解析改进

## [0.1.0] - 2025-03-20

### 新增
- 初始发布
- 实时Markdown渲染
- 基于SQLite FTS的全文搜索
- 密码保护（站点和文档级别）
- 响应式Web界面
- 100+种语言的语法高亮
- 文件监控和自动重载
- 跨平台支持（Windows、Linux、macOS）
- 嵌入静态资源，单文件部署
- 文档树导航
- Frontmatter支持（title、order、password）
- 基于Cookie的认证

### 功能
- **实时预览**：编辑Markdown即时查看更改
- **全文搜索**：所有文档快速搜索
- **密码保护**：保护敏感文档
- **单文件**：轻松部署，嵌入资源
- **移动端适配**：适配所有设备
- **代码高亮**：精美语法高亮

## 版本管理

DocsGo遵循[语义化版本](https://semver.org/lang/zh-CN/)：

- **主版本**：不兼容的API更改
- **次版本**：新功能（向后兼容）
- **修订版本**：Bug修复（向后兼容）

## 发布计划

- **修订版本**：按需进行Bug修复
- **次版本**：每月新功能
- **主版本**：每年或破坏性变更

## 升级指南

### 从 0.1.x 到 0.2.x

1. 备份 `docs/` 目录
2. 下载新二进制文件或从源码重建
3. 按需更新 `data/.env`
4. 重启服务器

配置无破坏性变更。

## 路线图

### 计划 0.3.0
- [ ] 插件系统
- [ ] 主题定制
- [ ] 高级搜索过滤器
- [ ] 导出PDF
- [ ] 多用户支持

### 计划 1.0.0
- [ ] 稳定API
- [ ] 插件市场
- [ ] 高级分析
- [ ] 企业功能

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
