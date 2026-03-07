package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func E3DCRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("e3dc", []modbus.RegisterDef{
		// Magic identifier (holding, register 40000 = 0xE3DC)
		{Address: 40000, Name: "E3DC Magic", SemanticName: "serial_number", Description: "E3DC identification (0xE3DC = 58332)", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// PV (holding registers, word-swapped int32)
		{Address: 40067, Name: "PV Power", SemanticName: "pv_power", Description: "Total PV power", Unit: "W", Category: "pv", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little, UseHolding: true},

		// Battery (holding registers)
		{Address: 40069, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery DC power", Unit: "W", Category: "battery", DataType: modbus.I32, Scale: -1.0, Words: 2, Endianness: modbus.Little, UseHolding: true},
		{Address: 40082, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Grid (holding registers)
		{Address: 40073, Name: "Grid Power", SemanticName: "meter_power", Description: "Grid/home power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little, UseHolding: true},
		{Address: 40075, Name: "Grid Feed-in Power", SemanticName: "grid_export_power", Description: "External feed-in power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: -1.0, Words: 2, Endianness: modbus.Little, UseHolding: true},
	}, nil)
}
