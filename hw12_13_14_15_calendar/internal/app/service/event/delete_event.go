package event

import (
	"context"
	"database/sql"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/app/model/err"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteEvent ...
func (s *Service) DeleteEvent(ctx context.Context, eventID sql.NullInt64) error {
	if !eventID.Valid {
		return status.Error(codes.InvalidArgument, err.ErrInvalidEventID)
	}

	return s.eventRepository.DeleteEvent(ctx, eventID.Int64)
}
