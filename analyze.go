package electricity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mtenrero/electricity-bill-analyzer/analyzers"
	"github.com/mtenrero/electricity-bill-analyzer/parser"
)

func FullMeanAnalysis(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Could not read request", http.StatusBadRequest)
	}

	consumptions, err := parser.ParseCSVBytes(b)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Error processing CSV data file. It may not have the appropriate format"))
	}

	weekDaysHourReport := analyzers.ReportWeekDaysHourly(&consumptions)
	hourlyReport := analyzers.ReportHourly(&consumptions)
	weekDaysReport := analyzers.ReportWeekDays(&consumptions)

	fullReport := analyzers.ResponseReport{
		WeekDaysReport:     weekDaysReport,
		HourlyReport:       hourlyReport,
		WeekDayHoursReport: weekDaysHourReport,
	}

	jsonReport, err := json.Marshal(fullReport)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating JSON report"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonReport)

}
