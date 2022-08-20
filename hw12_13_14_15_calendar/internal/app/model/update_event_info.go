package model

import (
	"database/sql"
)

// UpdateEventInfo ...
type UpdateEventInfo struct {
	Title                   sql.NullString `db:"title"`
	StartDate               sql.NullTime   `db:"start_date"`
	EndDate                 sql.NullTime   `db:"end_date"`
	NotificationIntervalMin sql.NullInt64  `db:"notification_interval_min"`
	Description             sql.NullString `db:"description"`
	OwnerID                 sql.NullInt64  `db:"owner_id"`
}
