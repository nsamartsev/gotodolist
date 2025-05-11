package main

import "time"

type Todo struct {
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
