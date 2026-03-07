# Modbus Debug Tool

## Overview
Downloadable Go binary with embedded web UI for diagnosing Modbus TCP connections to inverters.
Users run the binary, browser opens to `http://127.0.0.1:8765`, and a guided wizard diagnoses their setup.

## Architecture
- Go binary with embedded HTML/JS (Alpine.js) served via `net/http`
- Uses `github.com/goburrow/modbus` (same library as production Zap)
- Read-only tool - never writes to inverter registers
- Listens on `127.0.0.1` only (not exposed to network)

## Key Directories
- `cmd/modbus-debug/` - Entry point + HTTP handlers
- `internal/modbus/` - Register definitions, device profiles, TCP client
- `internal/scanner/` - Network subnet detection + port 502 scanning
- `internal/diagnostic/` - Detection, register reading, report generation
- `cmd/modbus-debug/web/` - Embedded web UI (Alpine.js single-page app, co-located for go:embed)

## Building
```bash
make build        # Build for current platform
make build-all    # Cross-compile for Windows/macOS/Linux
go test ./...     # Run tests
```

## Register Profiles
Copied from `device-simulator` - 10 supported brands:
sungrow, sdm630, solis, huawei, fronius, fronius-smart-meter, sma, deye, solaredge, pixii

## Important Constraints
- Never add write operations (read-only diagnostic tool)
- Keep standalone (no imports from other Sourceful repos)
- CGO_ENABLED=0 for all builds (pure Go, cross-compilable)
