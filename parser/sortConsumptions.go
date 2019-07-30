package parser

import (
	"strconv"

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
		return consumption.Hora == strconv.Itoa(hour)
	}).(Consumptions)

	return &filtered
}

// SplitConsumptionsHourly split the consumptions by hour in a map
func SplitConsumptionsHourly(consumptions *Consumptions) *ConsumptionsByHour {

	splitHoursConsumptions := make(ConsumptionsByHour, 25)

	for hour := 1; hour < 25; hour++ {
		filtered := filterConsumptionsHourly(consumptions, hour)
		splitHoursConsumptions[hour] = filtered
	}

	return &splitHoursConsumptions
}

// SplitConsumptionsWeekDays split the given consumptions by weekDay, 0=SUNDAY
func SplitConsumptionsWeekDays(consumptions *Consumptions) *ConsumptionsByWeekDay {
	splitWeekDaysConsumptions := make(ConsumptionsByWeekDay, 7)

	for weekDay := 0; weekDay < 7; weekDay++ {
		filtered := filterConsumptionsWeekDay(consumptions, weekDay)
		splitWeekDaysConsumptions[weekDay] = filtered
	}

	return &splitWeekDaysConsumptions
}
