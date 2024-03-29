package utils

import (
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// DT common values ...
const (
	beginningOfDayHour   = 0
	beginningOfDayMinute = 0
	beginningOfDaySecond = 0

	endOfDayHour   = 23
	endOfDayMinute = 59
	endOfDaySecond = 59

	dateTimeLayout = "2006-01-02T15:04:05"
)

// Date ...
const (
	DateLayout  = "2006-01-02"
	DaysInWeek  = 7
	DaysInMonth = 30
)

// BeginningOfDay returns the beginning of the day (00:00:00:00000) ...
func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d,
		beginningOfDayHour, beginningOfDayMinute, beginningOfDaySecond, beginningOfDaySecond, t.Location())
}

// EndOfDay returns the end of the day (23:59:59:99999) ...
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d,
		endOfDayHour, endOfDayMinute, endOfDaySecond, int(time.Second-time.Nanosecond), t.Location())
}

// MapToSlice ...
func MapToSlice(m map[int64]*model.Event) []*model.Event {
	events := make([]*model.Event, 0, len(m))
	for _, event := range m {
		events = append(events, event)
	}

	return events
}

// RoundUpToMinutes ...
func RoundUpToMinutes(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, date.Location())
}
