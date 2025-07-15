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

type InstitutionLicenceUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionLicenceRequest) (*dtos.InstitutionLicenceResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionLicenceRequest) (*dtos.InstitutionLicenceResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionLicenceResponse, error)
}

type institutionLicenceUsecase struct {
	institutionLicenceRepo repositories.InstitutionLicenceRepository
	validator              *validator.Validate
}

func NewInstitutionLicenceUsecase(
	institutionLicenceRepo repositories.InstitutionLicenceRepository,
	validator *validator.Validate,
) InstitutionLicenceUsecase {
	return &institutionLicenceUsecase{
		institutionLicenceRepo: institutionLicenceRepo,
		validator:              validator,
	}
}

func (uc *institutionLicenceUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionLicenceRequest) (*dtos.InstitutionLicenceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution licence: %w", err))
	}

	institutionLicence := &domain.InstitutionLicence{
		InstitutionID: req.InstitutionID,
		LanguageCode:  req.LanguageCode,
		LicenceTitle:  req.LicenceTitle,
		LicenceType:   req.LicenceType,
		GivenBy:       req.GivenBy,
		LinkToFile:    req.LinkToFile,
		DateStart:     req.DateStart,
		DateEnd:       req.DateEnd,
	}

	createdInstitutionLicence, err := uc.institutionLicenceRepo.Create(ctx, institutionLicence)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionLicenceDomainToResponseDTO(createdInstitutionLicence)
	return resp, nil
}

func (uc *institutionLicenceUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionLicenceRequest) (*dtos.InstitutionLicenceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution licence: %w", err))
	}

	institutionLicence := &domain.InstitutionLicence{
		ID: req.ID,
	}

	if req.LicenceType != nil {
		institutionLicence.LicenceType = *req.LicenceType
	}

	if req.LicenceTitle != nil {
		institutionLicence.LicenceTitle = *req.LicenceTitle
	}

	if req.GivenBy != nil {
		institutionLicence.GivenBy = *req.GivenBy
	}

	if req.DateStart != nil {
		institutionLicence.DateStart = *req.DateStart
	}

	if req.DateEnd != nil {
		institutionLicence.DateEnd = *req.DateEnd
	}

	if req.LinkToFile != nil {
		institutionLicence.LinkToFile = *req.LinkToFile
	}

	updatedInstitutionLicence, err := uc.institutionLicenceRepo.Update(ctx, institutionLicence)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionLicenceDomainToResponseDTO(updatedInstitutionLicence)
	return resp, nil
}

func (uc *institutionLicenceUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution licence"))
	}

	if err := uc.institutionLicenceRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionLicenceUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionLicenceResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution licence", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution licence", langCode))
	}

	institutionLicences, err := uc.institutionLicenceRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionLicenceResponse, len(institutionLicences))
	for index, degree := range institutionLicences {
		resp[index] = MapInstitutionLicenceDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionLicenceDomainToResponseDTO(institutionLicence *domain.InstitutionLicence) *dtos.InstitutionLicenceResponse {
	if institutionLicence == nil {
		return nil
	}

	return &dtos.InstitutionLicenceResponse{
		ID:           institutionLicence.ID,
		LicenceTitle: institutionLicence.LicenceTitle,
		LicenceType:  institutionLicence.LicenceType,
		GivenBy:      institutionLicence.GivenBy,
		LinkToFile:   institutionLicence.LinkToFile,
		DateStart:    institutionLicence.DateStart,
		DateEnd:      institutionLicence.DateEnd,
		CreatedAt:    institutionLicence.CreatedAt,
		UpdatedAt:    institutionLicence.UpdatedAt,
	}
}
