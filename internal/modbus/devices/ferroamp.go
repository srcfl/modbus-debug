package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

func FerroampRegisters() *modbus.RegisterSet {
	return modbus.NewRegisterSet("ferroamp", []modbus.RegisterDef{
		// Detection — use Inverter Status (Ferroamp has no serial register; status 1-3 = valid device)
		{Address: 2000, Name: "Inverter Status", SemanticName: "serial_number", Description: "0=unavailable, 1=idle, 2=running, 3=fault", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},

		// General
		{Address: 1004, Name: "Modbus Version", SemanticName: "modbus_version", Description: "Modbus major version", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},

		// Inverter Data (section 6.2) — F32 uses Little-endian word order per spec
		{Address: 2100, Name: "Inverter Active Power", SemanticName: "inverter_power", Description: "Power based on inverter active currents", Unit: "kW", Category: "pv", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},

		// PV Data (section 6.5)
		{Address: 5000, Name: "Idle SSOs", SemanticName: "idle_sso", Description: "Number of SSOs not producing energy", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 5002, Name: "Running SSOs", SemanticName: "running_sso", Description: "Number of running SSOs", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 5004, Name: "Faulty SSOs", SemanticName: "faulty_sso", Description: "Number of SSOs not working properly", Unit: "", Category: "pv", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 5064, Name: "Energy Produced", SemanticName: "total_pv_gen", Description: "Outgoing energy from PV panels to DC-Nanogrid", Unit: "kWh", Category: "pv", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 5100, Name: "Solar Output Power", SemanticName: "pv_power", Description: "Outgoing power from PV panels to DC-Nanogrid", Unit: "kW", Category: "pv", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},

		// Grid (section 6.2)
		{Address: 2016, Name: "Grid Frequency", SemanticName: "grid_frequency", Description: "Grid frequency", Unit: "Hz", Category: "grid", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 2032, Name: "Grid Voltage L1", SemanticName: "grid_voltage_l1", Description: "Grid voltage phase L1", Unit: "Vrms", Category: "grid", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 2036, Name: "Grid Voltage L2", SemanticName: "grid_voltage_l2", Description: "Grid voltage phase L2", Unit: "Vrms", Category: "grid", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 2040, Name: "Grid Voltage L3", SemanticName: "grid_voltage_l3", Description: "Grid voltage phase L3", Unit: "Vrms", Category: "grid", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},

		// Meter / Grid Power (section 6.3)
		{Address: 3100, Name: "Grid Active Power", SemanticName: "meter_power", Description: "Power based on grid active currents", Unit: "kW", Category: "meter", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 3064, Name: "Energy Exported To Grid", SemanticName: "total_export", Description: "Outgoing energy from facility to AC Grid", Unit: "kWh", Category: "meter", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 3068, Name: "Energy Imported From Grid", SemanticName: "total_import", Description: "Incoming energy from AC Grid to facility", Unit: "kWh", Category: "meter", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},

		// Load (section 6.4)
		{Address: 4100, Name: "Load Active Power", SemanticName: "load_power", Description: "Power based on load active currents", Unit: "kW", Category: "meter", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},

		// Battery Data (section 6.7)
		{Address: 6000, Name: "Idle Batteries", SemanticName: "idle_batteries", Description: "Number of batteries not configured for operation", Unit: "", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 6002, Name: "Running Batteries", SemanticName: "running_batteries", Description: "Number of running batteries", Unit: "", Category: "battery", DataType: modbus.U16, Scale: 1.0, Words: 1, Endianness: modbus.Big},
		{Address: 6008, Name: "Rated Capacity", SemanticName: "battery_capacity", Description: "Aggregated rated capacity of all connected batteries", Unit: "kWh", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 6012, Name: "State of Health", SemanticName: "battery_soh", Description: "Aggregated state of health", Unit: "%", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 6016, Name: "State of Charge", SemanticName: "battery_soc", Description: "Aggregated state of charge", Unit: "%", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 6064, Name: "Energy from Battery", SemanticName: "total_discharge", Description: "Outgoing energy from discharging batteries", Unit: "kWh", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 6068, Name: "Energy to Battery", SemanticName: "total_charge", Description: "Incoming energy to charging batteries", Unit: "kWh", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
		{Address: 6100, Name: "Battery Output Power", SemanticName: "battery_power", Description: "Power from all batteries to DC-Nanogrid", Unit: "kW", Category: "battery", DataType: modbus.F32, Scale: 1.0, Words: 2, Endianness: modbus.Little},
	}, nil)
}
