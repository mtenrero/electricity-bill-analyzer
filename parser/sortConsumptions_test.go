package parser

import "testing"

func TestFilterHourly(t *testing.T) {

	var sampleCSV = []byte("CUPS;Fecha;Hora;Consumo_kWh;Metodo_obtencion\nES9234823489237489EE;16/05/2019;1;0,141;R\nES3282748273489237EA;16/05/2019;2;0,111;R\nES9234823489237489EE;18/05/2019;1;0,222;R")

	var consumptionsParsed []ConsumptionEntry
	consumptionsParsed, _ = parseCSVBytes(sampleCSV)

	consumptionsHour1 := *filterConsumptionsHourly(&consumptionsParsed, 1)
	if len(consumptionsHour1) != 2 {
		t.Fail()
	}

	if consumptionsHour1[0].CUPS != "ES9234823489237489EE" {
		t.Fail()
	}

	consumptionsHour2 := *filterConsumptionsHourly(&consumptionsParsed, 2)
	if len(consumptionsHour2) != 1 {
		t.Fail()
	}
	if consumptionsHour2[0].CUPS != "ES3282748273489237EA" {
		t.Fail()
	}
}

func TestFilterWeekDay(t *testing.T) {

	var sampleCSV = []byte("CUPS;Fecha;Hora;Consumo_kWh;Metodo_obtencion\nES9234823489237489EE;16/05/2019;1;0,141;R\nES3282748273489237EA;16/05/2019;2;0,111;R\nES9234823489237489EE;18/05/2019;1;0,222;R")

	var consumptionsParsed []ConsumptionEntry
	consumptionsParsed, _ = parseCSVBytes(sampleCSV)

	filteredWeekDay := *filterConsumptionsWeekDay(&consumptionsParsed, 6)

	if len(filteredWeekDay) != 1 {
		t.Fail()
	}
	if filteredWeekDay[0].CUPS != "ES9234823489237489EE" {
		t.Fail()
	}
}
