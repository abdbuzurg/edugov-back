package postgres

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
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

func (r *pgEmployeeRepository) GetPersonnelIDsPaginated(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) ([]int64, error) {
	personnel, err := r.queries.GetPersonnelPaginated(ctx, sqlc.GetPersonnelPaginatedParams{
		LanguageCode: filter.LanguageCode,
		Uid: pgtype.Text{
			Valid:  true,
			String: filter.UID,
		},
		Name: pgtype.Text{
			Valid:  true,
			String: filter.Name,
		},
		Surname: pgtype.Text{
			Valid:  true,
			String: filter.Surname,
		},
		Middlename: pgtype.Text{
			Valid:  true,
			String: filter.Middleware,
		},
		Speciality: pgtype.Text{
			Valid:  true,
			String: filter.Speciality,
		},
		Page:  int32(filter.Page),
		Limit: int32(filter.Limit),
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive paginated personnel data (page - %d, limit - %d): %w", filter.Page, filter.Limit, err))
	}

	return personnel, nil
}
