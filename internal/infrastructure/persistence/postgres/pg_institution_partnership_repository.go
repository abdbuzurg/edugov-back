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

type pgInstitutionPartnershipRepository struct {
	store *Store
}

func NewPgInstitutionPartnershipRepository(store *Store) repositories.InstitutionPartnershipRepository {
	return &pgInstitutionPartnershipRepository{
		store: store,
	}
}

func (r *pgInstitutionPartnershipRepository) Create(ctx context.Context, institutionPartnership *domain.InstitutionPartnership) (*domain.InstitutionPartnership, error) {
	if institutionPartnership.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership instituion_id is provided"))
	}

	if institutionPartnership.LanguageCode != "en" && institutionPartnership.LanguageCode != "ru" && institutionPartnership.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership language_Code is provided"))
	}

	institutionPartnershipResult, err := r.store.Queries.CreateInstitutionPartnership(ctx, sqlc.CreateInstitutionPartnershipParams{
		InstitutionID: institutionPartnership.InstitutionID,
		LanguageCode:  institutionPartnership.LanguageCode,
		PartnerName:   institutionPartnership.PartnerName,
		PartnerType:   institutionPartnership.PartnerType,
		DateOfContract: pgtype.Date{
			Time:  institutionPartnership.DateOfContract,
			Valid: !institutionPartnership.DateOfContract.IsZero(),
		},
		LinkToPartner: institutionPartnership.LinkToPartner,
		Goal:          institutionPartnership.Goal,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution partnership: %w", err))
	}

	institutionPartnership.ID = institutionPartnershipResult.ID
	institutionPartnership.CreatedAt = institutionPartnershipResult.CreatedAt.Time
	institutionPartnership.UpdatedAt = institutionPartnershipResult.UpdatedAt.Time

	return institutionPartnership, nil
}

func (r *pgInstitutionPartnershipRepository) Update(ctx context.Context, institutionPartnership *domain.InstitutionPartnership) (*domain.InstitutionPartnership, error) {
	if institutionPartnership.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionPartnership, err := q.GetInstitutionPartnershipByID(ctx, institutionPartnership.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution partnership with given ID(%d) is not found", institutionPartnership.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionPartnership.ID, err))
		}

		updateInstitutionPartnershipParams := sqlc.UpdateInstitutionPartnershipParams{
			ID:             existingInstitutionPartnership.ID,
			PartnerName:    existingInstitutionPartnership.PartnerName,
			PartnerType:    existingInstitutionPartnership.PartnerType,
			DateOfContract: existingInstitutionPartnership.DateOfContract,
			LinkToPartner:  existingInstitutionPartnership.LinkToPartner,
			Goal:           existingInstitutionPartnership.Goal,
		}

		if institutionPartnership.PartnerName != "" {
			updateInstitutionPartnershipParams.PartnerName = institutionPartnership.PartnerName
		}

		if institutionPartnership.PartnerType != "" {
			updateInstitutionPartnershipParams.PartnerType = institutionPartnership.PartnerType
		}

		if !institutionPartnership.DateOfContract.IsZero() {
			updateInstitutionPartnershipParams.DateOfContract = pgtype.Date{
				Time:  institutionPartnership.DateOfContract,
				Valid: true,
			}
		}

		if institutionPartnership.LinkToPartner != "" {
			updateInstitutionPartnershipParams.LinkToPartner = institutionPartnership.LinkToPartner
		}

		if institutionPartnership.Goal != "" {
			updateInstitutionPartnershipParams.Goal = institutionPartnership.Goal
		}

		updateInstitutionPartnershipResult, err := q.UpdateInstitutionPartnership(ctx, updateInstitutionPartnershipParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution partnership: %w", err))
		}

		institutionPartnership.CreatedAt = updateInstitutionPartnershipResult.CreatedAt.Time
		institutionPartnership.UpdatedAt = updateInstitutionPartnershipResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution partnership update: %w", err))
	}

	return institutionPartnership, nil
}

func (r *pgInstitutionPartnershipRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution partnership ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionPartnershipByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution partnership with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution partnership with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionPartnership(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution partnership: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution partnership: %w", err))
	}

	return nil
}

func (r *pgInstitutionPartnershipRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionPartnership, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership ID is provided"))
	}

	institutionPartnershipResult, err := r.store.Queries.GetInstitutionPartnershipByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution partnership with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionPartnership{
		ID:             institutionPartnershipResult.ID,
		InstitutionID:  institutionPartnershipResult.InstitutionID,
		LanguageCode:   institutionPartnershipResult.LanguageCode,
		PartnerName:    institutionPartnershipResult.PartnerName,
		PartnerType:    institutionPartnershipResult.PartnerType,
		Goal:           institutionPartnershipResult.Goal,
		LinkToPartner:  institutionPartnershipResult.LinkToPartner,
		DateOfContract: institutionPartnershipResult.DateOfContract.Time,
		CreatedAt:      institutionPartnershipResult.CreatedAt.Time,
		UpdatedAt:      institutionPartnershipResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionPartnershipRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionPartnership, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution partnership language_code is provided"))
	}

	institutionPartnershipsResult, err := r.store.Queries.GetInstitutionPartnershipsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionPartnershipsByInstitutionIDAndLanguageCodeParams{
		InstitutionID: institutionID,
		LanguageCode:  langCode,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution partnership with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionPartnerships := make([]*domain.InstitutionPartnership, len(institutionPartnershipsResult))
	for index, partnership := range institutionPartnershipsResult {
		institutionPartnerships[index] = &domain.InstitutionPartnership{
			ID:             partnership.ID,
			InstitutionID:  partnership.InstitutionID,
			LanguageCode:   partnership.LanguageCode,
			PartnerName:    partnership.PartnerName,
			PartnerType:    partnership.PartnerType,
			Goal:           partnership.Goal,
			LinkToPartner:  partnership.LinkToPartner,
			DateOfContract: partnership.DateOfContract.Time,
			CreatedAt:      partnership.CreatedAt.Time,
			UpdatedAt:      partnership.UpdatedAt.Time,
		}
	}

	return institutionPartnerships, nil
}
