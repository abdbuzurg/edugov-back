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

type pgInstitutionAccreditationRepository struct {
	store *Store
}

func NewPgInstitutionAccreditatitonRepository(store *Store) repositories.InstitutionAccreditationRepository {
	return &pgInstitutionAccreditationRepository{
		store: store,
	}
}

func (r *pgInstitutionAccreditationRepository) Create(ctx context.Context, institutionAccreditation *domain.InstitutionAccreditation) (*domain.InstitutionAccreditation, error) {
	if institutionAccreditation.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("institution accreditation requires to have institution_id"))
	}

	institutionAccreditationResult, err := r.store.Queries.CreateInstitutionAccreditation(ctx, sqlc.CreateInstitutionAccreditationParams{
		InstitutionID:     institutionAccreditation.InstitutionID,
		LanguageCode:      institutionAccreditation.LanguageCode,
		AccreditationType: institutionAccreditation.AccreditationType,
		GivenBy:           institutionAccreditation.GivenBy,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution accreditation: %w", err))
	}

	institutionAccreditation.ID = institutionAccreditationResult.ID
	institutionAccreditation.CreatedAt = institutionAccreditationResult.CreatedAt.Time
	institutionAccreditation.UpdatedAt = institutionAccreditationResult.UpdatedAt.Time

	return institutionAccreditation, nil
}

func (r *pgInstitutionAccreditationRepository) Update(ctx context.Context, institutionAccreditation *domain.InstitutionAccreditation) (*domain.InstitutionAccreditation, error) {
	if institutionAccreditation.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution accreditation ID provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionAccreditation, err := q.GetInstitutionAccreditationByID(ctx, institutionAccreditation.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution accreditation with given ID(%d) not found", institutionAccreditation.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution accreditation with given ID(%d): %w", institutionAccreditation.ID, err))
		}

		updateInstitutionAccreditationParams := sqlc.UpdateInstitutionAccreditationParams{
			ID:                existingInstitutionAccreditation.ID,
			GivenBy:           existingInstitutionAccreditation.GivenBy,
			AccreditationType: existingInstitutionAccreditation.AccreditationType,
		}

		if institutionAccreditation.AccreditationType != "" {
			updateInstitutionAccreditationParams.AccreditationType = institutionAccreditation.AccreditationType
		}

		if institutionAccreditation.GivenBy != "" {
			updateInstitutionAccreditationParams.GivenBy = institutionAccreditation.GivenBy
		}

		updateInstitutionAccreditationResult, err := q.UpdateInstitutionAccreditation(ctx, updateInstitutionAccreditationParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution accreditation: %w", err))
		}

		institutionAccreditation.CreatedAt = updateInstitutionAccreditationResult.CreatedAt.Time
		institutionAccreditation.UpdatedAt = updateInstitutionAccreditationResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to perform update institution accreditation transaction: %w", err))
	}

	return institutionAccreditation, nil
}

func (r *pgInstitutionAccreditationRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution accreditation ID is provided for deletion"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionAccreditationByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution accreditation with given ID(%d) not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to check existance of institution accreditation by ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionAccreditation(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution accreditation: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to perform deletion institution accreditation transaction: %w", err))
	}

	return nil
}

func (r *pgInstitutionAccreditationRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionAccreditation, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution accreditation provided for retrival"))
	}

	institutionAccreditation, err := r.store.Queries.GetInstitutionAccreditationByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution accreditation with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionAccreditation{
		ID:            institutionAccreditation.ID,
		InstitutionID: institutionAccreditation.InstitutionID,
		LanguageCode:  institutionAccreditation.LanguageCode,
		AccreditationType:          institutionAccreditation.AccreditationType,
		GivenBy:       institutionAccreditation.GivenBy,
		CreatedAt:     institutionAccreditation.CreatedAt.Time,
		UpdatedAt:     institutionAccreditation.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionAccreditationRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionAccreditation, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution accreditation institution_id(%d) is provided", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution accreditation language_code(%s) is provided", langCode))
	}

	institutionAccreditationsResult, err := r.store.Queries.GetInstitutionAccreditationsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionAccreditationsByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution accreditation with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionAccreditations := make([]*domain.InstitutionAccreditation, len(institutionAccreditationsResult))
	for index, accreditation := range institutionAccreditationsResult {
		institutionAccreditations[index] = &domain.InstitutionAccreditation{
			ID:                accreditation.ID,
			InstitutionID:     accreditation.InstitutionID,
			LanguageCode:      accreditation.LanguageCode,
			AccreditationType: accreditation.AccreditationType,
			GivenBy:           accreditation.GivenBy,
			CreatedAt:         accreditation.CreatedAt.Time,
			UpdatedAt:         accreditation.UpdatedAt.Time,
		}
	}

	return institutionAccreditations, nil
}
