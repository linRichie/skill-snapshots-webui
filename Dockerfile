# Multi-stage Dockerfile for Skill Snapshots WebUI
# Builds both frontend and backend for production deployment

# Stage 1: Build frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /app

# Copy frontend package files
COPY package*.json ./
COPY tsconfig*.json ./
COPY vite.config.ts ./
COPY postcss.config.js ./
COPY tailwind.config.js ./
COPY index.html ./
COPY src ./src

# Install dependencies
RUN npm ci

# Build frontend
RUN npm run build

# Stage 2: Build backend
FROM golang:1.21-alpine AS backend-builder

WORKDIR /app

# Install build tools
RUN apk add --no-cache git gcc musl-dev

# Copy go.mod files
COPY server/go.mod ./server/
WORKDIR /app/server
RUN go mod download

# Copy backend source
COPY server/main.go ./

# Build backend
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/skill-snapshots-api .

# Stage 3: Final image
FROM alpine:3.18

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata git

# Set timezone
ENV TZ=Asia/Shanghai

# Copy from build stages
COPY --from=backend-builder /app/skill-snapshots-api ./
COPY --from=frontend-builder /app/dist ./dist

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /app

USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8000/health || exit 1

# Start server
CMD ["./skill-snapshots-api"]
