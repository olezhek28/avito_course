package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_event_repository.go -package=mocks . EventRepository

import (
	"context"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// EventRepository ...
type EventRepository interface {
	CreateEvent(ctx context.Context, event *model.Event) error
	UpdateEvent(ctx context.Context, eventID int64, updateEvent *model.UpdateEvent) error
	DeleteEvent(ctx context.Context, eventID int64) error
	GetEventListForDay(ctx context.Context, date time.Time) ([]*model.Event, error)
	GetEventListForWeek(ctx context.Context, weekStart time.Time) ([]*model.Event, error)
	GetEventListForMonth(ctx context.Context, monthStart time.Time) ([]*model.Event, error)
}
