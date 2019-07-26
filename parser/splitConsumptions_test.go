package parser

import "testing"

func TestSplitByHour(t *testing.T) {
	var sampleCSV = []byte("CUPS;Fecha;Hora;Consumo_kWh;Metodo_obtencion\nES9234823489237489EE;16/05/2019;1;0,141;R\nES3282748273489237EA;16/05/2019;2;0,111;R\nES9234823489237489EE;18/05/2019;1;0,222;R")

	var consumptionsParsed []ConsumptionEntry
	consumptionsParsed, _ = ParseCSVBytes(sampleCSV)

	splitConsumptions := *SplitConsumptionsHourly(&consumptionsParsed)

	if len(*splitConsumptions[1]) != 2 {
		t.Fail()
	}

	if len(*splitConsumptions[2]) != 1 {
		t.Fail()
	}
}

func TestSplitByWeekDay(t *testing.T) {
	var sampleCSV = []byte("CUPS;Fecha;Hora;Consumo_kWh;Metodo_obtencion\nES9234823489237489EE;16/05/2019;1;0,141;R\nES3282748273489237EA;16/05/2019;2;0,111;R\nES9234823489237489EE;18/05/2019;1;0,222;R")

	var consumptionsParsed []ConsumptionEntry
	consumptionsParsed, _ = ParseCSVBytes(sampleCSV)

	splitConsumptions := *SplitConsumptionsWeekDays(&consumptionsParsed)

	if len(*splitConsumptions[6]) != 1 {
		t.Fail()
	}

	if len(*splitConsumptions[4]) != 2 {
		t.Fail()
	}
}
