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

type EmployeeResearchActivityUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeResearchActivityRequest) (*dtos.EmployeeResearchActivityResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeResearchActivityRequest) (*dtos.EmployeeResearchActivityResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeResearchActivityResponse, error)
}

type employeeResearchActivityUsecase struct {
	employeeResearchActivityRepo repositories.EmployeeResearchActivityRepository
	validator                    *validator.Validate
}

func NewEmployeeResearchActivityUsecase(
	employeeResearchActivityRepo repositories.EmployeeResearchActivityRepository,
	validator *validator.Validate,
) EmployeeResearchActivityUsecase {
	return &employeeResearchActivityUsecase{
		employeeResearchActivityRepo: employeeResearchActivityRepo,
		validator:                    validator,
	}
}

func (uc *employeeResearchActivityUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeResearchActivityRequest) (*dtos.EmployeeResearchActivityResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee research activity: %w", err))
	}

	employeeResearchActivity := &domain.EmployeeResearchActivity{
		EmployeeID:            req.EmployeeID,
		LanguageCode:          req.LanguageCode,
		ResearchActivityTitle: req.ResearchActivityTitle,
		EmployeeRole:          req.EmployeeRole,
	}

	createdEmployeeResearchActivity, err := uc.employeeResearchActivityRepo.Create(ctx, employeeResearchActivity)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeResearchActivityDomainToResponseDTO(createdEmployeeResearchActivity)
	return resp, nil
}

func (uc *employeeResearchActivityUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeResearchActivityRequest) (*dtos.EmployeeResearchActivityResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee research activity: %w", err))
	}

	employeeResearchActivity := &domain.EmployeeResearchActivity{
		ID: req.ID,
	}

	if req.ResearchActivityTitle != nil {
		employeeResearchActivity.ResearchActivityTitle = *req.ResearchActivityTitle
	}

	if req.EmployeeRole != nil {
		employeeResearchActivity.EmployeeRole = *req.EmployeeRole
	}

	updatedEmployeeResearchActivity, err := uc.employeeResearchActivityRepo.Update(ctx, employeeResearchActivity)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeResearchActivityDomainToResponseDTO(updatedEmployeeResearchActivity)
	return resp, nil
}

func (uc *employeeResearchActivityUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee research activity"))
	}

	if err := uc.employeeResearchActivityRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeResearchActivityUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeResearchActivityResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee research activity", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee research activity", langCode))
	}

	employeeResearchActivitys, err := uc.employeeResearchActivityRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeResearchActivityResponse, len(employeeResearchActivitys))
	for index, degree := range employeeResearchActivitys {
		resp[index] = mappers.MapEmployeeResearchActivityDomainToResponseDTO(degree)
	}

	return resp, nil
}


