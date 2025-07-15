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

type pgInstitutionResearchSupportInfrastructureRepository struct {
	store *Store
}

func NewPgInstitutionResearchSupportInfrastructureRepository(store *Store) repositories.InstitutionResearchSupportInfrastructureRepository {
	return &pgInstitutionResearchSupportInfrastructureRepository{
		store: store,
	}
}

func (r *pgInstitutionResearchSupportInfrastructureRepository) Create(ctx context.Context, institutionRSI *domain.InstitutionResearchSupportInfrastructure) (*domain.InstitutionResearchSupportInfrastructure, error) {
	if institutionRSI.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure instituion_id is provided"))
	}

	if institutionRSI.LanguageCode != "en" && institutionRSI.LanguageCode != "ru" && institutionRSI.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure language_Code is provided"))
	}

	institutionRSIResult, err := r.store.Queries.CreateInstitutionResearchSupportInfrastructure(ctx, sqlc.CreateInstitutionResearchSupportInfrastructureParams{
		InstitutionID:                      institutionRSI.InstitutionID,
		LanguageCode:                       institutionRSI.LanguageCode,
		ResearchSupportInfrastructureTitle: institutionRSI.ResearchSupportInfrastructureTitle,
		ResearchSupportInfrastructureType:  institutionRSI.ResearchSupportInfrastructureType,
		TinOfLegalEntity:                   institutionRSI.TINOfLegalEntity,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution research support infrastructure: %w", err))
	}

	institutionRSI.ID = institutionRSIResult.ID
	institutionRSI.CreatedAt = institutionRSIResult.CreatedAt.Time
	institutionRSI.UpdatedAt = institutionRSIResult.UpdatedAt.Time

	return institutionRSI, nil
}

func (r *pgInstitutionResearchSupportInfrastructureRepository) Update(ctx context.Context, institutionRSI *domain.InstitutionResearchSupportInfrastructure) (*domain.InstitutionResearchSupportInfrastructure, error) {
	if institutionRSI.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionResearchSupportInfrastructure, err := q.GetInstitutionResearchSupportInfrastructureByID(ctx, institutionRSI.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution research support infrastructure with given ID(%d) is not found", institutionRSI.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionRSI.ID, err))
		}

		updateInstitutionResearchSupportInfrastructureParams := sqlc.UpdateInstitutionResearchSupportInfrastructureParams{
			ID:                                 existingInstitutionResearchSupportInfrastructure.ID,
			ResearchSupportInfrastructureTitle: existingInstitutionResearchSupportInfrastructure.ResearchSupportInfrastructureTitle,
			ResearchSupportInfrastructureType:  existingInstitutionResearchSupportInfrastructure.ResearchSupportInfrastructureType,
			TinOfLegalEntity:                   existingInstitutionResearchSupportInfrastructure.TinOfLegalEntity,
		}

		if institutionRSI.ResearchSupportInfrastructureTitle != "" {
			updateInstitutionResearchSupportInfrastructureParams.ResearchSupportInfrastructureTitle = institutionRSI.ResearchSupportInfrastructureTitle
		}

		if institutionRSI.ResearchSupportInfrastructureType != "" {
			updateInstitutionResearchSupportInfrastructureParams.ResearchSupportInfrastructureType = institutionRSI.ResearchSupportInfrastructureType
		}

		if institutionRSI.TINOfLegalEntity != "" {
			updateInstitutionResearchSupportInfrastructureParams.TinOfLegalEntity = institutionRSI.TINOfLegalEntity
		}

		updateInstitutionResearchSupportInfrastructureResult, err := q.UpdateInstitutionResearchSupportInfrastructure(ctx, updateInstitutionResearchSupportInfrastructureParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution research support infrastructure: %w", err))
		}

		institutionRSI.CreatedAt = updateInstitutionResearchSupportInfrastructureResult.CreatedAt.Time
		institutionRSI.UpdatedAt = updateInstitutionResearchSupportInfrastructureResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution research support infrastructure update: %w", err))
	}

	return institutionRSI, nil
}

func (r *pgInstitutionResearchSupportInfrastructureRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionResearchSupportInfrastructureByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution research support infrastructure with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution research support infrastructure with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionResearchSupportInfrastructure(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution research support infrastructure: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution research support infrastructure: %w", err))
	}

	return nil
}

func (r *pgInstitutionResearchSupportInfrastructureRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionResearchSupportInfrastructure, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure ID is provided"))
	}

	institutionRSIResult, err := r.store.Queries.GetInstitutionResearchSupportInfrastructureByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution research support infrastructure with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionResearchSupportInfrastructure{
		ID:                                 institutionRSIResult.ID,
		InstitutionID:                      institutionRSIResult.InstitutionID,
		LanguageCode:                       institutionRSIResult.LanguageCode,
		ResearchSupportInfrastructureTitle: institutionRSIResult.ResearchSupportInfrastructureTitle,
		ResearchSupportInfrastructureType:  institutionRSIResult.ResearchSupportInfrastructureType,
		TINOfLegalEntity:                   institutionRSIResult.TinOfLegalEntity,
		CreatedAt:                          institutionRSIResult.CreatedAt.Time,
		UpdatedAt:                          institutionRSIResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionResearchSupportInfrastructureRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionResearchSupportInfrastructure, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution research support infrastructure language_code is provided"))
	}

	institutionRSIsResult, err := r.store.Queries.GetInstitutionResearchSupportInfrastructuresByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionResearchSupportInfrastructuresByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution research support infrastructure with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionRSIs := make([]*domain.InstitutionResearchSupportInfrastructure, len(institutionRSIsResult))
	for index, rsi := range institutionRSIsResult {
		institutionRSIs[index] = &domain.InstitutionResearchSupportInfrastructure{
			ID:                                 rsi.ID,
			InstitutionID:                      rsi.InstitutionID,
			LanguageCode:                       rsi.LanguageCode,
			ResearchSupportInfrastructureTitle: rsi.ResearchSupportInfrastructureTitle,
			ResearchSupportInfrastructureType:  rsi.ResearchSupportInfrastructureType,
			TINOfLegalEntity:                   rsi.TinOfLegalEntity,
			CreatedAt:                          rsi.CreatedAt.Time,
			UpdatedAt:                          rsi.UpdatedAt.Time,
		}
	}

	return institutionRSIs, nil
}
