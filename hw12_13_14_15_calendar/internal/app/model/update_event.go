package model

import (
	"database/sql"
)

// UpdateEvent ...
type UpdateEvent struct {
	Title sql.NullString `db:"title"`
	Date  sql.NullTime   `db:"date"`
	Owner sql.NullString `db:"owner"`
}
