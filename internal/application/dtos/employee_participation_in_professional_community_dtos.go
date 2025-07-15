package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeParticipationInProfessionalCommunityRequest struct {
	EmployeeID                  int64  `json:"employeeID" validate:"required,min=1"`
	LanguageCode                string `json:"-" validate:"required,len=2"`
	ProfessionalCommunityTitle  string `json:"professionalCommunityTitle" validate:"required"`
	RoleInProfessionalCommunity string `json:"roleInProfessionalCommunity" validate:"required"`
}

type UpdateEmployeeParticipationInProfessionalCommunityRequest struct {
	ID                          int64   `json:"id" validate:"required,min=1"`
	ProfessionalCommunityTitle  *string `json:"professionalCommunityTitle" validate:"omitempty"`
	RoleInProfessionalCommunity *string `json:"roleInProfessionalCommunity" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeParticipationInProfessionalCommunityResponse struct {
	ID                          int64     `json:"id"`
	ProfessionalCommunityTitle  string    `json:"professionalCommunityTitle"`
	RoleInProfessionalCommunity string    `json:"roleInProfessionalCommunity"`
	CreatedAt                   time.Time `json:"createdAt"`
	UpdatedAt                   time.Time `json:"updatedAt"`
}
