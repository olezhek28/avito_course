package event

import (
	"context"
	"database/sql"
)

// DeleteEvent ...
func (s *Service) DeleteEvent(ctx context.Context, eventID sql.NullInt64) error {
	return nil
}
