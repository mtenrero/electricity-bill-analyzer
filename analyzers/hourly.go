package analyzers

import (
	"strconv"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
	"github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat"
)

// ReportHourly makes calculations (Mean) by hour based on the electricity consumption
func ReportHourly(consumptions *parser.Consumptions) Report {

	sumKwH := make([]ReportValue, 24)

	consumptionsByHour := parser.SplitConsumptionsHourly(consumptions)

	for hour, consumptions := range *consumptionsByHour {

		listConsumptions := funk.Map(*consumptions, func(consumption parser.ConsumptionEntry) float64 {
			float, _ := strconv.ParseFloat(consumption.Consumo, 64)
			return float
		}).([]float64)

		hourMean := stat.Mean(listConsumptions, nil)

		sumKwH[hour] = ReportValue(hourMean)
	}

	return sumKwH
}
