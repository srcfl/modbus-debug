package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func SAJRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("saj", []modbus.RegisterDef{
		// Serial Number (input, 8 regs = 16 chars)
		{Address: 36611, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 8, Endianness: modbus.Big},

		// PV (holding registers — H2 series)
		{Address: 16549, Name: "PV Total Power", SemanticName: "pv_power", Description: "Total PV power", Unit: "W", Category: "pv", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 16581, Name: "PV Energy Total", SemanticName: "total_pv_gen", Description: "Total PV generation", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Battery (holding registers)
		{Address: 16550, Name: "Battery Power", SemanticName: "battery_power", Description: "Total battery power", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 40972, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 16597, Name: "Battery Discharge Energy", SemanticName: "total_discharge", Description: "Total battery discharge energy", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Grid (holding registers)
		{Address: 16557, Name: "Grid Power", SemanticName: "meter_power", Description: "Grid power (meter)", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 16637, Name: "Grid Feed-in Energy", SemanticName: "total_export", Description: "Total feed-in energy", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.01, Words: 2, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
