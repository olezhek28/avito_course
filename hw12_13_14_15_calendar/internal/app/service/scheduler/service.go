package scheduler

import "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/pkg/rabbit"

type Service struct {
	rabbitClient rabbit.Client
}

// NewService ...
func NewService(rabbitClient rabbit.Client) *Service {
	return &Service{
		rabbitClient: rabbitClient,
	}
}
