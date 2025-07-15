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

type pgEmployeeDegreeRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeDegreeRepository(store *Store) repositories.EmployeeDegreeRepository {
	return &pgEmployeeDegreeRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeDegreeRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeDegreeRepository {
	return &pgEmployeeDegreeRepository{
		queries: q,
	}
}

func (r *pgEmployeeDegreeRepository) Create(ctx context.Context, employeeDegree *domain.EmployeeDegree) (*domain.EmployeeDegree, error) {
	employeeDegreeResult, err := r.queries.CreateEmployeeDegree(ctx, sqlc.CreateEmployeeDegreeParams{
		EmployeeID:     employeeDegree.EmployeeID,
		LanguageCode:   employeeDegree.LanguageCode,
		DegreeLevel:    employeeDegree.DegreeLevel,
		UniversityName: employeeDegree.UniversityName,
		Speciality:     employeeDegree.Speciality,
		DateStart: pgtype.Date{
			Time:  employeeDegree.DateStart,
			Valid: !employeeDegree.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeDegree.DateEnd,
			Valid: !employeeDegree.DateEnd.IsZero(),
		},
		GivenBy: pgtype.Text{
			Valid:  true,
			String: employeeDegree.GivenBy,
		},
		DateDegreeRecieved: pgtype.Date{
			Time:  employeeDegree.DateDegreeRecieved,
			Valid: !employeeDegree.DateDegreeRecieved.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee degree: %w", err))
	}

	employeeDegree.ID = employeeDegreeResult.ID
	employeeDegree.CreatedAt = employeeDegreeResult.CreatedAt.Time
	employeeDegree.UpdatedAt = employeeDegreeResult.UpdatedAt.Time

	return employeeDegree, nil
}

func (r *pgEmployeeDegreeRepository) Update(ctx context.Context, employeeDegree *domain.EmployeeDegree) (*domain.EmployeeDegree, error) {
	updateEmployeeDegreeResult, err := r.queries.UpdateEmployeeDegree(ctx, sqlc.UpdateEmployeeDegreeParams{
		ID:             employeeDegree.ID,
		DegreeLevel:    employeeDegree.DegreeLevel,
		UniversityName: employeeDegree.UniversityName,
		Speciality:     employeeDegree.Speciality,
		DateStart: pgtype.Date{
			Time:  employeeDegree.DateStart,
			Valid: !employeeDegree.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeDegree.DateEnd,
			Valid: !employeeDegree.DateEnd.IsZero(),
		},
		GivenBy: pgtype.Text{
			Valid:  true,
			String: employeeDegree.GivenBy,
		},
		DateDegreeRecieved: pgtype.Date{
			Time:  employeeDegree.DateDegreeRecieved,
			Valid: !employeeDegree.DateDegreeRecieved.IsZero(),
		},
	})

	employeeDegree.CreatedAt = updateEmployeeDegreeResult.CreatedAt.Time
	employeeDegree.UpdatedAt = updateEmployeeDegreeResult.UpdatedAt.Time

	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee degree: %w", err))
	}

	return employeeDegree, nil
}

func (r *pgEmployeeDegreeRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeDegree(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee degree: %w", err))
	}

	return nil
}

func (r *pgEmployeeDegreeRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeDegree, error) {
	employeeDegreeResult, err := r.queries.GetEmployeeDegreeByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee degree with give ID(%d): %w", id, err))
	}

	return &domain.EmployeeDegree{
		ID:                 employeeDegreeResult.ID,
		EmployeeID:         employeeDegreeResult.EmployeeID,
		LanguageCode:       employeeDegreeResult.LanguageCode,
		DegreeLevel:        employeeDegreeResult.DegreeLevel,
		UniversityName:     employeeDegreeResult.UniversityName,
		Speciality:         employeeDegreeResult.Speciality,
		DateStart:          employeeDegreeResult.DateStart.Time,
		DateEnd:            employeeDegreeResult.DateEnd.Time,
		GivenBy:            employeeDegreeResult.GivenBy.String,
		DateDegreeRecieved: employeeDegreeResult.DateDegreeRecieved.Time,
		CreatedAt:          employeeDegreeResult.CreatedAt.Time,
		UpdatedAt:          employeeDegreeResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeDegreeRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeDegree, error) {
	employeeDegreesResult, err := r.queries.GetEmployeeDegreesByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeDegreesByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee degrees with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	} 

	employeeDegrees := make([]*domain.EmployeeDegree, len(employeeDegreesResult))
	for index, degree := range employeeDegreesResult {
		employeeDegrees[index] = &domain.EmployeeDegree{
			ID:                 degree.ID,
			EmployeeID:         degree.EmployeeID,
			LanguageCode:       degree.LanguageCode,
			DegreeLevel:        degree.DegreeLevel,
			UniversityName:     degree.UniversityName,
			Speciality:         degree.Speciality,
			DateStart:          degree.DateStart.Time,
			DateEnd:            degree.DateEnd.Time,
			GivenBy:            degree.UniversityName,
			DateDegreeRecieved: degree.DateDegreeRecieved.Time,
			CreatedAt:          degree.CreatedAt.Time,
			UpdatedAt:          degree.UpdatedAt.Time,
		}
	}

	return employeeDegrees, nil
}
