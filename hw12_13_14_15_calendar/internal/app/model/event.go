package model

import "time"

// EventInfo ...
type EventInfo struct {
	Title string     `db:"title"`
	Date  *time.Time `db:"date"`
	Owner string     `db:"owner"`
}

// Event ...
type Event struct {
	ID        int64      `db:"id"`
	EventInfo *EventInfo `db:""`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
