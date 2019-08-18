package parser

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
)

func ParseCSVBytes(content []byte) ([]ConsumptionEntry, error) {
	reader := csv.NewReader(bytes.NewReader(content))
	reader.Comma = ';'

	var consumptions = make([]ConsumptionEntry, 0)

	var header []string
	var glerror error

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			glerror = err
			break
		}
		if header == nil {
			header = record
		} else {
			hora, _ := strconv.Atoi(record[2])
			consumptionEntry := ConsumptionEntry{
				CUPS:    record[0],
				Fecha:   record[1],
				Hora:    hora,
				Consumo: record[3],
			}

			consumptions = append(consumptions, consumptionEntry)
		}
	}

	return consumptions, glerror
}
