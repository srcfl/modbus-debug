package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func SMARegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("sma", []modbus.RegisterDef{
		// Identity
		{Address: 30057, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// PV Registers
		{Address: 30775, Name: "PV Power", SemanticName: "pv_power", Description: "DC power input", Unit: "W", Category: "pv", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 30773, Name: "PV1 Voltage", SemanticName: "pv1_voltage", Description: "DC input voltage string A", Unit: "V", Category: "pv", DataType: modbus.I32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 30769, Name: "PV1 Current", SemanticName: "pv1_current", Description: "DC input current string A", Unit: "A", Category: "pv", DataType: modbus.I32, Scale: 0.001, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery Registers
		{Address: 30845, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 30847, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery charge/discharge power", Unit: "W", Category: "battery", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 30849, Name: "Battery Status", SemanticName: "battery_flags", Description: "Battery operating status", Unit: "", Category: "battery", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Meter / Grid
		{Address: 30865, Name: "Meter Power", SemanticName: "meter_power", Description: "Grid power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 30803, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Energy Totals (U64 = 4 words)
		{Address: 30529, Name: "Total PV Generation", SemanticName: "total_pv_gen", Description: "Total yield", Unit: "Wh", Category: "pv", DataType: modbus.U64, Scale: 1.0, Words: 4, Endianness: modbus.Big, UseHolding: true},
		{Address: 30583, Name: "Total Import", SemanticName: "total_import", Description: "Total grid import", Unit: "Wh", Category: "meter", DataType: modbus.U64, Scale: 1.0, Words: 4, Endianness: modbus.Big, UseHolding: true},
		{Address: 30579, Name: "Total Export", SemanticName: "total_export", Description: "Total grid feed-in", Unit: "Wh", Category: "meter", DataType: modbus.U64, Scale: 1.0, Words: 4, Endianness: modbus.Big, UseHolding: true},

		// Nominal
		{Address: 30231, Name: "Nominal Power", SemanticName: "nominal_power", Description: "Nominal active power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
