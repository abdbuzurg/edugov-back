package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/shared/custom_errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type InstitutionProjectUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionProjectRequest) (*dtos.InstitutionProjectResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionProjectRequest) (*dtos.InstitutionProjectResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institution int64, langCode string) ([]*dtos.InstitutionProjectResponse, error)
}

type institutionProjectUsecase struct {
	institutionProjectRepo repositories.InstitutionProjectRepository
	validator              *validator.Validate
}

func NewInstitutionProjectUsecase(
	institutionProjectRepo repositories.InstitutionProjectRepository,
	validator *validator.Validate,
) InstitutionProjectUsecase {
	return &institutionProjectUsecase{
		institutionProjectRepo: institutionProjectRepo,
		validator:              validator,
	}
}

func (uc *institutionProjectUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionProjectRequest) (*dtos.InstitutionProjectResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution project: %w", err))
	}

	institutionProject := &domain.InstitutionProject{
		InstitutionID:   req.InstitutionID,
		LanguageCode:    req.LanguageCode,
		ProjectType:     req.ProjectType,
		ProjectTitle:    req.ProjectTitle,
		DateStart:       req.DateStart,
		DateEnd:         req.DateEnd,
		Fund:            req.Fund,
		InstitutionRole: req.InstitutionRole,
		Coordinator:     req.Coordinator,
		Partners:        make([]*domain.InstitutionProjectPartner, len(req.Partners)),
	}

	for index, partner := range req.Partners {
		institutionProject.Partners[index] = &domain.InstitutionProjectPartner{
			LanguageCode:  partner.LanguageCode,
			PartnerType:   partner.PartnerType,
			PartnerName:   partner.PartnerName,
			LinkToPartner: partner.LinkToPartner,
		}
	}

	createdInstitutionProject, err := uc.institutionProjectRepo.Create(ctx, institutionProject)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionProjectDomainToResponseDTO(createdInstitutionProject)
	return resp, nil
}

func (uc *institutionProjectUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionProjectRequest) (*dtos.InstitutionProjectResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution project: %w", err))
	}

	institutionProject := &domain.InstitutionProject{
		ID:       req.ID,
		Partners: make([]*domain.InstitutionProjectPartner, len(req.Partners)),
	}

	if req.ProjectType != nil {
		institutionProject.ProjectType = *req.ProjectType
	}

	if req.ProjectTitle != nil {
		institutionProject.ProjectTitle = *req.ProjectTitle
	}

	if req.DateStart != nil {
		institutionProject.DateStart = *req.DateStart
	}

	if req.DateEnd != nil {
		institutionProject.DateEnd = *req.DateEnd
	}

	if req.Fund != nil {
		institutionProject.Fund = *req.Fund
	}

	if req.InstitutionRole != nil {
		institutionProject.InstitutionRole = *req.InstitutionRole
	}

	if req.Coordinator != nil {
		institutionProject.Coordinator = *req.Coordinator
	}

	for index, partner := range req.Partners {
		institutionProject.Partners[index] = &domain.InstitutionProjectPartner{
			ID: partner.ID,
		}

		if partner.PartnerName != nil {
			institutionProject.Partners[index].PartnerName = *partner.PartnerName
		}

		if partner.PartnerType != nil {
			institutionProject.Partners[index].PartnerType = *partner.PartnerType
		}

		if partner.LinkToPartner != nil {
			institutionProject.Partners[index].LinkToPartner = *partner.LinkToPartner
		}
	}

	updatedInstitutionProject, err := uc.institutionProjectRepo.Update(ctx, institutionProject)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionProjectDomainToResponseDTO(updatedInstitutionProject)
	return resp, nil
}

func (uc *institutionProjectUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution project"))
	}

	if err := uc.institutionProjectRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionProjectUsecase) GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionProjectResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution project", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution project", langCode))
	}

	institutionProjects, err := uc.institutionProjectRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionProjectResponse, len(institutionProjects))
	for index, degree := range institutionProjects {
		resp[index] = MapInstitutionProjectDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionProjectDomainToResponseDTO(institutionProject *domain.InstitutionProject) *dtos.InstitutionProjectResponse {
	if institutionProject == nil {
		return nil
	}

	partners := make([]*dtos.InstitutionProjectPartnerResponse, len(institutionProject.Partners))
	for index, partner := range institutionProject.Partners {
		partners[index] = &dtos.InstitutionProjectPartnerResponse{
			ID:            partner.ID,
			PartnerType:   partner.PartnerType,
			PartnerName:   partner.PartnerName,
			LinkToPartner: partner.LinkToPartner,
			CreatedAt:     partner.CreatedAt,
			UpdatedAt:     partner.UpdatedAt,
		}
	}

	return &dtos.InstitutionProjectResponse{
		ID:              institutionProject.ID,
		ProjectType:     institutionProject.ProjectType,
		ProjectTitle:    institutionProject.ProjectTitle,
		DateStart:       institutionProject.DateStart,
		DateEnd:         institutionProject.DateEnd,
		Fund:            institutionProject.Fund,
		InstitutionRole: institutionProject.InstitutionRole,
		Coordinator:     institutionProject.Coordinator,
		Partners:        partners,
		CreatedAt:       institutionProject.CreatedAt,
		UpdatedAt:       institutionProject.UpdatedAt,
	}
}
