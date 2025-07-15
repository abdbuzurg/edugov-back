package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeResearchActivityRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeResearchActivityRepository(store *Store) repositories.EmployeeResearchActivityRepository {
	return &pgEmployeeResearchActivityRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeResearchActivityRepositoryWithQueries(q *sqlc.Queries) repositories.EmployeeResearchActivityRepository {
	return &pgEmployeeResearchActivityRepository{
		queries: q,
	}
}

func (r *pgEmployeeResearchActivityRepository) Create(ctx context.Context, employeeResearchActivity *domain.EmployeeResearchActivity) (*domain.EmployeeResearchActivity, error) {
	employeeResearchActivityResult, err := r.queries.CreateEmployeeResearchActivity(ctx, sqlc.CreateEmployeeResearchActivityParams{
		EmployeeID:            employeeResearchActivity.EmployeeID,
		LanguageCode:          employeeResearchActivity.LanguageCode,
		ResearchActivityTitle: employeeResearchActivity.ResearchActivityTitle,
		EmployeeRole:          employeeResearchActivity.EmployeeRole,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee research activity: %w", err))
	}

	employeeResearchActivity.ID = employeeResearchActivityResult.ID
	employeeResearchActivity.CreatedAt = employeeResearchActivityResult.CreatedAt.Time
	employeeResearchActivity.UpdatedAt = employeeResearchActivityResult.UpdatedAt.Time

	return employeeResearchActivity, nil
}

func (r *pgEmployeeResearchActivityRepository) Update(ctx context.Context, employeeResearchActivity *domain.EmployeeResearchActivity) (*domain.EmployeeResearchActivity, error) {
	updateEmployeeResearchActivityResult, err := r.queries.UpdateEmployeeResearchActivity(ctx, sqlc.UpdateEmployeeResearchActivityParams{
		ID:                    employeeResearchActivity.ID,
		ResearchActivityTitle: employeeResearchActivity.ResearchActivityTitle,
		EmployeeRole:          employeeResearchActivity.EmployeeRole,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee research activity: %w", err))
	}

	employeeResearchActivity.CreatedAt = updateEmployeeResearchActivityResult.CreatedAt.Time
	employeeResearchActivity.UpdatedAt = updateEmployeeResearchActivityResult.UpdatedAt.Time

	return employeeResearchActivity, nil
}

func (r *pgEmployeeResearchActivityRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeResearchActivity(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee research activity: %w", err))
	}

	return nil
}

func (r *pgEmployeeResearchActivityRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeResearchActivity, error) {
	employeeResearchActivityResult, err := r.queries.GetEmployeeResearchActivityByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee research activity with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeeResearchActivity{
		ID:                    employeeResearchActivityResult.ID,
		EmployeeID:            employeeResearchActivityResult.EmployeeID,
		LanguageCode:          employeeResearchActivityResult.LanguageCode,
		ResearchActivityTitle: employeeResearchActivityResult.ResearchActivityTitle,
		EmployeeRole:          employeeResearchActivityResult.EmployeeRole,
		CreatedAt:             employeeResearchActivityResult.CreatedAt.Time,
		UpdatedAt:             employeeResearchActivityResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeResearchActivityRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeResearchActivity, error) {
	employeeResearchActivitysResult, err := r.queries.GetEmployeeResearchActivitiesByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeResearchActivitiesByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee research activitys with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	} 

	employeeResearchActivitys := make([]*domain.EmployeeResearchActivity, len(employeeResearchActivitysResult))
	for index, researchActivity := range employeeResearchActivitysResult {
		employeeResearchActivitys[index] = &domain.EmployeeResearchActivity{
			ID:                    researchActivity.ID,
			EmployeeID:            researchActivity.EmployeeID,
			LanguageCode:          researchActivity.LanguageCode,
			ResearchActivityTitle: researchActivity.ResearchActivityTitle,
			EmployeeRole:          researchActivity.EmployeeRole,
			CreatedAt:             researchActivity.CreatedAt.Time,
			UpdatedAt:             researchActivity.UpdatedAt.Time,
		}
	}

	return employeeResearchActivitys, nil
}
