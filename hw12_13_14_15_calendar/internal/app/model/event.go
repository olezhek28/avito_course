package model

import "time"

// Event ...
type Event struct {
	ID    int64     `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Owner string    `json:"owner"`
}
