VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

.PHONY: build build-all clean test

build:
	CGO_ENABLED=0 go build $(LDFLAGS) -o bin/modbus-debug ./cmd/modbus-debug/

build-all:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/modbus-debug-windows-amd64.exe ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o bin/modbus-debug-darwin-arm64 ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build $(LDFLAGS) -o bin/modbus-debug-darwin-amd64 ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o bin/modbus-debug-linux-amd64 ./cmd/modbus-debug/
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build $(LDFLAGS) -o bin/modbus-debug-linux-arm64 ./cmd/modbus-debug/

test:
	go test ./...

clean:
	rm -rf bin/
