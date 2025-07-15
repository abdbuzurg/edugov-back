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

type EmployeeSocialUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeSocialRequest) (*dtos.EmployeeSocialResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeSocialRequest) (*dtos.EmployeeSocialResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeID(ctx context.Context, employeeID int64) ([]*dtos.EmployeeSocialResponse, error)
}

type employeeSocialUsecase struct {
	employeeSocialRepo repositories.EmployeeSocialRepository
	validator          *validator.Validate
}

func NewEmployeeSocialUsecase(
	employeeSocialRepo repositories.EmployeeSocialRepository,
	validator *validator.Validate,
) EmployeeSocialUsecase {
	return &employeeSocialUsecase{
		employeeSocialRepo: employeeSocialRepo,
		validator:          validator,
	}
}

func (uc *employeeSocialUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeSocialRequest) (*dtos.EmployeeSocialResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee social: %w", err))
	}

	employeeSocial := &domain.EmployeeSocial{
		EmployeeID:   req.EmployeeID,
		SocialName:   req.SocialName,
		LinkToSocial: req.LinkToSocial,
	}

	createdEmployeeSocial, err := uc.employeeSocialRepo.Create(ctx, employeeSocial)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeSocialDomainToResponseDTO(createdEmployeeSocial)
	return resp, nil
}

func (uc *employeeSocialUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeSocialRequest) (*dtos.EmployeeSocialResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee social: %w", err))
	}

	employeeSocial := &domain.EmployeeSocial{
		ID: req.ID,
	}

	if req.SocialName != nil {
		employeeSocial.SocialName = *req.SocialName
	}

	if req.LinkToSocial != nil {
		employeeSocial.LinkToSocial = *req.LinkToSocial
	}

	updatedEmployeeSocial, err := uc.employeeSocialRepo.Update(ctx, employeeSocial)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeSocialDomainToResponseDTO(updatedEmployeeSocial)
	return resp, nil
}

func (uc *employeeSocialUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee social"))
	}

	if err := uc.employeeSocialRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeSocialUsecase) GetByEmployeeID(ctx context.Context, employeeID int64) ([]*dtos.EmployeeSocialResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee social", employeeID))
	}

	employeeSocials, err := uc.employeeSocialRepo.GetByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeSocialResponse, len(employeeSocials))
	for index, degree := range employeeSocials {
		resp[index] = mappers.MapEmployeeSocialDomainToResponseDTO(degree)
	}

	return resp, nil
}


