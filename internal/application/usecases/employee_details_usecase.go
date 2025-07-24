package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/mappers"
	"backend/internal/shared/utils"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type EmployeeDetailsUsecase interface {
	GetByEmployeeIDAndLanguageCode(ctx context.Context, req int64) ([]*dtos.EmployeeDetailsResponse, error)
	Update(ctx context.Context, req []dtos.UpdateEmployeeDetailsRequest) ([]*dtos.EmployeeDetailsResponse, error)
}

type employeeDetailsUsecase struct {
	employeeDetailsRepo repositories.EmployeeDetailsRepository
	store               *postgres.Store
	validator           *validator.Validate
}

func NewEmployeeDetailsUsecase(
	employeeDetailsRepo repositories.EmployeeDetailsRepository,
	store *postgres.Store,
	validator *validator.Validate,
) EmployeeDetailsUsecase {
	return &employeeDetailsUsecase{
		employeeDetailsRepo: employeeDetailsRepo,
		store:               store,
		validator:           validator,
	}
}

func (uc *employeeDetailsUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, req int64) ([]*dtos.EmployeeDetailsResponse, error) {
	if req <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to retrive employee details: incorrect employee id"))
	}

	lang := utils.GetLanguageFromContext(ctx)
	employeeDetails, err := uc.employeeDetailsRepo.GetByEmployeeIDAndLanguageCode(ctx, req, lang)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeDetailsResponse, len(employeeDetails))
	for index, details := range employeeDetails {
		resp[index] = mappers.MapEmployeeDetailsDomainIntoResponseDTO(details)
	}

	return resp, nil
}

func (uc *employeeDetailsUsecase) Update(ctx context.Context, req []dtos.UpdateEmployeeDetailsRequest) ([]*dtos.EmployeeDetailsResponse, error) {
	for _, details := range req {
		if err := uc.validator.Struct(details); err != nil {
			return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee details: %w", err))
		}
	}

	resp := []*dtos.EmployeeDetailsResponse{}
	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeDetailsRepo := postgres.NewPGEmployeeDetailsRepositoryWithQueries(q)
		langCode := utils.GetLanguageFromContext(ctx)

		newDetails := make([]*domain.EmployeeDetails, len(req))
		for index, details := range req {
			newDetails[index] = &domain.EmployeeDetails{
				ID:           details.ID,
				EmployeeID:   details.EmployeeID,
				LanguageCode: langCode,
			}

			if details.Name != nil {
				newDetails[index].Name = *details.Name
			}

			if details.Surname != nil {
				newDetails[index].Surname = *details.Surname
			}

			if details.Middlename != nil {
				newDetails[index].Middlename = *details.Middlename
			}

			if details.IsEmployeeDetailsNew != nil {
				newDetails[index].IsEmployeeDetailsNew = *details.IsEmployeeDetailsNew
			}
		}

		lang := utils.GetLanguageFromContext(ctx)
		oldDetails, err := txEmployeeDetailsRepo.GetByEmployeeIDAndLanguageCode(ctx, newDetails[0].EmployeeID, lang)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}

		updatedDetails, newDetails, removeDetails := utils.CompareSlices(oldDetails, newDetails)
		for _, details := range updatedDetails {
			updatedDetailsResult, err := txEmployeeDetailsRepo.Update(ctx, details)
			if err != nil {
				return err
			}

			resp = append(resp, mappers.MapEmployeeDetailsDomainIntoResponseDTO(updatedDetailsResult))
		}

		for _, details := range newDetails {
			createdDetailsResult, err := txEmployeeDetailsRepo.Create(ctx, details)
			if err != nil {
				return err
			}

			resp = append(resp, mappers.MapEmployeeDetailsDomainIntoResponseDTO(createdDetailsResult))
		}

		for _, details := range removeDetails {
			if err := txEmployeeDetailsRepo.Delete(ctx, details.ID); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
