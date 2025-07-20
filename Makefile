# VCC Makefile

.PHONY: fmt build build-all clean test vet help

# Format code
fmt:
	go fmt ./...

# Build for current platform
build:
	go build -o vcc .

# Build for all platforms
build-all: clean
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/vcc-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/vcc-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o dist/vcc-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/vcc-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -o dist/vcc-windows-amd64.exe .
	@echo "Built for all platforms in dist/ folder"

# Clean build artifacts
clean:
	rm -f vcc
	rm -rf dist/

# Run tests
test:
	go test ./...

# Run go vet
vet:
	go vet ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build      - Build for current platform"
	@echo "  build-all  - Build for all platforms (Linux, macOS, Windows)"
	@echo "  clean      - Clean build artifacts"
	@echo "  fmt        - Format code"
	@echo "  test       - Run tests"
	@echo "  vet        - Run go vet"
	@echo "  help       - Show this help message"
