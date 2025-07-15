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

type pgInstitutionMainResearchDirectionRepository struct {
	store *Store
}

func NewPgInstitutionMainResearchDirectionRepository(store *Store) repositories.InstitutionMainResearchDirectionRepository {
	return &pgInstitutionMainResearchDirectionRepository{
		store: store,
	}
}

func (r *pgInstitutionMainResearchDirectionRepository) Create(ctx context.Context, institutionMRD *domain.InstitutionMainResearchDirection) (*domain.InstitutionMainResearchDirection, error) {
	if institutionMRD.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction instituion_id is provided"))
	}

	if institutionMRD.LanguageCode != "en" && institutionMRD.LanguageCode != "ru" && institutionMRD.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction language_Code is provided"))
	}

	institutionMRDResult, err := r.store.Queries.CreateInstitutionMainResearchDirection(ctx, sqlc.CreateInstitutionMainResearchDirectionParams{
		InstitutionID:          institutionMRD.InstitutionID,
		LanguageCode:           institutionMRD.LanguageCode,
		ResearchDirectionTitle: institutionMRD.ResearchDirectionTitle,
		Discipline:             institutionMRD.Discipline,
		AreaOfResearch: pgtype.Text{
			String: institutionMRD.AreaOfResearch,
			Valid:  institutionMRD.AreaOfResearch != "",
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create main research direction: %w", err))
	}

	institutionMRD.ID = institutionMRDResult.ID
	institutionMRD.CreatedAt = institutionMRDResult.CreatedAt.Time
	institutionMRD.UpdatedAt = institutionMRDResult.UpdatedAt.Time

	return institutionMRD, nil
}

func (r *pgInstitutionMainResearchDirectionRepository) Update(ctx context.Context, institutionMRD *domain.InstitutionMainResearchDirection) (*domain.InstitutionMainResearchDirection, error) {
	if institutionMRD.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionMRD, err := q.GetInstitutionMainResearchDirectionByID(ctx, institutionMRD.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {

				return custom_errors.BadRequest(fmt.Errorf("institution main research direction with given ID(%d) is not found", institutionMRD.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution main research direction with given ID(%d): %w", institutionMRD.ID, err))
		}

		updateInstitutionMRDParams := sqlc.UpdateInstitutionMainResearchDirectionParams{
			ID:                     existingInstitutionMRD.ID,
			ResearchDirectionTitle: existingInstitutionMRD.ResearchDirectionTitle,
			Discipline:             existingInstitutionMRD.Discipline,
			AreaOfResearch:         existingInstitutionMRD.AreaOfResearch,
		}

		if institutionMRD.ResearchDirectionTitle != "" {
			updateInstitutionMRDParams.ResearchDirectionTitle = institutionMRD.ResearchDirectionTitle
		}

		if institutionMRD.Discipline != "" {
			updateInstitutionMRDParams.Discipline = institutionMRD.Discipline
		}

		if institutionMRD.AreaOfResearch != "" {
			updateInstitutionMRDParams.AreaOfResearch = pgtype.Text{
				String: institutionMRD.AreaOfResearch,
				Valid:  true,
			}
		}

		updateInstitutionMRDResult, err := q.UpdateInstitutionMainResearchDirection(ctx, updateInstitutionMRDParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution main research direction: %w", err))
		}

		institutionMRD.CreatedAt = updateInstitutionMRDResult.CreatedAt.Time
		institutionMRD.UpdatedAt = updateInstitutionMRDResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution main research direction update: %w", err))
	}

	return institutionMRD, nil
}

func (r *pgInstitutionMainResearchDirectionRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionMainResearchDirectionByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution main research direction with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution main research direction with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionMainResearchDirection(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution main research direction: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution main research direction: %w", err))
	}

	return nil
}

func (r *pgInstitutionMainResearchDirectionRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionMainResearchDirection, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction ID is provided"))
	}

	institutionMRDResult, err := r.store.Queries.GetInstitutionMainResearchDirectionByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution main research direction with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionMainResearchDirection{
		ID:                     institutionMRDResult.ID,
		InstitutionID:          institutionMRDResult.InstitutionID,
		LanguageCode:           institutionMRDResult.LanguageCode,
		ResearchDirectionTitle: institutionMRDResult.ResearchDirectionTitle,
		Discipline:             institutionMRDResult.Discipline,
		AreaOfResearch:         institutionMRDResult.AreaOfResearch.String,
		CreatedAt:              institutionMRDResult.CreatedAt.Time,
		UpdatedAt:              institutionMRDResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionMainResearchDirectionRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionMainResearchDirection, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution main research direction language_code is provided"))
	}

	institutionMRDsResult, err := r.store.Queries.GetInstitutionMainResearchDirectionsByInstitutionIDAndLanguage(ctx, sqlc.GetInstitutionMainResearchDirectionsByInstitutionIDAndLanguageParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution main research direction with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionMRDs := make([]*domain.InstitutionMainResearchDirection, len(institutionMRDsResult))
	for index, mrd := range institutionMRDsResult {
		institutionMRDs[index] = &domain.InstitutionMainResearchDirection{
			ID:                     mrd.ID,
			InstitutionID:          mrd.InstitutionID,
			LanguageCode:           mrd.LanguageCode,
			ResearchDirectionTitle: mrd.ResearchDirectionTitle,
			Discipline:             mrd.Discipline,
			AreaOfResearch:         mrd.AreaOfResearch.String,
			CreatedAt:              mrd.CreatedAt.Time,
			UpdatedAt:              mrd.UpdatedAt.Time,
		}
	}

	return institutionMRDs, nil
}
