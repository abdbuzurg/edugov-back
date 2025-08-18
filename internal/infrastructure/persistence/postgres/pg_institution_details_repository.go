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

type pgInstitutionDetailsRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPGInstitutionDetailsRepository(store *Store) repositories.InstitutionDetailsRepository {
	return &pgInstitutionDetailsRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPGInstitutionDetailsRepositoryWithQuery(q *sqlc.Queries) repositories.InstitutionDetailsRepository {
	return &pgInstitutionDetailsRepository{
		queries: q,
	}
}

func (r *pgInstitutionDetailsRepository) Create(ctx context.Context, institutionDetails *domain.InstitutionDetails) (*domain.InstitutionDetails, error) {
	institutionDetailsResult, err := r.queries.CreateInstitutionDetails(ctx, sqlc.CreateInstitutionDetailsParams{
		InstitutionID:         institutionDetails.ID,
		LanguageCode:          institutionDetails.LanguageCode,
		InstitutionTitleShort: institutionDetails.InstitutionTitleShort,
		InstitutionTitleLong:  institutionDetails.InstitutionTitleLong,
		InstitutionType:       institutionDetails.InstitutionType,
		LegalStatus: pgtype.Text{
			Valid:  true,
			String: institutionDetails.LegalStatus,
		},
		Mission: pgtype.Text{
			Valid:  true,
			String: institutionDetails.Mission,
		},
		Founder: pgtype.Text{
			Valid:  true,
			String: institutionDetails.Founder,
		},
		LegalAddress: institutionDetails.LegalAddress,
		FactualAddress: pgtype.Text{
			String: *institutionDetails.FactualAddress,
			Valid:  institutionDetails.FactualAddress != nil,
		},
		City: pgtype.Text{
			Valid:  true,
			String: institutionDetails.City,
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution_details: %w", err))
	}

	institutionDetails.CreatedAt = institutionDetailsResult.CreatedAt.Time
	institutionDetails.UpdatedAt = institutionDetailsResult.UpdatedAt.Time

	return institutionDetails, nil
}

func (r *pgInstitutionDetailsRepository) Update(ctx context.Context, institutionDetails *domain.InstitutionDetails) (*domain.InstitutionDetails, error) {
	institutionDetailsResult, err := r.queries.UpdateInsitutionDetails(ctx, sqlc.UpdateInsitutionDetailsParams{
		InstitutionTitleShort: institutionDetails.InstitutionTitleShort,
		InstitutionTitleLong:  institutionDetails.InstitutionTitleLong,
		InstitutionType:       institutionDetails.InstitutionType,
		LegalStatus: pgtype.Text{
			Valid:  true,
			String: institutionDetails.LegalStatus,
		},
		Mission: pgtype.Text{
			Valid:  true,
			String: institutionDetails.Mission,
		},
		Founder: pgtype.Text{
			Valid:  true,
			String: institutionDetails.Founder,
		},
		LegalAddress: institutionDetails.LegalAddress,
		FactualAddress: pgtype.Text{
			String: *institutionDetails.FactualAddress,
			Valid:  institutionDetails.FactualAddress != nil,
		},
		City: pgtype.Text{
			Valid:  true,
			String: institutionDetails.City,
		},
		ID: institutionDetails.ID,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update institution_details: %w", err))
	}

	institutionDetails.CreatedAt = institutionDetailsResult.CreatedAt.Time
	institutionDetails.UpdatedAt = institutionDetailsResult.UpdatedAt.Time

	return institutionDetails, nil
}

func (r *pgInstitutionDetailsRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteInstitutionDetails(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution details by ID(%d): %w", id, err))
	}

	return nil
}
func (r *pgInstitutionDetailsRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) (*domain.InstitutionDetails, error) {
	institutionResult, err := r.queries.GetInstitutionDetailsByInstitutionIDAndLanguage(ctx, sqlc.GetInstitutionDetailsByInstitutionIDAndLanguageParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution_details by ID(%d) and language_code(%s): %w", institutionID, langCode, err))
	}

	return &domain.InstitutionDetails{
		ID:                    institutionResult.ID,
		InstitutionID:         institutionResult.InstitutionID,
		InstitutionTitleShort: institutionResult.InstitutionTitleShort,
		InstitutionTitleLong:  institutionResult.InstitutionTitleLong,
		InstitutionType:       institutionResult.InstitutionType,
		LanguageCode:          institutionResult.LanguageCode,
		LegalStatus:           institutionResult.LegalStatus.String,
		Mission:               institutionResult.Mission.String,
		Founder:               institutionResult.Founder.String,
		LegalAddress:          institutionResult.LegalAddress,
		FactualAddress:        &institutionResult.FactualAddress.String,
		City:                  institutionResult.City.String,
		CreatedAt:             institutionResult.CreatedAt.Time,
		UpdatedAt:             institutionResult.UpdatedAt.Time,
	}, nil
}
