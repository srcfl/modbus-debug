package diagnostic

import (
	"fmt"
	"strings"
	"time"
)

// ReportData contains all data needed to generate a diagnostic report.
type ReportData struct {
	Version   string
	Host      string
	Port      int
	TCP       *TCPResult
	Detection *DetectionResult
	Read      *ReadResult
}

// GenerateReport creates a text diagnostic report formatted for Discord.
func GenerateReport(data *ReportData) string {
	var b strings.Builder

	b.WriteString("=== Sourceful Modbus Diagnostic Report ===\n")
	b.WriteString(fmt.Sprintf("Date: %s\n", time.Now().Format("2006-01-02 15:04")))
	b.WriteString(fmt.Sprintf("Tool: %s\n", data.Version))
	b.WriteString("\n")

	// Network section
	b.WriteString("[NETWORK]\n")
	tcpStatus := "FAIL"
	tcpDuration := "N/A"
	if data.TCP != nil {
		tcpDuration = data.TCP.Duration
		if data.TCP.Success {
			tcpStatus = "OK"
		}
	}
	slaveID := byte(0)
	if data.Detection != nil {
		slaveID = data.Detection.SlaveID
	} else if data.Read != nil {
		slaveID = data.Read.SlaveID
	}
	b.WriteString(fmt.Sprintf("  Target: %s:%d  |  TCP: %s (%s)  |  Slave ID: %d\n",
		data.Host, data.Port, tcpStatus, tcpDuration, slaveID))
	b.WriteString("\n")

	// Detection section
	b.WriteString("[DETECTION]\n")
	if data.Detection != nil && data.Detection.Detected {
		b.WriteString(fmt.Sprintf("  Brand: %s  |  Serial: %s\n",
			data.Detection.DisplayName, data.Detection.Serial))
	} else {
		errMsg := "Unknown device"
		if data.Detection != nil && data.Detection.Error != "" {
			errMsg = data.Detection.Error
		}
		b.WriteString(fmt.Sprintf("  %s\n", errMsg))
	}
	b.WriteString("\n")

	// Registers section
	if data.Read == nil {
		b.WriteString("[REGISTERS]\n")
		b.WriteString("  No register data (detection failed or skipped).\n")
		b.WriteString("\n")
		b.WriteString("[ISSUES]\n")
		b.WriteString("  Device not identified. Could not read registers without a matching profile.\n")
		b.WriteString("  Please share this report with Sourceful support.\n")
	} else {
		b.WriteString(fmt.Sprintf("[REGISTERS] %d/%d OK\n", data.Read.Success, data.Read.Total))
		var issues []string

		for _, reg := range data.Read.Registers {
			if !reg.OK {
				issues = append(issues, fmt.Sprintf("  FAIL: %s (addr %d) - %s", reg.Name, reg.Address, reg.Error))
				continue
			}
			status := "OK"
			issue := checkValue(reg)
			if issue != "" {
				status = "WARN"
				issues = append(issues, fmt.Sprintf("  WARN: %s - %s", reg.Name, issue))
			}

			// Only show key registers in the summary
			if isKeyRegister(reg.SemanticName) {
				b.WriteString(fmt.Sprintf("  %-20s %s  [%s]\n", reg.Name+":", reg.ValueStr, status))
			}
		}

		b.WriteString("\n")
		b.WriteString("[ISSUES]\n")
		if len(issues) == 0 {
			b.WriteString("  None detected. Modbus TCP is working correctly.\n")
		} else {
			for _, issue := range issues {
				b.WriteString(issue + "\n")
			}
		}
	}

	b.WriteString("=== End Report ===\n")
	return b.String()
}

func isKeyRegister(semanticName string) bool {
	key := map[string]bool{
		"pv_power":        true,
		"battery_power":   true,
		"battery_soc":     true,
		"meter_power":     true,
		"grid_frequency":  true,
		"load_power":      true,
		"serial_number":   true,
		"nominal_power":   true,
		"total_pv_gen":    true,
		"total_import":    true,
		"total_export":    true,
		"total_charge":    true,
		"total_discharge": true,
	}
	return key[semanticName]
}

func checkValue(reg RegisterResult) string {
	switch reg.SemanticName {
	case "grid_frequency":
		if reg.Value < 45 || reg.Value > 55 {
			return fmt.Sprintf("Grid frequency %.2f Hz is outside normal range (45-55 Hz)", reg.Value)
		}
	case "battery_soc":
		if reg.Value < 0 || reg.Value > 100 {
			return fmt.Sprintf("Battery SoC %.1f%% is outside valid range (0-100%%)", reg.Value)
		}
	}
	return ""
}
