package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/shared/custom_errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type InstitutionUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionRequest) (*dtos.InstitutionResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionRequest) (*dtos.InstitutionResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, employeeID int64, langCode string) (*dtos.InstitutionResponse, error)
}

type institutionUsecase struct {
	institutionRepo repositories.InstitutionRepository
	validator       *validator.Validate
}

func NewInstitutionUsecase(
	institutionRepo repositories.InstitutionRepository,
	validator *validator.Validate,
) InstitutionUsecase {
	return &institutionUsecase{
		institutionRepo: institutionRepo,
		validator:       validator,
	}
}

func (uc *institutionUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionRequest) (*dtos.InstitutionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institutin: %w", err))
	}

	institution := &domain.Institution{
		YearOfEstablishment: req.YearOfEstablishment,
		Email:               req.Email,
		Fax:                 req.Fax,
		OfficialWebsite:     req.OfficialWebsite,
		Details: &domain.InstitutionDetails{
			LanguageCode:     req.Details.LanguageCode,
			InstitutionTitle: req.Details.InstitutionTitle,
			InstitutionType:  req.Details.InstitutionType,
			LegalStatus:      req.Details.LegalAddress,
			Mission:          req.Details.Mission,
			Founder:          req.Details.Founder,
			LegalAddress:     req.Details.LegalAddress,
		},
	}

	if req.Details.FactualAddress != nil {
		institution.Details.FactualAddress = req.Details.FactualAddress
	}

	createdInstitution, err := uc.institutionRepo.Create(ctx, institution)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionDomainToResponseDTO(createdInstitution)
	return resp, nil
}

func (uc *institutionUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionRequest) (*dtos.InstitutionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institutin: %w", err))
	}

	institution := &domain.Institution{
		ID: req.ID,
	}

	if req.YearOfEstablishment != nil {
		institution.YearOfEstablishment = *req.YearOfEstablishment
	}

	if req.Email != nil {
		institution.Email = *req.Email
	}

	if req.Fax != nil {
		institution.Fax = *req.Fax
	}

	if req.OfficialWebsite != nil {
		institution.OfficialWebsite = *req.OfficialWebsite
	}

	if req.Details != nil {
		details := &domain.InstitutionDetails{
			ID:            req.Details.ID,
			InstitutionID: req.ID,
		}

		if req.Details.InstitutionTitle != nil {
			details.InstitutionTitle = *req.Details.InstitutionTitle
		}

		if req.Details.InstitutionType != nil {
			details.InstitutionType = *req.Details.InstitutionType

		}
		if req.Details.LegalStatus != nil {
			details.LegalStatus = *req.Details.LegalStatus
		}

		if req.Details.Mission != nil {
			details.Mission = *req.Details.Mission
		}

		if req.Details.Founder != nil {
			details.Founder = *req.Details.Founder
		}

		if req.Details.LegalAddress != nil {
			details.LegalAddress = *req.Details.LegalAddress
		}

		if req.Details.FactualAddress != nil {
			details.FactualAddress = req.Details.FactualAddress
		}

    institution.Details = details
	}

	updatedInstitution, err := uc.institutionRepo.Update(ctx, institution)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionDomainToResponseDTO(updatedInstitution)
	return resp, nil
}

func (uc *institutionUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institutin"))
	}

	if err := uc.institutionRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionUsecase) GetByID(ctx context.Context, id int64, langCode string) (*dtos.InstitutionResponse, error) {
	institution, err := uc.institutionRepo.GetByID(ctx, id, langCode)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionDomainToResponseDTO(institution)
	return resp, nil
}

func MapInstitutionDomainToResponseDTO(institution *domain.Institution) *dtos.InstitutionResponse {
	if institution == nil {
		return nil
	}

	var details *dtos.InstitutionDetailsResponse
	if institution.Details != nil {
		details = &dtos.InstitutionDetailsResponse{
			ID:               institution.Details.ID,
			InstitutionTitle: institution.Details.InstitutionTitle,
			InstitutionType:  institution.Details.InstitutionType,
			LegalStatus:      institution.Details.LegalAddress,
			Mission:          institution.Details.Mission,
			Founder:          institution.Details.Founder,
			LegalAddress:     institution.Details.LegalAddress,
			FactualAddress:   institution.Details.FactualAddress,
			CreatedAt:        institution.Details.CreatedAt,
			UpdatedAt:        institution.Details.UpdatedAt,
		}
	}

	return &dtos.InstitutionResponse{
		ID:                  institution.ID,
		YearOfEstablishment: institution.YearOfEstablishment,
		Email:               institution.Email,
		Fax:                 institution.Fax,
		Details:             details,
		CreatedAt:           institution.CreatedAt,
		UpdatedAt:           institution.UpdatedAt,
	}
}
