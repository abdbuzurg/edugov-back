package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeePublicationRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeePublicationRepository(store *Store) repositories.EmployeePublicationRepository {
	return &pgEmployeePublicationRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeePublicationRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeePublicationRepository {
	return &pgEmployeePublicationRepository{
		queries: q,
	}
}

func (r *pgEmployeePublicationRepository) Create(ctx context.Context, employeePublication *domain.EmployeePublication) (*domain.EmployeePublication, error) {
	employeePublicationResult, err := r.queries.CreateEmployeePublication(ctx, sqlc.CreateEmployeePublicationParams{
		EmployeeID:        employeePublication.EmployeeID,
		LanguageCode:      employeePublication.LanguageCode,
		PublicationTitle:  employeePublication.PublicationTitle,
		LinkToPublication: employeePublication.LinkToPublication,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee publicaiton: %w", err))
	}

	employeePublication.ID = employeePublicationResult.ID
	employeePublication.CreatedAt = employeePublicationResult.CreatedAt.Time
	employeePublication.UpdatedAt = employeePublicationResult.UpdatedAt.Time

	return employeePublication, nil
}

func (r *pgEmployeePublicationRepository) Update(ctx context.Context, employeePublication *domain.EmployeePublication) (*domain.EmployeePublication, error) {
	updateEmployeePublicationResult, err := r.queries.UpdateEmployeePublication(ctx, sqlc.UpdateEmployeePublicationParams{
		ID:                employeePublication.ID,
		PublicationTitle:  employeePublication.PublicationTitle,
		LinkToPublication: employeePublication.LinkToPublication,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee publicaiton: %w", err))
	}

	employeePublication.CreatedAt = updateEmployeePublicationResult.CreatedAt.Time
	employeePublication.UpdatedAt = updateEmployeePublicationResult.UpdatedAt.Time

	return employeePublication, nil
}

func (r *pgEmployeePublicationRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeePublication(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee publicaiton: %w", err))
	}

	return nil
}

func (r *pgEmployeePublicationRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeePublication, error) {
	employeePublicationResult, err := r.queries.GetEmployeePublicationByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee publicaiton with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeePublication{
		ID:               employeePublicationResult.ID,
		EmployeeID:       employeePublicationResult.EmployeeID,
		LanguageCode:     employeePublicationResult.LanguageCode,
		PublicationTitle: employeePublicationResult.PublicationTitle,
		CreatedAt:        employeePublicationResult.CreatedAt.Time,
		UpdatedAt:        employeePublicationResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeePublicationRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeePublication, error) {
	employeePublicationsResult, err := r.queries.GetEmployeePublicationsByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeePublicationsByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee publicaitons with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	}

	employeePublications := make([]*domain.EmployeePublication, len(employeePublicationsResult))
	for index, publication := range employeePublicationsResult {
		employeePublications[index] = &domain.EmployeePublication{
			ID:                publication.ID,
			EmployeeID:        publication.EmployeeID,
			LanguageCode:      publication.LanguageCode,
			PublicationTitle:  publication.PublicationTitle,
			LinkToPublication: publication.LinkToPublication,
			CreatedAt:         publication.CreatedAt.Time,
			UpdatedAt:         publication.UpdatedAt.Time,
		}
	}

	return employeePublications, nil
}
