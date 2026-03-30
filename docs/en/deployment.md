---
title: "Deployment Guide"
title_dir: "Deployment"
description: "Deploy DocsGo to production environments"
keywords: "deployment, production, server, hosting"
order: 40
lang: "en"
i18n:
  lang: "English"
  alternate: "/zh/deployment"
---

# Deployment Guide

Deploy DocsGo to production environments with confidence.

## Production Checklist

Before deploying to production:

- [ ] Set strong `PASSWORD_SITE` if needed
- [ ] Configure `EMBED_WEB=true` for single-file deployment
- [ ] Set appropriate `PORT` (80, 443, or custom)
- [ ] Test all documentation links
- [ ] Verify search functionality
- [ ] Set up HTTPS (recommended)

## Deployment Options

### 1. Standalone Server

```bash
# Configure production settings
cat > data/.env << 'EOF'
PORT=8080
DOCS_DIR=./docs
SITE_TITLE="Production Docs"
PASSWORD_SITE=YourSecurePassword
EMBED_WEB=true
EOF

# Run in background
nohup ./docs-go > app.log 2>&1 &
```

### 2. With Nginx (Recommended)

Nginx configuration:

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

Run DocsGo on localhost:

```bash
# data/.env
PORT=8080
EMBED_WEB=true
```

### 3. With HTTPS

Using Let's Encrypt with Certbot:

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Obtain certificate
sudo certbot --nginx -d docs.yourdomain.com

# Auto-renewal is configured automatically
```

### 4. Docker Deployment

Create `Dockerfile`:

```dockerfile
FROM alpine:latest

WORKDIR /app

# Copy binary and docs
COPY docs-go .
COPY docs/ ./docs/
COPY data/ ./data/

# Expose port
EXPOSE 8080

# Run
CMD ["./docs-go"]
```

Build and run:

```bash
docker build -t docs-go .
docker run -d -p 8080:8080 docs-go
```

### 5. Docker Compose

Create `docker-compose.yml`:

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

Run:

```bash
docker-compose up -d
```

## Cloud Deployment

### AWS EC2

```bash
# Launch EC2 instance
# Copy binary and docs
scp -i key.pem docs-go docs/* ec2-user@instance-ip:/home/ec2-user/

# SSH and run
ssh -i key.pem ec2-user@instance-ip
nohup ./docs-go > app.log 2>&1 &
```

### DigitalOcean

Use the One-Click App or manual setup:

```bash
# Create droplet
# SSH in
curl -LO https://github.com/jinsi-ai/docs-go/releases/latest/download/docs-go-linux-amd64
chmod +x docs-go-linux-amd64
./docs-go-linux-amd64
```

## Monitoring

### Health Check

```bash
# Check if service is running
curl http://localhost:8080
```

### Log Monitoring

```bash
# View logs
tail -f app.log

# Monitor errors
grep -i error app.log
```

## Backup Strategy

### Document Backup

```bash
# Backup docs directory
tar -czf docs-backup-$(date +%Y%m%d).tar.gz docs/

# Backup data (search index)
tar -czf data-backup-$(date +%Y%m%d).tar.gz data/
```

### Automated Backup

Cron job for daily backup:

```bash
# Edit crontab
crontab -e

# Add line
0 2 * * * tar -czf /backup/docs-$(date +\%Y\%m\%d).tar.gz /path/to/docs/
```

---

<div align="center" style="margin-top: 40px; padding: 20px; border-top: 1px solid #e5e7eb;">
  <p>
    <strong>DocsGo</strong> by <strong><a href="./brand">JinSi AI</a></strong> | 
    近思切问，AI务实 | NearThink AI, Pragmatic Tech
  </p>
</div>
