# Skill Snapshots WebUI

> Claude Code 技能快照可视化管理平台 - 浏览、管理和版本控制你的 AI 技能

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Vue](https://img.shields.io/badge/Vue-3.4-brightgreen.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5.3-blue.svg)

[English](./README_EN.md) | 简体中文

## ✨ 功能特性

- 📊 **可视化仪表板** - 展示所有技能的统计信息
- 🔍 **搜索与筛选** - 按名称、分类或标签查找技能
- 📝 **详情查看** - 阅读 SKILL.md 文档，支持 Markdown 渲染
- 📜 **版本历史** - 通过 Git 集成追踪所有快照版本
- ⚙️ **管理后台** - 管理技能、导出数据、配置设置
- 🌙 **深色模式** - 支持浅色/深色主题切换
- 📱 **响应式设计** - 完美支持桌面、平板和移动设备
- 🚀 **极速体验** - 基于 Vite 构建，性能优异

## 🎯 适用场景

- **Claude Code 用户**：管理和版本控制你的自定义技能
- **团队协作**：共享和协作编辑技能定义
- **开源项目**：发布你的技能集合
- **备份归档**：归档不同版本的技能快照

## 📸 功能截图

### 技能列表

![技能列表](https://github.com/linRichie/skill-snapshots-webui/raw/main/public/List%20of%20skills.png)

### 技能详情

![技能详情](https://github.com/linRichie/skill-snapshots-webui/raw/main/public/Skill%20Details.png)

### 版本历史

![版本历史](https://github.com/linRichie/skill-snapshots-webui/raw/main/public/Version%20History.png)

### 管理后台

![管理后台](https://github.com/linRichie/skill-snapshots-webui/raw/main/public/Admin%20Backend.png)

## 🚀 快速开始

### 前置要求

- Node.js 18+
- Go 1.21+（后端服务）
- Git（版本追踪）

### 安装

```bash
# 克隆仓库
git clone https://github.com/linRichie/skill-snapshots-webui.git
cd skill-snapshots-webui

# 安装前端依赖
npm install

# 安装后端依赖
cd server && go mod download
```

### 配置文件（重要）

首次运行前，需要创建配置文件指定技能目录：

```bash
# 进入 server 目录
cd server

# 复制配置模板
cp config.example.yaml config.yaml

# 编辑 config.yaml，设置你的技能目录路径
# skills_dir: "/path/to/your/skills"
```

配置文件说明：

```yaml
# 技能目录路径（必需）
# 支持多种格式：
skills_dir: "../skill-snapshots"           # 相对路径
skills_dir: "/absolute/path/to/skills"     # 绝对路径
skills_dir: "~/Documents/skills"           # 波浪号展开
skills_dir: "${HOME}/Documents/skills"     # 环境变量 ${VAR}
skills_dir: "$HOME/Documents/skills"       # 环境变量 $VAR

# 技能分类映射（可选）
categories:
  pdf-text-extraction: "文档处理"
  skill-creator: "开发工具"

# 技能描述（可选）
descriptions:
  skill-creator: "创建新技能"

# 服务器配置（可选）
server:
  port: 8000        # API 服务端口
  mode: "release"   # Gin 运行模式 (debug | release)
```

### 开发模式

```bash
# 启动前端开发服务器（http://localhost:3000）
npm run dev

# 启动后端 API 服务（http://localhost:8000）
cd server && go run main.go
```

### 生产构建

```bash
# 构建前端
npm run build

# 构建后端
cd server && go build -o skill-snapshots-api main.go

# 运行生产服务器
./server/skill-snapshots-api
```

### Docker 部署

```bash
# 使用 Docker Compose 构建并启动
docker-compose up -d

# 访问 http://localhost:8000
```

## 📁 项目结构

```
skill-snapshots-webui/
├── src/                      # 前端源码
│   ├── api/                   # API 客户端
│   ├── assets/                # 静态资源（样式、图片）
│   ├── components/            # Vue 组件
│   ├── router/                # Vue Router 配置
│   ├── stores/                # Pinia 状态管理
│   ├── types/                 # TypeScript 类型定义
│   ├── views/                 # 页面组件
│   ├── App.vue                # 根组件
│   └── main.ts                # 入口文件
├── server/                    # Go 后端 API
│   ├── main.go                # 服务入口
│   └── go.mod                 # Go 依赖
├── docs/                      # 文档
├── public/                    # 公共资源
├── scripts/                   # 工具脚本
├── .env.example               # 环境变量模板
├── docker-compose.yml         # Docker Compose 配置
├── Dockerfile                 # Docker 构建配置
├── Makefile                   # 命令快捷方式
└── vite.config.ts             # Vite 配置
```

## ⚙️ 配置说明

### 环境变量

创建 `.env.local` 文件：

```bash
# API 基础地址（默认：http://localhost:8000）
VITE_API_BASE_URL=http://localhost:8000

# 应用标题
VITE_APP_TITLE=Skill Snapshots

# 环境
NODE_ENV=development
```

### 后端配置

后端服务从 `config.yaml` 文件读取配置（推荐方式）：

```bash
# 在 server 目录下创建配置文件
cd server
cp config.example.yaml config.yaml

# 编辑 config.yaml 设置技能目录路径
```

配置文件参数说明：

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| `skills_dir` | string | ✅ | 技能仓库根目录路径 |
| `categories` | map | ❌ | 技能分类映射 |
| `descriptions` | map | ❌ | 技能描述信息 |
| `server.port` | string | ❌ | API 服务端口（默认：8000） |
| `server.mode` | string | ❌ | Gin 运行模式（debug/release） |

## 🔌 API 端点

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/health` | 健康检查 |
| GET | `/api/skills` | 获取所有技能 |
| GET | `/api/skills/:name` | 获取技能详情 |
| GET | `/api/skills/:name/versions` | 获取技能版本 |
| GET | `/api/versions` | 获取所有版本 |
| GET | `/api/categories` | 获取分类列表 |
| GET | `/api/system/info` | 系统信息 |
| GET | `/api/system/stats` | 统计数据 |
| GET | `/api/config` | 获取当前配置 |

## 🛠️ 开发指南

### 添加新页面

1. 在 `src/views/` 创建组件
2. 在 `src/router/index.ts` 添加路由
3. 在 `src/components/AppHeader.vue` 添加导航链接

### 修改样式

- 全局样式：`src/assets/styles/main.css`
- Tailwind 配置：`tailwind.config.js`
- 组件样式：使用 Tailwind 工具类

### 后端开发

Go 后端使用 git 读取版本标签。确保你的技能仓库包含以下格式的标签：

```
<技能名称>/v1
<技能名称>/v2
...
```

## 📦 部署

### Vercel / Netlify

1. 构建前端：`npm run build`
2. 部署 `dist` 文件夹
3. 配置 API 代理指向你的后端服务

### Docker

```bash
docker build -t skill-snapshots-webui .
docker run -p 8000:8000 skill-snapshots-webui
```

### Kubernetes

详见 `docs/k8s.md` 获取 Kubernetes 部署指南。

## 🤝 贡献指南

欢迎贡献！请随时提交 Pull Request。

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 开源协议

本项目基于 MIT 协议开源 - 详见 [LICENSE](LICENSE) 文件

## 🙏 致谢

- [Claude Code](https://docs.anthropic.com) by Anthropic
- [Vue.js](https://vuejs.org/)
- [Gin](https://gin-gonic.com/)
- [TailwindCSS](https://tailwindcss.com/)

## 📮 获取帮助

- 🐛 问题反馈：[GitHub Issues](https://github.com/your-username/skill-snapshots-webui/issues)
- 💬 讨论交流：[GitHub Discussions](https://github.com/your-username/skill-snapshots-webui/discussions)

## 🔗 相关链接

- [变更日志](CHANGELOG.md)
- [贡献指南](CONTRIBUTING.md)
- [行为准则](CODE_OF_CONDUCT.md)

---

用 ❤️ 构建
