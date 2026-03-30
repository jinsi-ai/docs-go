# GitHub项目包装技能 (Project Packaging Skill)

> 基于近思AI项目包装实践总结
> 面向全球用户的标准化项目包装流程

---

## 概述

本技能文档提供一套可复用的GitHub项目包装标准流程，特别针对IP品牌整合和国际化文档架构。

**适用场景**：
- 开源项目首次发布
- 项目国际化升级
- 品牌重塑和文档重构
- 多语言文档体系建设

**核心理念**：近思切问，务实落地

---

## 第一部分：IP品牌层 (IP Branding)

### 1.1 品牌基础建设

- [ ] **确定IP名称**
  - 中英文名称对照
  - 域名/账号可用性检查（GitHub、Twitter、Discord等）
  - 参考：近思AI / JinSi AI

- [ ] **提炼品牌故事**（3分钟精简版）
  - 出处/灵感来源
  - 核心寓意解释
  - 为什么选择这个名字
  - 参考模板见下方

- [ ] **设计品牌口号**（双语）
  - 中文：近思切问，AI务实
  - 英文：NearThink AI, Pragmatic Tech
  - 要求：简短、有力、易记

- [ ] **定义核心价值观**（4个关键词）
  - 务实 | Pragmatic
  - 深度 | Deep Thinking
  - 渐进 | Progressive
  - 真诚 | Genuine

- [ ] **明确品牌定位**
  - 身份定位（如：务实派AI技术专家）
  - 态度定位（如：不追风口，只讲干货）
  - 风格定位（如：深度思考，落地实操）
  - 目标受众（核心/次要/延伸）

### 1.2 视觉识别系统

- [ ] **Logo设计**
  - 主Logo：200x200px SVG
  - Favicon：100x100px SVG
  - 配色方案（主色+辅色）
  - 风格：简约、专业、有辨识度

- [ ] **平台统一命名**
  - GitHub：项目名
  - Twitter/X：同名或相关名
  - Discord：同名服务器
  - 文档站点：统一子域名

### 1.3 品牌故事模板

```
## 品牌故事（精简版）

在[领域]浪潮中，有人在追逐风口，有人在贩卖焦虑，有人在讲高深概念。

[你的品牌名]，选择另一条路。

"[品牌名含义]"——这是[出处]的智慧，也是今天的行动指南。

[品牌名]，意味着从身边的问题出发，思考那些真正重要的事。
不空谈未来，只解决当下。
不追逐热点，只深挖本质。

[品牌名]，带你在[领域]时代务实前行。
一步一个脚印，从近处走向远方。

---

## 口号体系

主口号：
- 中文：[你的口号]
- 英文：[Your Slogan]

场景口号：
- 内容结尾：[品牌名]，解决真问题
- 品牌介绍：不追风口，只讲干货
```

---

## 第二部分：文档架构层 (Documentation Architecture)

### 2.1 目录结构设计

```
docs/
├── index.md              # 语言选择页或品牌首页
├── logo.svg              # Logo文件（便于引用）
├── favicon.svg           # Favicon
├── en/                   # 英文文档
│   ├── index.md         # 英文首页
│   ├── quickstart.md    # 快速入门
│   ├── config.md        # 配置指南
│   ├── usage.md         # 使用指南
│   ├── deployment.md    # 部署指南
│   ├── api.md           # API参考
│   ├── best-practices.md# 最佳实践
│   ├── faq.md           # 常见问题
│   ├── changelog.md     # 更新日志
│   ├── develop.md       # 开发指南
│   └── brand.md         # 品牌介绍
└── zh/                   # 中文文档（结构同上）
```

### 2.2 Frontmatter标准配置

每篇文档必须包含：

```yaml
---
title: "页面标题"                    # 页面标题
title_dir: "导航显示名"              # 目录导航显示名称（可选）
description: "页面描述"              # SEO描述
keywords: "关键词1,关键词2"          # SEO关键词
order: 10                            # 排序权重（10-100）
lang: "zh"                          # 语言代码
i18n:
  lang: "简体中文"                   # 显示语言名称
  alternate: "/en/page"             # 翻译版本链接
---
```

### 2.3 首页设计（中文品牌首页模板）

**结构**：
1. Logo + 品牌口号（居中）
2. 项目简介（1-2段）
3. 核心功能列表（5-6个）
4. 品牌故事区块（IP内容）
5. 快速开始代码示例
6. 双语文档导航
7. 品牌Footer（渐变色背景）

**参考链接**：
- 本项目的 docs/index.md

### 2.4 品牌页脚标准化

每篇文档底部必须包含：

```markdown
---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>[项目名]</strong> by <strong>[IP品牌名]</strong> | 
    [中文口号] | [英文口号]
  </p>
  <p>
    <a href="[GitHub]">GitHub</a> • 
    <a href="[Issues]">问题反馈</a> • 
    <a href="[Alternate Lang]">[语言]</a>
  </p>
</div>
```

