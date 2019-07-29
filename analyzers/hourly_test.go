package analyzers

import (
	"testing"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
)

func TestHourlyReport(t *testing.T) {

	consumptions := make(parser.Consumptions, 0)
	cons1 := parser.ConsumptionEntry{
		Hora:    "3",
		Consumo: "0.7",
	}
	cons2 := parser.ConsumptionEntry{
		Hora:    "3",
		Consumo: "3.8",
	}
	cons3 := parser.ConsumptionEntry{
		Hora:    "3",
		Consumo: "5",
	}
	consumptions = append(consumptions, cons1)
	consumptions = append(consumptions, cons2)
	consumptions = append(consumptions, cons3)

	report := ReportHourly(&consumptions)

	if report[3] != 3.1666666666666665 {
		t.Fail()
	}
}
