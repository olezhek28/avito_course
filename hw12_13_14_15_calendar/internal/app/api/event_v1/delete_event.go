package event_v1

import (
	"context"
	"database/sql"

	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteEvent ...
func (i *Implementation) DeleteEvent(ctx context.Context, req *desc.DeleteEventRequest) (*emptypb.Empty, error) {
	err := i.eventService.DeleteEvent(ctx, sql.NullInt64{
		Int64: req.GetId().GetValue(),
		Valid: req.GetId() != nil,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
