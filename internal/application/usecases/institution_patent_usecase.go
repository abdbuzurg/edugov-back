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

type InstitutionPatentUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionPatentRequest) (*dtos.InstitutionPatentResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionPatentRequest) (*dtos.InstitutionPatentResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionPatentResponse, error)
}

type institutionPatentUsecase struct {
	institutionPatentRepo repositories.InstitutionPatentRepository
	validator             *validator.Validate
}

func NewInstitutionPatentUsecase(
	institutionPatentRepo repositories.InstitutionPatentRepository,
	validator *validator.Validate,
) InstitutionPatentUsecase {
	return &institutionPatentUsecase{
		institutionPatentRepo: institutionPatentRepo,
		validator:             validator,
	}
}

func (uc *institutionPatentUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionPatentRequest) (*dtos.InstitutionPatentResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution patent: %w", err))
	}

	institutionPatent := &domain.InstitutionPatent{
		InstitutionID:    req.InstitutionID,
		LanguageCode:     req.LanguageCode,
		PatentTitle:      req.PatentTitle,
		Discipline:       req.Discipline,
		ImplementedIn:    req.ImplementedIn,
		LinkToPatentFile: req.LinkToPartnerFile,
		Description:      req.Description,
	}

	createdInstitutionPatent, err := uc.institutionPatentRepo.Create(ctx, institutionPatent)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionPatentDomainToResponseDTO(createdInstitutionPatent)
	return resp, nil
}

func (uc *institutionPatentUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionPatentRequest) (*dtos.InstitutionPatentResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution patent: %w", err))
	}

	institutionPatent := &domain.InstitutionPatent{
		ID: req.ID,
	}

	if req.PatentTitle != nil {
		institutionPatent.PatentTitle = *req.PatentTitle
	}

	if req.Discipline != nil {
		institutionPatent.Discipline = *req.Discipline
	}

	if req.Description != nil {
		institutionPatent.Description = *req.Description
	}

	if req.ImplementedIn != nil {
		institutionPatent.ImplementedIn = *req.ImplementedIn
	}

	if req.LinkToPartnerFile != nil {
		institutionPatent.LinkToPatentFile = *req.LinkToPartnerFile
	}

	updatedInstitutionPatent, err := uc.institutionPatentRepo.Update(ctx, institutionPatent)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionPatentDomainToResponseDTO(updatedInstitutionPatent)
	return resp, nil
}

func (uc *institutionPatentUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution patent"))
	}

	if err := uc.institutionPatentRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionPatentUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionPatentResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution patent", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution patent", langCode))
	}

	institutionPatents, err := uc.institutionPatentRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionPatentResponse, len(institutionPatents))
	for index, degree := range institutionPatents {
		resp[index] = MapInstitutionPatentDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionPatentDomainToResponseDTO(institutionPatent *domain.InstitutionPatent) *dtos.InstitutionPatentResponse {
	if institutionPatent == nil {
		return nil
	}

	return &dtos.InstitutionPatentResponse{
		ID:                institutionPatent.ID,
		PatentTitle:       institutionPatent.PatentTitle,
		Discipline:        institutionPatent.Discipline,
		ImplementedIn:     institutionPatent.ImplementedIn,
		LinkToPartnerFile: institutionPatent.LinkToPatentFile,
		Description:       institutionPatent.Description,
		CreatedAt:         institutionPatent.CreatedAt,
		UpdatedAt:         institutionPatent.UpdatedAt,
	}
}
