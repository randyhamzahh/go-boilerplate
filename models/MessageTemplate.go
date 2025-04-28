package model

import "time"

type MessageTemplate struct {
	ID        int
	Name      string
	Message   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
