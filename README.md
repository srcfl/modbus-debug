# Modbus TCP Debug Tool

A free, open-source diagnostic tool for testing Modbus TCP connections to solar inverters and energy meters. Download a single binary, run it, and get a clear answer about whether your device's Modbus TCP is working — no installation required.

Built by [Sourceful Energy](https://sourceful.energy) — the makers of the [Zap](https://sourceful.energy), a gateway for local energy coordination. This tool speaks the same Modbus TCP protocol as the Zap, making it the perfect companion for verifying your inverter setup before or during deployment.

**This is a read-only tool** — it never writes to your inverter registers. Safe to use at any time.

## Download

Grab the latest binary for your platform from **[Releases](https://github.com/srcfl/modbus-debug/releases)**.

| Platform | File | Notes |
|----------|------|-------|
| **macOS (Apple Silicon)** | `modbus-debug-darwin-arm64` | M1/M2/M3/M4 Macs |
| **macOS (Intel)** | `modbus-debug-darwin-amd64` | Older Macs |
| **Windows** | `modbus-debug-windows-amd64.exe` | 64-bit Windows |
| **Linux** | `modbus-debug-linux-amd64` | 64-bit x86 |
| **Linux (ARM)** | `modbus-debug-linux-arm64` | Raspberry Pi 4+, etc. |

## How to Run

### macOS

macOS blocks unsigned binaries by default. Open Terminal and run:

```bash
# Remove the quarantine flag (one-time)
xattr -d com.apple.quarantine modbus-debug-darwin-arm64

# Make it executable and run
chmod +x modbus-debug-darwin-arm64
./modbus-debug-darwin-arm64
```

Alternatively: right-click the file in Finder → Open → click "Open" in the dialog.

### Windows

Double-click `modbus-debug-windows-amd64.exe` or run it from Command Prompt.

Windows Firewall will ask to allow network access — **click "Allow"**. The tool needs to reach devices on your local network. It only listens on `127.0.0.1` (your machine only), so it's not accessible from other computers.

### Linux

```bash
chmod +x modbus-debug-linux-amd64
./modbus-debug-linux-amd64
```

No root required — the tool connects *to* port 502, it doesn't listen on it.

## Using the Tool

When you run the binary, your browser automatically opens to `http://127.0.0.1:8765`. You'll see a guided wizard with two modes:

### Guided Wizard (recommended)

Step-by-step diagnostic — perfect for troubleshooting:

1. **Network Scan** — Scans your local network for devices with Modbus TCP ports open (502, 1502, 8899, 6607). Pick a discovered device or enter an IP manually.

2. **TCP Test** — Verifies that a TCP connection to the device succeeds and measures response time.

3. **Inverter Detection** — Reads identification registers across 20 known inverter profiles to determine the device brand and model. Tries slave IDs 1, 0, and 247 automatically.

4. **Register Read** — Reads all registers for the detected profile (PV production, battery state, grid power, meter values) and decodes them with the correct data types and scale factors.

5. **Diagnostic Report** — Generates a text report you can copy/paste into Discord, email, or a support ticket:

```
=== Sourceful Modbus Diagnostic Report ===
Date: 2026-03-07 15:30
Tool: v0.1.0

[NETWORK]
  Target: 192.168.1.50:502  |  TCP: OK (12ms)  |  Slave ID: 1

[DETECTION]
  Brand: Sungrow SH Hybrid  |  Serial: SG12345678

[REGISTERS] 17/17 OK
  PV Power:       3250 W      [OK]
  Battery SoC:    67.3 %      [OK]
  Grid Frequency: 50.02 Hz    [OK]
  ...
=== End Report ===
```

### Manual Mode (power users)

Read any Modbus register directly:

- Connect to any IP:port with any slave ID
- Read holding or input registers at any address
- Choose register count and data type (U16, I16, U32, I32, F32, STR)
- See raw values and decoded results
- Useful for unsupported brands or debugging specific registers

## Supported Devices

20 inverter and meter profiles with full register maps:

| Brand | Model | Detection Register | Default Port |
|-------|-------|--------------------|--------------|
| Sungrow | SH Hybrid | 4989 (holding) | 502 |
| Huawei | SUN2000 | 30015 (holding) | 6607 |
| Solis | Hybrid | 33004 (input) | 8899 |
| Fronius | GEN24 Plus | 40052 (holding, SunSpec) | 502 |
| Deye | Hybrid | 3 (holding) | 8899 |
| SolarEdge | StorEdge | 40052 (holding, SunSpec) | 502 |
| SMA | Sunny Tripower | 30057 (holding) | 502 |
| Pixii | PowerShaper 2 | 40052 (holding, SunSpec) | 502 |
| Eastron | SDM630 (meter) | 64512 (input) | 502 |
| Fronius | Smart Meter TS 65A-3 | 40053 (holding, SunSpec) | 502 |
| Ferroamp | EnergyHub | 2000 (input) | 502 |
| GoodWe | ET/EH/BH/BT Hybrid | 35003 (holding) | 502 |
| Growatt | Hybrid | 23 (holding) | 502 |
| SolaX | X1/X3 | 768 (holding) | 502 |
| SofarSolar | HYD G3 | 1093 (holding) | 502 |
| Victron | Energy (GX) | 800 (input) | 502 |
| FoxESS | H3 | 30000 (holding) | 502 |
| Alpha ESS | SMILE | 1610 (holding) | 502 |
| E3DC | S10/E | 40000 (holding) | 502 |
| SAJ | H2 | 36611 (input) | 502 |

## Scanned Ports

| Port | Devices |
|------|---------|
| 502 | Standard Modbus TCP (most inverters) |
| 1502 | Kostal, some alternative setups |
| 8899 | Deye, Solis, Sofar Wi-Fi dongles |
| 6607 | Huawei SmartDongle |

## Troubleshooting

### No devices found during network scan
- Make sure your computer is on the **same network** as the inverter
- Verify Modbus TCP is **enabled** in your inverter's settings menu
- Some Wi-Fi dongles take 30-60 seconds to boot — wait and retry

### "Connection refused"
- The inverter's Modbus TCP service may not be enabled
- Some inverters allow only **one TCP connection at a time** — disconnect the Zap or any other Modbus client first, then retry

### "Connection timed out"
- Double-check the IP address (look in your router's DHCP client list)
- A firewall may be blocking connections to port 502
- Wi-Fi dongles (Deye, Solis) can be slow — the tool automatically adjusts timeouts for slow connections

### Can't detect inverter type
- Your inverter may use a non-standard slave ID — the wizard tries 1, 0, and 247 automatically
- Use **Manual Mode** to try other slave IDs or read specific registers
- If your brand isn't listed, the tool still generates a report with what it found — share it with support

### Register values look wrong
- Check that the correct profile is selected — a wrong profile will decode registers incorrectly
- Some inverters require specific slave IDs (e.g., Victron uses 100, Alpha ESS uses 85)
- Values of 0 may mean the inverter is in standby/night mode

## CLI Options

```
Usage: modbus-debug [options]

Options:
  -port int        HTTP server port (default 8765)
  -no-browser      Don't auto-open browser on start
```

Example: run on a different port without opening the browser:
```bash
./modbus-debug -port 9090 -no-browser
```

## Building from Source

Requires Go 1.22+.

```bash
# Clone the repo
git clone https://github.com/srcfl/modbus-debug.git
cd modbus-debug

# Build for your current platform
make build

# Cross-compile all platforms
make build-all

# Run tests
make test

# Run directly
make run
```

Output binaries are in the `bin/` directory. All builds are statically linked (`CGO_ENABLED=0`) with no external dependencies (~8 MB).

## Contributing

Found a bug or want to add support for a new inverter brand? Open an issue or submit a PR.

To add a new device profile, create a file in `internal/modbus/devices/` following the pattern of existing profiles (e.g., `sungrow.go`). Each profile defines register addresses, data types, endianness, and scale factors.

## License

MIT — see [LICENSE](LICENSE)
