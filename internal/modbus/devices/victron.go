package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

// VictronRegisters returns the register set for Victron Energy GX devices.
// Note: Victron uses unit ID 100 for system-level registers.
func VictronRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("victron", []modbus.RegisterDef{
		// Serial Number (input, 6 regs = 12 chars, unit ID 100)
		{Address: 800, Name: "Serial Number", SemanticName: "serial_number", Description: "System serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 6, Endianness: modbus.Big},

		// PV (input registers, unit ID 100)
		{Address: 850, Name: "DC PV Power", SemanticName: "pv_power", Description: "DC-coupled PV power", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 808, Name: "AC-out PV L1", SemanticName: "pv_ac_out_l1", Description: "AC-out PV power L1", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 809, Name: "AC-out PV L2", SemanticName: "pv_ac_out_l2", Description: "AC-out PV power L2", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 810, Name: "AC-out PV L3", SemanticName: "pv_ac_out_l3", Description: "AC-out PV power L3", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},

		// Battery (input registers, unit ID 100)
		{Address: 842, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery DC power (negative=charging)", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: -1.0, Words: 1, Endianness: modbus.Big},
		{Address: 843, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},

		// Grid (input registers, unit ID 100)
		{Address: 820, Name: "Grid Power L1", SemanticName: "grid_power_l1", Description: "Grid power L1", Unit: "W", Category: "grid", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 821, Name: "Grid Power L2", SemanticName: "grid_power_l2", Description: "Grid power L2", Unit: "W", Category: "grid", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 822, Name: "Grid Power L3", SemanticName: "grid_power_l3", Description: "Grid power L3", Unit: "W", Category: "grid", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
	}, nil)
}
