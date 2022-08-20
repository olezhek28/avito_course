package model

import (
	"database/sql"
)

// UpdateEventInfo ...
type UpdateEventInfo struct {
	Title sql.NullString `db:"title"`
	Date  sql.NullTime   `db:"date"`
	Owner sql.NullString `db:"owner"`
}
