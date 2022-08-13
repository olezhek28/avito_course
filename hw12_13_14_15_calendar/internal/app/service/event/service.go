package event

import "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"

type Service struct {
	eventRepository repository.EventRepository
}

// NewService ...
func NewService(
	eventRepository repository.EventRepository,
) *Service {
	return &Service{
		eventRepository: eventRepository,
	}
}
