---
title: "API参考"
title_dir: "API"
description: "DocsGo API接口文档"
keywords: "API, 接口, 参考, REST"
order: 50
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/api"
---

# API参考

DocsGo提供RESTful API供程序访问。

## 基础URL

```
http://localhost:8080
```

## 认证

如果设置了 `PASSWORD_SITE`，需包含会话cookie：

```bash
curl -b "session=your-session-id" http://localhost:8080/api/...
```

## 接口列表

### 文档接口

#### 获取文档树

```http
GET /api/tree
```

返回完整文档树结构。

**响应：**
```json
{
  "name": "文档根目录",
  "path": "/",
  "isDir": true,
  "children": [
    {
      "name": "指南",
      "path": "/guide",
      "isDir": true,
      "children": [...]
    }
  ]
}
```

#### 获取文档内容

```http
GET /api/doc?path=/guide/quickstart
```

返回文档内容和元数据。

**响应：**
```json
{
  "title": "快速入门",
  "content": "<html>...",
  "path": "/guide/quickstart",
  "lastModified": "2025-01-01T00:00:00Z"
}
```

### 搜索接口

#### 搜索文档

```http
POST /api/search
Content-Type: application/json

{
  "query": "installation"
}
```

**响应：**
```json
{
  "results": [
    {
      "title": "安装指南",
      "path": "/guide/install",
      "snippet": "...安装说明..."
    }
  ],
  "total": 1,
  "query": "installation"
}
```

#### 获取搜索索引状态

```http
GET /api/search/status
```

**响应：**
```json
{
  "indexed": true,
  "documentCount": 42,
  "lastUpdate": "2025-01-01T00:00:00Z"
}
```

### 认证接口

#### 检查站点密码

```http
POST /api/auth/site-password
Content-Type: application/json

{
  "password": "your-password"
}
```

**响应：**
```json
{
  "success": true,
  "message": "认证成功"
}
```

#### 检查文档密码

```http
POST /api/auth/doc-password
Content-Type: application/json

{
  "path": "/secret-doc",
  "password": "doc-password"
}
```

## 响应格式

所有API响应遵循此格式：

```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

错误响应：

```json
{
  "code": 400,
  "message": "Bad request",
  "data": null
}
```

## HTTP状态码

| 状态码 | 含义 |
|--------|------|
| 200 | 成功 |
| 400 | 错误请求 |
| 401 | 未授权 |
| 404 | 未找到 |
| 500 | 服务器错误 |

## 限流

当前未实现限流。生产环境建议添加：

- Nginx限流
- API网关
- 应用层限流

## SDK

暂无官方SDK。使用标准HTTP客户端：

### cURL示例

```bash
# 获取文档树
curl http://localhost:8080/api/tree

# 搜索文档
curl -X POST http://localhost:8080/api/search \
  -H "Content-Type: application/json" \
  -d '{"query": "config"}'
```

### Python示例

```python
import requests

# 搜索
response = requests.post(
    'http://localhost:8080/api/search',
    json={'query': 'deployment'}
)
results = response.json()
print(results['data']['results'])
```

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
