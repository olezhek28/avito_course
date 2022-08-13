package event

import (
	"context"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// GetEventListForMonth ...
func (s *Service) GetEventListForMonth(ctx context.Context, date time.Time) ([]*model.Event, error) {
	return s.eventRepository.GetEventListForMonth(ctx, date)
}
