package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func GoodWeRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("goodwe", []modbus.RegisterDef{
		// Serial Number (holding, 8 regs = 16 chars)
		{Address: 35003, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 8, Endianness: modbus.Big, UseHolding: true},

		// PV
		{Address: 35105, Name: "PV1 Power", SemanticName: "pv1_power", Description: "PV string 1 power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 35109, Name: "PV2 Power", SemanticName: "pv2_power", Description: "PV string 2 power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 35191, Name: "PV Energy Total", SemanticName: "total_pv_gen", Description: "Total PV generation", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery
		{Address: 35182, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power (pos=charge, neg=discharge)", Unit: "W", Category: "battery", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 37007, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 35209, Name: "Battery Discharge Energy", SemanticName: "total_discharge", Description: "Total battery discharge energy", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Grid
		{Address: 35123, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Meter
		{Address: 36025, Name: "Grid Power", SemanticName: "meter_power", Description: "Grid total active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: -1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 36017, Name: "Grid Buy Energy", SemanticName: "total_import", Description: "Total energy bought from grid", Unit: "kWh", Category: "meter", DataType: modbus.F32, Scale: 0.001, Words: 2, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
