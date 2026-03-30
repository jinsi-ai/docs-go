---
title: "部署指南"
title_dir: "部署"
description: "将DocsGo部署到生产环境"
keywords: "部署, 生产, 服务器, 托管"
order: 40
lang: "zh"
i18n:
  lang: "简体中文"
  alternate: "/en/deployment"
---

# 部署指南

将DocsGo部署到生产环境。

## 生产检查清单

部署到生产环境前：

- [ ] 如需设置强 `PASSWORD_SITE`
- [ ] 配置 `EMBED_WEB=true` 单文件部署
- [ ] 设置合适的 `PORT`（80、443或自定义）
- [ ] 测试所有文档链接
- [ ] 验证搜索功能
- [ ] 设置HTTPS（推荐）

## 部署选项

### 1. 独立服务器

```bash
# 配置生产环境
cat > data/.env << 'EOF'
PORT=8080
DOCS_DIR=./docs
SITE_TITLE="生产文档"
PASSWORD_SITE=YourSecurePassword
EMBED_WEB=true
EOF

# 后台运行
nohup ./docs-go > app.log 2>&1 &
```

### 2. 使用 Nginx（推荐）

Nginx配置：

```nginx
server {
    listen 80;
    server_name docs.yourdomain.com;
    
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

在localhost运行DocsGo：

```bash
# data/.env
PORT=8080
EMBED_WEB=true
```

### 3. 使用 HTTPS

使用Let's Encrypt和Certbot：

```bash
# 安装certbot
sudo apt install certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d docs.yourdomain.com

# 自动续期已配置
```

### 4. Docker部署

创建 `Dockerfile`：

```dockerfile
FROM alpine:latest

WORKDIR /app

# 复制二进制和文档
COPY docs-go .
COPY docs/ ./docs/
COPY data/ ./data/

# 暴露端口
EXPOSE 8080

# 运行
CMD ["./docs-go"]
```

构建和运行：

```bash
docker build -t docs-go .
docker run -d -p 8080:8080 docs-go
```

### 5. Docker Compose

创建 `docker-compose.yml`：

```yaml
version: '3'
services:
  docs:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./docs:/app/docs
      - ./data:/app/data
    restart: unless-stopped
```

运行：

```bash
docker-compose up -d
```

## 云部署

### AWS EC2

```bash
# 启动EC2实例
# 复制二进制和文档
scp -i key.pem docs-go docs/* ec2-user@instance-ip:/home/ec2-user/

# SSH并运行
ssh -i key.pem ec2-user@instance-ip
nohup ./docs-go > app.log 2>&1 &
```

### DigitalOcean

使用一键应用或手动设置：

```bash
# 创建droplet
# SSH登录
curl -LO https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64
```

## 监控

### 健康检查

```bash
# 检查服务是否运行
curl http://localhost:8080
```

### 日志监控

```bash
# 查看日志
tail -f app.log

# 监控错误
grep -i error app.log
```

## 备份策略

### 文档备份

```bash
# 备份docs目录
tar -czf docs-backup-$(date +%Y%m%d).tar.gz docs/

# 备份数据（搜索索引）
tar -czf data-backup-$(date +%Y%m%d).tar.gz data/
```

### 自动备份

每日备份的Cron任务：

```bash
# 编辑crontab
crontab -e

# 添加行
0 2 * * * tar -czf /backup/docs-$(date +\%Y\%m\%d).tar.gz /path/to/docs/
```

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">近思AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
