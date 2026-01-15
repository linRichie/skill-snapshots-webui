# Skill Snapshots WebUI

> A visual management platform for Claude Code skill snapshots - browse, manage, and version control your AI skills

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Vue](https://img.shields.io/badge/Vue-3.4-brightgreen.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5.3-blue.svg)

English | [简体中文](./README.md)

## ✨ Features

- 📊 **Dashboard** - Visual overview of all skills with statistics
- 🔍 **Search & Filter** - Find skills by name, category, or tags
- 📝 **Detail View** - Read SKILL.md documentation with Markdown rendering
- 📜 **Version History** - Track all snapshot versions with git integration
- ⚙️ **Admin Panel** - Manage skills, export data, configure settings
- 🌙 **Dark Mode** - Toggle between light and dark themes
- 📱 **Responsive** - Works on desktop, tablet, and mobile devices
- 🚀 **Fast** - Built with Vite for optimal performance

## 🎯 Use Cases

- **Claude Code Users**: Manage and version your custom skills
- **Teams**: Share and collaborate on skill definitions
- **Open Source**: Publish your skill collections
- **Backup**: Archive different versions of your skills

## 📸 Screenshots

### Skill List
![Skill List](docs/images/skill-list.png)

### Skill Detail
![Skill Detail](docs/images/skill-detail.png)

### Version History
![Version History](docs/images/version-history.png)

### Admin Panel
![Admin Panel](docs/images/admin-panel.png)

## 🚀 Quick Start

### Prerequisites

- Node.js 18+
- Go 1.21+ (for backend)
- Git (for version tracking)

### Installation

```bash
# Clone the repository
git clone https://github.com/your-username/skill-snapshots-webui.git
cd skill-snapshots-webui

# Install frontend dependencies
npm install

# Install backend dependencies
cd server && go mod download
```

### Configuration File (Important)

Before first run, create a configuration file to specify your skills directory:

```bash
# Enter server directory
cd server

# Copy configuration template
cp config.example.yaml config.yaml

# Edit config.yaml to set your skills directory path
# skills_dir: "/path/to/your/skills"
```

Configuration file explanation:

```yaml
# Skills directory path (required)
# Supports multiple formats:
skills_dir: "../skill-snapshots"           # Relative path
skills_dir: "/absolute/path/to/skills"     # Absolute path
skills_dir: "~/Documents/skills"           # Tilde expansion
skills_dir: "${HOME}/Documents/skills"     # Environment variable ${VAR}
skills_dir: "$HOME/Documents/skills"       # Environment variable $VAR

# Skill category mapping (optional)
categories:
  pdf-text-extraction: "Document processing"
  skill-creator: "Development tools"

# Skill descriptions (optional)
descriptions:
  skill-creator: "Create new skills"

# Server configuration (optional)
server:
  port: 8000        # API service port
  mode: "release"   # Gin mode (debug | release)
```

### Development

```bash
# Start frontend dev server (http://localhost:3000)
npm run dev

# Start backend API server (http://localhost:8000)
cd server && go run main.go
```

### Production Build

```bash
# Build frontend
npm run build

# Build backend
cd server && go build -o skill-snapshots-api main.go

# Run production server
./server/skill-snapshots-api
```

### Docker Deployment

```bash
# Build and start with Docker Compose
docker-compose up -d

# Access at http://localhost:8000
```

## 📁 Project Structure

```
skill-snapshots-webui/
├── src/
│   ├── api/           # API client
│   ├── assets/        # Static assets (styles, images)
│   ├── components/    # Vue components
│   ├── router/        # Vue Router configuration
│   ├── stores/        # Pinia state management
│   ├── types/         # TypeScript definitions
│   ├── views/         # Page components
│   ├── App.vue        # Root component
│   └── main.ts        # Entry point
├── server/            # Go backend API
│   ├── main.go        # Server entry point
│   └── go.mod         # Go dependencies
├── docs/              # Documentation
├── public/            # Public assets
├── scripts/           # Utility scripts
├── .env.example       # Environment variables template
├── docker-compose.yml # Docker Compose configuration
├── Dockerfile         # Docker build configuration
├── Makefile           # Command shortcuts
└── vite.config.ts     # Vite configuration
```

## ⚙️ Configuration

### Environment Variables

Create a `.env.local` file:

```bash
# API Base URL (default: http://localhost:8000)
VITE_API_BASE_URL=http://localhost:8000

# Application Title
VITE_APP_TITLE=Skill Snapshots

# Environment
NODE_ENV=development
```

### Backend Configuration

The backend reads configuration from `config.yaml` file (recommended):

```bash
# Create configuration file in server directory
cd server
cp config.example.yaml config.yaml

# Edit config.yaml to set skills directory path
```

Configuration parameters:

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `skills_dir` | string | ✅ | Root directory of skills repository |
| `categories` | map | ❌ | Skill category mapping |
| `descriptions` | map | ❌ | Skill description information |
| `server.port` | string | ❌ | API service port (default: 8000) |
| `server.mode` | string | ❌ | Gin mode (debug/release) |

## 🔌 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/skills` | List all skills |
| GET | `/api/skills/:name` | Get skill details |
| GET | `/api/skills/:name/versions` | Get skill versions |
| GET | `/api/versions` | Get all versions |
| GET | `/api/categories` | Get categories |
| GET | `/api/system/info` | System information |
| GET | `/api/system/stats` | Statistics |
| GET | `/api/config` | Current configuration |

## 🛠️ Development

### Adding New Pages

1. Create component in `src/views/`
2. Add route in `src/router/index.ts`
3. Add navigation link in `src/components/AppHeader.vue`

### Modifying Styles

- Global styles: `src/assets/styles/main.css`
- Tailwind config: `tailwind.config.js`
- Component styles: Use Tailwind utility classes

### Building Backend

The Go backend uses git to read version tags. Ensure your skills repository has tags following the pattern:

```
<skill-name>/v1
<skill-name>/v2
...
```

## 📦 Deployment

### Vercel / Netlify

1. Build the frontend: `npm run build`
2. Deploy the `dist` folder
3. Configure API proxy to your backend

### Docker

```bash
docker build -t skill-snapshots-webui .
docker run -p 8000:8000 skill-snapshots-webui
```

### Kubernetes

See `docs/k8s.md` for Kubernetes deployment guide.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Claude Code](https://docs.anthropic.com) by Anthropic
- [Vue.js](https://vuejs.org/)
- [Gin](https://gin-gonic.com/)
- [TailwindCSS](https://tailwindcss.com/)

## 📮 Support

- 📧 Email: support@example.com
- 🐛 Issues: [GitHub Issues](https://github.com/your-username/skill-snapshots-webui/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/your-username/skill-snapshots-webui/discussions)

## 🔗 Links

- [Documentation](https://your-username.github.io/skill-snapshots-webui)
- [Change Log](CHANGELOG.md)
- [Contributing Guide](CONTRIBUTING.md)

---

Made with ❤️ by the community
