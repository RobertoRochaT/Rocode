#!/bin/bash

# RoCode Quick Start Script
# This script helps you get started with RoCode quickly

set -e

echo "🚀 RoCode - Quick Start Script"
echo "================================"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Check if Docker is installed
check_docker() {
    if command -v docker &> /dev/null && command -v docker-compose &> /dev/null; then
        return 0
    else
        return 1
    fi
}

# Docker setup
setup_with_docker() {
    print_info "Starting RoCode with Docker..."
    echo ""

    # Check if containers are already running
    if docker-compose ps | grep -q "Up"; then
        print_warning "Containers are already running. Stopping them first..."
        docker-compose down
    fi

    print_info "Building and starting containers..."
    docker-compose up -d --build

    echo ""
    print_success "RoCode is starting up!"
    echo ""
    print_info "Waiting for services to be ready..."
    sleep 5

    echo ""
    echo "================================"
    print_success "🎉 RoCode is running!"
    echo "================================"
    echo ""
    echo "  📱 Frontend:  ${GREEN}http://localhost:3000${NC}"
    echo "  🔌 Backend:   ${GREEN}http://localhost:8080${NC}"
    echo "  🗄️  Database:  ${GREEN}localhost:5432${NC}"
    echo ""
    echo "To view logs:"
    echo "  ${BLUE}docker-compose logs -f${NC}"
    echo ""
    echo "To stop the application:"
    echo "  ${BLUE}docker-compose down${NC}"
    echo ""
}

# Manual setup
setup_manually() {
    print_info "Manual setup instructions:"
    echo ""

    # Check PostgreSQL
    print_info "1. Ensure PostgreSQL is running on localhost:5432"
    echo "   Database name: rocode"
    echo "   Username: postgres"
    echo "   Password: postgres"
    echo ""

    # Backend setup
    print_info "2. Start the backend (in a new terminal):"
    echo "   ${BLUE}cd backend${NC}"
    echo "   ${BLUE}cp .env.example .env${NC}"
    echo "   ${BLUE}go mod download${NC}"
    echo "   ${BLUE}go run .${NC}"
    echo ""

    # Frontend setup
    print_info "3. Start the frontend (in another terminal):"
    echo "   ${BLUE}cd frontend${NC}"
    echo "   ${BLUE}cp .env.example .env${NC}"
    echo "   ${BLUE}npm install${NC}"
    echo "   ${BLUE}npm run dev${NC}"
    echo ""

    print_success "After both servers are running, visit http://localhost:3000"
    echo ""
}

# Main menu
main() {
    echo "Choose your setup method:"
    echo ""
    echo "  1) 🐳 Docker (Recommended - Everything in one command)"
    echo "  2) 🔧 Manual (Requires Go, Node.js, and PostgreSQL)"
    echo "  3) ❌ Exit"
    echo ""

    read -p "Enter your choice (1-3): " choice
    echo ""

    case $choice in
        1)
            if check_docker; then
                setup_with_docker
            else
                print_error "Docker or Docker Compose is not installed!"
                echo ""
                print_info "Please install Docker Desktop from:"
                echo "  https://www.docker.com/products/docker-desktop"
                echo ""
                print_info "Or install manually with:"
                echo "  macOS: ${BLUE}brew install docker docker-compose${NC}"
                echo "  Linux: Follow official docs at docker.com"
                exit 1
            fi
            ;;
        2)
            setup_manually
            ;;
        3)
            print_info "Goodbye!"
            exit 0
            ;;
        *)
            print_error "Invalid choice. Please run the script again."
            exit 1
            ;;
    esac
}

# Run main function
main
