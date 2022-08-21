package sender

import "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"

type Service struct {
	rabbitConsumer rabbit.Consumer
}

// NewService ...
func NewService(rabbitConsumer rabbit.Consumer) *Service {
	return &Service{
		rabbitConsumer: rabbitConsumer,
	}
}
