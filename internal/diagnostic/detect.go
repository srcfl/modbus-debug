package diagnostic

import (
	"fmt"
	"net"
	"strings"
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
	Method      string `json:"method,omitempty"` // "fc43", "sunspec", or "serial"
	Vendor      string `json:"vendor,omitempty"` // Raw vendor string from FC43/SunSpec
	Model       string `json:"model,omitempty"`  // Raw model string from FC43/SunSpec
	Error       string `json:"error,omitempty"`
}

// vendorProfileMap maps lowercase vendor/manufacturer substrings to profile names.
var vendorProfileMap = []struct {
	substring   string
	profileName string
}{
	{"sungrow", "sungrow"},
	{"huawei", "huawei"},
	{"solis", "solis"},
	{"fronius", "fronius"},
	{"deye", "deye"},
	{"solaredge", "solaredge"},
	{"solar edge", "solaredge"},
	{"sma", "sma"},
	{"pixii", "pixii"},
	{"eastron", "sdm630"},
	{"sdm630", "sdm630"},
	{"ferroamp", "ferroamp"},
	{"goodwe", "goodwe"},
	{"growatt", "growatt"},
	{"solax", "solax"},
	{"sofar", "sofar"},
	{"victron", "victron"},
	{"foxess", "foxess"},
	{"fox ess", "foxess"},
	{"alpha", "alpha-ess"},
	{"e3dc", "e3dc"},
	{"e3/dc", "e3dc"},
	{"saj", "saj"},
	{"abb", "fronius"}, // ABB/FIMER rebrands of Fronius
	{"fimer", "fronius"},
}

// matchVendorToProfile tries to match a vendor/manufacturer string to a known profile.
func matchVendorToProfile(vendor string) (profileName string, found bool) {
	lower := strings.ToLower(vendor)
	for _, entry := range vendorProfileMap {
		if strings.Contains(lower, entry.substring) {
			return entry.profileName, true
		}
	}
	return "", false
}

// portProfileHints maps well-known Modbus TCP ports to profiles likely found on that port.
// Profiles listed here are tried first during Phase 3 serial probing.
var portProfileHints = map[int][]string{
	8899: {"deye", "solis", "sofar"},
	6607: {"huawei"},
}

