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

type InstitutionMainResearchDirectionUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionMainResearchDirectionRequest) (*dtos.InstitutionMainResearchDirectionResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionMainResearchDirectionRequest) (*dtos.InstitutionMainResearchDirectionResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionMainResearchDirectionResponse, error)
}

type institutionMainResearchDirectionUsecase struct {
	institutionMainResearchDirectionRepo repositories.InstitutionMainResearchDirectionRepository
	validator                            *validator.Validate
}

func NewInstitutionMainResearchDirectionUsecase(
	institutionMainResearchDirectionRepo repositories.InstitutionMainResearchDirectionRepository,
	validator *validator.Validate,
) InstitutionMainResearchDirectionUsecase {
	return &institutionMainResearchDirectionUsecase{
		institutionMainResearchDirectionRepo: institutionMainResearchDirectionRepo,
		validator:                            validator,
	}
}

func (uc *institutionMainResearchDirectionUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionMainResearchDirectionRequest) (*dtos.InstitutionMainResearchDirectionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution main research direction: %w", err))
	}

	institutionMainResearchDirection := &domain.InstitutionMainResearchDirection{
		InstitutionID:          req.InstitutionID,
		LanguageCode:           req.LanguageCode,
		ResearchDirectionTitle: req.ResearchDirectionTitle,
		Discipline:             req.Discipline,
		AreaOfResearch:         req.AreaOfResearch,
	}

	createdInstitutionMainResearchDirection, err := uc.institutionMainResearchDirectionRepo.Create(ctx, institutionMainResearchDirection)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionMainResearchDirectionDomainToResponseDTO(createdInstitutionMainResearchDirection)
	return resp, nil
}

func (uc *institutionMainResearchDirectionUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionMainResearchDirectionRequest) (*dtos.InstitutionMainResearchDirectionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution main research direction: %w", err))
	}

	institutionMainResearchDirection := &domain.InstitutionMainResearchDirection{
		ID: req.ID,
	}

	if req.ResearchDirectionTitle != nil {
		institutionMainResearchDirection.ResearchDirectionTitle = *req.ResearchDirectionTitle
	}

	if req.Discipline != nil {
		institutionMainResearchDirection.Discipline = *req.Discipline
	}

	if req.AreaOfResearch != nil {
		institutionMainResearchDirection.AreaOfResearch = *req.AreaOfResearch
	}

	updatedInstitutionMainResearchDirection, err := uc.institutionMainResearchDirectionRepo.Update(ctx, institutionMainResearchDirection)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionMainResearchDirectionDomainToResponseDTO(updatedInstitutionMainResearchDirection)
	return resp, nil
}

func (uc *institutionMainResearchDirectionUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution main research direction"))
	}

	if err := uc.institutionMainResearchDirectionRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionMainResearchDirectionUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionMainResearchDirectionResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution main research direction", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution main research direction", langCode))
	}

	institutionMainResearchDirections, err := uc.institutionMainResearchDirectionRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionMainResearchDirectionResponse, len(institutionMainResearchDirections))
	for index, degree := range institutionMainResearchDirections {
		resp[index] = MapInstitutionMainResearchDirectionDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionMainResearchDirectionDomainToResponseDTO(institutionMainResearchDirection *domain.InstitutionMainResearchDirection) *dtos.InstitutionMainResearchDirectionResponse {
	if institutionMainResearchDirection == nil {
		return nil
	}

	return &dtos.InstitutionMainResearchDirectionResponse{
		ID:                     institutionMainResearchDirection.ID,
		ResearchDirectionTitle: institutionMainResearchDirection.ResearchDirectionTitle,
		Discipline:             institutionMainResearchDirection.Discipline,
		AreaOfResearch:         institutionMainResearchDirection.AreaOfResearch,
		CreatedAt:              institutionMainResearchDirection.CreatedAt,
		UpdatedAt:              institutionMainResearchDirection.UpdatedAt,
	}
}
