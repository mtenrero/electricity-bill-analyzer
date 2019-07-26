package parser

import "time"

// ConsumptionsByWeekDay represents the energy consumptions split by weekdays
type ConsumptionsByWeekDay = map[int]*Consumptions

// ConsumptionsByHour represents energy consumptions split by day hours
type ConsumptionsByHour = map[int]*Consumptions

// Consumptions is a collection with multiple Consumption Entries
type Consumptions = []ConsumptionEntry

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

func (entry *ConsumptionEntry) weekDay() time.Weekday {
	timeConsumption, _ := entry.timeConsumption()

	return timeConsumption.Weekday()
}
