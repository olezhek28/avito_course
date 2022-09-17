package model

import (
	"database/sql"
	"time"
)

// UpdateEventInfo ...
type UpdateEventInfo struct {
	Title                sql.NullString `db:"title"`
	NotificationDate     sql.NullTime   `db:"notification_date"`
	StartDate            sql.NullTime   `db:"start_date"`
	EndDate              sql.NullTime   `db:"end_date"`
	NotificationInterval *time.Duration `db:"notification_interval"`
	Description          sql.NullString `db:"description"`
	OwnerID              sql.NullInt64  `db:"owner_id"`
}
