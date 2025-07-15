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

type pgEmployeeParticipationInEventRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeParticipationInEventRepository(store *Store) repositories.EmployeeParticipationInEventRepository {
	return &pgEmployeeParticipationInEventRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeParticipationInEventRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeParticipationInEventRepository {
	return &pgEmployeeParticipationInEventRepository{
		queries: q,
	}
}

func (r *pgEmployeeParticipationInEventRepository) Create(ctx context.Context, employeePIE *domain.EmployeeParticipationInEvent) (*domain.EmployeeParticipationInEvent, error) {
	employeePIEResult, err := r.queries.CreateEmployeeParticipationInEvent(ctx, sqlc.CreateEmployeeParticipationInEventParams{
		EmployeeID:   employeePIE.EmployeeID,
		LanguageCode: employeePIE.LanguageCode,
		EventTitle:   employeePIE.EventTitle,
		EventDate: pgtype.Date{
			Time:  employeePIE.EventDate,
			Valid: !employeePIE.EventDate.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee participation in event: %w", err))
	}

	employeePIE.ID = employeePIEResult.ID
	employeePIE.CreatedAt = employeePIEResult.CreatedAt.Time
	employeePIE.UpdatedAt = employeePIEResult.UpdatedAt.Time

	return employeePIE, nil
}

func (r *pgEmployeeParticipationInEventRepository) Update(ctx context.Context, employeePIE *domain.EmployeeParticipationInEvent) (*domain.EmployeeParticipationInEvent, error) {
	updateEmployeeParticipationInEventResult, err := r.queries.UpdateEmployeeParticipationInEvent(ctx, sqlc.UpdateEmployeeParticipationInEventParams{
		ID:         employeePIE.ID,
		EventTitle: employeePIE.EventTitle,
		EventDate: pgtype.Date{
			Time:  employeePIE.EventDate,
			Valid: !employeePIE.EventDate.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee participation in event: %w", err))
	}

	employeePIE.CreatedAt = updateEmployeeParticipationInEventResult.CreatedAt.Time
	employeePIE.UpdatedAt = updateEmployeeParticipationInEventResult.UpdatedAt.Time

	return employeePIE, nil
}

func (r *pgEmployeeParticipationInEventRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeParticipationInEvent(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee participation in event: %w", err))
	}

	return nil
}

func (r *pgEmployeeParticipationInEventRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeParticipationInEvent, error) {
	employeePIE, err := r.queries.GetEmployeeParticipationInEventByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee participation in event with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeeParticipationInEvent{
		ID:                             employeePIE.ID,
		EmployeeID:                     employeePIE.EmployeeID,
		LanguageCode:                   employeePIE.LanguageCode,
		EventTitle:                     employeePIE.EventTitle,
		EventDate:                      employeePIE.EventDate.Time,
		CreatedAt:                      employeePIE.CreatedAt.Time,
		UpdatedAt:                      employeePIE.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeParticipationInEventRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeParticipationInEvent, error) {
	employeePIEsResult, err := r.queries.GetEmployeeParticipationInEventsByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeParticipationInEventsByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee participation in events with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	} 

	employeePIEs := make([]*domain.EmployeeParticipationInEvent, len(employeePIEsResult))
	for index, employeePIE := range employeePIEsResult {
		employeePIEs[index] = &domain.EmployeeParticipationInEvent{
			ID:                             employeePIE.ID,
			EmployeeID:                     employeePIE.EmployeeID,
			LanguageCode:                   employeePIE.LanguageCode,
			EventTitle:                     employeePIE.EventTitle,
			EventDate:                      employeePIE.EventDate.Time,
			CreatedAt:                      employeePIE.CreatedAt.Time,
			UpdatedAt:                      employeePIE.UpdatedAt.Time,
		}
	}

	return employeePIEs, nil
}