// profilesForPort returns all profiles reordered so port-hinted profiles come first.
func profilesForPort(port int) []devices.Profile {
	hints, ok := portProfileHints[port]
	if !ok {
		return devices.AllProfiles()
	}

	hintSet := make(map[string]bool, len(hints))
	for _, h := range hints {
		hintSet[h] = true
	}

	all := devices.AllProfiles()
	reordered := make([]devices.Profile, 0, len(all))

	// Add hinted profiles first (in hint order)
	for _, name := range hints {
		for _, p := range all {
			if p.Name == name {
				reordered = append(reordered, p)
				break
			}
		}
	}
	// Add remaining profiles
	for _, p := range all {
		if !hintSet[p.Name] {
			reordered = append(reordered, p)
		}
	}
	return reordered
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

// DetectInverter tries to identify the inverter brand using a three-phase approach:
//  1. FC 43/14 Read Device Identification (cleanest when supported)
//  2. SunSpec discovery (probe for "SunS" header at well-known addresses)
//  3. Serial number register probing (fallback, tries all profiles)
//
// Tries slave IDs 1, 0, 247 in order using a single TCP connection.
func DetectInverter(host string, port int) DetectionResult {
	client, err := modbus.NewTCPClientWithTimeout(host, port, 1, 500*time.Millisecond)
	if err != nil {
		return DetectionResult{
			Detected: false,
			Error:    fmt.Sprintf("Connection failed: %s", err.Error()),
		}
	}
	defer client.Close()

	slaveIDs := []byte{1, 0, 247}

	for _, slaveID := range slaveIDs {
		client.SetSlaveID(slaveID)

		// Phase 1: Try FC 43/14 Read Device Identification
		if result := tryFC43Detection(client, slaveID); result.Detected {
			return result
		}

		// Phase 2: Try SunSpec discovery
		if result := trySunSpecDetection(client, slaveID); result.Detected {
			return result
		}
	}

	// Phase 3: Fall back to serial number register probing (port-hinted order)
	orderedProfiles := profilesForPort(port)
	for _, slaveID := range slaveIDs {
		client.SetSlaveID(slaveID)
		for _, profile := range orderedProfiles {
			result := tryProfile(client, &profile, slaveID)
			if result.Detected {
				result.Method = "serial"
				return result
			}
		}
	}

	return DetectionResult{
		Detected: false,
		Error:    "Could not detect inverter type. No profile matched on slave IDs 1, 0, or 247.",
	}
}

// tryFC43Detection attempts FC 43/14 Read Device Identification.
func tryFC43Detection(client *modbus.Client, slaveID byte) DetectionResult {
	devID, err := client.ReadDeviceIdentification()
	if err != nil {
		return DetectionResult{Detected: false, SlaveID: slaveID}
	}

	// Try to match vendor to a known profile
	vendorStr := devID.VendorName
	if vendorStr == "" {
		vendorStr = devID.ProductName
	}
	if vendorStr == "" {
		return DetectionResult{Detected: false, SlaveID: slaveID}
	}

	profileName, found := matchVendorToProfile(vendorStr)
	if !found {
		// We got device info but can't map to a profile - still useful
		return DetectionResult{
			Detected: false,
			SlaveID:  slaveID,
			Method:   "fc43",
			Vendor:   devID.VendorName,
			Model:    devID.ProductCode,
		}
	}

	// Find the display name
	displayName := profileName
	for _, p := range devices.AllProfiles() {
		if p.Name == profileName {
			displayName = p.DisplayName
			break
		}
	}

	return DetectionResult{
		Detected:    true,
		ProfileName: profileName,
		DisplayName: displayName,
		SlaveID:     slaveID,
		Method:      "fc43",
		Vendor:      devID.VendorName,
		Model:       devID.ProductCode,
		Serial:      "", // FC43 doesn't always provide serial
	}
}

// trySunSpecDetection attempts SunSpec common model discovery.
func trySunSpecDetection(client *modbus.Client, slaveID byte) DetectionResult {
	info, err := client.DiscoverSunSpec()
	if err != nil {
		return DetectionResult{Detected: false, SlaveID: slaveID}
	}

	// Try to match manufacturer to a known profile
	profileName, found := matchVendorToProfile(info.Manufacturer)
	if !found {
		// Try model string as fallback
		profileName, found = matchVendorToProfile(info.Model)
	}
	if !found {
		return DetectionResult{
			Detected: false,
			SlaveID:  slaveID,
			Method:   "sunspec",
			Vendor:   info.Manufacturer,
			Model:    info.Model,
		}
	}

	displayName := profileName
	for _, p := range devices.AllProfiles() {
		if p.Name == profileName {
			displayName = p.DisplayName
			break
		}
	}

	return DetectionResult{
		Detected:    true,
		ProfileName: profileName,
		DisplayName: displayName,
		Serial:      info.Serial,
		SlaveID:     slaveID,
		Method:      "sunspec",
		Vendor:      info.Manufacturer,
		Model:       info.Model,
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
func DetectAllProfiles(host string, port int, slaveID byte, timeout time.Duration) ([]ProfileResult, error) {
	client, err := modbus.NewTCPClientWithTimeout(host, port, slaveID, timeout)
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

	// Phase 1: Try FC 43/14
	if result := tryFC43Detection(client, slaveID); result.Detected {
		return result
	}

	// Phase 2: Try SunSpec
	if result := trySunSpecDetection(client, slaveID); result.Detected {
		return result
	}

	// Phase 3: Serial number probing (port-hinted order)
	for _, profile := range profilesForPort(port) {
		result := tryProfile(client, &profile, slaveID)
		if result.Detected {
			result.Method = "serial"
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
