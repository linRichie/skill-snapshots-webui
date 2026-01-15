# Skill Snapshots WebUI - Makefile
# Simplify development and deployment workflows

.PHONY: help install dev build clean test docker-build docker-up docker-down

# Default target
.DEFAULT_GOAL := help

# Variables
NODE_BIN := node_modules/.bin
FRONTEND_DIR := .
BACKEND_DIR := server

# Colors for output
CYAN := \033[0;36m
GREEN := \033[0;32m
YELLOW := \033[1;33m
NC := \033[0m

help: ## Show help information
	@echo "$(CYAN)Skill Snapshots WebUI - Development Commands$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'

install: ## Install all dependencies
	@echo "$(YELLOW)Installing frontend dependencies...$(NC)"
	cd $(FRONTEND_DIR) && npm install
	@echo "$(YELLOW)Installing backend dependencies...$(NC)"
	cd $(BACKEND_DIR) && go mod download
	@echo "$(GREEN)✓ Dependencies installed$(NC)"

dev-frontend: ## Start frontend development server
	@echo "$(CYAN)Starting frontend dev server...$(NC)"
	cd $(FRONTEND_DIR) && npm run dev

dev-backend: ## Start backend development server
	@echo "$(CYAN)Starting backend dev server...$(NC)"
	cd $(BACKEND_DIR) && go run main.go

dev: ## Start both frontend and backend dev servers
	@echo "$(CYAN)Starting development servers...$(NC)"
	@make -j2 dev-frontend dev-backend

build-frontend: ## Build frontend for production
	@echo "$(YELLOW)Building frontend...$(NC)"
	cd $(FRONTEND_DIR) && npm run build
	@echo "$(GREEN)✓ Frontend built$(NC)"

build-backend: ## Build backend for production
	@echo "$(YELLOW)Building backend...$(NC)"
	cd $(BACKEND_DIR) && CGO_ENABLED=0 go build -ldflags="-s -w" -o skill-snapshots-api .
	@echo "$(GREEN)✓ Backend built$(NC)"

build: build-frontend build-backend ## Build frontend and backend

type-check: ## Run TypeScript type checking
	@echo "$(YELLOW)Running type check...$(NC)"
	cd $(FRONTEND_DIR) && npm run type-check

clean: ## Clean build artifacts and dependencies
	@echo "$(YELLOW)Cleaning...$(NC)"
	rm -rf $(FRONTEND_DIR)/dist
	rm -rf $(FRONTEND_DIR)/node_modules
	rm -rf $(BACKEND_DIR)/skill-snapshots-api
	@echo "$(GREEN)✓ Cleaned$(NC)"

docker-build: ## Build Docker image
	@echo "$(YELLOW)Building Docker image...$(NC)"
	docker-compose build

docker-up: ## Start Docker containers
	@echo "$(YELLOW)Starting containers...$(NC)"
	docker-compose up -d
	@echo "$(GREEN)✓ Running at http://localhost:8000$(NC)"

docker-down: ## Stop Docker containers
	@echo "$(YELLOW)Stopping containers...$(NC)"
	docker-compose down

docker-restart: docker-down docker-up ## Restart Docker containers

docker-logs: ## Show Docker logs
	docker-compose logs -f

lint: ## Run linting
	@echo "$(YELLOW)Linting...$(NC)"
	cd $(FRONTEND_DIR) && npm run lint

fmt-backend: ## Format backend code
	@echo "$(YELLOW)Formatting backend...$(NC)"
	cd $(BACKEND_DIR) && go fmt ./...

install-tools: ## Install development tools
	@echo "$(YELLOW)Installing tools...$(NC)"
	npm install -g @vue/cli
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
