package analyzers

import (
	"strconv"
	"strings"
	"time"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
	"github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat"
)

// WeekDayString returns the weekDay English string
func WeekDayString(weekDay int) *string {
	weekDayString := time.Weekday(weekDay).String()
	return &weekDayString
}

// ReportWeekDays Analyzes the Consumption done based on weeekdays intervals and returns the consumption avg for each weekDay
func ReportWeekDays(consumptions *parser.Consumptions) Report {
	analysis := make([]ReportValue, 7)

	consumptionsByWeekDay := *parser.SplitConsumptionsWeekDays(consumptions)

	for weekDay := 0; weekDay < 7; weekDay++ {

		weekDayToSet := weekDay
		listConsumptions := funk.Map(*consumptionsByWeekDay[weekDay], func(consumption parser.ConsumptionEntry) float64 {
			consumoSanitized := strings.Replace(consumption.Consumo, ",", ".", -1)
			floatValue, _ := strconv.ParseFloat(consumoSanitized, 64)
			return floatValue
		}).([]float64)

		weekDayMean := stat.Mean(listConsumptions, nil)
		analysis[weekDay] = ReportValue(ReportValue{
			WeekDay:    &weekDayToSet,
			WeekString: WeekDayString(weekDayToSet),
			Mean:       &weekDayMean,
		})
	}

	return Report{Analysis: analysis}
}
