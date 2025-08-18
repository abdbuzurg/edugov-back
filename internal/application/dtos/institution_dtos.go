package dtos

import "time"

// --- Request DTOs ---

// CreateInstitutionRequest represents the input for creating a new Institution
// Note - the institution details are required with creating a new Institution
type CreateInstitutionRequest struct {
	YearOfEstablishment int32  `json:"yearOfEstablishment" validate:"required,gte=1000,lte=2100"`
	Email               string `json:"email" validate:"required,email"`
	Fax                 string `json:"fax" validate:"required"`
	OfficialWebsite     string `json:"officialWebsite" validate:"required"`
	PhoneNumber         string `json:"phoneNumber" validate:"required"`
	MailIndex           string `json:"mailIndex" validate:"required"`
}

// UpdateInstitutionRequest represents the input for uptaing existing institution.
// Note - the input will also might have update to institution details.
type UpdateInstitutionRequest struct {
	ID                  int64   `json:"id" validate:"required,min=1"`
	YearOfEstablishment *int32  `json:"yearOfEstablishment,omitempty"`
	Email               *string `json:"email,omitempty"`
	Fax                 *string `json:"fax,omitempty"`
	OfficialWebsite     *string `json:"officialWebsite" validate:"omitempty"`
	PhoneNumber         *string `json:"phoneNumber" validate:"omitempty"`
	MailIndex           *string `json:"mailIndex" validate:"omitempty"`
}

// --- Response DTOs ---
// InstitutionResponse represents the output for Institution (fully if filled)
// Note - InstitutionDetails field should always be filled
type InstitutionResponse struct {
	ID                  int64     `json:"id"`
	YearOfEstablishment int32     `json:"yearOfEstablishment"`
	Email               string    `json:"email"`
	Fax                 string    `json:"fax"`
	PhoneNumber         string    `json:"phoneNumber"`
	MailIndex           string    `json:"mailIndex"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type AllInstitutionResponse struct {
	InstitutitonTitleShort string `json:"institutionTitleShort"`
	InstitutitonTitleLong  string `json:"institutionTitleLong"`
	MailIndex              string `json:"mainIndex"`
	City                   string `json:"city"`
	Address                string `json:"address"`
	OfficialWebsite        string `json:"officialWebsite"`
	Email                  string `json:"email"`
	PhoneNumber            string `json:"phoneNumber"`
}
