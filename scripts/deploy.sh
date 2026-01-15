#!/bin/bash
# Deployment script for Skill Snapshots WebUI

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  Skill Snapshots WebUI Deployment${NC}"
echo -e "${GREEN}========================================${NC}"

# Check Docker
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker not found${NC}"
    exit 1
fi

# Parse command
COMMAND=${1:-"help"}

case "$COMMAND" in
    build)
        echo -e "${YELLOW}Building Docker image...${NC}"
        docker-compose build
        echo -e "${GREEN}✓ Build complete${NC}"
        ;;

    start|up)
        echo -e "${YELLOW}Starting services...${NC}"
        docker-compose up -d
        echo -e "${GREEN}✓ Services started${NC}"
        echo -e "Access: ${GREEN}http://localhost:8000${NC}"
        ;;

    stop|down)
        echo -e "${YELLOW}Stopping services...${NC}"
        docker-compose down
        echo -e "${GREEN}✓ Services stopped${NC}"
        ;;

    restart)
        echo -e "${YELLOW}Restarting services...${NC}"
        docker-compose restart
        echo -e "${GREEN}✓ Services restarted${NC}"
        ;;

    logs)
        echo -e "${YELLOW}Showing logs...${NC}"
        docker-compose logs -f
        ;;

    status|ps)
        echo -e "${YELLOW}Service status:${NC}"
        docker-compose ps
        ;;

    dev)
        echo -e "${CYAN}Development mode:${NC}"
        echo "  Frontend: npm run dev  (http://localhost:3000)"
        echo "  Backend:  cd server && go run main.go  (http://localhost:8000)"
        ;;

    clean)
        echo -e "${YELLOW}Cleaning containers and images...${NC}"
        docker-compose down -v --rmi all
        echo -e "${GREEN}✓ Cleaned${NC}"
        ;;

    install)
        echo -e "${YELLOW}Installing dependencies...${NC}"
        npm install
        cd server && go mod download
        echo -e "${GREEN}✓ Installed${NC}"
        ;;

    *)
        echo -e "${RED}Unknown command: $COMMAND${NC}"
        echo "Usage: ./scripts/deploy.sh [command]"
        echo ""
        echo "Commands:"
        echo "  build    - Build Docker image"
        echo "  start    - Start services"
        echo "  stop     - Stop services"
        echo "  restart  - Restart services"
        echo "  logs     - Show logs"
        echo "  status   - Show service status"
        echo "  dev      - Show dev mode instructions"
        echo "  clean    - Clean containers and images"
        echo "  install  - Install dependencies"
        echo "  help     - Show this help"
        exit 1
        ;;
esac
