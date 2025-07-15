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

type EmployeeParticipationInEventUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeParticipationInEventRequest) (*dtos.EmployeeParticipationInEventResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeParticipationInEventRequest) (*dtos.EmployeeParticipationInEventResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeParticipationInEventResponse, error)
}

type employeeParticipationInEventUsecase struct {
	employeeParticipationInEventRepo repositories.EmployeeParticipationInEventRepository
	validator                        *validator.Validate
}

func NewEmployeeParticipationInEventUsecase(
	employeeParticipationInEventRepo repositories.EmployeeParticipationInEventRepository,
	validator *validator.Validate,
) EmployeeParticipationInEventUsecase {
	return &employeeParticipationInEventUsecase{
		employeeParticipationInEventRepo: employeeParticipationInEventRepo,
		validator:                        validator,
	}
}

func (uc *employeeParticipationInEventUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeParticipationInEventRequest) (*dtos.EmployeeParticipationInEventResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee participation in event: %w", err))
	}

	employeeParticipationInEvent := &domain.EmployeeParticipationInEvent{
		EmployeeID:                     req.EmployeeID,
		LanguageCode:                   req.LanguageCode,
		EventTitle:                     req.EventTitle,
		EventDate:                      req.EventDate,
	}

	createdEmployeeParticipationInEvent, err := uc.employeeParticipationInEventRepo.Create(ctx, employeeParticipationInEvent)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeParticipationInEventDomainToResponseDTO(createdEmployeeParticipationInEvent)
	return resp, nil
}

func (uc *employeeParticipationInEventUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeParticipationInEventRequest) (*dtos.EmployeeParticipationInEventResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee participation in event: %w", err))
	}

	employeeParticipationInEvent := &domain.EmployeeParticipationInEvent{
		ID: req.ID,
	}

	if req.EventTitle != nil {
		employeeParticipationInEvent.EventTitle = *req.EventTitle
	}

	if req.EventDate != nil {
		employeeParticipationInEvent.EventDate = *req.EventDate
	}

	updatedEmployeeParticipationInEvent, err := uc.employeeParticipationInEventRepo.Update(ctx, employeeParticipationInEvent)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeParticipationInEventDomainToResponseDTO(updatedEmployeeParticipationInEvent)
	return resp, nil
}

func (uc *employeeParticipationInEventUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee participation in event"))
	}

	if err := uc.employeeParticipationInEventRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeParticipationInEventUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeParticipationInEventResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee participation in event", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee participation in event", langCode))
	}

	employeeParticipationInEvents, err := uc.employeeParticipationInEventRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeParticipationInEventResponse, len(employeeParticipationInEvents))
	for index, degree := range employeeParticipationInEvents {
		resp[index] = mappers.MapEmployeeParticipationInEventDomainToResponseDTO(degree)
	}

	return resp, nil
}


