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

type pgInstitutionProjectRepository struct {
	store *Store
}

func NewPgInstitutionProjectRepository(store *Store) repositories.InstitutionProjectRepository {
	return &pgInstitutionProjectRepository{
		store: store,
	}
}

func (r *pgInstitutionProjectRepository) Create(ctx context.Context, institutionProject *domain.InstitutionProject) (*domain.InstitutionProject, error) {
	if institutionProject.InstitutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project instituion_id is provided"))
	}

	if institutionProject.LanguageCode != "en" && institutionProject.LanguageCode != "ru" && institutionProject.LanguageCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project language_Code is provided"))
	}

	if len(institutionProject.Partners) == 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project partners is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		institutionProjectResult, err := q.CreateInstitutionProject(ctx, sqlc.CreateInstitutionProjectParams{
			InstitutionID: institutionProject.InstitutionID,
			LanguageCode:  institutionProject.LanguageCode,
			ProjectType:   institutionProject.ProjectType,
			ProjectTitle:  institutionProject.ProjectTitle,
			DateStart: pgtype.Date{
				Time:  institutionProject.DateStart,
				Valid: !institutionProject.DateStart.IsZero(),
			},
			DateEnd: pgtype.Date{
				Time:  institutionProject.DateEnd,
				Valid: !institutionProject.DateEnd.IsZero(),
			},
			Fund:            institutionProject.Fund,
			InstitutionRole: institutionProject.InstitutionRole,
			Coordinator:     institutionProject.Coordinator,
		})
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to create institution project: %w", err))
		}

		institutionProject.ID = institutionProjectResult.ID
		institutionProject.CreatedAt = institutionProjectResult.CreatedAt.Time
		institutionProject.UpdatedAt = institutionProjectResult.UpdatedAt.Time

		for index, projectPartner := range institutionProject.Partners {
			institutionProjectPartnerResult, err := q.CreateInstitutionProjectPartner(ctx, sqlc.CreateInstitutionProjectPartnerParams{
				InstitutionProjectID: institutionProject.ID,
				LanguageCode:         projectPartner.LanguageCode,
				PartnerName:          projectPartner.PartnerName,
				PartnerType:          projectPartner.PartnerType,
				LinkToPartner:        projectPartner.LinkToPartner,
			})
			if err != nil {
				return custom_errors.InternalServerError(fmt.Errorf("failed to create institution project partner: %w", err))
			}

			institutionProject.Partners[index].ID = institutionProjectPartnerResult.ID
			institutionProject.Partners[index].CreatedAt = institutionProjectPartnerResult.CreatedAt.Time
			institutionProject.Partners[index].UpdatedAt = institutionProjectPartnerResult.UpdatedAt.Time
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed creating institution project: %w", err))
	}

	return institutionProject, nil
}

