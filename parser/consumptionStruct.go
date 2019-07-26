package parser

import "time"

// ConsumptionEntry represents the data given by IDE Spanish Ligth distributor
type ConsumptionEntry struct {
	CUPS    string `json:"CUPS"`
	Fecha   string `json:"Fecha"`
	Hora    string `json:"Hora"`
	Consumo string `json:"Consumo_kWh"`
}

func (entry *ConsumptionEntry) timeConsumption() (time.Time, error) {
	timeStr := entry.Fecha + " " + entry.Hora
	return time.Parse("02/01/2006 15", timeStr)
}
