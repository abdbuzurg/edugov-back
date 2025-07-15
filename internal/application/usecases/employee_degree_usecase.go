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

type EmployeeDegreeUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeDegreeRequest) (*dtos.EmployeeDegreeResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeDegreeRequest) (*dtos.EmployeeDegreeResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeDegreeResponse, error)
}

type employeeDegreeUsecase struct {
	employeeDegreeRepo repositories.EmployeeDegreeRepository
	validator          *validator.Validate
}

func NewEmployeeDegreeUsecase(
	employeeDegreeRepo repositories.EmployeeDegreeRepository,
	validator *validator.Validate,
) EmployeeDegreeUsecase {
	return &employeeDegreeUsecase{
		employeeDegreeRepo: employeeDegreeRepo,
		validator:          validator,
	}
}

func (uc *employeeDegreeUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeDegreeRequest) (*dtos.EmployeeDegreeResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee degree: %w", err))
	}

	employeeDegree := &domain.EmployeeDegree{
		EmployeeID:         req.EmployeeID,
		LanguageCode:       req.LanguageCode,
		DegreeLevel:        req.DegreeLevel,
		UniversityName:     req.UniversityName,
		Speciality:         req.Speciality,
		DateStart:          req.DateStart,
		DateEnd:            req.DateEnd,
		GivenBy:            req.GivenBy,
		DateDegreeRecieved: req.DateDegreeRecieved,
	}

	createdEmployeeDegree, err := uc.employeeDegreeRepo.Create(ctx, employeeDegree)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeDegreeDomainToResponseDTO(createdEmployeeDegree)
	return resp, nil
}

func (uc *employeeDegreeUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeDegreeRequest) (*dtos.EmployeeDegreeResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee degree: %w", err))
	}

	employeeDegree := &domain.EmployeeDegree{
		ID: req.ID,
	}

	if req.UniversityName != nil {
		employeeDegree.UniversityName = *req.UniversityName
	}

	if req.DegreeLevel != nil {
		employeeDegree.DegreeLevel = *req.DegreeLevel
	}

	if req.Speciality != nil {
		employeeDegree.Speciality = *req.Speciality
	}

	if req.DateStart != nil {
		employeeDegree.DateStart = *req.DateStart
	}

	if req.DateEnd != nil {
		employeeDegree.DateEnd = *req.DateEnd
	}

	if req.GivenBy != nil {
		employeeDegree.GivenBy = *req.GivenBy
	}

	if req.DateDegreeRecieved != nil {
		employeeDegree.DateDegreeRecieved = *req.DateDegreeRecieved
	}

	updatedEmployeeDegree, err := uc.employeeDegreeRepo.Update(ctx, employeeDegree)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeDegreeDomainToResponseDTO(updatedEmployeeDegree)
	return resp, nil
}

func (uc *employeeDegreeUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee degree"))
	}

	if err := uc.employeeDegreeRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeDegreeUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeDegreeResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee degree", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee degree", langCode))
	}

	employeeDegrees, err := uc.employeeDegreeRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeDegreeResponse, len(employeeDegrees))
	for index, degree := range employeeDegrees {
		resp[index] = mappers.MapEmployeeDegreeDomainToResponseDTO(degree)
	}

	return resp, nil
}


