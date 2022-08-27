package repository

//go:generate mkdir -p mocks
//go:generate rm -rf ./mocks/*_minimock.go
//go:generate minimock -i EventRepository -o ./mocks/ -s "_minimock.go"

import (
	"context"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// EventRepository ...
type EventRepository interface {
	CreateEvent(ctx context.Context, eventInfo *model.EventInfo) error
	UpdateEvent(ctx context.Context, eventID int64, updateEventInfo *model.UpdateEventInfo) error
	DeleteEvent(ctx context.Context, eventID int64) error
	GetEventListForDay(ctx context.Context, date time.Time) ([]*model.Event, error)
	GetEventListForWeek(ctx context.Context, weekStart time.Time) ([]*model.Event, error)
	GetEventListForMonth(ctx context.Context, monthStart time.Time) ([]*model.Event, error)
	GetEventListByDate(ctx context.Context, startDate time.Time) ([]*model.Event, error)
}
