package parser

import (
	"github.com/thoas/go-funk"
)

// filterConsumptionsWeekDay Filters the consumptions by WeekDay
func filterConsumptionsWeekDay(consumptions *Consumptions, weekDay int) *Consumptions {

	filtered := funk.Filter(*consumptions, func(consumption ConsumptionEntry) bool {
		return int(consumption.WeekDay()) == weekDay
	}).(Consumptions)

	return &filtered
}

// filterConsumptionsHourly Filters the consumptions by hour
func filterConsumptionsHourly(consumptions *Consumptions, hour int) *Consumptions {
	filtered := funk.Filter(*consumptions, func(consumption ConsumptionEntry) bool {
		consumptionTime, _ := consumption.TimeConsumption()
		return consumptionTime.Hour() == hour
	}).(Consumptions)

	return &filtered
}

// splitConsumptionsHourly split the consumptions by hour in a map
func SplitConsumptionsHourly(consumptions *Consumptions) *ConsumptionsByHour {

	splitHoursConsumptions := make(ConsumptionsByHour, 24)

	for hour := 0; hour < 24; hour++ {
		filtered := filterConsumptionsHourly(consumptions, hour)
		splitHoursConsumptions[hour] = filtered
	}

	return &splitHoursConsumptions
}

// splitConsumptionsWeekDays split the given consumptions by weekDay, 0=SUNDAY
func SplitConsumptionsWeekDays(consumptions *Consumptions) *ConsumptionsByWeekDay {
	splitWeekDaysConsumptions := make(ConsumptionsByWeekDay, 7)

	for weekDay := 0; weekDay < 7; weekDay++ {
		filtered := filterConsumptionsWeekDay(consumptions, weekDay)
		splitWeekDaysConsumptions[weekDay] = filtered
	}

	return &splitWeekDaysConsumptions
}
