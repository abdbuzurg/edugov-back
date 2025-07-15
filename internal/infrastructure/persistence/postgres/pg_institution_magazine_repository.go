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

type pgInstitutionMagazineRepository struct {
	store *Store
}

func NewPgInstitutionMagazineRepository(store *Store) repositories.InstitutionMagazineRepository {
	return &pgInstitutionMagazineRepository{
		store: store,
	}
}

func (r *pgInstitutionMagazineRepository) Create(ctx context.Context, institutionMagazine *domain.InstitutionMagazine) (*domain.InstitutionMagazine, error) {
	if institutionMagazine.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution magazine institution_id is provided"))
	}

	if institutionMagazine.LanguageCode != "en" && institutionMagazine.LanguageCode != "ru" && institutionMagazine.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution magazine langauge_code is provided"))
	}

	institutionMagazineResult, err := r.store.Queries.CreateInstitutionMagazine(ctx, sqlc.CreateInstitutionMagazineParams{
		InstitutionID: institutionMagazine.InstitutionID,
		LanguageCode:  institutionMagazine.LanguageCode,
		MagazineName:  institutionMagazine.MagazineName,
		Link:          institutionMagazine.Link,
		LinkToRinc: pgtype.Text{
			String: institutionMagazine.LinkToRINC,
			Valid:  institutionMagazine.LinkToRINC != "",
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institutio magazine: %w", err))
	}

	institutionMagazine.ID = institutionMagazineResult.ID
	institutionMagazine.CreatedAt = institutionMagazineResult.CreatedAt.Time
	institutionMagazine.UpdatedAt = institutionMagazineResult.UpdatedAt.Time

	return institutionMagazine, nil
}

func (r *pgInstitutionMagazineRepository) Update(ctx context.Context, institutionMagazine *domain.InstitutionMagazine) (*domain.InstitutionMagazine, error) {
	if institutionMagazine.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution magazine institution_id is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionMagazine, err := q.GetInstitutionMagazineByID(ctx, institutionMagazine.ID)
		if err != nil {
			return custom_errors.BadRequest(fmt.Errorf("institution magazine with given ID(%d) is not found", institutionMagazine.ID))
		}

		updateInstitutionMagazineParams := sqlc.UpdateInstitutionMagazineParams{
			MagazineName: existingInstitutionMagazine.MagazineName,
			Link:         existingInstitutionMagazine.Link,
			LinkToRinc:   existingInstitutionMagazine.LinkToRinc,
		}

		if institutionMagazine.MagazineName != "" {
			updateInstitutionMagazineParams.MagazineName = institutionMagazine.MagazineName
		}

		if institutionMagazine.Link != "" {
			updateInstitutionMagazineParams.Link = institutionMagazine.Link
		}

		if institutionMagazine.LinkToRINC != "" {
			updateInstitutionMagazineParams.LinkToRinc = pgtype.Text{
				String: institutionMagazine.LinkToRINC,
				Valid:  true,
			}
		}

		updateInstitutionMagazineResult, err := q.UpdateInstitutionMagazine(ctx, updateInstitutionMagazineParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution magazine"))
		}

		institutionMagazine.CreatedAt = updateInstitutionMagazineResult.CreatedAt.Time
		institutionMagazine.UpdatedAt = updateInstitutionMagazineResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while updating instituion magazine"))
	}

	return institutionMagazine, nil
}

func (r *pgInstitutionMagazineRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution magazine ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionMagazineByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution magazine with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution magazine with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionMagazine(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delte institution magazine: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution magazine: %w", err))
	}

	return nil
}

func (r *pgInstitutionMagazineRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionMagazine, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution magazine id is provided"))
	}

	institutionMagazineResult, err := r.store.Queries.GetInstitutionMagazineByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion magazine with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionMagazine{
		ID:            institutionMagazineResult.ID,
		InstitutionID: institutionMagazineResult.InstitutionID,
		LanguageCode:  institutionMagazineResult.LanguageCode,
		MagazineName:  institutionMagazineResult.MagazineName,
		Link:          institutionMagazineResult.Link,
		LinkToRINC:    institutionMagazineResult.LinkToRinc.String,
		CreatedAt:     institutionMagazineResult.CreatedAt.Time,
		UpdatedAt:     institutionMagazineResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionMagazineRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionMagazine, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid instition magazine institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution magazine language_code is provided"))
	}

	institutionMagazinesResult, err := r.store.Queries.GetInstitutionMagazinesByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionMagazinesByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution magazine with given instition_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionMagazines := make([]*domain.InstitutionMagazine, len(institutionMagazinesResult))
	for index, magazine := range institutionMagazinesResult {
		institutionMagazines[index] = &domain.InstitutionMagazine{
			ID:            magazine.ID,
			InstitutionID: magazine.InstitutionID,
			LanguageCode:  magazine.LanguageCode,
			MagazineName:  magazine.MagazineName,
			Link:          magazine.Link,
			LinkToRINC:    magazine.LinkToRinc.String,
			CreatedAt:     magazine.CreatedAt.Time,
			UpdatedAt:     magazine.UpdatedAt.Time,
		}
	}

	return institutionMagazines, nil
}
