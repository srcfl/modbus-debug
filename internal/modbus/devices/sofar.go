package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func SofarRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("sofar", []modbus.RegisterDef{
		// Serial Number (holding, 7 regs = 14 chars)
		{Address: 1093, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 7, Endianness: modbus.Big, UseHolding: true},

		// PV (holding registers — G3 variant)
		{Address: 1414, Name: "PV1 Power", SemanticName: "pv1_power", Description: "PV string 1 power", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 10.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 1417, Name: "PV2 Power", SemanticName: "pv2_power", Description: "PV string 2 power", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 10.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 1670, Name: "PV Generation Total", SemanticName: "total_pv_gen", Description: "Total PV generation", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery (holding registers)
		{Address: 1542, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power (negative=discharge)", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: -10.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 1544, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Grid / Meter (holding registers)
		{Address: 1160, Name: "Grid Power", SemanticName: "meter_power", Description: "Grid total active power", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: -10.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 1678, Name: "Grid Import Energy", SemanticName: "total_import", Description: "Total energy purchased from grid", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
