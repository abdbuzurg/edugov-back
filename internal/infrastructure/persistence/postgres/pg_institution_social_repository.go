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

type pgInstitutionSocialRepository struct {
	store *Store
}

func NewPgInstitutionSocialRepository(store *Store) repositories.InstitutionSocialRepository {
	return &pgInstitutionSocialRepository{
		store: store,
	}
}

func (r *pgInstitutionSocialRepository) Create(ctx context.Context, institutionSocial *domain.InstitutionSocial) (*domain.InstitutionSocial, error) {
	if institutionSocial.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution social instituion_id is provided"))
	}

	institutionSocialResult, err := r.store.Queries.CreateInstitutionSocial(ctx, sqlc.CreateInstitutionSocialParams{
		InstitutionID: institutionSocial.InstitutionID,
		LinkToSocial:  institutionSocial.LinkToSocial,
		SocialName:    institutionSocial.SocialName,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution social: %w", err))
	}

	institutionSocial.ID = institutionSocialResult.ID
	institutionSocial.CreatedAt = institutionSocialResult.CreatedAt.Time
	institutionSocial.UpdatedAt = institutionSocialResult.UpdatedAt.Time

	return institutionSocial, nil
}

func (r *pgInstitutionSocialRepository) Update(ctx context.Context, institutionSocial *domain.InstitutionSocial) (*domain.InstitutionSocial, error) {
	if institutionSocial.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution social ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionSocial, err := q.GetInstitutionSocialByID(ctx, institutionSocial.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution social with given ID(%d) is not found", institutionSocial.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionSocial.ID, err))
		}

		updateInstitutionSocialParams := sqlc.UpdateInstitutionSocialParams{
			ID:           existingInstitutionSocial.ID,
			LinkToSocial: existingInstitutionSocial.LinkToSocial,
			SocialName:   existingInstitutionSocial.SocialName,
		}

		if institutionSocial.LinkToSocial != "" {
			updateInstitutionSocialParams.LinkToSocial = institutionSocial.LinkToSocial
		}

		if institutionSocial.SocialName != "" {
			updateInstitutionSocialParams.SocialName = institutionSocial.SocialName
		}

		updateInstitutionSocialResult, err := q.UpdateInstitutionSocial(ctx, updateInstitutionSocialParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution social: %w", err))
		}

		institutionSocial.CreatedAt = updateInstitutionSocialResult.CreatedAt.Time
		institutionSocial.UpdatedAt = updateInstitutionSocialResult.UpdatedAt.Time

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution social update: %w", err))
	}

	return institutionSocial, nil
}

func (r *pgInstitutionSocialRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution social ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionSocialByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution social with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution social with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionSocial(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution social: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution social: %w", err))
	}

	return nil
}

func (r *pgInstitutionSocialRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionSocial, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution social ID is provided"))
	}

	institutionSocialResult, err := r.store.Queries.GetInstitutionSocialByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution social with given ID(%d): %w", id, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &domain.InstitutionSocial{
		ID:            institutionSocialResult.ID,
		InstitutionID: institutionSocialResult.InstitutionID,
		LinkToSocial:  institutionSocialResult.LinkToSocial,
		SocialName:    institutionSocialResult.SocialName,
		CreatedAt:     institutionSocialResult.CreatedAt.Time,
		UpdatedAt:     institutionSocialResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionSocialRepository) GetByInstitutionID(ctx context.Context, institutionID int64) ([]*domain.InstitutionSocial, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution social institution_id is provided"))
	}

	institutionSocialsResult, err := r.store.Queries.GetInstitutionSocialsByInstitutionID(ctx, institutionID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution social with given institution_id(%d): %w", institutionID, err))
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	institutionSocials := make([]*domain.InstitutionSocial, len(institutionSocialsResult))
	for index, social := range institutionSocialsResult {
		institutionSocials[index] = &domain.InstitutionSocial{
			ID:            social.ID,
			InstitutionID: social.InstitutionID,
			LinkToSocial:  social.LinkToSocial,
			SocialName:    social.SocialName,
			CreatedAt:     social.CreatedAt.Time,
			UpdatedAt:     social.UpdatedAt.Time,
		}
	}

	return institutionSocials, nil
}
