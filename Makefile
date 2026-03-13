##
## COREL PROJECT, 2026
## Corel
## File description:
## Makefile
##

NAME = corel
GO = go
BUILD_ENV = CGO_ENABLED=0

all: windows linux macos

linux-amd64: $(SRC)
	@echo "Building $(NAME) for Linux - AMD64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=linux GOARCH=amd64 $(GO) build -o bin/$(NAME)-linux-amd64 .
	@echo "Build completed: bin/$(NAME)-linux-amd64"

linux-arm64: $(SRC)
	@echo "Building $(NAME) for Linux - ARM64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=linux GOARCH=arm64 $(GO) build -o bin/$(NAME)-linux-arm64 .
	@echo "Build completed: bin/$(NAME)-linux-arm64"

linux: linux-amd64 linux-arm64

windows-amd64: $(SRC)
	@echo "Building $(NAME) for Windows - AMD64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=windows GOARCH=amd64 $(GO) build -o bin/$(NAME)-windows-amd64.exe .
	@echo "Build completed: bin/$(NAME)-windows-amd64.exe"

windows-arm64: $(SRC)
	@echo "Building $(NAME) for Windows - ARM64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=windows GOARCH=arm64 $(GO) build -o bin/$(NAME)-windows-arm64.exe .
	@echo "Build completed: bin/$(NAME)-windows-arm64.exe"

windows: windows-amd64 windows-arm64

macos-amd64: $(SRC)
	@echo "Building $(NAME) for macOS - AMD64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=darwin GOARCH=amd64 $(GO) build -o bin/$(NAME)-macos-amd64 .
	@echo "Build completed: bin/$(NAME)-macos-amd64"

macos-arm64: $(SRC)
	@echo "Building $(NAME) for macOS - ARM64..."
	@mkdir -p bin
	@$(BUILD_ENV) GOOS=darwin GOARCH=arm64 $(GO) build -o bin/$(NAME)-macos-arm64 .
	@echo "Build completed: bin/$(NAME)-macos-arm64"

macos: macos-amd64 macos-arm64

clean: bin/$(NAME)
	rm -f bin/$(NAME)
	rm -rf bin

re: clean all

.PHONY: all clean re