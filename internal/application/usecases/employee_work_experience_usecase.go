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

type EmployeeWorkExperienceUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeWorkExperienceRequest) (*dtos.EmployeeWorkExperienceResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeWorkExperienceRequest) (*dtos.EmployeeWorkExperienceResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeWorkExperienceResponse, error)
}

type employeeWorkExperienceUsecase struct {
	employeeWorkExperienceRepo repositories.EmployeeWorkExperienceRepository
	validator                  *validator.Validate
}

func NewEmployeeWorkExperienceUsecase(
	employeeWorkExperienceRepo repositories.EmployeeWorkExperienceRepository,
	validator *validator.Validate,
) EmployeeWorkExperienceUsecase {
	return &employeeWorkExperienceUsecase{
		employeeWorkExperienceRepo: employeeWorkExperienceRepo,
		validator:                  validator,
	}
}

func (uc *employeeWorkExperienceUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeWorkExperienceRequest) (*dtos.EmployeeWorkExperienceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee professional activity in education: %w", err))
	}

	employeeWorkExperience := &domain.EmployeeWorkExperience{
		EmployeeID:   req.EmployeeID,
		LanguageCode: req.LanguageCode,
		Workplace:    req.Workplace,
		JobTitle:     req.JobTitle,
		Description:  req.Description,
		Ongoing:      req.Ongoing,
		DateStart:    req.DateStart,
		DateEnd:      req.DateEnd,
	}

	createdEmployeeWorkExperience, err := uc.employeeWorkExperienceRepo.Create(ctx, employeeWorkExperience)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeWorkExperienceDomainToResponseDTO(createdEmployeeWorkExperience)
	return resp, nil
}

func (uc *employeeWorkExperienceUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeWorkExperienceRequest) (*dtos.EmployeeWorkExperienceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee professional activity in education: %w", err))
	}

	employeeWorkExperience := &domain.EmployeeWorkExperience{
		ID:      req.ID,
		Ongoing: req.Ongoing,
	}

	if req.Workplace != nil {
		employeeWorkExperience.Workplace = *req.Workplace
	}
	if req.Description != nil {
		employeeWorkExperience.Description = *req.Description
	}
	if req.JobTitle != nil {
		employeeWorkExperience.JobTitle = *req.JobTitle
	}
	if req.DateStart != nil {
		employeeWorkExperience.DateStart = *req.DateStart
	}
	if req.DateEnd != nil {
		employeeWorkExperience.DateEnd = *req.DateEnd
	}

	updatedEmployeeWorkExperience, err := uc.employeeWorkExperienceRepo.Update(ctx, employeeWorkExperience)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeWorkExperienceDomainToResponseDTO(updatedEmployeeWorkExperience)
	return resp, nil
}

func (uc *employeeWorkExperienceUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee professional activity in education"))
	}

	if err := uc.employeeWorkExperienceRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeWorkExperienceUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeWorkExperienceResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee professional activity in education", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee professional activity in education", langCode))
	}

	employeeWorkExperiences, err := uc.employeeWorkExperienceRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeWorkExperienceResponse, len(employeeWorkExperiences))
	for index, degree := range employeeWorkExperiences {
		resp[index] = mappers.MapEmployeeWorkExperienceDomainToResponseDTO(degree)
	}

	return resp, nil
}
