package model

import "time"

// Todo (an Entity)
type Todo struct {
	ID      int       `json:id`
	Title   string    `json:title`
	Content string    `json:content`
	Due     time.Time `json:due`
}
