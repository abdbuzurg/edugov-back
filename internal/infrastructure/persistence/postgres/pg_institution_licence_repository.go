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

type pgInstitutionLicenceRepository struct {
	store *Store
}

func NewPgInstitutionLicenceRepository(store *Store) repositories.InstitutionLicenceRepository {
	return &pgInstitutionLicenceRepository{
		store: store,
	}
}

func (r *pgInstitutionLicenceRepository) Create(ctx context.Context, institutionLicence *domain.InstitutionLicence) (*domain.InstitutionLicence, error) {
	if institutionLicence.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence instituion_id is provided"))
	}

	if institutionLicence.LanguageCode != "en" && institutionLicence.LanguageCode != "ru" && institutionLicence.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence language_Code is provided"))
	}

	institutionLicenceResult, err := r.store.Queries.CreateInstitutionLicence(ctx, sqlc.CreateInstitutionLicenceParams{
		InstitutionID: institutionLicence.InstitutionID,
		LanguageCode:  institutionLicence.LanguageCode,
		LicenceTitle:  institutionLicence.LicenceTitle,
		LicenceType:   institutionLicence.LicenceType,
		LinkToFile:    institutionLicence.LinkToFile,
		GivenBy:       institutionLicence.GivenBy,
		DateStart: pgtype.Date{
			Time:  institutionLicence.DateStart,
			Valid: !institutionLicence.DateStart.IsZero(),
		},
		DateEnd: pgtype.Date{
			Time:  institutionLicence.DateEnd,
			Valid: !institutionLicence.DateEnd.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution licence: %w", err))
	}

	institutionLicence.ID = institutionLicenceResult.ID
	institutionLicence.CreatedAt = institutionLicenceResult.CreatedAt.Time
	institutionLicence.UpdatedAt = institutionLicenceResult.UpdatedAt.Time

	return institutionLicence, nil
}

func (r *pgInstitutionLicenceRepository) Update(ctx context.Context, institutionLicence *domain.InstitutionLicence) (*domain.InstitutionLicence, error) {
	if institutionLicence.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionLicence, err := q.GetInstitutionLicenceByID(ctx, institutionLicence.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution licence with given ID(%d) is not found", institutionLicence.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionLicence.ID, err))
		}

		updateInstitutionLicenceParams := sqlc.UpdateInstitutionLicenceParams{
			ID:           existingInstitutionLicence.ID,
			LicenceTitle: existingInstitutionLicence.LicenceTitle,
			LicenceType:  existingInstitutionLicence.LicenceType,
			LinkToFile:   existingInstitutionLicence.LinkToFile,
			GivenBy:      existingInstitutionLicence.GivenBy,
			DateStart:    existingInstitutionLicence.DateStart,
			DateEnd:      existingInstitutionLicence.DateEnd,
		}

		if institutionLicence.LicenceTitle != "" {
			updateInstitutionLicenceParams.LicenceTitle = institutionLicence.LicenceTitle
		}

		if institutionLicence.LicenceType != "" {
			updateInstitutionLicenceParams.LicenceType = institutionLicence.LicenceTitle
		}

		if institutionLicence.LinkToFile != "" {
			updateInstitutionLicenceParams.LinkToFile = institutionLicence.LinkToFile
		}

		if institutionLicence.GivenBy != "" {
			updateInstitutionLicenceParams.GivenBy = institutionLicence.GivenBy
		}

		if !institutionLicence.DateStart.IsZero() {
			updateInstitutionLicenceParams.DateStart = pgtype.Date{
				Time:  institutionLicence.DateStart,
				Valid: true,
			}
		}

		if !institutionLicence.DateEnd.IsZero() {
			updateInstitutionLicenceParams.DateEnd = pgtype.Date{
				Time:  institutionLicence.DateEnd,
				Valid: true,
			}
		}

		updateInstitutionLicenceResult, err := q.UpdateInstitutionLicence(ctx, updateInstitutionLicenceParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution licence: %w", err))
		}

		institutionLicence.CreatedAt = updateInstitutionLicenceResult.CreatedAt.Time
		institutionLicence.UpdatedAt = updateInstitutionLicenceResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution licence update: %w", err))
	}

	return institutionLicence, nil
}

func (r *pgInstitutionLicenceRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution licence ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionLicenceByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution licence with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution licence with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionLicence(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution licence: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution licence: %w", err))
	}

	return nil
}

func (r *pgInstitutionLicenceRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionLicence, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence ID is provided"))
	}

	institutionLicenceResult, err := r.store.Queries.GetInstitutionLicenceByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution licence with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionLicence{
		ID:            institutionLicenceResult.ID,
		InstitutionID: institutionLicenceResult.InstitutionID,
		LanguageCode:  institutionLicenceResult.LanguageCode,
		LicenceTitle:  institutionLicenceResult.LicenceTitle,
		LicenceType:   institutionLicenceResult.LicenceType,
		GivenBy:       institutionLicenceResult.GivenBy,
		LinkToFile:    institutionLicenceResult.LinkToFile,
		DateStart:     institutionLicenceResult.DateStart.Time,
		DateEnd:       institutionLicenceResult.DateEnd.Time,
		CreatedAt:     institutionLicenceResult.CreatedAt.Time,
		UpdatedAt:     institutionLicenceResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionLicenceRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionLicence, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution licence language_code is provided"))
	}

	institutionLicencesResult, err := r.store.Queries.GetInstitutionLicencesByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionLicencesByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution licence with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

  institutionLicences := make([]*domain.InstitutionLicence, len(institutionLicencesResult))
  for index, licence := range institutionLicencesResult {
    institutionLicences[index] = &domain.InstitutionLicence{
		ID:           licence.ID,
		LanguageCode: licence.LanguageCode,
		LicenceTitle: licence.LicenceTitle,
		LicenceType:  licence.LicenceTitle,
		GivenBy:      licence.GivenBy,
		LinkToFile:   licence.LinkToFile,
		DateStart:    licence.DateStart.Time,
		DateEnd:      licence.DateEnd.Time,
		CreatedAt:    licence.CreatedAt.Time,
		UpdatedAt:    licence.UpdatedAt.Time,
	} 
  }

	return institutionLicences, nil
}
