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

type pgInstitutionRankingRepository struct {
	store *Store
}

func NewPgInstitutionRankingRepository(store *Store) repositories.InstitutionRankingRepository {
	return &pgInstitutionRankingRepository{
		store: store,
	}
}

func (r *pgInstitutionRankingRepository) Create(ctx context.Context, institutionRanking *domain.InstitutionRanking) (*domain.InstitutionRanking, error) {
	if institutionRanking.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking instituion_id is provided"))
	}

	if institutionRanking.LanguageCode != "en" && institutionRanking.LanguageCode != "ru" && institutionRanking.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking language_Code is provided"))
	}

	institutionRankingResult, err := r.store.Queries.CreateInstitutionRanking(ctx, sqlc.CreateInstitutionRankingParams{
		InstitutionID: institutionRanking.InstitutionID,
		LanguageCode:  institutionRanking.LanguageCode,
		RankingTitle:  institutionRanking.RankingTitle,
		RankingType:   institutionRanking.RankingType,
		DateRecieved: pgtype.Date{
			Time:  institutionRanking.DateReceived,
			Valid: !institutionRanking.DateReceived.IsZero(),
		},
		RankingAgency:     institutionRanking.RankingAgency,
		LinkToRankingFile: institutionRanking.LinkToRankingFile,
		Description: pgtype.Text{
			String: institutionRanking.Description,
			Valid:  institutionRanking.Description != "",
		},
		LinkToRankingAgency: institutionRanking.LinkToRankingAgency,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution ranking: %w", err))
	}

	institutionRanking.ID = institutionRankingResult.ID
	institutionRanking.CreatedAt = institutionRankingResult.CreatedAt.Time
	institutionRanking.UpdatedAt = institutionRankingResult.UpdatedAt.Time

	return institutionRanking, nil
}

func (r *pgInstitutionRankingRepository) Update(ctx context.Context, institutionRanking *domain.InstitutionRanking) (*domain.InstitutionRanking, error) {
	if institutionRanking.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionRanking, err := q.GetInstitutionRankingByID(ctx, institutionRanking.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution ranking with given ID(%d) is not found", institutionRanking.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionRanking.ID, err))
		}

		updateInstitutionRankingParams := sqlc.UpdateInstitutionRankingParams{
			ID:                  existingInstitutionRanking.ID,
			RankingTitle:        existingInstitutionRanking.RankingTitle,
			RankingAgency:       existingInstitutionRanking.RankingAgency,
			DateRecieved:        existingInstitutionRanking.DateRecieved,
			LinkToRankingFile:   existingInstitutionRanking.LinkToRankingFile,
			Description:         existingInstitutionRanking.Description,
			LinkToRankingAgency: existingInstitutionRanking.LinkToRankingAgency,
		}

		if institutionRanking.RankingTitle != "" {
			updateInstitutionRankingParams.RankingTitle = institutionRanking.RankingTitle
		}

		if institutionRanking.RankingAgency != "" {
			updateInstitutionRankingParams.RankingAgency = institutionRanking.RankingAgency
		}

		if !institutionRanking.DateReceived.IsZero() {
			updateInstitutionRankingParams.DateRecieved = pgtype.Date{
				Time:  institutionRanking.DateReceived,
				Valid: true,
			}
		}

		if institutionRanking.RankingAgency != "" {
			updateInstitutionRankingParams.RankingAgency = institutionRanking.RankingAgency
		}

		if institutionRanking.LinkToRankingFile != "" {
			updateInstitutionRankingParams.LinkToRankingFile = institutionRanking.LinkToRankingFile
		}

		if institutionRanking.Description != "" {
			updateInstitutionRankingParams.Description = pgtype.Text{
				String: institutionRanking.Description,
				Valid:  true,
			}
		}

		if institutionRanking.LinkToRankingAgency != "" {
			updateInstitutionRankingParams.LinkToRankingAgency = institutionRanking.LinkToRankingAgency
		}

		updateInstitutionRankingResult, err := q.UpdateInstitutionRanking(ctx, updateInstitutionRankingParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution ranking: %w", err))
		}

		institutionRanking.CreatedAt = updateInstitutionRankingResult.CreatedAt.Time
		institutionRanking.UpdatedAt = updateInstitutionRankingResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution ranking update: %w", err))
	}

	return institutionRanking, nil
}

func (r *pgInstitutionRankingRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution ranking ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionRankingByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution ranking with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution ranking with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionRanking(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution ranking: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution ranking: %w", err))
	}

	return nil
}

func (r *pgInstitutionRankingRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionRanking, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking ID is provided"))
	}

	institutionRankingResult, err := r.store.Queries.GetInstitutionRankingByID(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution ranking with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionRanking{
		ID:                  institutionRankingResult.ID,
		InstitutionID:       institutionRankingResult.InstitutionID,
		LanguageCode:        institutionRankingResult.LanguageCode,
		RankingTitle:        institutionRankingResult.RankingTitle,
		RankingType:         institutionRankingResult.RankingType,
		DateReceived:        institutionRankingResult.DateRecieved.Time,
		RankingAgency:       institutionRankingResult.RankingAgency,
		Description:         institutionRankingResult.Description.String,
		LinkToRankingFile:   institutionRankingResult.LinkToRankingFile,
		LinkToRankingAgency: institutionRankingResult.LinkToRankingAgency,
		CreatedAt:           institutionRankingResult.CreatedAt.Time,
		UpdatedAt:           institutionRankingResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionRankingRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionRanking, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ranking language_code is provided"))
	}

	institutionRankingsResult, err := r.store.Queries.GetInstitutionRankingsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionRankingsByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution ranking with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionRankings := make([]*domain.InstitutionRanking, len(institutionRankingsResult))
	for index, ranking := range institutionRankingsResult {
		institutionRankings[index] = &domain.InstitutionRanking{
			ID:                  ranking.ID,
			InstitutionID:       ranking.InstitutionID,
			LanguageCode:        ranking.LanguageCode,
			RankingTitle:        ranking.RankingTitle,
			RankingType:         ranking.RankingType,
			DateReceived:        ranking.DateRecieved.Time,
			RankingAgency:       ranking.RankingAgency,
			Description:         ranking.Description.String,
			LinkToRankingFile:   ranking.LinkToRankingFile,
			LinkToRankingAgency: ranking.LinkToRankingAgency,
			CreatedAt:           ranking.CreatedAt.Time,
			UpdatedAt:           ranking.UpdatedAt.Time,
		}
	}

	return institutionRankings, nil
}
