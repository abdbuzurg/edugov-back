package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeParticipationInEventRequest struct {
	EmployeeID                     int64     `json:"employeeID" validate:"required,min=1"`
	LanguageCode                   string    `json:"-" validate:"required,len=2"`
	EventTitle                     string    `json:"eventTitle" validate:"required"`
	EventDate                      time.Time `json:"eventDate" validate:"required"`
}

type UpdateEmployeeParticipationInEventRequest struct {
	ID                             int64      `json:"id" validate:"required,min=1"`
	EventTitle                     *string    `json:"eventTitle" validate:"omitempty"`
	EventDate                      *time.Time `json:"eventDate" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeParticipationInEventResponse struct {
	ID                             int64     `json:"id"`
	EventTitle                     string    `json:"eventTitle"`
	EventDate                      time.Time `json:"eventDate"`
	CreatedAt                      time.Time `json:"createdAt"`
	UpdatedAt                      time.Time `json:"updatedAt"`
}
