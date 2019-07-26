package parser

import "testing"

func TestCSVParse(t *testing.T) {
	sampleCSV := []byte("CUPS;Fecha;Hora;Consumo_kWh;Metodo_obtencion\nES9234823489237489EE;16/05/2019;1;0,141;R\nES3282748273489237EA;16/05/2019;2;0,111;R")

	var consumptionsParsed []ConsumptionEntry

	consumptionsParsed, err := parseCSVBytes(sampleCSV)
	if err != nil {
		t.Error("parseCSVBytes errored :(")
	}

	if consumptionsParsed[0].CUPS != "ES9234823489237489EE" {
		t.Error("Not parsed correctly")
	}

	if consumptionsParsed[0].Consumo != "0,141" {
		t.Error("Not parsed correctly")
	}

	if consumptionsParsed[1].Fecha != "16/05/2019" {
		t.Error("Not parsed correctly")
	}

	if consumptionsParsed[1].Hora != "2" {
		t.Error("Not parsed correctly")
	}
}
