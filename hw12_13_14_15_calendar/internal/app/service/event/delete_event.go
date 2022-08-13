package event

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteEvent ...
func (s *Service) DeleteEvent(ctx context.Context, eventID sql.NullInt64) error {
	if !eventID.Valid {
		return status.Error(codes.InvalidArgument, "eventID is empty")
	}

	return s.eventRepository.DeleteEvent(ctx, eventID.Int64)
}
