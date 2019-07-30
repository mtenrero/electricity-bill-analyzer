package analyzers

import (
	"strconv"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
	"github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat"
)

// Analyzes the Consumption done based on weeekdays intervals and returns the consumption avg for each weekDay
func ReportWeekDays(consumptions *parser.Consumptions) Report {
	analysis := make([]ReportValue, 7)

	consumptionsByWeekDay := *parser.SplitConsumptionsWeekDays(consumptions)

	for weekDay := 0; weekDay < 7; weekDay++ {
		listConsumptions := funk.Map(*consumptionsByWeekDay[weekDay], func(consumption parser.ConsumptionEntry) float64 {
			float, _ := strconv.ParseFloat(consumption.Consumo, 64)
			return float
		}).([]float64)

		weekDayMean := stat.Mean(listConsumptions, nil)
		analysis[weekDay] = ReportValue(ReportValue{Mean: weekDayMean})
	}

	return Report{Analysis: analysis}
}
