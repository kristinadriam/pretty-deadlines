package models

import "time"

type Deadline struct {
	Title       string
	Description string
	DueDate     time.Time
}
