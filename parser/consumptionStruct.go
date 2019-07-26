package parser

// ConsumptionEntry represents the data given by IDE Spanish Ligth distributor
type ConsumptionEntry struct {
	CUPS    string `json:"CUPS"`
	Fecha   string `json:"Fecha"`
	Hora    string `json:"Hora"`
	Consumo string `json:"Consumo_kWh"`
}
