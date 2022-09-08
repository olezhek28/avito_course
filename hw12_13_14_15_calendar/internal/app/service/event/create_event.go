package event

import (
	"context"
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// CreateEvent ...
func (s *Service) CreateEvent(ctx context.Context, eventInfo *model.EventInfo) (int64, error) {
	eventInfo.NotificationDate = sql.NullTime{
		Time:  eventInfo.StartDate.Time,
		Valid: eventInfo.StartDate.Valid,
	}
	if eventInfo.NotificationInterval != nil {
		eventInfo.NotificationDate.Time = eventInfo.StartDate.Time.Add(-*eventInfo.NotificationInterval)
	}

	return s.eventRepository.CreateEvent(ctx, eventInfo)
}
