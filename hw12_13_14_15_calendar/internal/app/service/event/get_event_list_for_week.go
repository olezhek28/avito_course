package event

import (
	"context"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// GetEventListForWeek ...
func (s *Service) GetEventListForWeek(ctx context.Context, date time.Time) ([]*model.Event, error) {
	return s.eventRepository.GetEventListForWeek(ctx, date)
}
