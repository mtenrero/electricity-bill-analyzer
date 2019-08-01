package analyzers

import (
	"strconv"
	"strings"

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
			consumoSanitized := strings.Replace(consumption.Consumo, ",", ".", -1)
			floatValue, _ := strconv.ParseFloat(consumoSanitized, 64)
			return floatValue
		}).([]float64)

		weekDayMean := stat.Mean(listConsumptions, nil)
		analysis[weekDay] = ReportValue(ReportValue{Mean: weekDayMean})
	}

	return Report{Analysis: analysis}
}
