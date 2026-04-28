# Makefile for SERVER CLI tool

# Load environment variables
-include .env_dev
export

# Go parameters
GOCMD=$(shell which go 2>/dev/null || echo /usr/local/go/bin/go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOPATH=$(shell $(GOCMD) env GOPATH)
GOBIN=$(GOPATH)/bin
export PATH := $(GOBIN):/usr/local/go/bin:$(PATH)

BINARY_NAME=server-cli
MAIN_RUN = ./cmd/server/main.go
SWAG=$(GOBIN)/swag
AIR=$(GOBIN)/air
DLV=$(GOBIN)/dlv
GOOSE=$(GOBIN)/goose

# Database connection for migration
DB_URL="host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable"
GOOSE_DIR=db/sql
GOOSE_CMD=$(GOOSE) -dir $(GOOSE_DIR) postgres $(DB_URL)

# Default target is to build the binary
all: build

start:
	@echo "Starting development server..."
	$(GOCMD) run $(MAIN_RUN)

dev:
	@echo "Starting development server with Air..."
	$(AIR) -c .air.toml

debug:
	@echo "Starting development server in debug mode with Air..."
	@mkdir -p tmp
	$(AIR) -c .air.toml --build.cmd "go build -gcflags='all=-N -l' -o ./tmp/main $(MAIN_RUN)" --build.full_bin "dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./tmp/main --continue"

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_RUN)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
swag:
	@echo "Generating Swagger documentation..."
	$(SWAG) init -g $(MAIN_RUN) -o ./cmd/swag/docs

# Migration commands
create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -dir $(GOOSE_DIR) create $$name sql

up:
	$(GOOSE_CMD) up

down:
	$(GOOSE_CMD) down

status:
	$(GOOSE_CMD) status

reset:
	$(GOOSE_CMD) reset
	
# Cross-platform builds
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux-amd64 -v $(MAIN_RUN)

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows-amd64.exe -v $(MAIN_RUN)

build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-darwin-amd64 -v $(MAIN_RUN)

# Build for all platforms
build-all: build-linux build-windows build-mac

# Install to GOPATH/bin
install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) -v $(MAIN_RUN)

.PHONY: all build clean build-linux build-windows build-mac build-all install start dev create up down status reset