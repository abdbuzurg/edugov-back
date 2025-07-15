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

type EmployeePatentUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeePatentRequest) (*dtos.EmployeePatentResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeePatentRequest) (*dtos.EmployeePatentResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeePatentResponse, error)
}

type employeePatentUsecase struct {
	employeePatentRepo repositories.EmployeePatentRepository
	validator          *validator.Validate
}

func NewEmployeePatentUsecase(
	employeePatentRepo repositories.EmployeePatentRepository,
	validator *validator.Validate,
) EmployeePatentUsecase {
	return &employeePatentUsecase{
		employeePatentRepo: employeePatentRepo,
		validator:          validator,
	}
}

func (uc *employeePatentUsecase) Create(ctx context.Context, req *dtos.CreateEmployeePatentRequest) (*dtos.EmployeePatentResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee patent: %w", err))
	}

	employeePatent := &domain.EmployeePatent{
		EmployeeID:       req.EmployeeID,
		LanguageCode:     req.LanguageCode,
		PatentTitle:      req.PatentTitle,
		Description:      req.Description,
	}

	createdEmployeePatent, err := uc.employeePatentRepo.Create(ctx, employeePatent)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeePatentDomainToResponseDTO(createdEmployeePatent)
	return resp, nil
}

func (uc *employeePatentUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeePatentRequest) (*dtos.EmployeePatentResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee patent: %w", err))
	}

	employeePatent := &domain.EmployeePatent{
		ID: req.ID,
	}

	if req.PatentTitle != nil {
		employeePatent.PatentTitle = *req.PatentTitle
	}

	if req.Description != nil {
		employeePatent.Description = *req.Description
	}

	updatedEmployeePatent, err := uc.employeePatentRepo.Update(ctx, employeePatent)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeePatentDomainToResponseDTO(updatedEmployeePatent)
	return resp, nil
}

func (uc *employeePatentUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee patent"))
	}

	if err := uc.employeePatentRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeePatentUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeePatentResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee patent", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee patent", langCode))
	}

	employeePatents, err := uc.employeePatentRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeePatentResponse, len(employeePatents))
	for index, degree := range employeePatents {
		resp[index] = mappers.MapEmployeePatentDomainToResponseDTO(degree)
	}

	return resp, nil
}


