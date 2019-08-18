package analyzers

type ResponseReport struct {
	WeekDaysReport     Report                `json:"weekDaysReport"`
	HourlyReport       Report                `json:"hourlyReport"`
	WeekDayHoursReport []ReportWeekDayHourly `json:"weekdayHoursReport"`
}

type Report struct {
	Analysis []ReportValue `json:"analysis"`
}

type ReportWeekDayHourly struct {
	WeekDay    *int    `json:"weekDay,omitempty"`
	WeekString *string `json:"weekString,omitempty`
	Analysis   Report  `json:"report,omitempty"`
}

type ReportValue struct {
	Hour       *int     `json:"hour,omitempty"`
	WeekDay    *int     `json:"weekDay,omitempty"`
	WeekString *string  `json:"weekString,omitempty"`
	Mean       *float64 `json:"mean,omitempty"`
}
