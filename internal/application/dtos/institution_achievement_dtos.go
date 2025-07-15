package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionAchievementRequest struct {
	InstitutionID    int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode     string    `json:"-" validate:"required,len=2"`
	AchievementTitle string    `json:"achievementTitle" validate:"required"`
	AchievementType  string    `json:"achievementType" validate:"required"`
	DateReceived     time.Time `json:"dateReceived" validate:"required"`
	GivenBy          string    `json:"givenBy" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	LinkToFile       string    `json:"linkToFile" validate:"required"`
}

type UpdateInstitutionAchievementRequest struct {
	ID               int64      `json:"id" validate:"required,min=1"`
	AchievementTitle *string    `json:"achievementTitle" validate:"omitempty"`
	AchievementType  *string    `json:"achievementType" validate:"omitempty"`
	DateReceived     *time.Time `json:"dateReceived" validate:"omitempty"`
	GivenBy          *string    `json:"givenBy" validate:"omitempty"`
	Description      *string    `json:"description" validate:"omitempty"`
	LinkToFile       *string    `json:"linkToFile" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionAchievementResponse struct {
	ID               int64     `json:"id"`
	AchievementTitle string    `json:"achievementTitle"`
	AchievementType  string    `json:"achievementType"`
	DateReceived     time.Time `json:"dateReceived"`
	GivenBy          string    `json:"givenBy"`
	Description      string    `json:"description"`
	LinkToFile       string    `json:"linkToFile"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
