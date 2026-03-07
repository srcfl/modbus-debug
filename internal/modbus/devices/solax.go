package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func SolaXRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("solax", []modbus.RegisterDef{
		// Serial Number (holding, 7 regs = 14 chars)
		{Address: 768, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 7, Endianness: modbus.Big, UseHolding: true},

		// PV (input registers)
		{Address: 10, Name: "PV1 Power", SemanticName: "pv1_power", Description: "PV string 1 DC power", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 11, Name: "PV2 Power", SemanticName: "pv2_power", Description: "PV string 2 DC power", Unit: "W", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 148, Name: "Solar Energy Total", SemanticName: "total_pv_gen", Description: "Total solar energy", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Little},

		// Battery (input registers)
		{Address: 22, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power (negative=discharge)", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: -1.0, Words: 1, Endianness: modbus.Big},
		{Address: 28, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},

		// Grid (input registers) — word-swapped int32
		{Address: 70, Name: "Grid Power", SemanticName: "meter_power", Description: "Feed-in power (negative=import)", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: -1.0, Words: 2, Endianness: modbus.Little},
		{Address: 74, Name: "Consumed Energy Total", SemanticName: "total_import", Description: "Total consumed energy from grid", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Little},

		// Grid voltages (input registers)
		{Address: 202, Name: "Grid Voltage R", SemanticName: "grid_voltage_l1", Description: "Grid voltage phase R", Unit: "V", Category: "grid", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 203, Name: "Grid Voltage S", SemanticName: "grid_voltage_l2", Description: "Grid voltage phase S", Unit: "V", Category: "grid", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 204, Name: "Grid Voltage T", SemanticName: "grid_voltage_l3", Description: "Grid voltage phase T", Unit: "V", Category: "grid", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
	}, nil)
}
