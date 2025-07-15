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

type pgEmployeeRefresherCourseRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeRefresherCourseRepository(store *Store) repositories.EmployeeRefresherCourseRepository {
	return &pgEmployeeRefresherCourseRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeRefresherCourseRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeRefresherCourseRepository {
	return &pgEmployeeRefresherCourseRepository{
		queries: q,
	}
}

func (r *pgEmployeeRefresherCourseRepository) Create(ctx context.Context, employeeRefresherCourse *domain.EmployeeRefresherCourse) (*domain.EmployeeRefresherCourse, error) {
	employeeRefresherCourseResult, err := r.queries.CreateEmployeeRefresherCourse(ctx, sqlc.CreateEmployeeRefresherCourseParams{
		EmployeeID:   employeeRefresherCourse.EmployeeID,
		LanguageCode: employeeRefresherCourse.LanguageCode,
		CourseTitle:  employeeRefresherCourse.CourseTitle,
		DateStart: pgtype.Date{
			Time:  employeeRefresherCourse.DateStart,
			Valid: !employeeRefresherCourse.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeRefresherCourse.DateEnd,
			Valid: !employeeRefresherCourse.DateEnd.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee refresher course: %w", err))
	}

	employeeRefresherCourse.ID = employeeRefresherCourseResult.ID
	employeeRefresherCourse.CreatedAt = employeeRefresherCourseResult.CreatedAt.Time
	employeeRefresherCourse.UpdatedAt = employeeRefresherCourseResult.UpdatedAt.Time

	return employeeRefresherCourse, nil
}

func (r *pgEmployeeRefresherCourseRepository) Update(ctx context.Context, employeeRefresherCourse *domain.EmployeeRefresherCourse) (*domain.EmployeeRefresherCourse, error) {
	updateEmployeeRefresherCourseResult, err := r.queries.UpdateEmployeeRefresherCourse(ctx, sqlc.UpdateEmployeeRefresherCourseParams{
		ID:          employeeRefresherCourse.ID,
		CourseTitle: employeeRefresherCourse.CourseTitle,
		DateStart: pgtype.Date{
			Time:  employeeRefresherCourse.DateStart,
			Valid: !employeeRefresherCourse.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  employeeRefresherCourse.DateEnd,
			Valid: !employeeRefresherCourse.DateEnd.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee refresher course: %w", err))
	}

	employeeRefresherCourse.CreatedAt = updateEmployeeRefresherCourseResult.CreatedAt.Time
	employeeRefresherCourse.UpdatedAt = updateEmployeeRefresherCourseResult.UpdatedAt.Time

	return employeeRefresherCourse, nil
}

func (r *pgEmployeeRefresherCourseRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeRefresherCourse(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee refresher course: %w", err))
	}

	return nil
}

func (r *pgEmployeeRefresherCourseRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeRefresherCourse, error) {
	employeeRefresherCourseResult, err := r.queries.GetEmployeeRefresherCourseByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee refresher course with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeeRefresherCourse{
		ID:           employeeRefresherCourseResult.ID,
		EmployeeID:   employeeRefresherCourseResult.EmployeeID,
		LanguageCode: employeeRefresherCourseResult.LanguageCode,
		CourseTitle:  employeeRefresherCourseResult.CourseTitle,
		DateStart:    employeeRefresherCourseResult.DateStart.Time,
		DateEnd:      employeeRefresherCourseResult.DateEnd.Time,
		CreatedAt:    employeeRefresherCourseResult.CreatedAt.Time,
		UpdatedAt:    employeeRefresherCourseResult.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeRefresherCourseRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeRefresherCourse, error) {
	employeeRefresherCoursesResult, err := r.queries.GetEmployeeRefresherCoursesByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeRefresherCoursesByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee refresher courses with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
  }

	employeeRefresherCourses := make([]*domain.EmployeeRefresherCourse, len(employeeRefresherCoursesResult))
	for index, degree := range employeeRefresherCoursesResult {
		employeeRefresherCourses[index] = &domain.EmployeeRefresherCourse{
			ID:           degree.ID,
			EmployeeID:   degree.EmployeeID,
			LanguageCode: degree.LanguageCode,
			CourseTitle:  degree.CourseTitle,
			DateStart:    degree.DateStart.Time,
			DateEnd:      degree.DateEnd.Time,
			CreatedAt:    degree.CreatedAt.Time,
			UpdatedAt:    degree.UpdatedAt.Time,
		}
	}

	return employeeRefresherCourses, nil
}
