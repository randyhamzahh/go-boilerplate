package model

import (
	"time"
)

type User struct {
	ID        uint      // Standard field for the primary key
	Name      string    // A regular string field
	Email     string    // A pointer to a string, allowing for null values
	Password  string    // A pointer to a string, allowing for null values
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
