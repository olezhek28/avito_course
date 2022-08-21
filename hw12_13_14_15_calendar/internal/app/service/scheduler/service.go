package scheduler

import "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"

type Service struct {
	rabbitProducer rabbit.Producer
}

// NewService ...
func NewService(rabbitProducer rabbit.Producer) *Service {
	return &Service{
		rabbitProducer: rabbitProducer,
	}
}
