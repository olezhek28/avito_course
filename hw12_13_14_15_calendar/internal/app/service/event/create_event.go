package event

import (
	"context"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
)

// CreateEvent ...
func (s *Service) CreateEvent(ctx context.Context, event *model.Event) error {
	return s.eventRepository.CreateEvent(ctx, event)
}
