package dtos

import "time"

// ---- REQUEST DTOs ----


// ---- RESPONSE DTOs ----

type UserResponse struct {
	ID                 int64     `json:"id"`
	Email              string    `json:"email"`
	PasswordHash       string    `json:"-"`
	UserType           string    `json:"userType"`
	EntityID           time.Time `json:"entityID"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
