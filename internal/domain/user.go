package domain

import "time"

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	UserType     string
	EntityID     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
