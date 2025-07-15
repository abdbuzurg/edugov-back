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

type pgInstitutionRepository struct {
	store *Store
}

func NewPgInstitutionRepository(store *Store) repositories.InstitutionRepository {
	return &pgInstitutionRepository{
		store: store,
	}
}

func (r *pgInstitutionRepository) Create(ctx context.Context, institution *domain.Institution) (*domain.Institution, error) {
	if institution.Details == nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("institution details are required for creation"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		institutionResult, err := q.CreateInstitution(ctx, sqlc.CreateInstitutionParams{
			YearOfEstablishment: institution.YearOfEstablishment,
			Email:               institution.Email,
			Fax:                 institution.Fax,
			OfficialWebsite:     institution.OfficialWebsite,
		})

		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to create institution: %v", err))
		}

		institution.ID = institutionResult.ID
		institution.CreatedAt = institutionResult.CreatedAt.Time
		institution.UpdatedAt = institutionResult.UpdatedAt.Time

		institutionDetails, err := q.CreateInstitutionDetails(ctx, sqlc.CreateInstitutionDetailsParams{
			InstitutionID:    institution.ID,
			LanguageCode:     institution.Details.LanguageCode,
			InstitutionTitle: institution.Details.InstitutionTitle,
			InstitutionType:  institution.Details.InstitutionType,
			LegalStatus:      institution.Details.LegalStatus,
			Mission:          institution.Details.Mission,
			Founder:          institution.Details.Founder,
			LegalAddress:     institution.Details.LegalAddress,
			FactualAddress: pgtype.Text{
				String: *institution.Details.FactualAddress,
				Valid:  institution.Details.FactualAddress != nil,
			},
		})
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to create institution details: %w", err))
		}

		institution.Details.ID = institutionDetails.ID
		institution.Details.CreatedAt = institutionDetails.CreatedAt.Time
		institution.Details.UpdatedAt = institutionDetails.UpdatedAt.Time

		return nil
	})

	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed during institution creation: %w", err))
	}

	return institution, nil
}

func (r *pgInstitutionRepository) Update(ctx context.Context, institution *domain.Institution) (*domain.Institution, error) {
	if institution.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution ID is provided for update"))
	}

	if institution.Details == nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("institution details are required for update"))
	}

	if institution.Details.ID <= 0 {
		return nil, custom_errors.InternalServerError(fmt.Errorf("invalid institution details ID is provided for update"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitution, err := q.GetInstitutionByID(ctx, institution.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution with given ID(%d) not found", institution.ID))
			}

      return custom_errors.InternalServerError(fmt.Errorf("failed to retrived institution by ID(%d): %w", institution.ID, err))
		}

		updateInstitutionParams := sqlc.UpdateInstitutionParams{
			ID:                  existingInstitution.ID,
			YearOfEstablishment: existingInstitution.YearOfEstablishment,
			Fax:                 existingInstitution.Fax,
			Email:               existingInstitution.Email,
			OfficialWebsite:     existingInstitution.OfficialWebsite,
		}
		if institution.YearOfEstablishment != 0 {
			updateInstitutionParams.YearOfEstablishment = institution.YearOfEstablishment
		}

		if institution.Email != "" {
			updateInstitutionParams.Email = institution.Email
		}

		if institution.Fax != "" {
			updateInstitutionParams.Fax = institution.Fax
		}

		if institution.OfficialWebsite != "" {
			updateInstitutionParams.OfficialWebsite = institution.OfficialWebsite
		}

		updateInstitutionResult, err := q.UpdateInstitution(ctx, updateInstitutionParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution: %w", err))
		}

		institution.CreatedAt = updateInstitutionResult.CreatedAt.Time
		institution.UpdatedAt = updateInstitutionResult.UpdatedAt.Time

		existingInstitutionDetails, err := q.GetInstitutionDetailsByID(ctx, institution.Details.ID)
		if err != nil {
      if errors.Is(err, sql.ErrNoRows) {
        return custom_errors.BadRequest(fmt.Errorf("institution details with given ID(%d) not found", institution.Details.ID))
      }
       
      return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution details by ID(%d): %w", institution.Details.ID, err))
		}

		updateInstitutionDetailsParams := sqlc.UpdateInsitutionDetailsParams{
			ID:               existingInstitutionDetails.ID,
			InstitutionType:  existingInstitutionDetails.InstitutionType,
			InstitutionTitle: existingInstitutionDetails.InstitutionTitle,
			LegalStatus:      existingInstitutionDetails.LegalStatus,
			Mission:          existingInstitutionDetails.Mission,
			Founder:          existingInstitutionDetails.Founder,
			LegalAddress:     existingInstitutionDetails.LegalAddress,
			FactualAddress: pgtype.Text{
				String: *institution.Details.FactualAddress,
				Valid:  institution.Details.FactualAddress != nil,
			},
		}

		if institution.Details.InstitutionType != "" {
			updateInstitutionDetailsParams.InstitutionType = institution.Details.InstitutionType
		}

		if institution.Details.InstitutionTitle != "" {
			updateInstitutionDetailsParams.InstitutionTitle = institution.Details.InstitutionTitle
		}

		if institution.Details.LegalStatus != "" {
			updateInstitutionDetailsParams.LegalAddress = institution.Details.LegalStatus
		}

		if institution.Details.Mission != "" {
			updateInstitutionDetailsParams.Mission = institution.Details.Mission
		}

		if institution.Details.Founder != "" {
			updateInstitutionDetailsParams.Founder = institution.Details.Founder
		}

		if institution.Details.LegalAddress != "" {
			updateInstitutionDetailsParams.LegalAddress = institution.Details.LegalAddress
		}

		updateInstitutionDetailsResult, err := q.UpdateInsitutionDetails(ctx, updateInstitutionDetailsParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution details: %v", err))
		}

		institution.Details.CreatedAt = updateInstitutionDetailsResult.CreatedAt.Time
		institution.Details.UpdatedAt = updateInstitutionDetailsResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed during institution update: %w", err))
	}

	return institution, nil
}

