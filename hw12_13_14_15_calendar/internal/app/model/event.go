package model

import "time"

// EventInfo ...
// type EventInfo struct {
//	Title     string     `db:"title"`
//	Date      *time.Time `db:"date"`
//	Owner     string     `db:"owner"`
//	CreatedAt *time.Time `db:"created_at"`
// }

// Event ...
type Event struct {
	ID int64 `db:"id"`
	// EventInfo EventInfo
	Title     string     `db:"title"`
	Date      *time.Time `db:"date"`
	Owner     string     `db:"owner"`
	CreatedAt *time.Time `db:"created_at"`
}
