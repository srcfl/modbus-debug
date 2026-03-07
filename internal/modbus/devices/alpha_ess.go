package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func AlphaESSRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("alpha-ess", []modbus.RegisterDef{
		// Serial Number (holding, 8 regs = 16 chars, slave ID 85/0x55)
		{Address: 1610, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 8, Endianness: modbus.Big, UseHolding: true},

		// Grid frequency
		{Address: 1052, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// PV (holding registers)
		{Address: 1055, Name: "PV1 Power", SemanticName: "pv1_power", Description: "PV string 1 power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 1059, Name: "PV2 Power", SemanticName: "pv2_power", Description: "PV string 2 power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery (holding registers)
		{Address: 294, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 258, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 290, Name: "Battery Discharge Energy", SemanticName: "total_discharge", Description: "Total battery discharge energy", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Grid (holding registers)
		{Address: 33, Name: "Grid Power", SemanticName: "meter_power", Description: "Grid total active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 18, Name: "Grid Import Energy", SemanticName: "total_import", Description: "Total energy from grid", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 16, Name: "Grid Export Energy", SemanticName: "total_export", Description: "Total energy to grid", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
