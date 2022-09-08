package event_v1

import (
	"context"

	converter "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/converter/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// CreateEvent ...
func (i *Implementation) CreateEvent(ctx context.Context, req *desc.CreateEventRequest) (*desc.CreateEventResponse, error) {
	id, err := i.eventService.CreateEvent(ctx, converter.ToEventInfo(req.GetEventInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateEventResponse{
		Result: &desc.CreateEventResponse_Result{
			Id: id,
		},
	}, nil
}
