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

type EmployeeScientificAwardUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeScientificAwardRequest) (*dtos.EmployeeScientificAwardResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeScientificAwardRequest) (*dtos.EmployeeScientificAwardResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeScientificAwardResponse, error)
}

type employeeScientificAwardUsecase struct {
	employeeScientificAwardRepo repositories.EmployeeScientificAwardRepository
	validator                   *validator.Validate
}

func NewEmployeeScientificAwardUsecase(
	employeeScientificAwardRepo repositories.EmployeeScientificAwardRepository,
	validator *validator.Validate,
) EmployeeScientificAwardUsecase {
	return &employeeScientificAwardUsecase{
		employeeScientificAwardRepo: employeeScientificAwardRepo,
		validator:                   validator,
	}
}

func (uc *employeeScientificAwardUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeScientificAwardRequest) (*dtos.EmployeeScientificAwardResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee scientific award: %w", err))
	}

	employeeScientificAward := &domain.EmployeeScientificAward{
		EmployeeID:                req.EmployeeID,
		LanguageCode:              req.LanguageCode,
		ScientificAwardTitle:      req.ScientificAwardTitle,
		GivenBy:                   req.GivenBy,
	}

	createdEmployeeScientificAward, err := uc.employeeScientificAwardRepo.Create(ctx, employeeScientificAward)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeScientificAwardDomainToResponseDTO(createdEmployeeScientificAward)
	return resp, nil
}

func (uc *employeeScientificAwardUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeScientificAwardRequest) (*dtos.EmployeeScientificAwardResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee scientific award: %w", err))
	}

	employeeScientificAward := &domain.EmployeeScientificAward{
		ID: req.ID,
	}

	if req.ScientificAwardTitle != nil {
		employeeScientificAward.ScientificAwardTitle = *req.ScientificAwardTitle
	}

	if req.GivenBy != nil {
		employeeScientificAward.GivenBy = *req.GivenBy
	}

	updatedEmployeeScientificAward, err := uc.employeeScientificAwardRepo.Update(ctx, employeeScientificAward)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeScientificAwardDomainToResponseDTO(updatedEmployeeScientificAward)
	return resp, nil
}

func (uc *employeeScientificAwardUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee scientific award"))
	}

	if err := uc.employeeScientificAwardRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeScientificAwardUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeScientificAwardResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee scientific award", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee scientific award", langCode))
	}

	employeeScientificAwards, err := uc.employeeScientificAwardRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeScientificAwardResponse, len(employeeScientificAwards))
	for index, degree := range employeeScientificAwards {
		resp[index] = mappers.MapEmployeeScientificAwardDomainToResponseDTO(degree)
	}

	return resp, nil
}


