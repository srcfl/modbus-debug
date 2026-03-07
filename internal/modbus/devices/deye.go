package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func DeyeRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("deye", []modbus.RegisterDef{
		// Identity
		{Address: 3, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 5, Endianness: modbus.Big, UseHolding: true},

		// PV Registers
		{Address: 672, Name: "PV Power", SemanticName: "pv_power", Description: "Total PV power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 109, Name: "PV1 Voltage", SemanticName: "pv1_voltage", Description: "PV1 voltage", Unit: "V", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 110, Name: "PV1 Current", SemanticName: "pv1_current", Description: "PV1 current", Unit: "A", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 111, Name: "PV2 Voltage", SemanticName: "pv2_voltage", Description: "PV2 voltage", Unit: "V", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 112, Name: "PV2 Current", SemanticName: "pv2_current", Description: "PV2 current", Unit: "A", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Battery Registers
		{Address: 590, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power", Unit: "W", Category: "battery", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 588, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 587, Name: "Battery Voltage", SemanticName: "battery_voltage", Description: "Battery voltage", Unit: "V", Category: "battery", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Meter / Grid
		{Address: 619, Name: "Meter Power", SemanticName: "meter_power", Description: "Grid power", Unit: "W", Category: "meter", DataType: modbus.I16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
		{Address: 79, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big, UseHolding: true},

		// Energy Totals
		{Address: 534, Name: "Total PV Generation", SemanticName: "total_pv_gen", Description: "Total PV energy", Unit: "kWh", Category: "pv", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 520, Name: "Total Charge", SemanticName: "total_charge", Description: "Total battery charge", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 522, Name: "Total Discharge", SemanticName: "total_discharge", Description: "Total battery discharge", Unit: "kWh", Category: "battery", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 524, Name: "Total Import", SemanticName: "total_import", Description: "Total grid import", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},
		{Address: 526, Name: "Total Export", SemanticName: "total_export", Description: "Total grid export", Unit: "kWh", Category: "meter", DataType: modbus.U32, Scale: 0.1, Words: 2, Endianness: modbus.Big, UseHolding: true},

		// Load
		{Address: 653, Name: "Load Power", SemanticName: "load_power", Description: "Total load power", Unit: "W", Category: "control", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big, UseHolding: true},
	}, nil)
}
