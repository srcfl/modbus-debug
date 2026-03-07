package diagnostic

import (
	"fmt"

	"github.com/srcfl/modbus-debug/internal/modbus"
	"github.com/srcfl/modbus-debug/internal/modbus/devices"
)

// RegisterResult contains the result of reading a single register.
type RegisterResult struct {
	Address      uint16   `json:"address"`
	Name         string   `json:"name"`
	SemanticName string   `json:"semantic_name,omitempty"`
	Category     string   `json:"category"`
	Value        float64  `json:"value"`
	ValueStr     string   `json:"value_str"`
	Unit         string   `json:"unit"`
	Raw          []uint16 `json:"raw"`
	DataType     string   `json:"data_type"`
	UseHolding   bool     `json:"use_holding"`
	OK           bool     `json:"ok"`
	Error        string   `json:"error,omitempty"`
}

// ReadResult contains results of reading all registers for a profile.
type ReadResult struct {
	Profile   string           `json:"profile"`
	SlaveID   byte             `json:"slave_id"`
	Registers []RegisterResult `json:"registers"`
	Total     int              `json:"total"`
	Success   int              `json:"success"`
	Failed    int              `json:"failed"`
}

// ReadAllRegisters reads all registers for the given profile.
func ReadAllRegisters(host string, port int, profileName string, slaveID byte) (*ReadResult, error) {
	rs := devices.GetRegisterSet(profileName)
	if rs == nil {
		return nil, fmt.Errorf("unknown profile: %s", profileName)
	}

	client, err := modbus.NewTCPClient(host, port, slaveID)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer client.Close()

	result := &ReadResult{
		Profile: profileName,
		SlaveID: slaveID,
		Total:   len(rs.Definitions),
	}

	for _, def := range rs.Definitions {
		rr := readSingleRegister(client, &def)
		result.Registers = append(result.Registers, rr)
		if rr.OK {
			result.Success++
		} else {
			result.Failed++
		}
	}

	return result, nil
}

func readSingleRegister(client *modbus.Client, def *modbus.RegisterDef) RegisterResult {
	rr := RegisterResult{
		Address:      def.Address,
		Name:         def.Name,
		SemanticName: def.SemanticName,
		Category:     def.Category,
		Unit:         def.Unit,
		DataType:     def.DataType.String(),
		UseHolding:   def.UseHolding,
	}

	data, err := client.ReadRegisters(def.Address, uint16(def.Words), def.UseHolding)
	if err != nil {
		rr.Error = err.Error()
		return rr
	}

	regs := modbus.BytesToRegisters(data)
	rr.Raw = regs

	if def.DataType == modbus.STR {
		str := modbus.DecodeString(regs)
		rr.ValueStr = str
		rr.OK = true
		return rr
	}

	value, err := modbus.DecodeValue(regs, def)
	if err != nil {
		rr.Error = err.Error()
		return rr
	}

	rr.Value = value
	rr.ValueStr = formatValue(value, def.Unit)
	rr.OK = true
	return rr
}

func formatValue(value float64, unit string) string {
	if unit != "" {
		return fmt.Sprintf("%.2f %s", value, unit)
	}
	return fmt.Sprintf("%.2f", value)
}

// RawReadResult contains the result of reading arbitrary registers.
type RawReadResult struct {
	Address  uint16   `json:"address"`
	Count    uint16   `json:"count"`
	Holding  bool     `json:"holding"`
	Raw      []uint16 `json:"raw"`
	RawHex   []string `json:"raw_hex"`
	Decoded  float64  `json:"decoded"`
	ValueStr string   `json:"value_str"`
	OK       bool     `json:"ok"`
	Error    string   `json:"error,omitempty"`
}

// ReadRaw reads arbitrary registers and decodes them.
func ReadRaw(host string, port int, slaveID byte, address, count uint16, holding bool, dataType string) (*RawReadResult, error) {
	client, err := modbus.NewTCPClient(host, port, slaveID)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer client.Close()

	result := &RawReadResult{
		Address: address,
		Count:   count,
		Holding: holding,
	}

	data, err := client.ReadRegisters(address, count, holding)
	if err != nil {
		result.Error = err.Error()
		return result, nil
	}

	regs := modbus.BytesToRegisters(data)
	result.Raw = regs
	result.RawHex = make([]string, len(regs))
	for i, r := range regs {
		result.RawHex[i] = fmt.Sprintf("0x%04X", r)
	}

	dt, err := modbus.ParseDataType(dataType)
	if err != nil {
		result.OK = true
		result.ValueStr = "(raw only - unknown data type)"
		return result, nil
	}

	if dt == modbus.STR {
		result.ValueStr = modbus.DecodeString(regs)
		result.OK = true
		return result, nil
	}

	value, err := modbus.DecodeRawValue(regs, dt, modbus.Big)
	if err != nil {
		result.Error = err.Error()
		return result, nil
	}

	result.Decoded = value
	result.ValueStr = fmt.Sprintf("%.4f", value)
	result.OK = true
	return result, nil
}
