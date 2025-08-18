package mappers

import (
	"backend/internal/application/dtos"
	"backend/internal/domain"
)

func MapInstitutionDomainToResponseDTO(institution *domain.Institution) *dtos.InstitutionResponse {
	if institution == nil {
		return nil
	}

	return &dtos.InstitutionResponse{
		ID:                  institution.ID,
		YearOfEstablishment: institution.YearOfEstablishment,
		Email:               institution.Email,
		Fax:                 institution.Fax,
		PhoneNumber:         institution.PhoneNumber,
		MailIndex:           institution.MailIndex,
		CreatedAt:           institution.CreatedAt,
		UpdatedAt:           institution.UpdatedAt,
	}
}
