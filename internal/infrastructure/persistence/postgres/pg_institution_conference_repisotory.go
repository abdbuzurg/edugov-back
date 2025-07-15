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

type pgInstitutionConferenceRepository struct {
	store *Store
}

func NewPgInstitutionConferenceRepository(store *Store) repositories.InstitutionConferenceRepository {
	return &pgInstitutionConferenceRepository{
		store: store,
	}
}

func (r *pgInstitutionConferenceRepository) Create(ctx context.Context, institutionConference *domain.InstitutionConference) (*domain.InstitutionConference, error) {
	if institutionConference.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference institution_id is provided"))
	}

	if institutionConference.LanguageCode != "en" && institutionConference.LanguageCode != "ru" && institutionConference.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference langauge_code is provided"))
	}

	institutionConferenceResult, err := r.store.Queries.CreateInstitutionConference(ctx, sqlc.CreateInstitutionConferenceParams{
		InstitutionID:   institutionConference.InstitutionID,
		LanguageCode:    institutionConference.LanguageCode,
		ConferenceTitle: institutionConference.ConferenceTitle,
		Link:            institutionConference.Link,
		LinkToRinc: pgtype.Text{
			String: institutionConference.Link,
			Valid:  institutionConference.Link != "",
		},
		DateOfConference: pgtype.Date{
			Time:  institutionConference.DateOfConference,
			Valid: !institutionConference.DateOfConference.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution conference: %w", err))
	}

	institutionConference.ID = institutionConferenceResult.ID
	institutionConference.CreatedAt = institutionConferenceResult.CreatedAt.Time
	institutionConference.UpdatedAt = institutionConferenceResult.UpdatedAt.Time

	return institutionConference, nil
}

func (r *pgInstitutionConferenceRepository) Update(ctx context.Context, institutionConference *domain.InstitutionConference) (*domain.InstitutionConference, error) {
	if institutionConference.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionConference, err := q.GetInstitutionConferenceByID(ctx, institutionConference.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution conference with given ID(%d) is not found", institutionConference.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution conference with given ID(%d): %w", institutionConference.ID, err))
		}

		updateInstitutionConferenceParams := sqlc.UpdateInstitutionConferenceParams{
			ID:               existingInstitutionConference.ID,
			ConferenceTitle:  existingInstitutionConference.ConferenceTitle,
			Link:             existingInstitutionConference.Link,
			LinkToRinc:       existingInstitutionConference.LinkToRinc,
			DateOfConference: existingInstitutionConference.DateOfConference,
		}

		if institutionConference.ConferenceTitle != "" {
			existingInstitutionConference.ConferenceTitle = institutionConference.ConferenceTitle
		}

		if institutionConference.Link != "" {
			existingInstitutionConference.Link = institutionConference.Link
		}

		if institutionConference.LinkToRINC != "" {
			existingInstitutionConference.LinkToRinc = pgtype.Text{
				String: institutionConference.LinkToRINC,
				Valid:  true,
			}
		}

		if !institutionConference.DateOfConference.IsZero() {
			existingInstitutionConference.DateOfConference = pgtype.Date{
				Time:  institutionConference.DateOfConference,
				Valid: true,
			}
		}

		updateInstitutionConferenceResult, err := q.UpdateInstitutionConference(ctx, updateInstitutionConferenceParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution conference: %w", err))
		}

		institutionConference.CreatedAt = updateInstitutionConferenceResult.CreatedAt.Time
		institutionConference.UpdatedAt = updateInstitutionConferenceResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed during institution conference update: %w", err))
	}

	return institutionConference, nil
}

func (r *pgInstitutionConferenceRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution conference ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution conference with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution conference with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionConference(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution details: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transction failed during the deleting of institution conference: %w", err))
	}

	return nil
}

func (r *pgInstitutionConferenceRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionConference, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference ID is provided"))
	}

	institutionConferenceResult, err := r.store.Queries.GetInstitutionConferenceByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution conference with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionConference{
		ID:               institutionConferenceResult.ID,
		InstitutionID:    institutionConferenceResult.InstitutionID,
		LanguageCode:     institutionConferenceResult.LanguageCode,
		ConferenceTitle:  institutionConferenceResult.ConferenceTitle,
		Link:             institutionConferenceResult.Link,
		LinkToRINC:       institutionConferenceResult.LinkToRinc.String,
		DateOfConference: institutionConferenceResult.DateOfConference.Time,
		CreatedAt:        institutionConferenceResult.CreatedAt.Time,
		UpdatedAt:        institutionConferenceResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionConferenceRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionConference, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference instution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution conference langauge_code is provided"))
	}

	institutionConferencesResult, err := r.store.Queries.GetInstitutionConferencesByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionConferencesByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution conference with given institutionID(%d) and languageCode(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

  institutionConferences := make([]*domain.InstitutionConference, len(institutionConferencesResult))
  for index, conference := range institutionConferencesResult {
    institutionConferences[index] = &domain.InstitutionConference{
		ID:               conference.ID,
		InstitutionID:    conference.InstitutionID,
		LanguageCode:     conference.LanguageCode,
		ConferenceTitle:  conference.ConferenceTitle,
		Link:             conference.Link,
		LinkToRINC:       conference.LinkToRinc.String,
		DateOfConference: conference.DateOfConference.Time,
		CreatedAt:        conference.CreatedAt.Time,
		UpdatedAt:        conference.UpdatedAt.Time,
	} 
  }

	return institutionConferences, nil
}
