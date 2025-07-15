package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/mappers"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type EmployeePublicationUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeePublicationRequest) (*dtos.EmployeePublicationResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeePublicationRequest) (*dtos.EmployeePublicationResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeePublicationResponse, error)
}

type employeePublicationUsecase struct {
	employeePublicationRepo repositories.EmployeePublicationRepository
	validator               *validator.Validate
}

func NewEmployeePublicationUsecase(
	employeePublicationRepo repositories.EmployeePublicationRepository,
	validator *validator.Validate,
) EmployeePublicationUsecase {
	return &employeePublicationUsecase{
		employeePublicationRepo: employeePublicationRepo,
		validator:               validator,
	}
}

func (uc *employeePublicationUsecase) Create(ctx context.Context, req *dtos.CreateEmployeePublicationRequest) (*dtos.EmployeePublicationResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee publication: %w", err))
	}

	employeePublication := &domain.EmployeePublication{
		EmployeeID:        req.EmployeeID,
		LanguageCode:      req.LanguageCode,
		PublicationTitle:  req.PublicationTitle,
		LinkToPublication: req.LinkToPublication,
	}

	createdEmployeePublication, err := uc.employeePublicationRepo.Create(ctx, employeePublication)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeePublicationDomainToResponseDTO(createdEmployeePublication)
	return resp, nil
}

func (uc *employeePublicationUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeePublicationRequest) (*dtos.EmployeePublicationResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee publication: %w", err))
	}

	employeePublication := &domain.EmployeePublication{
		ID: req.ID,
	}

	if req.PublicationTitle != nil {
		employeePublication.PublicationTitle = *req.PublicationTitle
	}

	if req.LinkToPublication != nil {
		employeePublication.LinkToPublication = *req.LinkToPublication
	}

	updatedEmployeePublication, err := uc.employeePublicationRepo.Update(ctx, employeePublication)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeePublicationDomainToResponseDTO(updatedEmployeePublication)
	return resp, nil
}

func (uc *employeePublicationUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee publication"))
	}

	return uc.employeePublicationRepo.Delete(ctx, id)
}

func (uc *employeePublicationUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeePublicationResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee publication", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee publication", langCode))
	}

	employeePublications, err := uc.employeePublicationRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeePublicationResponse, len(employeePublications))
	for index, degree := range employeePublications {
		resp[index] = mappers.MapEmployeePublicationDomainToResponseDTO(degree)
	}

	return resp, nil
}


