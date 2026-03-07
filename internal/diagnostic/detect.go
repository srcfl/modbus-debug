package diagnostic

import (
	"fmt"
	"net"
	"time"

	"github.com/srcfl/modbus-debug/internal/modbus"
	"github.com/srcfl/modbus-debug/internal/modbus/devices"
)

// TCPResult contains the result of a TCP connectivity test.
type TCPResult struct {
	Success  bool   `json:"success"`
	Duration string `json:"duration"`
	Error    string `json:"error,omitempty"`
}

// DetectionResult contains the result of inverter auto-detection.
type DetectionResult struct {
	Detected    bool   `json:"detected"`
	ProfileName string `json:"profile_name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Serial      string `json:"serial,omitempty"`
	SlaveID     byte   `json:"slave_id"`
	Error       string `json:"error,omitempty"`
}

// TestTCP tests TCP connectivity to the given host:port.
func TestTCP(host string, port int) TCPResult {
	addr := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	duration := time.Since(start)

	if err != nil {
		return TCPResult{
			Success:  false,
			Duration: fmt.Sprintf("%dms", duration.Milliseconds()),
			Error:    err.Error(),
		}
	}
	conn.Close()
	return TCPResult{
		Success:  true,
		Duration: fmt.Sprintf("%dms", duration.Milliseconds()),
	}
}

// DetectInverter tries to identify the inverter brand by reading the serial number register
// for each known profile. Tries slave IDs 1, 0, 247 in order.
func DetectInverter(host string, port int) DetectionResult {
	slaveIDs := []byte{1, 0, 247}

	for _, slaveID := range slaveIDs {
		result := detectWithSlaveID(host, port, slaveID)
		if result.Detected {
			return result
		}
	}

	return DetectionResult{
		Detected: false,
		Error:    "Could not detect inverter type. No profile matched on slave IDs 1, 0, or 247.",
	}
}

// DetectInverterWithSlaveID tries to identify the inverter with a specific slave ID.
func DetectInverterWithSlaveID(host string, port int, slaveID byte) DetectionResult {
	return detectWithSlaveID(host, port, slaveID)
}

// ProfileResult contains the detection result for a single profile.
type ProfileResult struct {
	Profile     string `json:"profile"`
	DisplayName string `json:"display_name"`
	Detected    bool   `json:"detected"`
	Serial      string `json:"serial,omitempty"`
	Error       string `json:"error,omitempty"`
	DurationMs  int64  `json:"duration_ms"`
}

// DetectAllProfiles tries all profiles with a single TCP connection for the given slave ID.
// Returns results for every profile.
func DetectAllProfiles(host string, port int, slaveID byte) ([]ProfileResult, error) {
	client, err := modbus.NewTCPClientWithTimeout(host, port, slaveID, 500*time.Millisecond)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer client.Close()

	var results []ProfileResult
	for _, profile := range devices.AllProfiles() {
		start := time.Now()
		result := tryProfile(client, &profile, slaveID)
		dur := time.Since(start).Milliseconds()

		pr := ProfileResult{
			Profile:     profile.Name,
			DisplayName: profile.DisplayName,
			Detected:    result.Detected,
			Serial:      result.Serial,
			DurationMs:  dur,
		}
		if result.Error != "" {
			pr.Error = result.Error
		}
		results = append(results, pr)
	}
	return results, nil
}

// DetectProfile tries a single profile with a given slave ID.
func DetectProfile(host string, port int, slaveID byte, profileName string) DetectionResult {
	var profile *devices.Profile
	for _, p := range devices.AllProfiles() {
		if p.Name == profileName {
			profile = &p
			break
		}
	}
	if profile == nil {
		return DetectionResult{Detected: false, SlaveID: slaveID, Error: "Unknown profile"}
	}

	client, err := modbus.NewTCPClientWithTimeout(host, port, slaveID, 500*time.Millisecond)
	if err != nil {
		return DetectionResult{
			Detected: false,
			SlaveID:  slaveID,
			Error:    fmt.Sprintf("Connection failed: %s", err.Error()),
		}
	}
	defer client.Close()

	return tryProfile(client, profile, slaveID)
}

func detectWithSlaveID(host string, port int, slaveID byte) DetectionResult {
	client, err := modbus.NewTCPClientWithTimeout(host, port, slaveID, 500*time.Millisecond)
	if err != nil {
		return DetectionResult{
			Detected: false,
			SlaveID:  slaveID,
			Error:    fmt.Sprintf("Connection failed: %s", err.Error()),
		}
	}
	defer client.Close()

	for _, profile := range devices.AllProfiles() {
		result := tryProfile(client, &profile, slaveID)
		if result.Detected {
			return result
		}
	}

	return DetectionResult{
		Detected: false,
		SlaveID:  slaveID,
	}
}

func tryProfile(client *modbus.Client, profile *devices.Profile, slaveID byte) DetectionResult {
	rs := profile.Registers()
	serialDef, ok := rs.Semantic["serial_number"]
	if !ok {
		return DetectionResult{Detected: false, SlaveID: slaveID}
	}

	data, err := client.ReadRegisters(serialDef.Address, uint16(serialDef.Words), serialDef.UseHolding)
	if err != nil {
		return DetectionResult{Detected: false, SlaveID: slaveID, Error: err.Error()}
	}

	regs := modbus.BytesToRegisters(data)
	if len(regs) < serialDef.Words {
		return DetectionResult{Detected: false, SlaveID: slaveID}
	}

	var serial string
	if serialDef.DataType == modbus.STR {
		serial = modbus.DecodeString(regs)
	} else {
		val, err := modbus.DecodeValue(regs, serialDef)
		if err != nil {
			return DetectionResult{Detected: false, SlaveID: slaveID}
		}
		serial = fmt.Sprintf("%.0f", val)
	}

	if serial != "" && serial != "0" {
		return DetectionResult{
			Detected:    true,
			ProfileName: profile.Name,
			DisplayName: profile.DisplayName,
			Serial:      serial,
			SlaveID:     slaveID,
		}
	}

	return DetectionResult{Detected: false, SlaveID: slaveID}
}
