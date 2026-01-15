# Contributing to Skill Snapshots WebUI

Thank you for your interest in contributing to Skill Snapshots WebUI! This document provides guidelines and instructions for contributing.

## 🤝 How to Contribute

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Title**: Clear and descriptive
- **Description**: Detailed explanation of the problem
- **Reproduction Steps**: Steps to reproduce the issue
- **Expected Behavior**: What you expected to happen
- **Actual Behavior**: What actually happened
- **Environment**: OS, browser, Node.js version
- **Screenshots**: If applicable

### Suggesting Enhancements

Enhancement suggestions are welcome! Please include:

- **Use Case**: What problem would this solve?
- **Proposed Solution**: How do you envision it working?
- **Alternatives**: What alternatives have you considered?
- **Impact**: Who would benefit from this feature?

## 🛠️ Development Setup

### Fork and Clone

```bash
# Fork the repository on GitHub
# Clone your fork
git clone https://github.com/YOUR_USERNAME/skill-snapshots-webui.git
cd skill-snapshots-webui

# Add upstream remote
git remote add upstream https://github.com/ORIGINAL_OWNER/skill-snapshots-webui.git
```

### Install Dependencies

```bash
npm install
cd server && go mod download
```

### Start Development

```bash
# Terminal 1: Frontend
npm run dev

# Terminal 2: Backend
cd server && go run main.go
```

## 📝 Coding Standards

### Frontend (Vue/TypeScript)

- Use TypeScript for all new code
- Follow Vue 3 Composition API style
- Use `<script setup>` syntax for components
- Follow the existing code structure
- Add comments for complex logic

```vue
<script setup lang="ts">
import { ref, computed } from 'vue'

// Use descriptive variable names
const isLoading = ref(false)
const filteredItems = computed(() => {
  // ...
})
</script>
```

### Backend (Go)

- Follow Go conventions and idioms
- Add meaningful comments
- Handle errors appropriately
- Write testable code

```go
// parseVersions parses version information from git tags
func parseVersions(skillName string) []Version {
    // ...
}
```

### Commit Messages

Follow conventional commit format:

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

Examples:

```
feat(api): add filter by category endpoint
fix(ui): correct version number display
docs(readme): update installation instructions
```

## 🧪 Testing

Before submitting, ensure:

- [ ] Code compiles without errors (`npm run type-check`)
- [ ] No ESLint warnings (`npm run lint`)
- [ ] Manual testing of your changes
- [ ] Documentation is updated

## 📤 Submitting Changes

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

### 2. Make Changes

- Write clear, focused commits
- Follow coding standards
- Update documentation

### 3. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 4. Create Pull Request

- Provide a clear description of changes
- Reference related issues
- Include screenshots for UI changes
- Ensure CI checks pass

## 🎨 Style Guide

### File Naming

- Components: `PascalCase.vue` (e.g., `SkillCard.vue`)
- Utilities: `kebab-case.ts` (e.g., `date-utils.ts`)
- Styles: `kebab-case.css`

### Component Structure

```vue
<template>
  <!-- Template -->
</template>

<script setup lang="ts">
// Imports
// Props
// Emits
// State
// Computed
// Methods
// Lifecycle
</script>

<style scoped>
/* Styles */
</style>
```

### CSS/Styles

- Use Tailwind utility classes when possible
- Add custom styles in `<style scoped>`
- Follow BEM for complex CSS

## 📚 Documentation

- Update README for user-facing changes
- Add comments for complex code
- Update API documentation for backend changes
- Add examples for new features

## 🗺️ Project Roadmap

See [ROADMAP.md](ROADMAP.md) for planned features and direction.

## 💬 Getting Help

- Open a discussion for questions
- Check existing issues first
- Join our community chat (coming soon)

## ⭐ Recognition

Contributors will be recognized in:
- CONTRIBUTORS.md file
- Release notes
- Project documentation

Thank you for contributing! 🎉
