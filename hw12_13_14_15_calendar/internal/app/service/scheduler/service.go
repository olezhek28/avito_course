package scheduler

import (
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/repository"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

type Service struct {
	rabbitProducer rabbit.Producer

	eventRepository repository.EventRepository

	checkPeriod time.Duration
}

// NewService ...
func NewService(
	rabbitProducer rabbit.Producer,

	eventRepository repository.EventRepository,

	checkPeriod time.Duration,
) *Service {
	return &Service{
		rabbitProducer: rabbitProducer,

		eventRepository: eventRepository,

		checkPeriod: checkPeriod,
	}
}
