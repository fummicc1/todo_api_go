package main

import "time"

// Todo (an Entity)
type Todo struct {
	ID      string
	Title   string
	Content string
	Due     time.Time
}
