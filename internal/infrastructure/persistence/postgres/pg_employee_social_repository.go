package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeSocialRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeSocialRepository(store *Store) repositories.EmployeeSocialRepository {
	return &pgEmployeeSocialRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeSocialRepositoryWithQueries(q *sqlc.Queries) repositories.EmployeeSocialRepository {
	return &pgEmployeeSocialRepository{
		queries: q,
	}
}

func (r *pgEmployeeSocialRepository) Create(ctx context.Context, employeeSocial *domain.EmployeeSocial) (*domain.EmployeeSocial, error) {
	employeeSocialResult, err := r.queries.CreateEmployeeSocial(ctx, sqlc.CreateEmployeeSocialParams{
		EmployeeID:   employeeSocial.EmployeeID,
		SocialName:   employeeSocial.SocialName,
		LinkToSocial: employeeSocial.LinkToSocial,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee social: %w", err))
	}

	employeeSocial.ID = employeeSocialResult.ID
	employeeSocial.CreatedAt = employeeSocialResult.CreatedAt.Time
	employeeSocial.UpdatedAt = employeeSocialResult.UpdatedAt.Time

	return employeeSocial, nil
}

func (r *pgEmployeeSocialRepository) Update(ctx context.Context, employeeSocial *domain.EmployeeSocial) (*domain.EmployeeSocial, error) {
	updateEmployeeSocialResult, err := r.queries.UpdateEmployeeSocial(ctx, sqlc.UpdateEmployeeSocialParams{
		ID:           employeeSocial.ID,
		SocialName:   employeeSocial.SocialName,
		LinkToSocial: employeeSocial.LinkToSocial,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee social: %w", err))
	}

	employeeSocial.CreatedAt = updateEmployeeSocialResult.CreatedAt.Time
	employeeSocial.UpdatedAt = updateEmployeeSocialResult.UpdatedAt.Time

	return employeeSocial, nil
}

func (r *pgEmployeeSocialRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeSocial(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee social: %w", err))
	}

	return nil
}

func (r *pgEmployeeSocialRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeSocial, error) {
	employeeSocialResult, err := r.queries.GetEmployeeSocialByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee social with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeeSocial{
		ID:           employeeSocialResult.ID,
		EmployeeID:   employeeSocialResult.EmployeeID,
		SocialName:   employeeSocialResult.SocialName,
		LinkToSocial: employeeSocialResult.LinkToSocial,
		CreatedAt:    employeeSocialResult.CreatedAt.Time,
		UpdatedAt:    employeeSocialResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeSocialRepository) GetByEmployeeID(ctx context.Context, employeeID int64) ([]*domain.EmployeeSocial, error) {
	employeeSocialsResult, err := r.queries.GetEmployeeSocialsByEmployeeIDAndLanguageCode(ctx, employeeID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee socials with given EmployeeID(%d): %w", employeeID, err))
	} 

	employeeSocials := make([]*domain.EmployeeSocial, len(employeeSocialsResult))
	for index, degree := range employeeSocialsResult {
		employeeSocials[index] = &domain.EmployeeSocial{
			ID:           degree.ID,
			EmployeeID:   degree.EmployeeID,
			SocialName:   degree.SocialName,
			LinkToSocial: degree.LinkToSocial,
			CreatedAt:    degree.CreatedAt.Time,
			UpdatedAt:    degree.UpdatedAt.Time,
		}
	}

	return employeeSocials, nil
}
