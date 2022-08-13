package event_v1

import (
	"context"

	converter "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/converter/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// GetEventListForMonth ...
func (i *Implementation) GetEventListForMonth(ctx context.Context, req *desc.GetEventListForMonthRequest) (*desc.GetEventListForMonthResponse, error) {
	res, err := i.eventService.GetEventListForMonth(ctx, req.GetMonthStart().AsTime())
	if err != nil {
		return nil, err
	}

	return &desc.GetEventListForMonthResponse{
		Result: &desc.GetEventListForMonthResponse_Result{
			Events: converter.ToDescEvents(res),
		},
	}, nil
}
