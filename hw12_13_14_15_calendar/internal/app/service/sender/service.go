package sender

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/logger"
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"
)

// Service ...
type Service struct {
	logger         *logger.Logger
	rabbitConsumer rabbit.Consumer
}

// NewService ...
func NewService(logger *logger.Logger, rabbitConsumer rabbit.Consumer) *Service {
	return &Service{
		logger:         logger,
		rabbitConsumer: rabbitConsumer,
	}
}
