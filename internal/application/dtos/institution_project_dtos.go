package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionProjectRequest struct {
	InstitutionID   int64                                     `json:"institutionID" validate:"required,min=1"`
	LanguageCode    string                                    `json:"-" validate:"required,len=2"`
	ProjectType     string                                    `json:"projectType" validate:"required"`
	ProjectTitle    string                                    `json:"projectTitle" validate:"required"`
	DateStart       time.Time                                 `json:"dateStart" validate:"required"`
	DateEnd         time.Time                                 `json:"dateEnd" validate:"required"`
	Fund            float64                                   `json:"fund" validate:"required"`
	InstitutionRole string                                    `json:"institutionRole" validate:"required"`
	Coordinator     string                                    `json:"coordinator" validate:"required"`
	Partners        []*CreateInstitutionProjectPartnerRequest `json:"partner" validate:"omitempty,dive"`
}

type CreateInstitutionProjectPartnerRequest struct {
	LanguageCode  string `json:"languageCode" validate:"required,len=2"`
	PartnerType   string `json:"partnerType" validate:"required"`
	PartnerName   string `json:"partnerName" validate:"required"`
	LinkToPartner string `json:"linkToPartner" validate:"required"`
}

type UpdateInstitutionProjectRequest struct {
	ID              int64                                     `json:"id" validate:"required,min=1"`
	ProjectType     *string                                   `json:"projectType" validate:"omitempty"`
	ProjectTitle    *string                                   `json:"projectTitle" validate:"omitempty"`
	DateStart       *time.Time                                `json:"dateStart" validate:"omitempty"`
	DateEnd         *time.Time                                `json:"dateEnd" validate:"omitempty"`
	Fund            *float64                                  `json:"fund" validate:"omitempty"`
	InstitutionRole *string                                   `json:"institutionRole" validate:"omitempty"`
	Coordinator     *string                                   `json:"coordinator" validate:"omitempty"`
	Partners        []*UpdateInstitutionProjectPartnerRequest `json:"partners" validate:"omitempty,dive"`
}

type UpdateInstitutionProjectPartnerRequest struct {
	ID            int64   `json:"id" validate:"required,min=1"`
	PartnerType   *string `json:"partnerType" validate:"omitempty"`
	PartnerName   *string `json:"partnerName" validate:"omitempty"`
	LinkToPartner *string `json:"linkToPartner" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionProjectResponse struct {
	ID              int64                               `json:"id"`
	LanguageCode    string                              `json:"-"`
	ProjectType     string                              `json:"projectType"`
	ProjectTitle    string                              `json:"projectTitle"`
	DateStart       time.Time                           `json:"dateStart"`
	DateEnd         time.Time                           `json:"dateEnd"`
	Fund            float64                             `json:"fund"`
	InstitutionRole string                              `json:"institutionRole"`
	Coordinator     string                              `json:"coordinator"`
	Partners        []*InstitutionProjectPartnerResponse `json:"partners,omitempty"`
	CreatedAt       time.Time                           `json:"createdAt"`
	UpdatedAt       time.Time                           `json:"updatedAt"`
}

type InstitutionProjectPartnerResponse struct {
	ID            int64    `json:"id"`
	PartnerType   string    `json:"partnerType"`
	PartnerName   string    `json:"partnerName"`
	LinkToPartner string    `json:"linkToPartner"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
