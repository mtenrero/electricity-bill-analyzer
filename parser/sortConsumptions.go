package parser

import (
	"github.com/thoas/go-funk"
)

// filterConsumptionsWeekDay Filters the consumptions by WeekDay
func filterConsumptionsWeekDay(consumptions *Consumptions, weekDay int) *Consumptions {

	filtered := funk.Filter(*consumptions, func(consumption ConsumptionEntry) bool {
		return int(consumption.weekDay()) == weekDay
	}).(Consumptions)

	return &filtered
}

// filterConsumptionsHourly Filters the consumptions by hour
func filterConsumptionsHourly(consumptions *Consumptions, hour int) *Consumptions {
	filtered := funk.Filter(*consumptions, func(consumption ConsumptionEntry) bool {
		consumptionTime, _ := consumption.timeConsumption()
		return consumptionTime.Hour() == hour
	}).(Consumptions)

	return &filtered
}

// splitConsumptionsHourly split the consumptions by hour in a map
func splitConsumptionsHourly(consumptions *Consumptions) *map[int]*Consumptions {

	splitHoursConsumptions := make(map[int]*Consumptions, 24)

	for hour := 0; hour < 24; hour++ {
		filtered := filterConsumptionsHourly(consumptions, hour)
		splitHoursConsumptions[hour] = filtered
	}

	return &splitHoursConsumptions
}
