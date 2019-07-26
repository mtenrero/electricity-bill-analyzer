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
