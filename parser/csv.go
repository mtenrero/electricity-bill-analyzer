package parser

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
)

func parseCSVBytes(content []byte) ([]ConsumptionEntry, error) {
	reader := csv.NewReader(bytes.NewReader(content))
	reader.Comma = ';'

	var consumptions = make([]ConsumptionEntry, 0)

	var header []string
	var err error

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			consumptionEntry := ConsumptionEntry{
				CUPS:    record[0],
				Fecha:   record[1],
				Hora:    record[2],
				Consumo: record[3],
			}

			consumptions = append(consumptions, consumptionEntry)
		}
	}

	return consumptions, err
}
