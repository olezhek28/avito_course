package event_v1

import (
	"context"

	converter "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/converter/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
)

// GetEventListForWeek ...
func (i *Implementation) GetEventListForWeek(ctx context.Context, req *desc.GetEventListForWeekRequest) (*desc.GetEventListForWeekResponse, error) {
	res, err := i.eventService.GetEventListForWeek(ctx, req.GetWeekStart().AsTime())
	if err != nil {
		return nil, err
	}

	return &desc.GetEventListForWeekResponse{
		Result: &desc.GetEventListForWeekResponse_Result{
			Events: converter.ToDescEvents(res),
		},
	}, nil
}
