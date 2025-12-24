package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type pgEmployeeWorkExperienceRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeWorkExperienceRepository(store *Store) repositories.EmployeeWorkExperienceRepository {
	return &pgEmployeeWorkExperienceRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeWorkExperienceRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeWorkExperienceRepository {
	return &pgEmployeeWorkExperienceRepository{
		queries: q,
	}
}

func (r *pgEmployeeWorkExperienceRepository) Create(ctx context.Context, employeeWorkExperience *domain.EmployeeWorkExperience) (*domain.EmployeeWorkExperience, error) {
	employeeWorkExperienceResult, err := r.queries.CreateEmployeeWorkExperience(ctx, sqlc.CreateEmployeeWorkExperienceParams{
		EmployeeID:   employeeWorkExperience.EmployeeID,
		LanguageCode: employeeWorkExperience.LanguageCode,
		JobTitle:     employeeWorkExperience.JobTitle,
		Workplace:    employeeWorkExperience.Workplace,
		Description:  employeeWorkExperience.Description,
		DateStart: pgtype.Date{
			Time:  employeeWorkExperience.DateStart,
			Valid: !employeeWorkExperience.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeWorkExperience.DateEnd,
			Valid: !employeeWorkExperience.DateEnd.IsZero(),
		},
		OnGoing: employeeWorkExperience.Ongoing,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee work experience: %w", err))
	}

	employeeWorkExperience.ID = employeeWorkExperienceResult.ID
	employeeWorkExperience.CreatedAt = employeeWorkExperienceResult.CreatedAt.Time
	employeeWorkExperience.UpdatedAt = employeeWorkExperienceResult.UpdatedAt.Time

	return employeeWorkExperience, nil
}

func (r *pgEmployeeWorkExperienceRepository) Update(ctx context.Context, employeeWorkExperience *domain.EmployeeWorkExperience) (*domain.EmployeeWorkExperience, error) {
	updateEmployeeWorkExperienceResult, err := r.queries.UpdateEmployeeWorkExperience(ctx, sqlc.UpdateEmployeeWorkExperienceParams{
		ID:          employeeWorkExperience.ID,
		JobTitle:    employeeWorkExperience.JobTitle,
		Workplace:   employeeWorkExperience.Workplace,
		Description: employeeWorkExperience.Description,
		DateStart: pgtype.Date{
			Time:  employeeWorkExperience.DateStart,
			Valid: !employeeWorkExperience.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeWorkExperience.DateEnd,
			Valid: !employeeWorkExperience.DateEnd.IsZero(),
		},
		OnGoing: employeeWorkExperience.Ongoing,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee work experience: %w", err))
	}

	employeeWorkExperience.CreatedAt = updateEmployeeWorkExperienceResult.CreatedAt.Time
	employeeWorkExperience.UpdatedAt = updateEmployeeWorkExperienceResult.UpdatedAt.Time

	return employeeWorkExperience, nil
}

func (r *pgEmployeeWorkExperienceRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeWorkExperience(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee work experience: %w", err))
	}

	return nil
}

func (r *pgEmployeeWorkExperienceRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeWorkExperience, error) {
	employeeWorkExperienceResult, err := r.queries.GetEmployeeWorkExperienceByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee work experience with give ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.EmployeeWorkExperience{
		ID:           employeeWorkExperienceResult.ID,
		EmployeeID:   employeeWorkExperienceResult.EmployeeID,
		LanguageCode: employeeWorkExperienceResult.LanguageCode,
		Workplace:    employeeWorkExperienceResult.Workplace,
		Description:  employeeWorkExperienceResult.Description,
		JobTitle:     employeeWorkExperienceResult.JobTitle,
		DateStart:    employeeWorkExperienceResult.DateStart.Time,
		DateEnd:      employeeWorkExperienceResult.DateEnd.Time,
		Ongoing:      employeeWorkExperienceResult.OnGoing,
		CreatedAt:    employeeWorkExperienceResult.CreatedAt.Time,
		UpdatedAt:    employeeWorkExperienceResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeWorkExperienceRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeWorkExperience, error) {
	employeeWorkExperiencesResult, err := r.queries.GetEmployeeWorkExperiencesByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeWorkExperiencesByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee work experiences with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	}

	employeeWorkExperiences := make([]*domain.EmployeeWorkExperience, len(employeeWorkExperiencesResult))
	for index, workExperience := range employeeWorkExperiencesResult {
		employeeWorkExperiences[index] = &domain.EmployeeWorkExperience{
			ID:           workExperience.ID,
			EmployeeID:   workExperience.EmployeeID,
			LanguageCode: workExperience.LanguageCode,
			Workplace:    workExperience.Workplace,
			Description:  workExperience.Description,
			JobTitle:     workExperience.JobTitle,
			DateStart:    workExperience.DateStart.Time,
			DateEnd:      workExperience.DateEnd.Time,
			Ongoing:      workExperience.OnGoing,
			CreatedAt:    workExperience.CreatedAt.Time,
			UpdatedAt:    workExperience.UpdatedAt.Time,
		}
	}

	return employeeWorkExperiences, nil
}

func (r *pgEmployeeWorkExperienceRepository) ListUniqueOngoingWorkplaces(ctx context.Context, langCode string) ([]string, error) {
	workplaces, err := r.queries.ListUniqueWorkplaces(ctx, langCode)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive unique workplaces: %w", err))
	}

	result := []string{}
	for index := range workplaces {
		if workplaces[index].Valid {
			result = append(result, workplaces[index].String)
		}
	}

	return result, nil
}
