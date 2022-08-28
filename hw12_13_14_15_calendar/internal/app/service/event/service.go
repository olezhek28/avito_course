package event

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
)

// Service ...
type Service struct {
	logger          *logger.Logger
	eventRepository repository.EventRepository
}

// NewService ...
func NewService(
	logger *logger.Logger,
	eventRepository repository.EventRepository,
) *Service {
	return &Service{
		logger:          logger,
		eventRepository: eventRepository,
	}
}
