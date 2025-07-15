package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionRankingRequest struct {
	InstitutionID           int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode            string    `json:"-" validate:"required,len=2"`
	RankingTitle            string    `json:"rankingTitle" validate:"required"`
	RankingType             string    `json:"rankingType" validate:"required"`
	DateReceived            time.Time `json:"dateReceived" validate:"required"`
	RankingAgency           string    `json:"rankingAgency" validate:"required"`
	Description             string    `json:"description" validate:"required"`
	LinkToRankingFile       string    `json:"linkToRankingFile" validate:"required"`
	LinkToRankingAgencyFile string    `json:"linkToRankingAgencyFile" validate:"required"`
}

type UpdateInstitutionRankingRequest struct {
	ID                      int64      `json:"id" validate:"required,min=1"`
	RankingTitle            *string    `json:"rankingTitle" validate:"omitempty"`
	RankingType             *string    `json:"rankingType" validate:"omitempty"`
	DateReceived            *time.Time `json:"dateReceived" validate:"omitempty"`
	RankingAgency           *string    `json:"rankingAgency" validate:"omitempty"`
	Description             *string    `json:"description" validate:"omitempty"`
	LinkToRankingFile       *string    `json:"linkToRankingFile" validate:"omitempty"`
	LinkToRankingAgencyFile *string    `json:"linkToRankingAgencyFile" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionRankingResponse struct {
	ID                      int64     `json:"id"`
	RankingTitle            string    `json:"rankingTitle"`
	RankingType             string    `json:"rankingType"`
	DateReceived            time.Time `json:"dateReceived"`
	RankingAgency           string    `json:"rankingAgency"`
	Description             string    `json:"description"`
	LinkToRankingFile       string    `json:"linkToRankingFile"`
	LinkToRankingAgencyFile string    `json:"linkToRankingAgencyFile"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}
