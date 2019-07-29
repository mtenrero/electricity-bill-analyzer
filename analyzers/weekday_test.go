package analyzers

import (
	"github.com/mtenrero/electricity-bill-analyzer/parser"
	"testing"
)

func TestReportWeekDays(t *testing.T) {
	consumptions := make(parser.Consumptions, 0)
	cons1 := parser.ConsumptionEntry{
		Fecha:   "28/7/2019",
		Hora:    "3",
		Consumo: "0.7",
	}
	cons2 := parser.ConsumptionEntry{
		Fecha:   "21/7/2019",
		Hora:    "3",
		Consumo: "3.8",
	}
	cons3 := parser.ConsumptionEntry{
		Fecha:   "29/07/2019",
		Hora:    "3",
		Consumo: "5",
	}
	consumptions = append(consumptions, cons1)
	consumptions = append(consumptions, cons2)
	consumptions = append(consumptions, cons3)

	report := ReportWeekDays(&consumptions)

	if report[0] != 2.25 {
		t.Fail()
	}

	if report[1] != 5 {
		t.Fail()
	}
}
