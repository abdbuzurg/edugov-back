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

type InstitutionAccreditationUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionAccreditationRequest) (*dtos.InstitutionAccreditationResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionAccreditationRequest) (*dtos.InstitutionAccreditationResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionAccreditationResponse, error)
}

type institutionAccreditationUsecase struct {
	institutionAccreditationRepo repositories.InstitutionAccreditationRepository
	validator                    *validator.Validate
}

func NewInstitutionAccreditationUsecase(
	institutionAccreditationRepo repositories.InstitutionAccreditationRepository,
	validator *validator.Validate,
) InstitutionAccreditationUsecase {
	return &institutionAccreditationUsecase{
		institutionAccreditationRepo: institutionAccreditationRepo,
		validator:                    validator,
	}
}

func (uc *institutionAccreditationUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionAccreditationRequest) (*dtos.InstitutionAccreditationResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution accreditation: %w", err))
	}

	institutionAccreditation := &domain.InstitutionAccreditation{
		InstitutionID: req.InstitutionID,
		LanguageCode:  req.LanguageCode,
		AccreditationType:          req.AccreditationType,
		GivenBy:       req.GivenBy,
	}

	createdInstitutionAccreditation, err := uc.institutionAccreditationRepo.Create(ctx, institutionAccreditation)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionAccreditationDomainToResponseDTO(createdInstitutionAccreditation)
	return resp, nil
}

func (uc *institutionAccreditationUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionAccreditationRequest) (*dtos.InstitutionAccreditationResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution accreditation: %w", err))
	}

	institutionAccreditation := &domain.InstitutionAccreditation{
		ID: req.ID,
	}

	if req.AccreditationType != nil {
		institutionAccreditation.AccreditationType = *req.AccreditationType
	}

	if req.GivenBy != nil {
		institutionAccreditation.GivenBy = *req.GivenBy
	}

	updatedInstitutionAccreditation, err := uc.institutionAccreditationRepo.Update(ctx, institutionAccreditation)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionAccreditationDomainToResponseDTO(updatedInstitutionAccreditation)
	return resp, nil
}

func (uc *institutionAccreditationUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution accreditation"))
	}

	if err := uc.institutionAccreditationRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionAccreditationUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionAccreditationResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution accreditation", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution accreditation", langCode))
	}

	institutionAccreditations, err := uc.institutionAccreditationRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionAccreditationResponse, len(institutionAccreditations))
	for index, degree := range institutionAccreditations {
		resp[index] = MapInstitutionAccreditationDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionAccreditationDomainToResponseDTO(institutionAccreditation *domain.InstitutionAccreditation) *dtos.InstitutionAccreditationResponse {
	if institutionAccreditation == nil {
		return nil
	}

	return &dtos.InstitutionAccreditationResponse{
		ID:                institutionAccreditation.ID,
		AccreditationType: institutionAccreditation.AccreditationType,
		GivenBy:           institutionAccreditation.GivenBy,
		CreatedAt:         institutionAccreditation.CreatedAt,
		UpdatedAt:         institutionAccreditation.UpdatedAt,
	}
}
