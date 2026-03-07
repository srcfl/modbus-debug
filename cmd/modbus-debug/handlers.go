package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/srcfl/modbus-debug/internal/diagnostic"
	"github.com/srcfl/modbus-debug/internal/modbus/devices"
	"github.com/srcfl/modbus-debug/internal/scanner"
)

func handleProfiles(w http.ResponseWriter, r *http.Request) {
	type profileInfo struct {
		Name           string `json:"name"`
		DisplayName    string `json:"display_name"`
		SerialAddress  int    `json:"serial_address,omitempty"`
		SerialHolding  bool   `json:"serial_holding,omitempty"`
	}

	var profiles []profileInfo
	for _, p := range devices.AllProfiles() {
		info := profileInfo{
			Name:        p.Name,
			DisplayName: p.DisplayName,
		}
		rs := p.Registers()
		if sd, ok := rs.Semantic["serial_number"]; ok {
			info.SerialAddress = int(sd.Address)
			info.SerialHolding = sd.UseHolding
		}
		profiles = append(profiles, info)
	}
	writeJSON(w, profiles)
}

func handleNetworkInterfaces(w http.ResponseWriter, r *http.Request) {
	ifaces, err := scanner.GetInterfaces()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, ifaces)
}

func handleNetworkScan(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Subnet  string `json:"subnet"`
		Port    int    `json:"port"`
		Timeout int    `json:"timeout_ms"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Subnet == "" {
		writeError(w, http.StatusBadRequest, "subnet is required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}
	timeout := 500 * time.Millisecond
	if req.Timeout > 0 {
		timeout = time.Duration(req.Timeout) * time.Millisecond
	}

	hosts, err := scanner.ScanSubnet(req.Subnet, req.Port, timeout)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, hosts)
}

func handleScanPorts(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, scanner.UniqueDefaultPorts())
}

func handleDiagnoseTCP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" {
		writeError(w, http.StatusBadRequest, "host is required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}

	result := diagnostic.TestTCP(req.Host, req.Port)
	writeJSON(w, result)
}

func handleDiagnoseDetect(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		SlaveID *int   `json:"slave_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" {
		writeError(w, http.StatusBadRequest, "host is required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}

	var result diagnostic.DetectionResult
	if req.SlaveID != nil {
		result = diagnostic.DetectInverterWithSlaveID(req.Host, req.Port, byte(*req.SlaveID))
	} else {
		result = diagnostic.DetectInverter(req.Host, req.Port)
	}
	writeJSON(w, result)
}

func handleDiagnoseDetectOne(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		SlaveID int    `json:"slave_id"`
		Profile string `json:"profile"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" || req.Profile == "" {
		writeError(w, http.StatusBadRequest, "host and profile are required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}

	result := diagnostic.DetectProfile(req.Host, req.Port, byte(req.SlaveID), req.Profile)
	writeJSON(w, result)
}

func handleDiagnoseDetectBatch(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		SlaveID int    `json:"slave_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" {
		writeError(w, http.StatusBadRequest, "host is required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}

	results, err := diagnostic.DetectAllProfiles(req.Host, req.Port, byte(req.SlaveID))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, results)
}

func handleDiagnoseRead(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Profile string `json:"profile"`
		SlaveID int    `json:"slave_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" || req.Profile == "" {
		writeError(w, http.StatusBadRequest, "host and profile are required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}

	result, err := diagnostic.ReadAllRegisters(req.Host, req.Port, req.Profile, byte(req.SlaveID))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, result)
}

// reportData stores the latest diagnostic state for report generation.
var reportData *diagnostic.ReportData

func handleDiagnoseReport(w http.ResponseWriter, r *http.Request) {
	if reportData == nil {
		writeError(w, http.StatusBadRequest, "No diagnostic data available. Run the diagnostic first.")
		return
	}
	report := diagnostic.GenerateReport(reportData)
	writeJSON(w, map[string]string{"report": report})
}

func handleRawRead(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		SlaveID  int    `json:"slave_id"`
		Address  int    `json:"address"`
		Count    int    `json:"count"`
		Holding  bool   `json:"holding"`
		DataType string `json:"data_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Host == "" {
		writeError(w, http.StatusBadRequest, "host is required")
		return
	}
	if req.Port == 0 {
		req.Port = 502
	}
	if req.Count == 0 {
		req.Count = 1
	}
	if req.DataType == "" {
		req.DataType = "U16"
	}

	result, err := diagnostic.ReadRaw(
		req.Host, req.Port, byte(req.SlaveID),
		uint16(req.Address), uint16(req.Count),
		req.Holding, req.DataType,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, result)
}

func handleUpdateReport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Host      string                    `json:"host"`
		Port      int                       `json:"port"`
		TCP       *diagnostic.TCPResult     `json:"tcp"`
		Detection *diagnostic.DetectionResult `json:"detection"`
		Read      *diagnostic.ReadResult    `json:"read"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	port := req.Port
	if port == 0 {
		port = 502
	}
	reportData = &diagnostic.ReportData{
		Version: version,
		Host:    req.Host,
		Port:    port,
		TCP:       req.TCP,
		Detection: req.Detection,
		Read:      req.Read,
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

