VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"
BIN     := bin

.PHONY: build build-all run test clean

## Build for current platform
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(BIN)/modbus-debug ./cmd/modbus-debug/

## Cross-compile for all supported platforms
build-all: clean
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BIN)/modbus-debug-windows-amd64.exe ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o $(BIN)/modbus-debug-darwin-arm64       ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build $(LDFLAGS) -o $(BIN)/modbus-debug-darwin-amd64       ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o $(BIN)/modbus-debug-linux-amd64        ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build $(LDFLAGS) -o $(BIN)/modbus-debug-linux-arm64        ./cmd/modbus-debug/
	@echo "Built $(shell ls -1 $(BIN)/ | wc -l | tr -d ' ') binaries in $(BIN)/"

## Build and run locally
run: build
	./$(BIN)/modbus-debug

## Run tests
test:
	go test ./...

## Remove build artifacts
clean:
	rm -rf $(BIN)/
