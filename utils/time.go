package utils

import "time"

func ParseDate(dateString string) time.Time {
	parseDateFormat := "2006-01-02"
	parsedDate, err := time.Parse(parseDateFormat, dateString)
	if err != nil {
		panic("parsing date error")
	}
	return parsedDate
}
