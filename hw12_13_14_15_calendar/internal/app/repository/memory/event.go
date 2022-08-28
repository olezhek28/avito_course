package memory_repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/utils"
)

var currentID int64

type eventRepository struct {
	mu           sync.RWMutex
	eventsByIDs  map[int64]*model.Event
	eventsByDate map[time.Time]map[int64]*model.Event
}

// NewEventRepository ...
func NewEventRepository() repository.EventRepository {
	return &eventRepository{
		eventsByIDs:  make(map[int64]*model.Event),
		eventsByDate: make(map[time.Time]map[int64]*model.Event),
	}
}

// CreateEvent ...
func (r *eventRepository) CreateEvent(_ context.Context, eventInfo *model.EventInfo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	currentID++
	now := time.Now()

	r.eventsByIDs[currentID] = &model.Event{
		ID:        currentID,
		EventInfo: eventInfo,
		CreatedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
	}

	beginDay := utils.BeginningOfDay(eventInfo.StartDate.Time)
	if _, found := r.eventsByDate[beginDay]; !found {
		r.eventsByDate[beginDay] = make(map[int64]*model.Event)
	}

	r.eventsByDate[beginDay][currentID] = r.eventsByIDs[currentID]

	return nil
}

// UpdateEvent ...
func (r *eventRepository) UpdateEvent(_ context.Context, eventID int64, updateEventInfo *model.UpdateEventInfo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	oldBeginDay := utils.BeginningOfDay(r.eventsByIDs[eventID].EventInfo.StartDate.Time)
	updEvent := r.eventsByIDs[eventID]

	if updateEventInfo.Title.Valid {
		updEvent.EventInfo.Title = updateEventInfo.Title.String
	}
	if updateEventInfo.NotificationDate.Valid {
		updEvent.EventInfo.NotificationDate = updateEventInfo.NotificationDate
	}
	if updateEventInfo.StartDate.Valid {
		updEvent.EventInfo.StartDate = updateEventInfo.StartDate
	}
	if updateEventInfo.EndDate.Valid {
		updEvent.EventInfo.EndDate = updateEventInfo.EndDate
	}
	if updateEventInfo.NotificationInterval != nil {
		updEvent.EventInfo.NotificationInterval = updateEventInfo.NotificationInterval
	}
	if updateEventInfo.Description.Valid {
		updEvent.EventInfo.Description = updateEventInfo.Description
	}
	if updateEventInfo.OwnerID.Valid {
		updEvent.EventInfo.OwnerID = updateEventInfo.OwnerID.Int64
	}

	updEvent.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	beginDay := utils.BeginningOfDay(updEvent.EventInfo.StartDate.Time)
	if utils.BeginningOfDay(updEvent.EventInfo.StartDate.Time) != oldBeginDay {
		delete(r.eventsByDate[oldBeginDay], eventID)
		if len(r.eventsByDate[oldBeginDay]) == 0 {
			delete(r.eventsByDate, oldBeginDay)
		}

		if _, found := r.eventsByDate[beginDay]; !found {
			r.eventsByDate[beginDay] = make(map[int64]*model.Event)
		}

		r.eventsByDate[beginDay][eventID] = updEvent
	}

	r.eventsByIDs[eventID] = updEvent

	return nil
}

// DeleteEvent ...
func (r *eventRepository) DeleteEvent(_ context.Context, eventID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	beginDay := utils.BeginningOfDay(r.eventsByIDs[eventID].EventInfo.StartDate.Time)
	delete(r.eventsByDate[beginDay], eventID)
	delete(r.eventsByIDs, eventID)

	return nil
}

// GetEventListForDay ...
func (r *eventRepository) GetEventListForDay(_ context.Context, date time.Time) ([]*model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	beginDay := utils.BeginningOfDay(date)
	if _, found := r.eventsByDate[beginDay]; !found {
		return nil, fmt.Errorf("no events for day %s", date.Format(utils.DateLayout))
	}

	return utils.MapToSlice(r.eventsByDate[beginDay]), nil
}

// GetEventListForWeek ...
func (r *eventRepository) GetEventListForWeek(_ context.Context, weekStart time.Time) ([]*model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	beginWeek := utils.BeginningOfDay(weekStart)
	endWeek := beginWeek.AddDate(0, 0, utils.DaysInWeek)

	var events []*model.Event
	for date, event := range r.eventsByDate {
		if (date.After(beginWeek) && date.Before(endWeek)) || date.Equal(beginWeek) || date.Equal(endWeek) {
			events = append(events, utils.MapToSlice(event)...)
		}
	}

	return events, nil
}

// GetEventListForMonth	...
func (r *eventRepository) GetEventListForMonth(_ context.Context, monthStart time.Time) ([]*model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	beginMonth := utils.BeginningOfDay(monthStart)
	endMonth := beginMonth.AddDate(0, 0, utils.DaysInMonth)

	var events []*model.Event
	for date, event := range r.eventsByDate {
		if date.After(beginMonth) && date.Before(endMonth) || date.Equal(beginMonth) || date.Equal(endMonth) {
			events = append(events, utils.MapToSlice(event)...)
		}
	}

	return events, nil
}

// GetEventListByDate ...
func (r *eventRepository) GetEventListByDate(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Event, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteEventsBeforeDate ...
func (r *eventRepository) DeleteEventsBeforeDate(ctx context.Context, date time.Time) error {
	// TODO implement me
	panic("implement me")
}
