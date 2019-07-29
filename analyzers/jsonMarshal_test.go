package analyzers

import (
	"encoding/json"
	"testing"

	"github.com/mtenrero/electricity-bill-analyzer/parser"
)

func TestMarshalReportValue(t *testing.T) {
	floatvalue := ReportValue(2.2222)

	stringCast := floatvalue.Marshal()

	if stringCast != "2.2222" {
		t.Fail()
	}
}

func TestJSONMarshal(t *testing.T) {

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

	reportJSON, err := json.Marshal(report.JSON())

	if err != nil {
		t.Error(err)
	}

	t.Log(reportJSON)

}
