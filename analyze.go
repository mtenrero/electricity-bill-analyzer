package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mtenrero/electricity-bill-analyzer/analyzers"
	"github.com/mtenrero/electricity-bill-analyzer/parser"
)

func FullMeanAnalysis(w http.ResponseWriter, r *http.Request) {

	const maxMemory = 5 * 1024 * 1024 // 5 megabytes

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request", http.StatusBadRequest)
	}

	consumptions, err := parser.ParseCSVBytes(b)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Error processing CSV data file. It may not have the appropriate format"))
	}

	report := analyzers.ReportWeekDaysHourly(&consumptions)

	jsonReport, err := json.Marshal(report)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Generating Report"))
	}

	w.Write(jsonReport)

}