---

## 第三部分：GitHub仓库层 (Repository Layer)

### 3.1 README标准化

**必备文件**：
- [ ] `README.md`（英文，默认显示）
- [ ] `README_CN.md`（中文，可选但推荐）
- [ ] `LICENSE`（MIT推荐）
- [ ] `AGENTS.md`（AI编码代理指南）

**README结构**：
1. Logo（居中，120x120px）
2. 项目标题 + 徽章（Badges）
3. 一句话描述
4. 为什么选择本项目（对比表格）
5. 核心功能（5-6点）
6. 快速开始（代码示例）
7. 安装方式
8. 配置说明
9. 技术栈
10. 许可证
11. **品牌章节**（About [IP Name]）

### 3.2 AGENTS.md内容清单

必须包含：

```markdown
# AGENTS.md

## 项目概述
- 技术栈
- 核心功能
- 目标用户

## 构建命令
- 编译命令
- 运行命令
- 交叉编译

## 测试命令
- 运行所有测试
- 单包测试
- 单函数测试
- 覆盖率测试

## 代码规范
- 导入顺序
- 命名规范
- 注释规范
- 错误处理
- 并发规范

## 项目结构
- 目录说明
- 关键文件
```

### 3.3 提交规范

使用 Conventional Commits：

```
feat: 新功能
fix: Bug修复
docs: 文档更新
style: 代码格式（不影响功能）
refactor: 重构
test: 测试相关
chore: 构建/工具相关
```

---

## 第四部分：技术实现层 (Technical Implementation)

### 4.1 文档格式支持

- [ ] **Markdown**：完整支持GFM语法
- [ ] **HTML**：直接放置 .html 文件
- [ ] **文件优先级**：index.html > README.md > index.md

### 4.2 国际化实现

- [ ] 语言标识：Frontmatter中的 `lang` 字段
- [ ] 翻译链接：Frontmatter中的 `i18n.alternate`
- [ ] 首页切换：语言选择器或直接进入默认语言
- [ ] SEO优化：每页独立的多语言meta标签

### 4.3 响应式设计

- [ ] 移动端适配
- [ ] 平板适配
- [ ] 桌面端适配
- [ ] 代码高亮（100+语言）

---

## 第五部分：发布前检查清单 (Pre-Release Checklist)

### 5.1 品牌检查

- [ ] Logo显示正常
- [ ] 品牌口号出现在所有关键页面
- [ ] 品牌故事完整且易懂
- [ ] 所有链接正确（没有404）

### 5.2 文档检查

- [ ] 首页内容完整
- [ ] 快速入门可执行（测试一遍）
- [ ] 所有文档有品牌页脚
- [ ] Frontmatter格式正确
- [ ] 中英文链接互相对应

### 5.3 技术检查

- [ ] README Logo路径正确
- [ ] AGENTS.md信息准确
- [ ] 代码示例可运行
- [ ] 搜索功能正常

### 5.4 SEO检查

- [ ] 每个页面有title和description
- [ ] 关键词布局合理
- [ ] GitHub Topics设置（中英文关键词）

---

## 第六部分：维护与迭代 (Maintenance)

### 6.1 定期更新

- [ ] 每月更新：文档时效性检查
- [ ] 每季更新：品牌故事微调
- [ ] 每年更新：整体架构评估

### 6.2 版本发布

- [ ] 更新 CHANGELOG.md
- [ ] 更新文档中的版本号
- [ ] 创建Git Tag
- [ ] 发布Release Notes（双语）

---

## 附录：参考链接

### 本项目参考

- 品牌故事：[docs/zh/brand.md](./zh/brand.md)
- 首页设计：[docs/index.md](../index.md)
- README模板：[README.md](../../../README.md)
- AGENTS模板：[AGENTS.md](../../../AGENTS.md)

### 外部资源

- [Conventional Commits](https://www.conventionalcommits.org/) - 提交规范
- [Keep a Changelog](https://keepachangelog.com/) - 更新日志规范
- [Semantic Versioning](https://semver.org/lang/zh-CN/) - 语义化版本
- [Markdown Guide](https://www.markdownguide.org/) - Markdown语法

### 设计资源

- Logo设计：[Canva](https://www.canva.com/)
- SVG编辑：[Figma](https://www.figma.com/) 或 [Inkscape](https://inkscape.org/)
- 配色方案：[Coolors](https://coolors.co/)

---

## 使用建议

1. **首次包装**：严格按照检查清单执行，不要跳过任何步骤
2. **品牌优先**：IP内容是差异化的关键，投入足够时间
3. **务实落地**：文档要实用，不要过度设计
4. **渐进完善**：不必一次完美，可以迭代优化
5. **双语并重**：面向全球，英文文档质量同样重要

---

> 近思切问，务实落地
> 
> 本技能由 近思AI 总结提炼 | JinSi AI
