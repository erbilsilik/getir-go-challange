package utilities

import (
	"time"
)

func ParseDate(layout string, dateString string) time.Time {
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		panic("parsing date error")
	}
	return parsedDate
}

