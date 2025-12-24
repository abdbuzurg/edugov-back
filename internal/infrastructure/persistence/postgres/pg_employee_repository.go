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
	employeeResult, err := r.queries.CreateEmployee(ctx, sqlc.CreateEmployeeParams{
		UniqueID: employee.UniqueID,
		UserID: pgtype.Int8{
			Int64: employee.UserID,
			Valid: true,
		},
		Gender: pgtype.Text{
			String: employee.Gender,
			Valid:  true,
		},
		Tin: pgtype.Text{
			String: employee.Tin,
			Valid:  true,
		},
	})
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
		Gender:    employeeResult.Gender.String,
		Tin:       employeeResult.Tin.String,
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
		Gender:    employeeResult.Gender.String,
		Tin:       employeeResult.Tin.String,
		CreatedAt: employeeResult.CreatedAt.Time,
		UpdatedAt: employeeResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeRepository) GetByUserID(ctx context.Context, userID int64) (*domain.Employee, error) {
	employeeResult, err := r.queries.GetEmployeeByUserID(ctx, pgtype.Int8{
		Int64: userID,
		Valid: true,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee by given user_id(%d): %w", userID, err))
	}

	return &domain.Employee{
		ID:        employeeResult.ID,
		UniqueID:  employeeResult.UniqueID,
		Gender:    employeeResult.Gender.String,
		Tin:       employeeResult.Tin.String,
		CreatedAt: employeeResult.CreatedAt.Time,
		UpdatedAt: employeeResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeRepository) GetPersonnelIDsPaginated(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) ([]*repositories.GetPersonnelPaginatedQueryResult, error) {
	personnelResult, err := r.queries.GetPersonnelPaginated(ctx, sqlc.GetPersonnelPaginatedParams{
		LanguageCode: filter.LanguageCode,
		Uid:          filter.UID,
		Name:         filter.Name,
		Surname:      filter.Surname,
		Middlename:   filter.Middlename,
		Workplace:    filter.Workplace,
		Page:         int32((filter.Page - 1) * filter.Limit),
		Limit:        int32(filter.Limit),
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive paginated personnel data (page - %d, limit - %d): %w", filter.Page, filter.Limit, err))
	}

	personnel := make([]*repositories.GetPersonnelPaginatedQueryResult, len(personnelResult))
	for index := range personnel {
		personnel[index] = &repositories.GetPersonnelPaginatedQueryResult{
			EmployeeID:            personnelResult[index].ID,
			Surname:               personnelResult[index].Surname,
			Name:                  personnelResult[index].Name,
			Middlename:            personnelResult[index].Middlename.String,
			Currentworkplace:      personnelResult[index].CurrentWorkplace.String,
			Highestacademicdegree: personnelResult[index].HighestAcademicDegree.String,
			Speciality:            personnelResult[index].Speciality.String,
			UniqueID:              personnelResult[index].UniqueID,
		}
	}

	return personnel, nil
}

func (r *pgEmployeeRepository) CountPersonnel(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) (int64, error) {
	count, err := r.queries.CountPersonnel(ctx, sqlc.CountPersonnelParams{
		LanguageCode: filter.LanguageCode,
		Uid:          filter.UID,
		Name:         filter.Name,
		Surname:      filter.Surname,
		Middlename:   filter.Middlename,
		Workplace:    filter.Workplace,
	})
	if err != nil {
		return 0, custom_errors.InternalServerError(fmt.Errorf("could not count personnel: %w", err))
	}

	return count, nil
}
