package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func PixiiRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("pixii", []modbus.RegisterDef{
		// Identity (SunSpec Common Model)
		{Address: 40052, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "battery", DataType: modbus.STR, Scale: 1.0, Words: 16, Endianness: modbus.Big, UseHolding: true},

		// Battery (SunSpec Storage Model)
		{Address: 40168, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery charge/discharge power", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 40132, Name: "Battery SoC", SemanticName: "battery_soc", Description: "State of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 40134, Name: "Battery Status", SemanticName: "battery_flags", Description: "Battery operating status", Unit: "", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Meter (SunSpec Meter Model)
		{Address: 40188, Name: "Meter Power", SemanticName: "meter_power", Description: "Total active power", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 40187, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
