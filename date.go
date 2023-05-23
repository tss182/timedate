package date

import (
	"time"
)

const (
	dateFormat = "2006-01-02"
	Week       = "week"
	Month      = "month"
)

func beginOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func GetStartEndDate(date time.Time, optionDate string) (startDate string, endDate string) {
	switch optionDate {
	case Week:
		weekday := int(date.Weekday())
		startDate = date.AddDate(0, 0, -weekday).Format(dateFormat)
		endDate = date.AddDate(0, 0, 7-weekday-1).Format(dateFormat)
	case Month:
		startDate = beginOfMonth(date).Format(dateFormat)
		endDate = endOfMonth(date).Format(dateFormat)
	default:
		startDate = date.Format(dateFormat)
		endDate = date.Format(dateFormat)
	}
	return
}
