package utils

import "time"

func ParseDate(dateString string) (time.Time, error) {
	parseDateFormat := "2006-01-02"
	parsedDate, err := time.Parse(parseDateFormat, dateString)
	if err != nil {
		return parsedDate, err
	}
	return parsedDate, nil
}
