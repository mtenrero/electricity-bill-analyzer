package analyzers

import (
	"github.com/mtenrero/electricity-bill-analyzer/parser"
)


// Analyzes the Consumption done based on weeekdays intervals and returns the consumption avg for each weekDay
func ReportWeekDaysHourly(consumptions *parser.Consumptions) [][]float64 {
	report := make([][]float64, 7)

	consumptionsByWeekDay := *parser.SplitConsumptionsWeekDays(consumptions)

	for weekDay := 0; weekDay < 7; weekDay++ {
		report[weekDay] = ReportHourly(consumptionsByWeekDay[weekDay])
	}

	return report
}
