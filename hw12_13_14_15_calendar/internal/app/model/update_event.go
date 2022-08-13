package model

import (
	"database/sql"
	"time"
)

// UpdateEvent ...
type UpdateEvent struct {
	Title     sql.NullString `json:"title"`
	Date      time.Time      `json:"date"`
	Owner     sql.NullString `json:"owner"`
	CreatedAt *time.Time     `db:"created_at"`
}
