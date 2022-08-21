package model

import (
	"database/sql"
	"time"
)

// EventInfo ...
type EventInfo struct {
	Title                   string         `db:"title", json:"title"`
	StartDate               *time.Time     `db:"start_date", json:"start_date"`
	EndDate                 *time.Time     `db:"end_date", json:"end_date"`
	NotificationIntervalMin sql.NullInt64  `db:"notification_interval_min", json:"notification_interval_min"`
	Description             sql.NullString `db:"description", json:"description"`
	OwnerID                 int64          `db:"owner_id", json:"owner_id"`
}

// Event ...
type Event struct {
	ID        int64      `db:"id", json:"id"`
	EventInfo *EventInfo `db:""`
	CreatedAt *time.Time `db:"created_at", json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at", json:"updated_at"`
}
