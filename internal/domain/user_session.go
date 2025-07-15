package domain

import "time"

type UserSession struct {
	ID           int64
	UserID       int64
	RefreshToken string
	ExpiresAt   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
