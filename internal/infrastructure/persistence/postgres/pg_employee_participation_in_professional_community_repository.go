package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgEmployeeParticipationInProfessionalCommunityRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPgEmployeeParticipationInProfessionalCommunityRepository(store *Store) repositories.EmployeeParticipationInProfessionalCommunityRepository {
	return &pgEmployeeParticipationInProfessionalCommunityRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgEmployeeParticipationInProfessionalCommunityRepositoryWithQuery(q *sqlc.Queries) repositories.EmployeeParticipationInProfessionalCommunityRepository {
	return &pgEmployeeParticipationInProfessionalCommunityRepository{
		queries: q,
	}
}

func (r *pgEmployeeParticipationInProfessionalCommunityRepository) Create(ctx context.Context, employeePIPC *domain.EmployeeParticipationInProfessionalCommunity) (*domain.EmployeeParticipationInProfessionalCommunity, error) {
	participationInProfessionalCommunity, err := r.queries.CreateEmployeeParticipationInProfessionalCommunity(ctx, sqlc.CreateEmployeeParticipationInProfessionalCommunityParams{
		EmployeeID:                  employeePIPC.EmployeeID,
		LanguageCode:                employeePIPC.LanguageCode,
		ProfessionalCommunityTitle:  employeePIPC.ProfessionalCommunityTitle,
		RoleInProfessionalCommunity: employeePIPC.RoleInProfessionalCommunity,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create employee participation in professional community: %w", err))
	}

	employeePIPC.ID = participationInProfessionalCommunity.ID
	employeePIPC.CreatedAt = participationInProfessionalCommunity.CreatedAt.Time
	employeePIPC.UpdatedAt = participationInProfessionalCommunity.UpdatedAt.Time

	return employeePIPC, nil
}

func (r *pgEmployeeParticipationInProfessionalCommunityRepository) Update(ctx context.Context, employeePIPC *domain.EmployeeParticipationInProfessionalCommunity) (*domain.EmployeeParticipationInProfessionalCommunity, error) {
	updateEmployeeParticipationInProfessionalCommunityResult, err := r.queries.UpdateEmployeeParticipationInProfessionalCommunity(ctx, sqlc.UpdateEmployeeParticipationInProfessionalCommunityParams{
		ID:                          employeePIPC.ID,
		ProfessionalCommunityTitle:  employeePIPC.ProfessionalCommunityTitle,
		RoleInProfessionalCommunity: employeePIPC.RoleInProfessionalCommunity,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update employee participation in professional community: %w", err))
	}

	employeePIPC.CreatedAt = updateEmployeeParticipationInProfessionalCommunityResult.CreatedAt.Time
	employeePIPC.UpdatedAt = updateEmployeeParticipationInProfessionalCommunityResult.UpdatedAt.Time

	return employeePIPC, nil
}

func (r *pgEmployeeParticipationInProfessionalCommunityRepository) Delete(ctx context.Context, id int64) error {
	if err := r.queries.DeleteEmployeeParticipationInProfessionalCommunity(ctx, id); err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete employee participation in professional community: %w", err))
	}

	return nil
}

func (r *pgEmployeeParticipationInProfessionalCommunityRepository) GetByID(ctx context.Context, id int64) (*domain.EmployeeParticipationInProfessionalCommunity, error) {
	employeePIPC, err := r.queries.GetEmployeeParticipationInProfessionalCommunityByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee participation in professional community with give ID(%d): %w", id, err))
	} 

	return &domain.EmployeeParticipationInProfessionalCommunity{
		ID:                          employeePIPC.ID,
		EmployeeID:                  employeePIPC.EmployeeID,
		LanguageCode:                employeePIPC.LanguageCode,
		ProfessionalCommunityTitle:  employeePIPC.ProfessionalCommunityTitle,
		RoleInProfessionalCommunity: employeePIPC.RoleInProfessionalCommunity,
		CreatedAt:                   employeePIPC.CreatedAt.Time,
		UpdatedAt:                   employeePIPC.UpdatedAt.Time,
	}, nil
}

func (r *pgEmployeeParticipationInProfessionalCommunityRepository) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeParticipationInProfessionalCommunity, error) {
	employeePIPCsResult, err := r.queries.GetEmployeeParticipationInProfessionalCommunitysByEmployeeIDAndLanguageCode(ctx, sqlc.GetEmployeeParticipationInProfessionalCommunitysByEmployeeIDAndLanguageCodeParams{
		EmployeeID:   employeeID,
		LanguageCode: langCode,
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive employee participation in professional communitys with given EmployeeID(%d) and language_code(%s): %w", employeeID, langCode, err))
	}

	employeePIPCs := make([]*domain.EmployeeParticipationInProfessionalCommunity, len(employeePIPCsResult))
	for index, participationInProfessionalCommunity := range employeePIPCsResult {
		employeePIPCs[index] = &domain.EmployeeParticipationInProfessionalCommunity{
			ID:                          participationInProfessionalCommunity.ID,
			EmployeeID:                  participationInProfessionalCommunity.EmployeeID,
			LanguageCode:                participationInProfessionalCommunity.LanguageCode,
			ProfessionalCommunityTitle:  participationInProfessionalCommunity.ProfessionalCommunityTitle,
			RoleInProfessionalCommunity: participationInProfessionalCommunity.RoleInProfessionalCommunity,
			CreatedAt:                   participationInProfessionalCommunity.CreatedAt.Time,
			UpdatedAt:                   participationInProfessionalCommunity.UpdatedAt.Time,
		}
	}

	return employeePIPCs, nil
}
