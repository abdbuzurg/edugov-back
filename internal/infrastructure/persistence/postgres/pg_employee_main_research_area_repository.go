package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeMainResearchAreaRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeMainResearchAreaRepository(store *Store) repositories.EmployeeMainResearchArea {
	return &pgEmployeeMainResearchAreaRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeMainResearchAreaRepositoryWithQueries(q *sqlc.Queries) repositories.EmployeeMainResearchArea {
	return &pgEmployeeMainResearchAreaRepository{
		queries: q,
	}
}

func (r *pgEmployeeMainResearchAreaRepository) CreateMRA(ctx context.Context, employeeMRA *domain.EmployeeMainResearchArea) (*domain.EmployeeMainResearchArea, error) {
	employeeMRAResult, err := r.queries.CreateEmployeeMainResearchArea(ctx, sqlc.CreateEmployeeMainResearchAreaParams{
		EmployeeID:   employeeMRA.EmployeeID,
		LanguageCode: employeeMRA.LanguageCode,
		Discipline:   employeeMRA.Discipline,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee main research area: %w", err))
	}

	employeeMRA.ID = employeeMRAResult.ID
	employeeMRA.CreatedAt = employeeMRAResult.CreatedAt.Time
	employeeMRA.UpdatedAt = employeeMRAResult.UpdatedAt.Time

	return employeeMRA, nil
}

func (r *pgEmployeeMainResearchAreaRepository) CreateRAKT(ctx context.Context, rakt *domain.ResearchAreaKeyTopic) (*domain.ResearchAreaKeyTopic, error) {
	raktResult, err := r.queries.CreateEmployeeMainResearchAreaKeyTopic(ctx, sqlc.CreateEmployeeMainResearchAreaKeyTopicParams{
		EmployeeMainResearchAreaID: rakt.EmployeeMainResearchAreaID,
		KeyTopicTitle:              rakt.KeyTopicTitle,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create main research area key topic: %w", err))
	}

	rakt.ID = raktResult.ID
	rakt.CreatedAt = raktResult.CreatedAt.Time
	rakt.UpdatedAt = raktResult.UpdatedAt.Time

	return rakt, nil
}

func (r *pgEmployeeMainResearchAreaRepository) UpdateMRA(ctx context.Context, employeeMRA *domain.EmployeeMainResearchArea) (*domain.EmployeeMainResearchArea, error) {
	updateEmployeeMRAResult, err := r.queries.UpdateEmployeeMainResearchArea(ctx, sqlc.UpdateEmployeeMainResearchAreaParams{
		ID:         employeeMRA.ID,
		Area:       employeeMRA.Area,
		Discipline: employeeMRA.Discipline,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee main research are: %w", err))
	}

	employeeMRA.CreatedAt = updateEmployeeMRAResult.CreatedAt.Time
	employeeMRA.UpdatedAt = updateEmployeeMRAResult.UpdatedAt.Time

	return employeeMRA, nil
}

func (r *pgEmployeeMainResearchAreaRepository) UpdateRAKT(ctx context.Context, rakt *domain.ResearchAreaKeyTopic) (*domain.ResearchAreaKeyTopic, error) {
	updatedRAKT, err := r.queries.UpdateEmployeeMainResearchAreaKeyTopic(ctx, sqlc.UpdateEmployeeMainResearchAreaKeyTopicParams{
		ID:            rakt.ID,
		KeyTopicTitle: rakt.KeyTopicTitle,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update main research area key topic: %w", err))
	}

	rakt.CreatedAt = updatedRAKT.CreatedAt.Time
	rakt.UpdatedAt = updatedRAKT.UpdatedAt.Time

	return rakt, nil
}

func (r *pgEmployeeMainResearchAreaRepository) DeleteMRA(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeMainResearchAreaKeyTopic(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee main research area: %w", err))
	}

	return nil
}

func (r *pgEmployeeMainResearchAreaRepository) DeleteRAKT(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeMainResearchAreaKeyTopic(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete main research area key topic: %w", err))
	}

	return nil
}

func (r *pgEmployeeMainResearchAreaRepository) GetMRAByID(ctx context.Context, id int64) (*domain.EmployeeMainResearchArea, error) {
	employeeMRAResult, err := r.queries.GetEmployeeMainResearchAreaByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee main research are with given ID(%d): %w", id, err))
	}

	return &domain.EmployeeMainResearchArea{
		ID:           employeeMRAResult.ID,
		EmployeeID:   employeeMRAResult.EmployeeID,
		LanguageCode: employeeMRAResult.LanguageCode,
		Area:         employeeMRAResult.Area,
		Discipline:   employeeMRAResult.Discipline,
		CreatedAt:    employeeMRAResult.CreatedAt.Time,
		UpdatedAt:    employeeMRAResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeMainResearchAreaRepository) GetRAKTByID(ctx context.Context, id int64) (*domain.ResearchAreaKeyTopic, error) {
	rakt, err := r.queries.GetEmployeeMainResearchAreaKeyTopicByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive main research are key topic by ID(%d): %w", id, err))
	}

	return &domain.ResearchAreaKeyTopic{
		ID:                         rakt.ID,
		EmployeeMainResearchAreaID: rakt.EmployeeMainResearchAreaID,
		KeyTopicTitle:              rakt.KeyTopicTitle,
		CreatedAt:                  rakt.CreatedAt.Time,
		UpdatedAt:                  rakt.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeMainResearchAreaRepository) GetMRAByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeMainResearchArea, error) {
	employeeMRAsResult, err := r.queries.GetEmployeeMainResearchAreasByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeMainResearchAreasByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee main research areas by given ID(%d) and language_code(%s): %w", employeeID, langCode, err))
	}

	employeeMRAs := make([]*domain.EmployeeMainResearchArea, len(employeeMRAsResult))
	for indexMRA, employeeMRA := range employeeMRAsResult {
		employeeMRAs[indexMRA] = &domain.EmployeeMainResearchArea{
			ID:           employeeMRA.ID,
			LanguageCode: employeeMRA.LanguageCode,
			Area:         employeeMRA.Area,
			Discipline:   employeeMRA.Discipline,
			CreatedAt:    employeeMRA.CreatedAt.Time,
			UpdatedAt:    employeeMRA.UpdatedAt.Time,
		}
	}

	return employeeMRAs, nil
}

func (r *pgEmployeeMainResearchAreaRepository) GetRAKTByMRAIDAndLanguageCode(ctx context.Context, employeeMRAID int64) ([]*domain.ResearchAreaKeyTopic, error) {
	raktsResult, err := r.queries.GetEmployeeMainResearchAreaKeyTopicsByEmployeeMainResearchAreaIDAndLanguageCode(ctx, employeeMRAID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive main research are key topic by employee MRA ID(%d): %w", employeeMRAID, err))
	}

	rakts := make([]*domain.ResearchAreaKeyTopic, len(raktsResult))
	for index, rakt := range raktsResult {
		rakts[index] = &domain.ResearchAreaKeyTopic{
			ID:            rakt.ID,
      EmployeeMainResearchAreaID: rakt.EmployeeMainResearchAreaID,
			KeyTopicTitle: rakt.KeyTopicTitle,
			CreatedAt:     rakt.CreatedAt.Time,
			UpdatedAt:     rakt.UpdatedAt.Time,
		}
	}

  return rakts, nil
}
