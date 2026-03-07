package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func FoxESSRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("foxess", []modbus.RegisterDef{
		// Serial Number (holding, 8 regs = 16 chars)
		{Address: 30000, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 8, Endianness: modbus.Big, UseHolding: true},

		// PV (holding registers — H3 series TCP)
		{Address: 31002, Name: "PV1 Power", SemanticName: "pv1_power", Description: "PV string 1 power", Unit: "W", Category: "pv", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 31005, Name: "PV2 Power", SemanticName: "pv2_power", Description: "PV string 2 power", Unit: "W", Category: "pv", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 32012, Name: "Grid Consumption Total", SemanticName: "total_import", Description: "Total grid consumption energy", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery (holding registers)
		{Address: 31036, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery charge/discharge power", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 31038, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Grid (holding registers)
		{Address: 31026, Name: "Meter Power R", SemanticName: "meter_power_l1", Description: "Meter power phase R", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: -1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 31027, Name: "Meter Power S", SemanticName: "meter_power_l2", Description: "Meter power phase S", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: -1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 31028, Name: "Meter Power T", SemanticName: "meter_power_l3", Description: "Meter power phase T", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: -1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
