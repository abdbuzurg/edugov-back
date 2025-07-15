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

type pgInstitutionAchievementRepository struct {
	store *Store
}

func NewPgInstitutionAchievementRepository(store *Store) repositories.InstitutionAchievementRepository {
	return &pgInstitutionAchievementRepository{
		store: store,
	}
}

func (r *pgInstitutionAchievementRepository) Create(ctx context.Context, institutionAchievement *domain.InstitutionAchievement) (*domain.InstitutionAchievement, error) {
	if institutionAchievement.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution achievement institutio_id(%d) is provided", institutionAchievement.InstitutionID))
	}

	institutionAchievementResult, err := r.store.Queries.CreateInstitutionAchievement(ctx, sqlc.CreateInstitutionAchievementParams{
		InstitutionID:    institutionAchievement.InstitutionID,
		LanguageCode:     institutionAchievement.LanguageCode,
		AchievementTitle: institutionAchievement.AchievementTitle,
		AchievementType:             institutionAchievement.AchievementType,
		DateRecieved: pgtype.Date{
			Time:  institutionAchievement.DateReceived,
			Valid: !institutionAchievement.DateReceived.IsZero(),
		},
		GivenBy:    institutionAchievement.GivenBy,
		LinkToFile: institutionAchievement.LinkToFile,
		Description: pgtype.Text{
			String: institutionAchievement.Description,
			Valid:  institutionAchievement.Description != "",
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution achievement: %w", err))
	}

	institutionAchievement.ID = institutionAchievementResult.ID
	institutionAchievement.CreatedAt = institutionAchievementResult.CreatedAt.Time
	institutionAchievement.UpdatedAt = institutionAchievementResult.UpdatedAt.Time

	return institutionAchievement, nil
}

func (r *pgInstitutionAchievementRepository) Update(ctx context.Context, institutionAchievement *domain.InstitutionAchievement) (*domain.InstitutionAchievement, error) {
	if institutionAchievement.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution achievement ID(%d) is provided", institutionAchievement.ID))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionAchievement, err := q.GetInstitutionAchievementByID(ctx, institutionAchievement.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution achievment with given ID(%d) is not found", institutionAchievement.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution achievement with given ID(%d): %w", institutionAchievement.ID, err))
		}

		updateInstitutionAchievementParams := sqlc.UpdateInstitutionAchievementParams{
			ID:               existingInstitutionAchievement.ID,
			AchievementTitle: existingInstitutionAchievement.AchievementTitle,
			AchievementType:             existingInstitutionAchievement.AchievementType,
			DateRecieved:     existingInstitutionAchievement.DateRecieved,
			GivenBy:          existingInstitutionAchievement.GivenBy,
			LinkToFile:       existingInstitutionAchievement.LinkToFile,
			Description:      existingInstitutionAchievement.Description,
		}

		if institutionAchievement.AchievementTitle != "" {
			updateInstitutionAchievementParams.AchievementTitle = institutionAchievement.AchievementTitle
		}

		if institutionAchievement.AchievementType != "" {
			updateInstitutionAchievementParams.AchievementType = institutionAchievement.AchievementType
		}

		if !institutionAchievement.DateReceived.IsZero() {
			updateInstitutionAchievementParams.DateRecieved = pgtype.Date{
				Time:  institutionAchievement.DateReceived,
				Valid: true,
			}
		}

		if institutionAchievement.GivenBy != "" {
			updateInstitutionAchievementParams.GivenBy = institutionAchievement.GivenBy
		}

		if institutionAchievement.LinkToFile != "" {
			updateInstitutionAchievementParams.LinkToFile = institutionAchievement.LinkToFile
		}

		if institutionAchievement.Description != "" {
			updateInstitutionAchievementParams.Description = pgtype.Text{
				String: institutionAchievement.Description,
				Valid:  true,
			}
		}

		updateInstitutionAchievementResult, err := q.UpdateInstitutionAchievement(ctx, updateInstitutionAchievementParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution achievement: %w", err))
		}

		institutionAchievement.CreatedAt = updateInstitutionAchievementResult.CreatedAt.Time
		institutionAchievement.UpdatedAt = updateInstitutionAchievementResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to perform institution achievement update transaction: %w", err))
	}

	return institutionAchievement, nil
}

func (r *pgInstitutionAchievementRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution achievement ID(%d) is provided", id))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionAchievementByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution achievement with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution achievement with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionAchivement(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution achievement with ID(%d): %w", id, err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to perfom institution achievement deletion transaction: %w", err))
	}

	return nil
}

func (r *pgInstitutionAchievementRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionAchievement, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution achievement ID(%d) is provided", id))
	}

	institutionAchiementResult, err := r.store.Queries.GetInstitutionAchievementByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrived institution achievement with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionAchievement{
		ID:               institutionAchiementResult.ID,
		InstitutionID:    institutionAchiementResult.InstitutionID,
		LanguageCode:     institutionAchiementResult.LanguageCode,
		AchievementTitle: institutionAchiementResult.AchievementTitle,
		AchievementType:             institutionAchiementResult.AchievementType,
		DateReceived:     institutionAchiementResult.DateRecieved.Time,
		GivenBy:          institutionAchiementResult.GivenBy,
		Description:      institutionAchiementResult.Description.String,
		CreatedAt:        institutionAchiementResult.CreatedAt.Time,
		UpdatedAt:        institutionAchiementResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionAchievementRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionAchievement, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution achievement institution_id(%d) is provided", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution achievement language_code(%s) is provided", langCode))
	}

	institutionAchievementsResult, err := r.store.Queries.GetInstitutionAchievementsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionAchievementsByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrieve institution achievement with given institutio_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionAchievements := make([]*domain.InstitutionAchievement, len(institutionAchievementsResult))
	for index, achievement := range institutionAchievementsResult {
		institutionAchievements[index] = &domain.InstitutionAchievement{
			ID:               achievement.ID,
			LanguageCode:     achievement.LanguageCode,
			AchievementTitle: achievement.AchievementTitle,
			AchievementType:  achievement.AchievementType,
			DateReceived:     achievement.DateRecieved.Time,
			GivenBy:          achievement.GivenBy,
			LinkToFile:       achievement.LinkToFile,
			Description:      achievement.Description.String,
			CreatedAt:        achievement.CreatedAt.Time,
			UpdatedAt:        achievement.UpdatedAt.Time,
		}
	}

	return institutionAchievements, nil
}
