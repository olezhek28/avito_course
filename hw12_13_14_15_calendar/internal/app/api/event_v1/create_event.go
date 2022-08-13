package event_v1

import (
	"context"

	desc "github.com/olezhek28/avito_course/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateEvent ...
func (i *Implementation) CreateEvent(ctx context.Context, req *desc.CreateEventRequest) (*emptypb.Empty, error) {
	return nil, nil
}

