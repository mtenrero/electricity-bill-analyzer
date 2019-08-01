package analyzers

type ResponseReport struct {
	WeekDaysReport     Report      `json:"weekDaysReport"`
	HourlyReport       Report      `json:"hourlyReport"`
	WeekDayHoursReport ReportGroup `json:"weekdayHoursReport"`
}

type Report struct {
	Analysis []ReportValue `json:"analysis"`
}

type ReportGroup []Report

type ReportValue struct {
	Hour    int     `json:"hour,omitempty"`
	WeekDay int     `json:"weekDay,omitempty"`
	Mean    float64 `json:"mean"`
}