func (r *pgInstitutionRepository) Delete(ctx context.Context, id int64) error {
	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution with given ID(%d) not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to check existing of institution by ID(%d)", id))
		}

		err = q.DeleteInsitution(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution: %v", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed during institution deletion: %v", err))
	}

	return nil
}

func (r *pgInstitutionRepository) GetByID(ctx context.Context, id int64, langCode string) (*domain.Institution, error) {
	institutionResult, err := r.store.Queries.GetInstitutionByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, custom_errors.BadRequest(fmt.Errorf("institution with given ID(%d) not found", id))
		}

    return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution with given ID(%d)", id))
	}

  institutionDetailsResult, err := r.store.Queries.GetInstitutionDetailsByID(ctx, institutionResult.ID)
  if err != nil && !errors.Is(err, sql.ErrNoRows) {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, custom_errors.BadRequest(fmt.Errorf("institution details with givin institution ID(%d) not found", institutionResult.ID))
    }

    return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution details with institution ID(%d): %w", institutionResult.ID, err))
  }

	return &domain.Institution{
    ID: institutionResult.ID,
    YearOfEstablishment: institutionResult.YearOfEstablishment,
    Fax: institutionResult.Fax,
    Email: institutionResult.Fax,
    OfficialWebsite: institutionResult.OfficialWebsite,
    CreatedAt: institutionResult.CreatedAt.Time,
    UpdatedAt: institutionDetailsResult.UpdatedAt.Time,
    Details: &domain.InstitutionDetails{
      ID: institutionDetailsResult.ID,
      InstitutionID: institutionDetailsResult.InstitutionID,
      InstitutionTitle: institutionDetailsResult.InstitutionTitle,
      InstitutionType: institutionDetailsResult.InstitutionType,
      LegalStatus: institutionDetailsResult.LegalStatus,
      Mission: institutionDetailsResult.Mission,
      Founder: institutionDetailsResult.Founder,
      LegalAddress: institutionDetailsResult.LegalAddress,
      FactualAddress: &institutionDetailsResult.FactualAddress.String,
    },
  }, nil
}
