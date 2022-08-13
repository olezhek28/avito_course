package event_v1

import (
	"context"
	"database/sql"

	converter "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/converter/event_v1"
	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdateEvent ...
func (i *Implementation) UpdateEvent(ctx context.Context, req *desc.UpdateEventRequest) (*emptypb.Empty, error) {
	err := i.eventService.UpdateEvent(ctx,
		sql.NullInt64{
			Int64: req.GetId().GetValue(),
			Valid: req.GetId() != nil,
		},
		converter.ToUpdateEvent(req.GetUpdateEventInfo()),
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
