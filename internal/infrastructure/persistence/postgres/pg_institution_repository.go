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

type pgInstitutionRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgInstitutionRepository(store *Store) repositories.InstitutionRepository {
	return &pgInstitutionRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgInstitutionRepositoryWithQueries(q *sqlc.Queries) repositories.InstitutionRepository {
	return &pgInstitutionRepository{
		queries: q,
	}
}

func (r *pgInstitutionRepository) Create(ctx context.Context, institution *domain.Institution) (*domain.Institution, error) {
	if institution.Details == nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("institution details are required for creation"))
	}

	institutionResult, err := r.queries.CreateInstitution(ctx, sqlc.CreateInstitutionParams{
		YearOfEstablishment: pgtype.Int4{
			Int32: institution.YearOfEstablishment,
			Valid: true,
		},
		Email: institution.Email,
		Fax: pgtype.Text{
			String: institution.Fax,
			Valid:  true,
		},
		OfficialWebsite: institution.OfficialWebsite,
		PhoneNumber: pgtype.Text{
			String: institution.PhoneNumber,
			Valid:  true,
		},
		MailIndex: pgtype.Text{
			String: institution.MailIndex,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create institution: %w", err))
	}

	institution.ID = institutionResult.ID
	institution.CreatedAt = institutionResult.CreatedAt.Time
	institution.UpdatedAt = institutionResult.UpdatedAt.Time

	return institution, nil
}

func (r *pgInstitutionRepository) Update(ctx context.Context, institution *domain.Institution) (*domain.Institution, error) {
	updateInstitutionResult, err := r.queries.UpdateInstitution(ctx, sqlc.UpdateInstitutionParams{
		YearOfEstablishment: pgtype.Int4{
			Int32: institution.YearOfEstablishment,
			Valid: true,
		},
		Email: institution.Email,
		Fax: pgtype.Text{
			String: institution.Fax,
			Valid:  true,
		},
		OfficialWebsite: institution.OfficialWebsite,
		PhoneNumber: pgtype.Text{
			String: institution.PhoneNumber,
			Valid:  true,
		},
		MailIndex: pgtype.Text{
			String: institution.MailIndex,
			Valid:  true,
		},
		ID: institution.ID,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update institution: %w", err))
	}

	institution.CreatedAt = updateInstitutionResult.CreatedAt.Time
	institution.UpdatedAt = updateInstitutionResult.UpdatedAt.Time

	return institution, nil
}

func (r *pgInstitutionRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteInsitution(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution: %v", err))
	}
	return nil

}

func (r *pgInstitutionRepository) GetByID(ctx context.Context, id int64, langCode string) (*domain.Institution, error) {
	institutionResult, err := r.store.Queries.GetInstitutionByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution with given ID(%d)", id))
	}

	return &domain.Institution{
		ID:                  institutionResult.ID,
		YearOfEstablishment: institutionResult.YearOfEstablishment.Int32,
		Fax:                 institutionResult.Fax.String,
		Email:               institutionResult.Email,
		OfficialWebsite:     institutionResult.OfficialWebsite,
		PhoneNumber:         institutionResult.PhoneNumber.String,
		MailIndex:           institutionResult.MailIndex.String,
		CreatedAt:           institutionResult.CreatedAt.Time,
		UpdatedAt:           institutionResult.UpdatedAt.Time,
	}, nil
}

func (r *pgInstitutionRepository) GetAllInstitutions(ctx context.Context) ([]*domain.Institution, error) {
	institutionsResult, err := r.queries.GetAllInstitutions(ctx)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive all institutions: %w", err))
	}

	institutions := make([]*domain.Institution, len(institutionsResult))
	for index := range institutions {
		institutions[index] = &domain.Institution{
			ID:                  institutionsResult[index].ID,
			YearOfEstablishment: institutionsResult[index].YearOfEstablishment.Int32,
			Fax:                 institutionsResult[index].Fax.String,
			Email:               institutionsResult[index].Email,
			OfficialWebsite:     institutionsResult[index].OfficialWebsite,
			PhoneNumber:         institutionsResult[index].PhoneNumber.String,
			MailIndex:           institutionsResult[index].MailIndex.String,
			CreatedAt:           institutionsResult[index].CreatedAt.Time,
			UpdatedAt:           institutionsResult[index].UpdatedAt.Time,
		}
	}

	return institutions, nil
}
