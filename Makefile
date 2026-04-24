# Makefile for SERVER CLI tool

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

# Default target is to build the binary
all: build

start:
	@echo "Starting development server..."
	$(GOCMD) run $(MAIN_RUN)

dev:
	@echo "Starting development server with Air..."
	$(AIR) -c .air.toml

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

.PHONY: all build clean build-linux build-windows build-mac build-all install start dev