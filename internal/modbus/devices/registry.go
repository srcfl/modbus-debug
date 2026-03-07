package devices

import "github.com/srcfl/modbus-debug/internal/modbus"

// Profile contains metadata for a device profile used in detection and display.
type Profile struct {
	Name        string
	DisplayName string
	Registers   func() *modbus.RegisterSet
}

var profiles = []Profile{
	{Name: "sungrow", DisplayName: "Sungrow SH Hybrid", Registers: SungrowRegisters},
	{Name: "huawei", DisplayName: "Huawei SUN2000", Registers: HuaweiRegisters},
	{Name: "solis", DisplayName: "Solis Hybrid", Registers: SolisRegisters},
	{Name: "fronius", DisplayName: "Fronius GEN24", Registers: FroniusRegisters},
	{Name: "deye", DisplayName: "Deye Hybrid", Registers: DeyeRegisters},
	{Name: "solaredge", DisplayName: "SolarEdge", Registers: SolarEdgeRegisters},
	{Name: "sma", DisplayName: "SMA Tripower", Registers: SMARegisters},
	{Name: "pixii", DisplayName: "Pixii PowerShaper", Registers: PixiiRegisters},
	{Name: "sdm630", DisplayName: "Eastron SDM630", Registers: SDM630Registers},
	{Name: "fronius-smart-meter", DisplayName: "Fronius Smart Meter", Registers: FroniusMeterRegisters},
	{Name: "ferroamp", DisplayName: "Ferroamp EnergyHub", Registers: FerroampRegisters},
	{Name: "goodwe", DisplayName: "GoodWe Hybrid", Registers: GoodWeRegisters},
	{Name: "growatt", DisplayName: "Growatt Hybrid", Registers: GrowattRegisters},
	{Name: "solax", DisplayName: "SolaX X1/X3", Registers: SolaXRegisters},
	{Name: "sofar", DisplayName: "SofarSolar HYD G3", Registers: SofarRegisters},
	{Name: "victron", DisplayName: "Victron Energy", Registers: VictronRegisters},
	{Name: "foxess", DisplayName: "FoxESS H3", Registers: FoxESSRegisters},
	{Name: "alpha-ess", DisplayName: "Alpha ESS SMILE", Registers: AlphaESSRegisters},
	{Name: "e3dc", DisplayName: "E3DC", Registers: E3DCRegisters},
}

// GetRegisterSet returns the RegisterSet for a given profile name.
func GetRegisterSet(name string) *modbus.RegisterSet {
	for _, p := range profiles {
		if p.Name == name {
			return p.Registers()
		}
	}
	return nil
}

// AllProfiles returns all registered device profiles.
func AllProfiles() []Profile {
	return profiles
}
