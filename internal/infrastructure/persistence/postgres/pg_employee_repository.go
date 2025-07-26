package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeRepository(store *Store) repositories.EmployeeRepository {
	return &pgEmployeeRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeRepository {
	return &pgEmployeeRepository{
		queries: q,
	}
}

func (r *pgEmployeeRepository) Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	employeeResult, err := r.queries.CreateEmployee(ctx, employee.UniqueID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee: %w", err))
	}

	employee.ID = employeeResult.ID
	employee.CreatedAt = employeeResult.CreatedAt.Time
	employee.UpdatedAt = employeeResult.UpdatedAt.Time

	return employee, nil
}

func (r *pgEmployeeRepository) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteEmployeeDetails(ctx, id)
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed deleting employee: %w", err))
	}

	return nil
}

func (r *pgEmployeeRepository) GetByID(ctx context.Context, id int64) (*domain.Employee, error) {
	employeeResult, err := r.queries.GetEmployeeByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee by given ID(%d): %w", id, err))
	}

	return &domain.Employee{
		ID:        employeeResult.ID,
		UniqueID:  employeeResult.UniqueID,
		CreatedAt: employeeResult.CreatedAt.Time,
		UpdatedAt: employeeResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeRepository) GetByUniqueID(ctx context.Context, uniqueIdentifer string) (*domain.Employee, error) {
	employeeResult, err := r.queries.GetEmployeeByUniqueIdentifier(ctx, uniqueIdentifer)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee by given uniqueIdentifer(%s): %w", uniqueIdentifer, err))
	}

	return &domain.Employee{
		ID:        employeeResult.ID,
		UniqueID:  employeeResult.UniqueID,
		CreatedAt: employeeResult.CreatedAt.Time,
		UpdatedAt: employeeResult.UpdatedAt.Time,
	}, nil
}
