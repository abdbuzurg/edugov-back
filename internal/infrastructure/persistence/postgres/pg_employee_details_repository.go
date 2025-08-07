package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type pgEmployeeDetailsRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPGEmployeeDetailsRepository(store *Store) repositories.EmployeeDetailsRepository {
	return &pgEmployeeDetailsRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPGEmployeeDetailsRepositoryWithQueries(q *sqlc.Queries) repositories.EmployeeDetailsRepository {
	return &pgEmployeeDetailsRepository{
		queries: q,
	}
}

func (r *pgEmployeeDetailsRepository) Create(ctx context.Context, employeeDetails *domain.EmployeeDetails) (*domain.EmployeeDetails, error) {
	createdEmployeeDetails, err := r.queries.CreateEmployeeDetails(ctx, sqlc.CreateEmployeeDetailsParams{
		EmployeeID:   employeeDetails.EmployeeID,
		LanguageCode: employeeDetails.LanguageCode,
		Surname:      employeeDetails.Surname,
		Name:         employeeDetails.Name,
		Middlename: pgtype.Text{
			Valid:  true,
			String: employeeDetails.Middlename,
		},
		IsEmployeeDetailsNew: employeeDetails.IsEmployeeDetailsNew,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create emplyoee details: %w", err))
	}

	employeeDetails.ID = createdEmployeeDetails.ID
	employeeDetails.CreatedAt = createdEmployeeDetails.CreatedAt.Time
	employeeDetails.UpdatedAt = createdEmployeeDetails.UpdatedAt.Time

	return employeeDetails, nil
}

func (r *pgEmployeeDetailsRepository) Update(ctx context.Context, employeeDetails *domain.EmployeeDetails) (*domain.EmployeeDetails, error) {
	updatedEmployeeDetails, err := r.queries.UpdateEmployeeDetails(ctx, sqlc.UpdateEmployeeDetailsParams{
		Surname: employeeDetails.Surname,
		Name:    employeeDetails.Name,
		Middlename: pgtype.Text{
			Valid:  true,
			String: employeeDetails.Middlename,
		},
		ID:                   employeeDetails.ID,
		IsEmployeeDetailsNew: employeeDetails.IsEmployeeDetailsNew,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee_detials: %w", err))
	}

	employeeDetails.CreatedAt = updatedEmployeeDetails.CreatedAt.Time
	employeeDetails.UpdatedAt = updatedEmployeeDetails.UpdatedAt.Time

	return employeeDetails, nil
}

func (r *pgEmployeeDetailsRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeDetails(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete emplyoee details: %w", err))
	}

	return nil
}

func (r *pgEmployeeDetailsRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeDetails, error) {
	employeeDetailsResult, err := r.queries.GetEmployeeDetailsByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive emplyoee details: %w", err))
	}

	return &domain.EmployeeDetails{
		ID:                   employeeDetailsResult.ID,
		EmployeeID:           employeeDetailsResult.EmployeeID,
		LanguageCode:         employeeDetailsResult.LanguageCode,
		Surname:              employeeDetailsResult.Surname,
		Name:                 employeeDetailsResult.Name,
		Middlename:           employeeDetailsResult.Middlename.String,
		IsEmployeeDetailsNew: employeeDetailsResult.IsEmployeeDetailsNew,
		CreatedAt:            employeeDetailsResult.CreatedAt.Time,
		UpdatedAt:            employeeDetailsResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeDetailsRepository) GetByEmployeeID(ctx context.Context, employeeID int64) ([]*domain.EmployeeDetails, error) {
	employeeDetailsResults, err := r.queries.GetEmployeeDetailsByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee details and employeeID(%d): %w", employeeID, err))
	}

	employeeDetails := make([]*domain.EmployeeDetails, len(employeeDetailsResults))
	for index, details := range employeeDetailsResults {
		employeeDetails[index] = &domain.EmployeeDetails{
			ID:                   details.ID,
			EmployeeID:           details.EmployeeID,
			LanguageCode:         details.LanguageCode,
			Surname:              details.Surname,
			Name:                 details.Name,
			Middlename:           details.Middlename.String,
			IsEmployeeDetailsNew: details.IsEmployeeDetailsNew,
			CreatedAt:            details.CreatedAt.Time,
			UpdatedAt:            details.UpdatedAt.Time,
		}
	}

	return employeeDetails, nil
}
