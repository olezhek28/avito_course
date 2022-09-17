package scheduler

import (
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

// Service ...
type Service struct {
	logger          *logger.Logger
	rabbitProducer  rabbit.Producer
	eventRepository repository.EventRepository
	checkPeriod     time.Duration
}

// NewService ...
func NewService(
	logger *logger.Logger,
	rabbitProducer rabbit.Producer,
	eventRepository repository.EventRepository,
	checkPeriod time.Duration,
) *Service {
	return &Service{
		logger:          logger,
		rabbitProducer:  rabbitProducer,
		eventRepository: eventRepository,
		checkPeriod:     checkPeriod,
	}
}