func (r *pgInstitutionProjectRepository) Update(ctx context.Context, institutionProject *domain.InstitutionProject) (*domain.InstitutionProject, error) {
	if institutionProject.ID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		existingInstitutionProject, err := q.GetInstitutionProjectByID(ctx, institutionProject.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution project with given ID(%d) is not found", institutionProject.ID))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive instituion licence with given ID(%d): %w", institutionProject.ID, err))
		}

		updateInstitutionProjectParams := sqlc.UpdateInstitutionProjectParams{
			ID:              existingInstitutionProject.ID,
			ProjectTitle:    existingInstitutionProject.ProjectTitle,
			ProjectType:     existingInstitutionProject.ProjectType,
			DateStart:       existingInstitutionProject.DateStart,
			DateEnd:         existingInstitutionProject.DateEnd,
			Fund:            existingInstitutionProject.Fund,
			InstitutionRole: existingInstitutionProject.InstitutionRole,
			Coordinator:     existingInstitutionProject.Coordinator,
		}

		if institutionProject.ProjectTitle != "" {
			updateInstitutionProjectParams.ProjectTitle = institutionProject.ProjectTitle
		}

		if institutionProject.ProjectType != "" {
			updateInstitutionProjectParams.ProjectType = institutionProject.ProjectType
		}

		if !institutionProject.DateStart.IsZero() {
			updateInstitutionProjectParams.DateStart = pgtype.Date{
				Time:  institutionProject.DateStart,
				Valid: true,
			}
		}

		if !institutionProject.DateEnd.IsZero() {
			updateInstitutionProjectParams.DateEnd = pgtype.Date{
				Time:  institutionProject.DateEnd,
				Valid: true,
			}
		}

		if institutionProject.Fund > 0 {
			updateInstitutionProjectParams.Fund = institutionProject.Fund
		}

		if institutionProject.InstitutionRole != "" {
			updateInstitutionProjectParams.InstitutionRole = institutionProject.InstitutionRole
		}

		if institutionProject.Coordinator != "" {
			updateInstitutionProjectParams.Coordinator = institutionProject.Coordinator
		}

		updateInstitutionProjectResult, err := q.UpdateInstitutionProject(ctx, updateInstitutionProjectParams)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to update institution project: %w", err))
		}

		institutionProject.CreatedAt = updateInstitutionProjectResult.CreatedAt.Time
		institutionProject.UpdatedAt = updateInstitutionProjectResult.UpdatedAt.Time

		if len(institutionProject.Partners) == 0 {
			return nil
		}

		for index, projectPartner := range institutionProject.Partners {
			if projectPartner.ID <= 0 {
				return custom_errors.BadRequest(fmt.Errorf("invalid institution project partner ID"))
			}

			existingInstitutionProjectPartner, err := q.GetInstitutionProjectPartnerByID(ctx, projectPartner.ID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return custom_errors.BadRequest(fmt.Errorf("institution project partner with given ID(%d) is not found", projectPartner.ID))
				}

				return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project partner with given ID(%d): %w", projectPartner.ID, err))
			}

			updateInstitutionProjectPartnerParams := sqlc.UpdateInstitutionProjectPartnerParams{
				ID:            existingInstitutionProjectPartner.ID,
				PartnerName:   existingInstitutionProjectPartner.PartnerName,
				PartnerType:   existingInstitutionProjectPartner.PartnerType,
				LinkToPartner: existingInstitutionProjectPartner.LinkToPartner,
			}

			if projectPartner.PartnerName != "" {
				updateInstitutionProjectPartnerParams.PartnerName = projectPartner.PartnerName
			}

			if projectPartner.PartnerType != "" {
				updateInstitutionProjectPartnerParams.PartnerType = projectPartner.PartnerType
			}

			if projectPartner.LinkToPartner != "" {
				updateInstitutionProjectPartnerParams.LinkToPartner = projectPartner.LinkToPartner
			}

			updateInstitutionProjectPartnerResult, err := q.UpdateInstitutionProjectPartner(ctx, updateInstitutionProjectPartnerParams)
			if err != nil {
				return custom_errors.InternalServerError(fmt.Errorf("failed to update institution project partner: %w", err))
			}

			institutionProject.Partners[index].CreatedAt = updateInstitutionProjectPartnerResult.CreatedAt.Time
			institutionProject.Partners[index].UpdatedAt = updateInstitutionProjectPartnerResult.UpdatedAt.Time
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed while perfoming institution project update: %w", err))
	}

	return institutionProject, nil
}

func (r *pgInstitutionProjectRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid institution project ID is provided"))
	}

	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		_, err := q.GetInstitutionProjectByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return custom_errors.BadRequest(fmt.Errorf("institution project with given ID(%d) is not found", id))
			}

			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project with given ID(%d): %w", id, err))
		}

		err = q.DeleteInstitutionProject(ctx, id)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to delete institution project: %w", err))
		}

		return nil
	})
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed while deliting institution project: %w", err))
	}

	return nil
}

