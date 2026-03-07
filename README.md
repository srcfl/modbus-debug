# Modbus TCP Debug Tool

A self-service diagnostic tool for verifying Modbus TCP connectivity with solar inverters. Download a single binary, run it, and get a clear answer about whether your inverter's Modbus TCP is working.

Built by [Sourceful Energy](https://sourceful.energy) - the same Modbus library used in production Zap gateways.

## Quick Start

1. Download the binary for your platform from [Releases](https://github.com/srcfl/modbus-debug/releases)
2. Run it (see platform-specific instructions below)
3. Your browser opens automatically to `http://127.0.0.1:8765`
4. Follow the guided wizard to diagnose your Modbus TCP setup

## Running on Each Platform

### macOS

macOS will block unsigned binaries by default. To run:

```bash
# Option 1: Remove the quarantine attribute
xattr -d com.apple.quarantine modbus-debug-darwin-arm64

# Option 2: Right-click the file in Finder → Open
# Click "Open" in the dialog that appears

# Then run:
chmod +x modbus-debug-darwin-arm64
./modbus-debug-darwin-arm64
```

**Apple Silicon (M1/M2/M3/M4):** Download `modbus-debug-darwin-arm64`
**Intel Mac:** Download `modbus-debug-darwin-amd64`

### Windows

```
modbus-debug-windows-amd64.exe
```

Windows Firewall will prompt you to allow network access. **Click "Allow"** - the tool needs to make TCP connections to devices on your local network. The tool only listens on `127.0.0.1` (localhost) and is not accessible from other machines.

### Linux

```bash
chmod +x modbus-debug-linux-amd64
./modbus-debug-linux-amd64
```

**Note:** Scanning port 502 does not require root privileges since the tool only connects to (not listens on) privileged ports.

## What It Does

The tool runs a local web server on `127.0.0.1:8765` and provides a browser-based wizard that:

1. **Scans your network** for devices with common Modbus TCP ports open (502, 1502, 8899, 6607)
2. **Tests TCP connectivity** to your inverter
3. **Auto-detects the inverter brand** by reading serial number registers across all known profiles
4. **Reads all registers** for the detected profile (PV, Battery, Meter, Grid)
5. **Generates a diagnostic report** you can copy/paste into Discord or support

There's also a **Manual Mode** for power users to read arbitrary registers with custom data type decoding.

**This is a read-only tool** - it never writes to inverter registers.

## Supported Devices

| Brand | Model | Serial Register |
|-------|-------|-----------------|
| Sungrow | SH Hybrid | 4989 |
| Huawei | SUN2000 | 30015 |
| Solis | Hybrid | 33004 |
| Fronius | GEN24 Plus | 40052 |
| Deye | Hybrid | 3 |
| SolarEdge | StorEdge | 40052 |
| SMA | Sunny Tripower | 30057 |
| Pixii | PowerShaper 2 | 40052 |
| Eastron | SDM630 (meter) | 64512 |
| Fronius | Smart Meter TS 65A-3 | 40053 |

## Common Ports Scanned

| Port | Usage |
|------|-------|
| 502 | Standard Modbus TCP |
| 1502 | Alternative Modbus TCP |
| 8899 | Deye/Solis Wi-Fi dongle |
| 6607 | Huawei SmartDongle |

## Troubleshooting

### "Connection refused" or no devices found
- Verify Modbus TCP is **enabled** in your inverter's settings
- Check that your computer is on the **same network/subnet** as the inverter
- Some inverters only allow **one TCP connection at a time** - disconnect the Zap or any other Modbus client first

### "Connection timed out"
- The inverter's IP may be wrong - check your router's DHCP client list
- A firewall may be blocking outgoing connections on port 502

### Can't detect inverter type
- The inverter may use a non-standard slave ID - try 0 or 247 manually
- Use **Manual Mode** to read the serial number register for your specific brand
- Your inverter model may not be in our profile database yet

## CLI Options

```
Usage: modbus-debug [options]

Options:
  -port int        HTTP server port (default 8765)
  -no-browser      Don't auto-open browser on start
```

## Building from Source

Requires Go 1.22+.

```bash
# Build for current platform
make build

# Cross-compile for all platforms (no CGO required)
make build-all

# Run tests
go test ./...
```

### Build Targets

| File | Platform |
|------|----------|
| `modbus-debug-darwin-arm64` | macOS Apple Silicon |
| `modbus-debug-darwin-amd64` | macOS Intel |
| `modbus-debug-windows-amd64.exe` | Windows x64 |
| `modbus-debug-linux-amd64` | Linux x64 |
| `modbus-debug-linux-arm64` | Linux ARM64 |

All binaries are statically linked (`CGO_ENABLED=0`) with no external dependencies. ~8MB binary size.

## Future Ideas

- "Send to Sourceful" button for unidentified devices to help expand profile support
- SunSpec auto-discovery (read model headers at 40000+)
- Connection keep-alive with periodic register polling
- Export results as JSON/CSV

## License

MIT - see [LICENSE](LICENSE)
