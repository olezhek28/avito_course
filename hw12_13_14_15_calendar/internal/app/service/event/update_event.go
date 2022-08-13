package event

import (
	"context"
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateEvent ...
func (s *Service) UpdateEvent(ctx context.Context, eventID sql.NullInt64, updateEvent *model.UpdateEvent) error {
	if updateEvent == nil {
		return status.Error(codes.InvalidArgument, "parameters for updating are not set")
	}
	if !eventID.Valid {
		return status.Error(codes.InvalidArgument, "eventID is empty")
	}

	return s.eventRepository.UpdateEvent(ctx, eventID.Int64, updateEvent)
}
