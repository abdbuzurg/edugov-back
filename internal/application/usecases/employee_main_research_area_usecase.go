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

type EmployeeMainResearchAreaUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeMainResearchAreaRequest) (*dtos.EmployeeMainResearchAreaResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeMainResearchAreaRequest) (*dtos.EmployeeMainResearchAreaResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeMainResearchAreaResponse, error)
}

type employeeMainResearchAreaUsecase struct {
	store           *postgres.Store
	employeeMRARepo repositories.EmployeeMainResearchArea
	validator       *validator.Validate
}

func NewEmployeeMainResearchAreaUsecase(
	employeeMRARepo repositories.EmployeeMainResearchArea,
	store *postgres.Store,
	validator *validator.Validate,
) EmployeeMainResearchAreaUsecase {
	return &employeeMainResearchAreaUsecase{
		employeeMRARepo: employeeMRARepo,
		store:           store,
		validator:       validator,
	}
}

func (uc *employeeMainResearchAreaUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeMainResearchAreaRequest) (*dtos.EmployeeMainResearchAreaResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee main research area: %w", err))
	}

	var employeeMRA *domain.EmployeeMainResearchArea
	var err error
	err = uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeMainResearchAreaRepo := postgres.NewPgEmployeeMainResearchAreaRepositoryWithQueries(q)

		employeeMRA, err = txEmployeeMainResearchAreaRepo.CreateMRA(ctx, &domain.EmployeeMainResearchArea{
			EmployeeID:   req.EmployeeID,
			LanguageCode: req.LanguageCode,
			Discipline:   req.Discipline,
			Area:         req.Area,
		})
		if err != nil {
			return err
		}

		mainResearchAreaKT := make([]*domain.ResearchAreaKeyTopic, len(req.KeyTopics))
		for index, rakt := range req.KeyTopics {
			mainResearchAreaKT[index], err = txEmployeeMainResearchAreaRepo.CreateRAKT(ctx, &domain.ResearchAreaKeyTopic{
				EmployeeMainResearchAreaID: employeeMRA.ID,
				KeyTopicTitle:              rakt.KeyTopicTitle,
			})
			if err != nil {
				return err
			}
		}

		employeeMRA.KeyTopics = mainResearchAreaKT

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed to create employee main research area: %w", err))
	}

	resp := mappers.MapEmployeeMainResearchAreaDomainToResponseDTO(employeeMRA)

	return resp, nil
}

func (uc *employeeMainResearchAreaUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeMainResearchAreaRequest) (*dtos.EmployeeMainResearchAreaResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee main research area: %w", err))
	}

	var employeeMRA *domain.EmployeeMainResearchArea
	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeMainResearchAreaRepo := postgres.NewPgEmployeeMainResearchAreaRepositoryWithQueries(q)

		_, err := txEmployeeMainResearchAreaRepo.GetMRAByID(ctx, req.ID)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}

		employeeMRA = &domain.EmployeeMainResearchArea{
			ID: req.ID,
		}

		if req.Discipline != nil {
			employeeMRA.Discipline = *req.Discipline
		}

		if req.Area != nil {
			employeeMRA.Area = *req.Area
		}

		employeeMRA, err = txEmployeeMainResearchAreaRepo.UpdateMRA(ctx, employeeMRA)
		if err != nil {
			return err
		}

		reqRakt := make([]*domain.ResearchAreaKeyTopic, len(req.KeyTopics))
		for index, kt := range req.KeyTopics {
			reqRakt[index] = &domain.ResearchAreaKeyTopic{
				ID:                         kt.ID,
				EmployeeMainResearchAreaID: req.ID,
			}

			if kt.KeyTopicTitle != nil {
				reqRakt[index].KeyTopicTitle = *kt.KeyTopicTitle
			}
		}

		oldRakt, err := txEmployeeMainResearchAreaRepo.GetRAKTByMRAIDAndLanguageCode(ctx, employeeMRA.ID)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}

		updatedRAKTs, newRAKTs, removeRAKTs := utils.CompareSlices(oldRakt, reqRakt)
		for _, rakt := range updatedRAKTs {
			updatedRAKT, err := txEmployeeMainResearchAreaRepo.UpdateRAKT(ctx, rakt)
			if err != nil {
				return err
			}

			employeeMRA.KeyTopics = append(employeeMRA.KeyTopics, updatedRAKT)
		}

		for _, rakt := range newRAKTs {
			createdRAKT, err := txEmployeeMainResearchAreaRepo.CreateRAKT(ctx, rakt)
			if err != nil {
				return err
			}

			employeeMRA.KeyTopics = append(employeeMRA.KeyTopics, createdRAKT)
		}

		for _, rakt := range removeRAKTs {
			if err := txEmployeeMainResearchAreaRepo.DeleteRAKT(ctx, rakt.ID); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed to update employee main research area"))
	}

	resp := mappers.MapEmployeeMainResearchAreaDomainToResponseDTO(employeeMRA)
	return resp, nil
}

func (uc *employeeMainResearchAreaUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input - ID(%d) to delete employee main research area", id))
	}

	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeMainResearchAreaRepo := postgres.NewPgEmployeeMainResearchAreaRepositoryWithQueries(q)

		if err := txEmployeeMainResearchAreaRepo.DeleteRAKTbyMraID(ctx, id); err != nil {
			return err
		}

		if err := txEmployeeMainResearchAreaRepo.DeleteMRA(ctx, id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return uc.employeeMRARepo.DeleteMRA(ctx, id)
}

func (uc *employeeMainResearchAreaUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeMainResearchAreaResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee main research area", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee main research area", langCode))
	}

	var employeeMRAs []*domain.EmployeeMainResearchArea
	var err error
	err = uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeMainResearchAreaRepo := postgres.NewPgEmployeeMainResearchAreaRepositoryWithQueries(q)

		employeeMRAs, err = txEmployeeMainResearchAreaRepo.GetMRAByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
		if err != nil {
			return err
		}

		for index, mra := range employeeMRAs {
			rakts, err := txEmployeeMainResearchAreaRepo.GetRAKTByMRAIDAndLanguageCode(ctx, mra.ID)
			if err != nil {
				return err
			}

			employeeMRAs[index].KeyTopics = rakts
		}

		return nil
	})

	resp := make([]*dtos.EmployeeMainResearchAreaResponse, len(employeeMRAs))
	for index, mra := range employeeMRAs {
		resp[index] = mappers.MapEmployeeMainResearchAreaDomainToResponseDTO(mra)
	}

	return resp, nil
}
