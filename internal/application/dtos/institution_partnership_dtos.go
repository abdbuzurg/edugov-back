package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionPartnershipRequest struct {
	InstitutionID  int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode   string    `json:"-" validate:"required,len=2"`
	PartnerName    string    `json:"partnerName" validate:"required"`
	PartnerType    string    `json:"partnerType" validate:"required"`
	Goal           string    `json:"goal" validate:"required"`
	LinkToPartner  string    `json:"linkToPartner" validate:"required"`
	DateOfContract time.Time `json:"dateOfContract" validate:"required"`
}

type UpdateInstitutionPartnershipRequest struct {
	ID             int64      `json:"id" validate:"required,min=1"`
	PartnerName    *string    `json:"partnerName" validate:"omitempty"`
	PartnerType    *string    `json:"partnerType" validate:"omitempty"`
	Goal           *string    `json:"goal" validate:"omitempty"`
	LinkToPartner  *string    `json:"linkToPartner" validate:"omitempty"`
	DateOfContract *time.Time `json:"dateOfContract" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionPartnershipResponse struct {
	ID             int64     `json:"id"`
	PartnerName    string    `json:"partnerName"`
	PartnerType    string    `json:"partnerType"`
	Goal           string    `json:"goal"`
	LinkToPartner  string    `json:"linkToPartner"`
	DateOfContract time.Time `json:"dateOfContract"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
