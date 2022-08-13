package event_v1

import (
	"context"

	converter "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/converter/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// GetEventListForDay ...
func (i *Implementation) GetEventListForDay(ctx context.Context, req *desc.GetEventListForDayRequest) (*desc.GetEventListForDayResponse, error) {
	res, err := i.eventService.GetEventListForDay(ctx, req.GetDate().AsTime())
	if err != nil {
		return nil, err
	}

	return &desc.GetEventListForDayResponse{
		Result: &desc.GetEventListForDayResponse_Result{
			Events: converter.ToDescEvents(res),
		},
	}, nil
}
