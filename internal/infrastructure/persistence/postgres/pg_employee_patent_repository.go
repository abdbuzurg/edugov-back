package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeePatentRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeePatentRepository(store *Store) repositories.EmployeePatentRepository {
	return &pgEmployeePatentRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeePatentRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeePatentRepository {
	return &pgEmployeePatentRepository{
		queries: q,
	}
}

func (r *pgEmployeePatentRepository) Create(ctx context.Context, employeePatent *domain.EmployeePatent) (*domain.EmployeePatent, error) {
	employeePatentResult, err := r.queries.CreateEmployeePatent(ctx, sqlc.CreateEmployeePatentParams{
		EmployeeID:       employeePatent.EmployeeID,
		LanguageCode:     employeePatent.LanguageCode,
		PatentTitle:      employeePatent.PatentTitle,
		Description:      employeePatent.Description,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee patent: %w", err))
	}

	employeePatent.ID = employeePatentResult.ID
	employeePatent.CreatedAt = employeePatentResult.CreatedAt.Time
	employeePatent.UpdatedAt = employeePatentResult.UpdatedAt.Time

	return employeePatent, nil
}

func (r *pgEmployeePatentRepository) Update(ctx context.Context, employeePatent *domain.EmployeePatent) (*domain.EmployeePatent, error) {

	updateEmployeePatentResult, err := r.queries.UpdateEmployeePatent(ctx, sqlc.UpdateEmployeePatentParams{
		ID:          employeePatent.ID,
		PatentTitle: employeePatent.PatentTitle,
		Description: employeePatent.Description,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee patent: %w", err))
	}

	employeePatent.CreatedAt = updateEmployeePatentResult.CreatedAt.Time
	employeePatent.UpdatedAt = updateEmployeePatentResult.UpdatedAt.Time

	return employeePatent, nil
}

func (r *pgEmployeePatentRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeePatent(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee patent: %w", err))
	}

	return nil
}

func (r *pgEmployeePatentRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeePatent, error) {
	employeePatentResult, err := r.queries.GetEmployeePatentByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee patent with give ID(%d): %w", id, err))
	}

	return &domain.EmployeePatent{
		ID:               employeePatentResult.ID,
		EmployeeID:       employeePatentResult.EmployeeID,
		LanguageCode:     employeePatentResult.LanguageCode,
		PatentTitle:      employeePatentResult.PatentTitle,
		Description:      employeePatentResult.Description,
		CreatedAt:        employeePatentResult.CreatedAt.Time,
		UpdatedAt:        employeePatentResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeePatentRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeePatent, error) {
	employeePatentsResult, err := r.queries.GetEmployeePatentsByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeePatentsByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee patents with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	} 

	employeePatents := make([]*domain.EmployeePatent, len(employeePatentsResult))
	for index, employeePatentResult := range employeePatentsResult {
		employeePatents[index] = &domain.EmployeePatent{
			ID:               employeePatentResult.ID,
			EmployeeID:       employeePatentResult.EmployeeID,
			LanguageCode:     employeePatentResult.LanguageCode,
			PatentTitle:      employeePatentResult.PatentTitle,
			Description:      employeePatentResult.Description,
			CreatedAt:        employeePatentResult.CreatedAt.Time,
			UpdatedAt:        employeePatentResult.UpdatedAt.Time,
		}
	}

	return employeePatents, nil
}
