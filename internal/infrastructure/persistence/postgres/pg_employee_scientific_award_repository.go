package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeScientificAwardRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeScientificAwardRepository(store *Store) repositories.EmployeeScientificAwardRepository {
	return &pgEmployeeScientificAwardRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeScientificAwardRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeScientificAwardRepository {
	return &pgEmployeeScientificAwardRepository{
		queries: q,
	}
}

func (r *pgEmployeeScientificAwardRepository) Create(ctx context.Context, employeeScientificAward *domain.EmployeeScientificAward) (*domain.EmployeeScientificAward, error) {
	employeeScientificAwardResult, err := r.queries.CreateEmployeeScientificAward(ctx, sqlc.CreateEmployeeScientificAwardParams{
		EmployeeID:           employeeScientificAward.EmployeeID,
		LanguageCode:         employeeScientificAward.LanguageCode,
		ScientificAwardTitle: employeeScientificAward.ScientificAwardTitle,
		GivenBy:              employeeScientificAward.GivenBy,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee scientific award: %w", err))
	}

	employeeScientificAward.ID = employeeScientificAwardResult.ID
	employeeScientificAward.CreatedAt = employeeScientificAwardResult.CreatedAt.Time
	employeeScientificAward.UpdatedAt = employeeScientificAwardResult.UpdatedAt.Time

	return employeeScientificAward, nil
}

func (r *pgEmployeeScientificAwardRepository) Update(ctx context.Context, employeeScientificAward *domain.EmployeeScientificAward) (*domain.EmployeeScientificAward, error) {
	updateEmployeeScientificAwardResult, err := r.queries.UpdateEmployeeScientificAward(ctx, sqlc.UpdateEmployeeScientificAwardParams{
		ID:                   employeeScientificAward.ID,
		ScientificAwardTitle: employeeScientificAward.ScientificAwardTitle,
		GivenBy:              employeeScientificAward.GivenBy,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee scientific award: %w", err))
	}

	employeeScientificAward.CreatedAt = updateEmployeeScientificAwardResult.CreatedAt.Time
	employeeScientificAward.UpdatedAt = updateEmployeeScientificAwardResult.UpdatedAt.Time

	return employeeScientificAward, nil
}

func (r *pgEmployeeScientificAwardRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeScientificAward(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee scientific award: %w", err))
	}

	return nil
}

func (r *pgEmployeeScientificAwardRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeScientificAward, error) {
	employeeScientificAwardResult, err := r.queries.GetEmployeeScientificAwardByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee scientific award with give ID(%d): %w", id, err))
	}

	return &domain.EmployeeScientificAward{
		ID:                   employeeScientificAwardResult.ID,
		EmployeeID:           employeeScientificAwardResult.EmployeeID,
		LanguageCode:         employeeScientificAwardResult.LanguageCode,
		ScientificAwardTitle: employeeScientificAwardResult.ScientificAwardTitle,
		GivenBy:              employeeScientificAwardResult.GivenBy,
		CreatedAt:            employeeScientificAwardResult.CreatedAt.Time,
		UpdatedAt:            employeeScientificAwardResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeScientificAwardRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeScientificAward, error) {
	employeeScientificAwardsResult, err := r.queries.GetEmployeeScientificAwardsByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeScientificAwardsByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee scientific awards with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	}

	employeeScientificAwards := make([]*domain.EmployeeScientificAward, len(employeeScientificAwardsResult))
	for index, scientificAward := range employeeScientificAwardsResult {
		employeeScientificAwards[index] = &domain.EmployeeScientificAward{
			ID:                   scientificAward.ID,
			EmployeeID:           scientificAward.EmployeeID,
			LanguageCode:         scientificAward.LanguageCode,
			ScientificAwardTitle: scientificAward.ScientificAwardTitle,
			GivenBy:              scientificAward.GivenBy,
			CreatedAt:            scientificAward.CreatedAt.Time,
			UpdatedAt:            scientificAward.UpdatedAt.Time,
		}
	}

	return employeeScientificAwards, nil
}
