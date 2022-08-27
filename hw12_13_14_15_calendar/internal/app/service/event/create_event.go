package event

import (
	"context"
	"database/sql"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/utils"
)

// CreateEvent ...
func (s *Service) CreateEvent(ctx context.Context, eventInfo *model.EventInfo) error {
	var startDate, endDate time.Time

	if eventInfo.StartDate.Valid {
		startDate = utils.RoundUpToMinutes(eventInfo.StartDate.Time)
	}
	if eventInfo.EndDate.Valid {
		endDate = utils.RoundUpToMinutes(eventInfo.EndDate.Time)
	}

	eventInfo.StartDate = sql.NullTime{
		Time:  startDate,
		Valid: eventInfo.StartDate.Valid,
	}

	eventInfo.EndDate = sql.NullTime{
		Time:  endDate,
		Valid: eventInfo.EndDate.Valid,
	}

	eventInfo.NotificationDate = sql.NullTime{
		Time:  startDate.Add(-*eventInfo.NotificationInterval),
		Valid: eventInfo.StartDate.Valid,
	}

	return s.eventRepository.CreateEvent(ctx, eventInfo)
}
