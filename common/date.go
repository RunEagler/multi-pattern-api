package common

import "time"

func Date(year, month, day int) time.Time {

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Datetime(year, month, day, hour, minute int) time.Time {

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}

func DateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func DatetimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
