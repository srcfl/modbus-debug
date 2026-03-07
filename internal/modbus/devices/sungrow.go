package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func SungrowRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("sungrow", []modbus.RegisterDef{
		// Serial Number
		{Address: 4989, Name: "Serial Number", SemanticName: "serial_number", Description: "Device serial number", Unit: "", Category: "pv", DataType: modbus.STR, Scale: 1.0, Words: 10, Endianness: modbus.Big},

		// PV Registers (Input Registers)
		{Address: 5000, Name: "Nominal Output Power", SemanticName: "nominal_power", Description: "Nominal output power", Unit: "kW", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5007, Name: "Inside Temperature", SemanticName: "inverter_temperature", Description: "Internal temperature", Unit: "C", Category: "pv", DataType: modbus.I16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5010, Name: "MPPT 1 Voltage", SemanticName: "pv1_voltage", Description: "MPPT 1 Voltage", Unit: "V", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5011, Name: "MPPT 1 Current", SemanticName: "pv1_current", Description: "MPPT 1 Current", Unit: "A", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5012, Name: "MPPT 2 Voltage", SemanticName: "pv2_voltage", Description: "MPPT 2 Voltage", Unit: "V", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5013, Name: "MPPT 2 Current", SemanticName: "pv2_current", Description: "MPPT 2 Current", Unit: "A", Category: "pv", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5016, Name: "PV Power", SemanticName: "pv_power", Description: "Total DC power", Unit: "W", Category: "pv", DataType: modbus.U32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 13002, Name: "Total PV Generation", SemanticName: "total_pv_gen", Description: "Total PV generation", Unit: "Wh", Category: "pv", DataType: modbus.U32, Scale: 100.0, Words: 2, Endianness: modbus.Little},

		// Battery Registers
		{Address: 13000, Name: "System State Flags", SemanticName: "battery_flags", Description: "Inverter running states", Unit: "", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 13019, Name: "Battery Voltage", SemanticName: "battery_voltage", Description: "Battery voltage", Unit: "V", Category: "battery", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 13020, Name: "Battery Current", SemanticName: "battery_current", Description: "Battery current", Unit: "A", Category: "battery", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 13021, Name: "Battery Power", SemanticName: "battery_power", Description: "Battery power", Unit: "W", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 13022, Name: "Battery SoC", SemanticName: "battery_soc", Description: "Battery state of charge", Unit: "%", Category: "battery", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 13024, Name: "Battery Temperature", SemanticName: "battery_temperature", Description: "Battery temperature", Unit: "C", Category: "battery", DataType: modbus.I16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 13026, Name: "Total Battery Discharge", SemanticName: "total_discharge", Description: "Total discharge energy", Unit: "Wh", Category: "battery", DataType: modbus.U32, Scale: 100.0, Words: 2, Endianness: modbus.Little},
		{Address: 13040, Name: "Total Charge Energy", SemanticName: "total_charge", Description: "Total charge energy", Unit: "Wh", Category: "battery", DataType: modbus.U32, Scale: 100.0, Words: 2, Endianness: modbus.Little},

		// Meter Registers
		{Address: 5600, Name: "Meter Power", SemanticName: "meter_power", Description: "Meter total active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 5602, Name: "Meter L1 Power", SemanticName: "meter_l1_power", Description: "Phase A active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 5604, Name: "Meter L2 Power", SemanticName: "meter_l2_power", Description: "Phase B active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 5606, Name: "Meter L3 Power", SemanticName: "meter_l3_power", Description: "Phase C active power", Unit: "W", Category: "meter", DataType: modbus.I32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 5740, Name: "Meter L1 Voltage", SemanticName: "meter_l1_voltage", Description: "Phase A voltage", Unit: "V", Category: "meter", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5741, Name: "Meter L2 Voltage", SemanticName: "meter_l2_voltage", Description: "Phase B voltage", Unit: "V", Category: "meter", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5742, Name: "Meter L3 Voltage", SemanticName: "meter_l3_voltage", Description: "Phase C voltage", Unit: "V", Category: "meter", DataType: modbus.U16, Scale: 0.1, Words: 1, Endianness: modbus.Big},
		{Address: 5743, Name: "Meter L1 Current", SemanticName: "meter_l1_current", Description: "Phase A current", Unit: "A", Category: "meter", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big},
		{Address: 5744, Name: "Meter L2 Current", SemanticName: "meter_l2_current", Description: "Phase B current", Unit: "A", Category: "meter", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big},
		{Address: 5745, Name: "Meter L3 Current", SemanticName: "meter_l3_current", Description: "Phase C current", Unit: "A", Category: "meter", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big},
		{Address: 13036, Name: "Total Import Energy", SemanticName: "total_import", Description: "Total import energy", Unit: "Wh", Category: "meter", DataType: modbus.U32, Scale: 100.0, Words: 2, Endianness: modbus.Little},
		{Address: 13045, Name: "Total Export Energy", SemanticName: "total_export", Description: "Total export energy", Unit: "Wh", Category: "meter", DataType: modbus.U32, Scale: 100.0, Words: 2, Endianness: modbus.Little},

		// Grid Registers
		{Address: 5241, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.U16, Scale: 0.01, Words: 1, Endianness: modbus.Big},
	}, nil)
}
