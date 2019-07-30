package analyzers

import (
	"math"
	"strconv"
	"strings"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
	"github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat"
)

// ReportHourly makes calculations (Mean) by hour based on the electricity consumption
func ReportHourly(consumptions *parser.Consumptions) Report {

	analysis := make([]ReportValue, 25)

	consumptionsByHour := parser.SplitConsumptionsHourly(consumptions)

	for hour, consumptions := range *consumptionsByHour {

		listConsumptions := funk.Map(*consumptions, func(consumption parser.ConsumptionEntry) float64 {
			consumoSanitized := strings.Replace(consumption.Consumo, ",", ".", -1)
			floatValue, _ := strconv.ParseFloat(consumoSanitized, 64)
			return floatValue
		}).([]float64)

		hourMean := stat.Mean(listConsumptions, nil)

		if math.IsNaN(hourMean) {
			analysis[hour] = ReportValue{
				Mean: 0,
			}
		} else {
			analysis[hour] = ReportValue{
				Mean: hourMean,
			}
		}
	}

	return Report{Analysis: analysis}
}
