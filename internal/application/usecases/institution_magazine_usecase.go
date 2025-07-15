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

type InstitutionMagazineUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionMagazineRequest) (*dtos.InstitutionMagazineResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionMagazineRequest) (*dtos.InstitutionMagazineResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionMagazineResponse, error)
}

type institutionMagazineUsecase struct {
	institutionMagazineRepo repositories.InstitutionMagazineRepository
	validator               *validator.Validate
}

func NewInstitutionMagazineUsecase(
	institutionMagazineRepo repositories.InstitutionMagazineRepository,
	validator *validator.Validate,
) InstitutionMagazineUsecase {
	return &institutionMagazineUsecase{
		institutionMagazineRepo: institutionMagazineRepo,
		validator:               validator,
	}
}

func (uc *institutionMagazineUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionMagazineRequest) (*dtos.InstitutionMagazineResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution magazine: %w", err))
	}

	institutionMagazine := &domain.InstitutionMagazine{
		InstitutionID: req.InstitutionID,
		LanguageCode:  req.LanguageCode,
		MagazineName:  req.MagazineName,
		Link:          req.Link,
		LinkToRINC:    req.LinkToRINC,
	}

	createdInstitutionMagazine, err := uc.institutionMagazineRepo.Create(ctx, institutionMagazine)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionMagazineDomainToResponseDTO(createdInstitutionMagazine)
	return resp, nil
}

func (uc *institutionMagazineUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionMagazineRequest) (*dtos.InstitutionMagazineResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution magazine: %w", err))
	}

	institutionMagazine := &domain.InstitutionMagazine{
		ID: req.ID,
	}

	if req.MagazineName != nil {
		institutionMagazine.MagazineName = *req.MagazineName
	}

	if req.Link != nil {
		institutionMagazine.Link = *req.Link
	}

	if req.LinkToRINC != nil {
		institutionMagazine.LinkToRINC = *req.LinkToRINC
	}

	updatedInstitutionMagazine, err := uc.institutionMagazineRepo.Update(ctx, institutionMagazine)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionMagazineDomainToResponseDTO(updatedInstitutionMagazine)
	return resp, nil
}

func (uc *institutionMagazineUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution magazine"))
	}

	if err := uc.institutionMagazineRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionMagazineUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionMagazineResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution magazine", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution magazine", langCode))
	}

	institutionMagazines, err := uc.institutionMagazineRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionMagazineResponse, len(institutionMagazines))
	for index, degree := range institutionMagazines {
		resp[index] = MapInstitutionMagazineDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionMagazineDomainToResponseDTO(institutionMagazine *domain.InstitutionMagazine) *dtos.InstitutionMagazineResponse {
	if institutionMagazine == nil {
		return nil
	}

	return &dtos.InstitutionMagazineResponse{
		ID:           institutionMagazine.ID,
		MagazineName: institutionMagazine.MagazineName,
		Link:         institutionMagazine.Link,
		LinkToRINC:   institutionMagazine.LinkToRINC,
		CreatedAt:    institutionMagazine.CreatedAt,
		UpdatedAt:    institutionMagazine.UpdatedAt,
	}
}
