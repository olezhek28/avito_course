package memory_repository

import (
	"context"
	"sync"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
)

type eventRepository struct {
	mu sync.RWMutex
}

// NewEventRepository ...
func NewEventRepository() repository.EventRepository {
	return &eventRepository{}
}

func (r *eventRepository) CreateEvent(ctx context.Context, event *model.EventInfo) error {
	// TODO implement me
	panic("implement me")
}

func (r *eventRepository) UpdateEvent(ctx context.Context, eventID int64, updateEvent *model.UpdateEvent) error {
	// TODO implement me
	panic("implement me")
}

func (r *eventRepository) DeleteEvent(ctx context.Context, eventID int64) error {
	// TODO implement me
	panic("implement me")
}

func (r *eventRepository) GetEventListForDay(ctx context.Context, date time.Time) ([]*model.Event, error) {
	// TODO implement me
	panic("implement me")
}

func (r *eventRepository) GetEventListForWeek(ctx context.Context, weekStart time.Time) ([]*model.Event, error) {
	// TODO implement me
	panic("implement me")
}

func (r *eventRepository) GetEventListForMonth(ctx context.Context, monthStart time.Time) ([]*model.Event, error) {
	// TODO implement me
	panic("implement me")
}
