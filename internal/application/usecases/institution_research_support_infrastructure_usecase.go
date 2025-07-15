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

type InstitutionResearchSupportInfrastructureUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionResearchSupportInfrastructureRequest) (*dtos.InstitutionResearchSupportInfrastructureResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionResearchSupportInfrastructureRequest) (*dtos.InstitutionResearchSupportInfrastructureResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionResearchSupportInfrastructureResponse, error)
}

type institutionResearchSupportInfrastructureUsecase struct {
	institutionResearchSupportInfrastructureRepo repositories.InstitutionResearchSupportInfrastructureRepository
	validator                                    *validator.Validate
}

func NewInstitutionResearchSupportInfrastructureUsecase(
	institutionResearchSupportInfrastructureRepo repositories.InstitutionResearchSupportInfrastructureRepository,
	validator *validator.Validate,
) InstitutionResearchSupportInfrastructureUsecase {
	return &institutionResearchSupportInfrastructureUsecase{
		institutionResearchSupportInfrastructureRepo: institutionResearchSupportInfrastructureRepo,
		validator: validator,
	}
}

func (uc *institutionResearchSupportInfrastructureUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionResearchSupportInfrastructureRequest) (*dtos.InstitutionResearchSupportInfrastructureResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution research support infrastructure: %w", err))
	}

	institutionResearchSupportInfrastructure := &domain.InstitutionResearchSupportInfrastructure{
		InstitutionID:                      req.InstitutionID,
		LanguageCode:                       req.LanguageCode,
		ResearchSupportInfrastructureTitle: req.ResearchSupportInfrastructureTitle,
		ResearchSupportInfrastructureType:  req.ResearchSupportInfrastructureType,
		TINOfLegalEntity:                   req.TINOfLegalEntity,
	}

	createdInstitutionResearchSupportInfrastructure, err := uc.institutionResearchSupportInfrastructureRepo.Create(ctx, institutionResearchSupportInfrastructure)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionResearchSupportInfrastructureDomainToResponseDTO(createdInstitutionResearchSupportInfrastructure)
	return resp, nil
}

func (uc *institutionResearchSupportInfrastructureUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionResearchSupportInfrastructureRequest) (*dtos.InstitutionResearchSupportInfrastructureResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution research support infrastructure: %w", err))
	}

	institutionResearchSupportInfrastructure := &domain.InstitutionResearchSupportInfrastructure{
		ID: req.ID,
	}

	if req.ResearchSupportInfrastructureTitle != nil {
		institutionResearchSupportInfrastructure.ResearchSupportInfrastructureTitle = *req.ResearchSupportInfrastructureTitle
	}

	if req.ResearchSupportInfrastructureType != nil {
		institutionResearchSupportInfrastructure.ResearchSupportInfrastructureType = *req.ResearchSupportInfrastructureType
	}

	if req.TINOfLegalEntity != nil {
		institutionResearchSupportInfrastructure.TINOfLegalEntity = *req.TINOfLegalEntity
	}

	updatedInstitutionResearchSupportInfrastructure, err := uc.institutionResearchSupportInfrastructureRepo.Update(ctx, institutionResearchSupportInfrastructure)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionResearchSupportInfrastructureDomainToResponseDTO(updatedInstitutionResearchSupportInfrastructure)
	return resp, nil
}

func (uc *institutionResearchSupportInfrastructureUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution research support infrastructure"))
	}

	if err := uc.institutionResearchSupportInfrastructureRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionResearchSupportInfrastructureUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionResearchSupportInfrastructureResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution research support infrastructure", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution research support infrastructure", langCode))
	}

	institutionResearchSupportInfrastructures, err := uc.institutionResearchSupportInfrastructureRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionResearchSupportInfrastructureResponse, len(institutionResearchSupportInfrastructures))
	for index, degree := range institutionResearchSupportInfrastructures {
		resp[index] = MapInstitutionResearchSupportInfrastructureDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionResearchSupportInfrastructureDomainToResponseDTO(institutionResearchSupportInfrastructure *domain.InstitutionResearchSupportInfrastructure) *dtos.InstitutionResearchSupportInfrastructureResponse {
	if institutionResearchSupportInfrastructure == nil {
		return nil
	}

	return &dtos.InstitutionResearchSupportInfrastructureResponse{
		ID:                                 institutionResearchSupportInfrastructure.ID,
		ResearchSupportInfrastructureType:  institutionResearchSupportInfrastructure.ResearchSupportInfrastructureType,
		ResearchSupportInfrastructureTitle: institutionResearchSupportInfrastructure.ResearchSupportInfrastructureTitle,
		TINOfLegalEntity:                   institutionResearchSupportInfrastructure.TINOfLegalEntity,
		CreatedAt:                          institutionResearchSupportInfrastructure.CreatedAt,
		UpdatedAt:                          institutionResearchSupportInfrastructure.UpdatedAt,
	}
}