func (r *pgInstitutionProjectRepository) GetByID(ctx context.Context, id int64) (*domain.InstitutionProject, error) {
	if id <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project ID is provided"))
	}

	var institutionProject *domain.InstitutionProject
	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		institutionProjectResult, err := q.GetInstitutionProjectByID(ctx, id)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project with given ID(%d): %w", id, err))
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		institutionProjectPartnersResult, err := q.GetInstitutionProjectPartnersByInstitutionProjectIDAndLanguageCode(ctx, sqlc.GetInstitutionProjectPartnersByInstitutionProjectIDAndLanguageCodeParams{
			InstitutionProjectID: institutionProjectResult.ID,
			LanguageCode:         institutionProjectResult.LanguageCode,
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project with given institutionProjectID(%d): %w", institutionProjectResult.ID, err))
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		institutionProjectPartners := make([]*domain.InstitutionProjectPartner, len(institutionProjectPartnersResult))
		for index, projectPartner := range institutionProjectPartnersResult {
			institutionProjectPartners[index] = &domain.InstitutionProjectPartner{
				ID:                   projectPartner.ID,
				InstitutionProjectID: projectPartner.InstitutionProjectID,
				LanguageCode:         projectPartner.LanguageCode,
				PartnerType:          projectPartner.PartnerType,
				PartnerName:          projectPartner.PartnerName,
				LinkToPartner:        projectPartner.LinkToPartner,
				CreatedAt:            projectPartner.CreatedAt.Time,
				UpdatedAt:            projectPartner.UpdatedAt.Time,
			}
		}

		institutionProject = &domain.InstitutionProject{
			ID:              institutionProjectResult.ID,
			InstitutionID:   institutionProjectResult.InstitutionID,
			LanguageCode:    institutionProjectResult.LanguageCode,
			ProjectType:     institutionProjectResult.ProjectType,
			ProjectTitle:    institutionProjectResult.ProjectTitle,
			DateStart:       institutionProjectResult.DateStart.Time,
			DateEnd:         institutionProjectResult.DateEnd.Time,
			Fund:            institutionProjectResult.Fund,
			InstitutionRole: institutionProjectResult.InstitutionRole,
			Coordinator:     institutionProjectResult.Coordinator,
			Partners:        institutionProjectPartners,
			CreatedAt:       institutionProjectResult.CreatedAt.Time,
			UpdatedAt:       institutionProjectResult.UpdatedAt.Time,
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed retriving data for institution project: %w", err))
	}

	return institutionProject, nil
}

func (r *pgInstitutionProjectRepository) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionProject, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project institution_id is provided"))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid institution project language_code is provided"))
	}

	var institutionProjects []*domain.InstitutionProject
	err := r.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		institutionProjectsResult, err := r.store.Queries.GetInstitutionProjectsByInstitutionIDAndLanguageCode(ctx, sqlc.GetInstitutionProjectsByInstitutionIDAndLanguageCodeParams{
			InstitutionID: institutionID,
			LanguageCode:  langCode,
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project with given institution_id(%d) and language_code(%s): %w", institutionID, langCode, err))
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

    institutionProjects = make([]*domain.InstitutionProject, len(institutionProjectsResult))
		for index, project := range institutionProjectsResult {
			institutionProjectPartnersResult, err := q.GetInstitutionProjectPartnersByInstitutionProjectIDAndLanguageCode(ctx, sqlc.GetInstitutionProjectPartnersByInstitutionProjectIDAndLanguageCodeParams{
				InstitutionProjectID: project.ID,
				LanguageCode:         project.LanguageCode,
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return custom_errors.InternalServerError(fmt.Errorf("failed to retrive institution project with given institutionProjectID(%d): %w", project.ID, err))
			} else if errors.Is(err, sql.ErrNoRows) {
				return nil
			}

			institutionProjectPartners := make([]*domain.InstitutionProjectPartner, len(institutionProjectPartnersResult))
			for subIndex, projectPartner := range institutionProjectPartnersResult {
				institutionProjectPartners[subIndex] = &domain.InstitutionProjectPartner{
					ID:                   projectPartner.ID,
					InstitutionProjectID: projectPartner.InstitutionProjectID,
					LanguageCode:         projectPartner.LanguageCode,
					PartnerType:          projectPartner.PartnerType,
					PartnerName:          projectPartner.PartnerName,
					LinkToPartner:        projectPartner.LinkToPartner,
					CreatedAt:            projectPartner.CreatedAt.Time,
					UpdatedAt:            projectPartner.UpdatedAt.Time,
				}
			}

			institutionProjects[index] = &domain.InstitutionProject{
				ID:              project.ID,
				InstitutionID:   project.InstitutionID,
				LanguageCode:    project.LanguageCode,
				ProjectType:     project.ProjectType,
				ProjectTitle:    project.ProjectTitle,
				DateStart:       project.DateStart.Time,
				DateEnd:         project.DateEnd.Time,
				Fund:            project.Fund,
				InstitutionRole: project.InstitutionRole,
				Coordinator:     project.Coordinator,
				Partners:        institutionProjectPartners,
				CreatedAt:       project.CreatedAt.Time,
				UpdatedAt:       project.UpdatedAt.Time,
			}
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed retriving institution project: %w", err))
	}

	return institutionProjects, nil
}
