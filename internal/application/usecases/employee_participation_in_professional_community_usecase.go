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

type EmployeeParticipationInProfessionalCommunityUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeParticipationInProfessionalCommunityRequest) (*dtos.EmployeeParticipationInProfessionalCommunityResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeParticipationInProfessionalCommunityRequest) (*dtos.EmployeeParticipationInProfessionalCommunityResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeParticipationInProfessionalCommunityResponse, error)
}

type employeeParticipationInProfessionalCommunityUsecase struct {
	employeeParticipationInProfessionalCommunityRepo repositories.EmployeeParticipationInProfessionalCommunityRepository
	validator                                        *validator.Validate
}

func NewEmployeeParticipationInProfessionalCommunityUsecase(
	employeeParticipationInProfessionalCommunityRepo repositories.EmployeeParticipationInProfessionalCommunityRepository,
	validator *validator.Validate,
) EmployeeParticipationInProfessionalCommunityUsecase {
	return &employeeParticipationInProfessionalCommunityUsecase{
		employeeParticipationInProfessionalCommunityRepo: employeeParticipationInProfessionalCommunityRepo,
		validator: validator,
	}
}

func (uc *employeeParticipationInProfessionalCommunityUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeParticipationInProfessionalCommunityRequest) (*dtos.EmployeeParticipationInProfessionalCommunityResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee participation in professional community: %w", err))
	}

	employeeParticipationInProfessionalCommunity := &domain.EmployeeParticipationInProfessionalCommunity{
		EmployeeID:                  req.EmployeeID,
		LanguageCode:                req.LanguageCode,
		ProfessionalCommunityTitle:  req.ProfessionalCommunityTitle,
		RoleInProfessionalCommunity: req.RoleInProfessionalCommunity,
	}

	createdEmployeeParticipationInProfessionalCommunity, err := uc.employeeParticipationInProfessionalCommunityRepo.Create(ctx, employeeParticipationInProfessionalCommunity)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeParticipationInProfessionalCommunityDomainToResponseDTO(createdEmployeeParticipationInProfessionalCommunity)
	return resp, nil
}

func (uc *employeeParticipationInProfessionalCommunityUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeParticipationInProfessionalCommunityRequest) (*dtos.EmployeeParticipationInProfessionalCommunityResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee participation in professional community: %w", err))
	}

	employeeParticipationInProfessionalCommunity := &domain.EmployeeParticipationInProfessionalCommunity{
		ID: req.ID,
	}

	if req.ProfessionalCommunityTitle != nil {
		employeeParticipationInProfessionalCommunity.ProfessionalCommunityTitle = *req.ProfessionalCommunityTitle
	}

	if req.RoleInProfessionalCommunity != nil {
		employeeParticipationInProfessionalCommunity.RoleInProfessionalCommunity = *req.RoleInProfessionalCommunity
	}

	updatedEmployeeParticipationInProfessionalCommunity, err := uc.employeeParticipationInProfessionalCommunityRepo.Update(ctx, employeeParticipationInProfessionalCommunity)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeParticipationInProfessionalCommunityDomainToResponseDTO(updatedEmployeeParticipationInProfessionalCommunity)
	return resp, nil
}

func (uc *employeeParticipationInProfessionalCommunityUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee participation in professional community"))
	}

	if err := uc.employeeParticipationInProfessionalCommunityRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeParticipationInProfessionalCommunityUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeParticipationInProfessionalCommunityResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee participation in professional community", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee participation in professional community", langCode))
	}

	employeeParticipationInProfessionalCommunitys, err := uc.employeeParticipationInProfessionalCommunityRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeParticipationInProfessionalCommunityResponse, len(employeeParticipationInProfessionalCommunitys))
	for index, degree := range employeeParticipationInProfessionalCommunitys {
		resp[index] = mappers.MapEmployeeParticipationInProfessionalCommunityDomainToResponseDTO(degree)
	}

	return resp, nil
}


