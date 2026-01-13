# feishu2md

[![Golang - feishu2md](https://img.shields.io/github/go-mod/go-version/wsine/feishu2md?color=%2376e1fe&logo=go)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/wsine/feishu2md?color=orange&logo=github)](https://github.com/Wsine/feishu2md/releases)
[![Docker - feishu2md](https://img.shields.io/badge/Docker-feishu2md-2496ed?logo=docker&logoColor=white)](https://hub.docker.com/r/wwwsine/feishu2md)
[![Render - feishu2md](https://img.shields.io/badge/Render-feishu2md-4cfac9?logo=render&logoColor=white)](https://feishu2md.onrender.com)
![Last Review](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fbadge-last-review.wsine.workers.dev%2FWsine%2Ffeishu2md&query=%24.reviewed_at&label=last%20review)

> 🚀 **飞书文档转换 Markdown 工具** - 支持文档、知识库批量下载，智能图片压缩，一键部署

这是一个功能强大的飞书文档转换工具，使用 Go 语言实现，支持将飞书文档转换为 Markdown 格式，并具备智能图片压缩、WebP 转换等高级功能。

**✨ 核心特性**

- 📄 **单文档/批量下载** - 支持单个文档、文件夹、知识库批量转换
- 🖼️ **智能图片压缩** - 自动压缩图片至 800px 宽度，WebP 格式转换，减少 70%+ 文件大小
- 🌐 **Web 界面** - 友好的 Web 界面，支持在线转换和下载
- 📦 **多种部署方式** - 支持 Docker、Render、Vercel、本地运行等多种部署方式
- 🔒 **安全可靠** - 使用官方 API，支持企业自建应用
- 🎨 **Markdown 优化** - 优化表格格式、图片处理，可直接粘贴到 MD 编辑器

---

## 📖 目录

- [快速开始](#快速开始)
- [功能特性](#功能特性)
- [获取 API 凭证](#获取-api-凭证)
- [使用方式](#使用方式)
  - [命令行版本](#命令行版本)
  - [Web 界面版本](#web-界面版本)
  - [Docker 部署](#docker-部署)
  - [Render 部署](#render-部署)
  - [Vercel 部署](#vercel-部署)
- [高级功能](#高级功能)
- [文档](#文档)
- [常见问题](#常见问题)
- [贡献指南](#贡献指南)
- [许可证](#许可证)

---

## 🚀 快速开始

### 方式一：使用 Web 界面（推荐）

1. 访问在线演示：https://feishu2md.onrender.com
2. 粘贴飞书文档链接
3. 勾选"嵌入图片到 Markdown"（推荐）
4. 点击"开始转换"
5. 等待转换完成，下载或复制 Markdown

### 方式二：使用 Docker

```bash
docker run -it --rm -p 8080:8080 \
  -e FEISHU_APP_ID=<your_id> \
  -e FEISHU_APP_SECRET=<your_secret> \
  -e GIN_MODE=release \
  wwwsine/feishu2md
```

然后访问 http://localhost:8080

### 方式三：命令行工具

下载最新版本的 [可执行文件](https://github.com/Wsine/feishu2md/releases)，配置 API 凭证后：

```bash
# 下载单个文档
feishu2md dl "https://domain.feishu.cn/docx/docxtoken"

# 批量下载文件夹
feishu2md dl --batch "https://domain.feishu.cn/drive/folder/foldertoken"

# 批量下载知识库
feishu2md dl --wiki "https://domain.feishu.cn/wiki/settings/123456789101112"
```

---

## ✨ 功能特性

### 🎯 核心功能

| 功能 | 描述 | 状态 |
|------|------|------|
| 单文档转换 | 将单个飞书文档转换为 Markdown | ✅ |
| 批量转换 | 支持文件夹、知识库批量转换 | ✅ |
| 图片下载 | 自动下载文档中的图片 | ✅ |
| 图片压缩 | 智能压缩图片，WebP 格式转换 | ✅ |
| Web 界面 | 友好的在线转换界面 | ✅ |
| 命令行工具 | 支持命令行批量操作 | ✅ |

### 🖼️ 图片智能压缩

- **自动调整尺寸**：超过 800px 宽度的图片自动缩放
- **WebP 转换**：将图片转换为 WebP 格式，减少 70%+ 文件大小
- **Base64 嵌入**：支持将图片直接嵌入 Markdown，方便复制粘贴
- **压缩日志**：实时显示压缩效果，方便调试

**压缩效果示例**：

```
[图片压缩] 原始: 2048576 bytes, 压缩后: 512000 bytes, 减少: 75.0%
```

### 🌐 Web 界面特性

- 🎨 **现代化 UI** - 简洁美观的界面设计
- 🌍 **中文本地化** - 完整的中文界面
- 📊 **状态指示** - 实时显示转换进度
- 🎯 **智能提示** - 清晰的操作指引
- 📱 **响应式设计** - 支持移动端访问

---

## 🔑 获取 API 凭证

### 步骤 1：创建飞书应用

1. 访问 [飞书开发者后台](https://open.feishu.cn/app)
2. 点击"创建企业自建应用"（个人版）
3. 填写应用信息（名称、描述等）

### 步骤 2：配置权限

打开"权限管理"，开通以下必要权限：

| 权限名称 | 权限代码 | 用途 |
|---------|---------|------|
| 查看新版文档 | `docx:document:readonly` | 获取文档基本信息和内容 |
| 下载云文档中的图片和附件 | `docs:document.media:download` | 下载文档中的图片 |
| 查看、评论、编辑和管理云空间中所有文件 | `drive:file:readonly` | 获取文件夹中的文件清单 |
| 查看知识库 | `wiki:wiki:readonly` | 获取知识空间节点信息 |

### 步骤 3：获取凭证

打开"凭证与基础信息"，获取：

- **App ID**：类似 `cli_a6727c4ffc71d00b`
- **App Secret**：类似 `dt7LyzH6HOexxH4z9ssXpghYgE8PIvSI`

---

## 📚 使用方式

### 命令行版本

#### 安装

从 [Release](https://github.com/Wsine/feishu2md/releases) 下载对应平台的可执行文件，放置到 PATH 路径中。

#### 配置

```bash
# 设置 API 凭证
feishu2md config --appId <your_id> --appSecret <your_secret>

# 查看配置
feishu2md config
```

#### 基本使用

```bash
# 查看帮助
feishu2md -h

# 下载单个文档
feishu2md dl "https://domain.feishu.cn/docx/docxtoken"

# 指定输出目录
feishu2md dl -o ./output "https://domain.feishu.cn/docx/docxtoken"

# 批量下载文件夹
feishu2md dl --batch "https://domain.feishu.cn/drive/folder/foldertoken"

# 批量下载知识库
feishu2md dl --wiki "https://domain.feishu.cn/wiki/settings/123456789101112"
```

#### 高级选项

```bash
# 导出 API 响应（调试用）
feishu2md dl --dump "https://domain.feishu.cn/docx/docxtoken"

# 查看版本
feishu2md -v
```

### Web 界面版本

#### 本地运行

```bash
# 克隆仓库
git clone https://github.com/Wsine/feishu2md.git
cd feishu2md

# 配置环境变量
cp .env.example .env
# 编辑 .env 文件，填入你的 FEISHU_APP_ID 和 FEISHU_APP_SECRET

# 启动服务
./start.sh

# 访问 http://localhost:8080
```

#### 使用 ngrok 公网访问

```bash
# 安装 ngrok
brew install ngrok  # macOS
# 或访问 https://ngrok.com 下载

# 启动 ngrok
ngrok http 8080

# 使用 ngrok URL 访问
```

### Docker 部署

#### 使用 Docker 命令

```bash
docker run -it --rm -p 8080:8080 \
  -e FEISHU_APP_ID=<your_id> \
  -e FEISHU_APP_SECRET=<your_secret> \
  -e GIN_MODE=release \
  wwwsine/feishu2md
```

#### 使用 Docker Compose

创建 `docker-compose.yml`：

```yaml
version: '3'
services:
  feishu2md:
    image: wwwsine/feishu2md
    environment:
      FEISHU_APP_ID: <your_id>
      FEISHU_APP_SECRET: <your_secret>
      GIN_MODE: release
    ports:
      - "8080:8080"
```

启动服务：

```bash
docker compose up -d
```

### Render 部署（免费）

Render 提供免费的 Web Service 部署，每月 750 小时免费额度。

#### 详细步骤

1. **Fork 项目到你的 GitHub**
   - 访问 https://github.com/Wsine/feishu2md
   - 点击右上角 "Fork" 按钮

2. **注册 Render 账号**
   - 访问 https://render.com
   - 使用 GitHub 账号登录

3. **创建 Web Service**
   - 点击 "New +" → "Web Service"
   - 连接你的 GitHub 仓库
   - 选择 `feishu2md` 仓库

4. **配置环境变量**
   ```
   FEISHU_APP_ID = <your_app_id>
   FEISHU_APP_SECRET = <your_app_secret>
   GIN_MODE = release
   PORT = 8080
   ```

5. **部署应用**
   - 点击 "Create Web Service"
   - 等待 2-5 分钟部署完成
   - 访问生成的 URL

**详细教程**：参考 [docs/本地部署使用指南_20251230_v1.1.md](docs/本地部署使用指南_20251230_v1.1.md)

### Vercel 部署

Vercel 提供免费的静态网站托管，适合前端部署。

#### 详细步骤

1. **安装 Vercel CLI**
   ```bash
   npm install -g vercel
   ```

2. **登录 Vercel**
   ```bash
   vercel login
   ```

3. **部署项目**
   ```bash
   vercel
   ```

4. **配置环境变量**
   在 Vercel Dashboard 中添加环境变量

**详细教程**：参考 [docs/本地部署使用指南_20251230_v1.1.md](docs/本地部署使用指南_20251230_v1.1.md)

---

## 🔧 高级功能

### 图片压缩配置

在 `core/client.go` 中可以调整压缩参数：

```go
const (
    maxImageWidth  = 800  // 最大图片宽度
    webpQuality    = 80   // WebP 质量 (0-100)
)
```

### Markdown 格式优化

工具会自动优化 Markdown 格式：

- **表格优化**：调整表格边框和对齐方式
- **图片处理**：支持本地路径和 Base64 嵌入
- **代码块**：保留语法高亮
- **链接处理**：自动转换飞书链接

### 批量下载优化

- **并发控制**：限制并发下载数量，避免 API 限流
- **进度显示**：实时显示下载进度
- **错误重试**：自动重试失败的下载

---

## 📖 文档

- [产品需求文档 (PRD)](docs/PRD.md) - 产品功能规划和需求分析
- [技术实现文档](docs/技术实现文档.md) - 技术架构和实现细节
- [本地部署使用指南](docs/本地部署使用指南_20251230_v1.1.md) - 详细的部署教程
- [文档导航](docs/README.md) - 所有文档的索引

---

## ❓ 常见问题

### Q: 部署后无法访问？

A: 检查以下几点：
1. 环境变量是否正确配置
2. 防火墙是否放行 8080 端口
3. Render 免费版有冷启动，等待 30-60 秒

### Q: 图片无法下载？

A: 确认以下权限已开通：
- `docs:document.media:download`
- `docx:document:readonly`

### Q: 转换后的文件很大？

A: 使用"嵌入图片到 Markdown"功能，工具会自动压缩图片并转换为 WebP 格式。

### Q: 如何更新到最新版本？

A:
- **Docker**: `docker pull wwwsine/feishu2md`
- **命令行**: 下载最新 Release
- **Render**: 推送代码到 GitHub，自动部署

### Q: 支持飞书旧版文档吗？

A: 旧版文档工具已不再维护，但分支 [v1_support](https://github.com/Wsine/feishu2md/tree/v1_support) 仍可使用。

---

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

**注意**：由于原作者已不再使用飞书文档，项目转为社区维护，欢迎 PR，有能力的维护者会被选择为主协调员。

---

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

---

## 🙏 致谢

- [chyroc/lark](https://github.com/chyroc/lark) - 飞书 Go SDK
- [chyroc/lark_docs_md](https://github.com/chyroc/lark_docs_md) - 飞书文档转换参考
- [Wsine](https://github.com/Wsine) - 原作者

---

## 📮 联系方式

- **GitHub Issues**: https://github.com/Wsine/feishu2md/issues
- **在线演示**: https://feishu2md.onrender.com

---

## 🌟 Star History

如果这个项目对你有帮助，请给它一个 Star ⭐

[![Star History Chart](https://api.star-history.com/svg?repos=Wsine/feishu2md&type=Date)](https://star-history.com/#Wsine/feishu2md&Date)

---

**Made with ❤️ by feishu2md Community**
