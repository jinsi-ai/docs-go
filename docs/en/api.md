---
title: "API Reference"
title_dir: "API"
description: "API documentation for DocsGo"
keywords: "API, reference, endpoints, REST"
order: 50
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/api"
---

# API Reference

DocsGo provides RESTful APIs for programmatic access.

## Base URL

```
http://localhost:8080
```

## Authentication

If `PASSWORD_SITE` is set, include the session cookie:

```bash
curl -b "session=your-session-id" http://localhost:8080/api/...
```

## Endpoints

### Document APIs

#### Get Document Tree

```http
GET /api/tree
```

Returns the complete document tree structure.

**Response:**
```json
{
  "name": "文档根目录",
  "path": "/",
  "isDir": true,
  "children": [
    {
      "name": "Guide",
      "path": "/guide",
      "isDir": true,
      "children": [...]
    }
  ]
}
```

#### Get Document Content

```http
GET /api/doc?path=/guide/quickstart
```

Returns document content and metadata.

**Response:**
```json
{
  "title": "Quick Start",
  "content": "<html>...",
  "path": "/guide/quickstart",
  "lastModified": "2025-01-01T00:00:00Z"
}
```

### Search APIs

#### Search Documents

```http
POST /api/search
Content-Type: application/json

{
  "query": "installation"
}
```

**Response:**
```json
{
  "results": [
    {
      "title": "Installation Guide",
      "path": "/guide/install",
      "snippet": "...installation instructions..."
    }
  ],
  "total": 1,
  "query": "installation"
}
```

#### Get Search Index Status

```http
GET /api/search/status
```

**Response:**
```json
{
  "indexed": true,
  "documentCount": 42,
  "lastUpdate": "2025-01-01T00:00:00Z"
}
```

### Authentication APIs

#### Check Site Password

```http
POST /api/auth/site-password
Content-Type: application/json

{
  "password": "your-password"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Authentication successful"
}
```

#### Check Document Password

```http
POST /api/auth/doc-password
Content-Type: application/json

{
  "path": "/secret-doc",
  "password": "doc-password"
}
```

## Response Format

All API responses follow this format:

```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

Error response:

```json
{
  "code": 400,
  "message": "Bad request",
  "data": null
}
```

## HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 400 | Bad request |
| 401 | Unauthorized |
| 404 | Not found |
| 500 | Server error |

## Rate Limiting

Currently no rate limiting is implemented. For production use, consider adding:

- Nginx rate limiting
- API gateway
- Application-level throttling

## SDKs

No official SDKs yet. Use standard HTTP clients:

### cURL Example

```bash
# Get document tree
curl http://localhost:8080/api/tree

# Search documents
curl -X POST http://localhost:8080/api/search \
  -H "Content-Type: application/json" \
  -d '{"query": "config"}'
```

### Python Example

```python
import requests

# Search
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
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
