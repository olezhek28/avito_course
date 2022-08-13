package event_v1

import (
	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/service/event"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedEventServiceV1Server

	eventService *event.Service
}

// NewEventV1 return new instance of Implementation.
func NewEventV1(eventService *event.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedEventServiceV1Server{},

		eventService,
	}
}
