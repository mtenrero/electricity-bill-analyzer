package analyzers

import (
	"github.com/mtenrero/electricity-bill-analyzer/parser"
)

// Analyzes the Consumption done based on weeekdays intervals and returns the consumption avg for each weekDay
func ReportWeekDaysHourly(consumptions *parser.Consumptions) []ReportWeekDayHourly {
	report := make([]ReportWeekDayHourly, 7)

	consumptionsByWeekDay := *parser.SplitConsumptionsWeekDays(consumptions)

	for weekDay := 0; weekDay < 7; weekDay++ {
		weekDayToSet := weekDay
		reportweekDay := ReportWeekDayHourly{
			WeekDay:    &weekDayToSet,
			WeekString: WeekDayString(weekDayToSet),
			Analysis:   ReportHourly(consumptionsByWeekDay[weekDay]),
		}
		report[weekDay] = reportweekDay
	}

	return report
}
