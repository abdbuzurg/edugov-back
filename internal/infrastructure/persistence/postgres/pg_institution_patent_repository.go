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
)

type pgInstitutionPatentRepository struct {
	store *Store
}

func NewPgInstitutionPatentRepository(store *Store) repositories.InstitutionPatentRepository {
	return &pgInstitutionPatentRepository{
		store: store,
	}
}

func (r *pgInstitutionPatentRepository) Create(ctx context.Context, institutionPatent *domain.InstitutionPatent) (*domain.InstitutionPatent, error) {
	if institutionPatent.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent instituion_id is provided"))
	}

	if institutionPatent.LanguageCode != "en" && institutionPatent.LanguageCode != "ru" && institutionPatent.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent language_Code is provided"))
	}

	institutionPatentResult, err := r.store.Queries.CreateInstitutionPatent(ctx, sqlc.CreateInstitutionPatentParams{
		InstitutionID:    institutionPatent.InstitutionID,
		LanguageCode:     institutionPatent.LanguageCode,
		PatentTitle:      institutionPatent.PatentTitle,
		Discipline:       institutionPatent.Discipline,
		Description:      institutionPatent.Description,
		ImplementedIn:    institutionPatent.ImplementedIn,
		LinkToPatentFile: institutionPatent.LinkToPatentFile,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution patent: %w", err))
	}

	institutionPatent.ID = institutionPatentResult.ID
	institutionPatent.CreatedAt = institutionPatentResult.CreatedAt.Time
	institutionPatent.UpdatedAt = institutionPatentResult.UpdatedAt.Time

	return institutionPatent, nil
}

func (r *pgInstitutionPatentRepository) Update(ctx context.Context, institutionPatent *domain.InstitutionPatent) (*domain.InstitutionPatent, error) {
	if institutionPatent.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionPatent, err := q.GetInstitutionPatentByID(ctx, institutionPatent.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution patent with given ID(%d) is not found", institutionPatent.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionPatent.ID, err))
		}

		updateInstitutionPatentParams := sqlc.UpdateInstitutionPatentParams{
			ID:               existingInstitutionPatent.ID,
			PatentTitle:      existingInstitutionPatent.PatentTitle,
			Discipline:       existingInstitutionPatent.Discipline,
			Description:      existingInstitutionPatent.Description,
			ImplementedIn:    existingInstitutionPatent.ImplementedIn,
			LinkToPatentFile: existingInstitutionPatent.LinkToPatentFile,
		}

		if institutionPatent.PatentTitle != "" {
			updateInstitutionPatentParams.PatentTitle = institutionPatent.PatentTitle
		}

		if institutionPatent.Discipline != "" {
			updateInstitutionPatentParams.Discipline = institutionPatent.Discipline
		}

		if institutionPatent.Description != "" {
			updateInstitutionPatentParams.Description = institutionPatent.Description
		}

		if institutionPatent.ImplementedIn != "" {
			updateInstitutionPatentParams.ImplementedIn = institutionPatent.ImplementedIn
		}

		if institutionPatent.LinkToPatentFile != "" {
			updateInstitutionPatentParams.LinkToPatentFile = institutionPatent.LinkToPatentFile
		}

		updateInstitutionPatentResult, err := q.UpdateInstitutionPatent(ctx, updateInstitutionPatentParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution patent: %w", err))
		}

		institutionPatent.CreatedAt = updateInstitutionPatentResult.CreatedAt.Time
		institutionPatent.UpdatedAt = updateInstitutionPatentResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution patent update: %w", err))
	}

	return institutionPatent, nil
}

func (r *pgInstitutionPatentRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution patent ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionPatentByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution patent with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution patent with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionPatent(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution patent: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution patent: %w", err))
	}

	return nil
}

func (r *pgInstitutionPatentRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionPatent, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent ID is provided"))
	}

	institutionPatentResult, err := r.store.Queries.GetInstitutionPatentByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution patent with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionPatent{
		ID:               institutionPatentResult.ID,
		InstitutionID:    institutionPatentResult.InstitutionID,
		LanguageCode:     institutionPatentResult.LanguageCode,
		PatentTitle:      institutionPatentResult.PatentTitle,
		Discipline:       institutionPatentResult.Discipline,
		Description:      institutionPatentResult.Discipline,
		ImplementedIn:    institutionPatentResult.ImplementedIn,
		LinkToPatentFile: institutionPatentResult.LinkToPatentFile,
		CreatedAt:        institutionPatentResult.CreatedAt.Time,
		UpdatedAt:        institutionPatentResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionPatentRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionPatent, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution patent language_code is provided"))
	}

	institutionPatentsResult, err := r.store.Queries.GetInstitutionPatentsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionPatentsByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution patent with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionPatents := make([]*domain.InstitutionPatent, len(institutionPatentsResult))
	for index, patent := range institutionPatentsResult {
		institutionPatents[index] = &domain.InstitutionPatent{
			ID:               patent.ID,
			InstitutionID:    patent.InstitutionID,
			LanguageCode:     patent.LanguageCode,
			PatentTitle:      patent.PatentTitle,
			Discipline:       patent.Discipline,
			Description:      patent.Discipline,
			ImplementedIn:    patent.ImplementedIn,
			LinkToPatentFile: patent.LinkToPatentFile,
			CreatedAt:        patent.CreatedAt.Time,
			UpdatedAt:        patent.UpdatedAt.Time,
		}
	}

	return institutionPatents, nil
}
