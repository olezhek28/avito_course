package model

import (
	"database/sql"
	"time"
)

// EventInfo ...
type EventInfo struct {
	Title                string         `db:"title" json:"title"`
	NotificationDate     sql.NullTime   `db:"notification_date" json:"notification_date"`
	StartDate            sql.NullTime   `db:"start_date" json:"start_date"`
	EndDate              sql.NullTime   `db:"end_date" json:"end_date"`
	NotificationInterval *time.Duration `db:"notification_interval" json:"notification_interval"`
	Description          sql.NullString `db:"description" json:"description"`
	OwnerID              int64          `db:"owner_id" json:"owner_id"`
}

// Event ...
type Event struct {
	ID        int64        `db:"id" json:"id"`
	EventInfo *EventInfo   `db:""`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}
