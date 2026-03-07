package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func GrowattRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("growatt", []modbus.RegisterDef{
		// Serial Number (holding, 5 regs = 10 chars)
		{Address: 23, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 5, Endianness: modbus.Big, UseHolding: true},

		// PV (input registers)
		{Address: 1, Name: "PV Input Power", SemanticName: "pv_power", Description: "Total PV input power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
		{Address: 91, Name: "PV Energy Total", SemanticName: "total_pv_gen", Description: "Total PV generation", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},

		// Battery (input registers)
		{Address: 1009, Name: "Battery Discharge Power", SemanticName: "battery_power", Description: "Battery discharge power", Unit: "W", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
		{Address: 1011, Name: "Battery Charge Power", SemanticName: "battery_charge_power", Description: "Battery charge power", Unit: "W", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
		{Address: 1014, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 1054, Name: "Total Discharge Energy", SemanticName: "total_discharge", Description: "Total battery discharge energy", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},

		// Grid (input registers)
		{Address: 1021, Name: "Grid Import Power", SemanticName: "meter_power", Description: "AC power to user total", Unit: "W", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
		{Address: 1029, Name: "Grid Export Power", SemanticName: "grid_export_power", Description: "AC power to grid total", Unit: "W", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
		{Address: 1046, Name: "Grid Import Energy", SemanticName: "total_import", Description: "Total energy bought from grid", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big},
	}, nil)
}
