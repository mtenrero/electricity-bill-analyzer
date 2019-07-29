package analyzers

import "strconv"

type Report []ReportValue
type ReportGroup []Report
type ReportValue float64

// Marshal the float64 ReportValue into a string
func (value ReportValue) Marshal() string {
	valueString := strconv.FormatFloat(float64(value), 'f', 4, 64)
	return valueString
}

func (reportGroup *ReportGroup) JSON() []byte {
	stringReport := []byte("{")

	for index, value := range *reportGroup {
		if index != 0 {
			stringReport = append(stringReport, ',')
		}

		stringReport = append(stringReport, '"')
		stringReport = append(stringReport, string(index)...)
		stringReport = append(stringReport, '"', ':')

		jsonReport := value.JSON()
		stringReport = append(stringReport, jsonReport...)

	}

	stringReport = append(stringReport, '}')
	return stringReport
}

func (report *Report) JSON() []byte {

	stringReport := []byte("{")

	for index, value := range *report {
		if index != 0 {
			stringReport = append(stringReport, ',')
		}
		stringReport = append(stringReport, '"')
		stringReport = append(stringReport, strconv.Itoa(index)...)
		stringReport = append(stringReport, '"', ':', '"')
		stringReport = append(stringReport, value.Marshal()...)
		stringReport = append(stringReport, '"')
	}

	stringReport = append(stringReport, '}')
	return stringReport

}
