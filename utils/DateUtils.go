package utils

import (
	"time"
)

var DATE_FORMAT_DATE string = "2019-01-01"

func DateFormatGetTimeByDateInLocation(datestr string) time.Time {
	time, err := time.ParseInLocation(DATE_FORMAT_DATE, datestr, time.Local)
	if err != nil {
		LogError(err)
	}
	return time
}

func DateFormatGetStrByISO8601Time(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
