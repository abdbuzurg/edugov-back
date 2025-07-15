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

type InstitutionPartnershipUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionPartnershipRequest) (*dtos.InstitutionPartnershipResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionPartnershipRequest) (*dtos.InstitutionPartnershipResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionPartnershipResponse, error)
}

type institutionPartnershipUsecase struct {
	institutionPartnershipRepo repositories.InstitutionPartnershipRepository
	validator                  *validator.Validate
}

func NewInstitutionPartnershipUsecase(
	institutionPartnershipRepo repositories.InstitutionPartnershipRepository,
	validator *validator.Validate,
) InstitutionPartnershipUsecase {
	return &institutionPartnershipUsecase{
		institutionPartnershipRepo: institutionPartnershipRepo,
		validator:                  validator,
	}
}

func (uc *institutionPartnershipUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionPartnershipRequest) (*dtos.InstitutionPartnershipResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution achievement: %w", err))
	}

	institutionPartnership := &domain.InstitutionPartnership{
		InstitutionID:  req.InstitutionID,
		LanguageCode:   req.LanguageCode,
		PartnerName:    req.PartnerName,
		PartnerType:    req.PartnerType,
		Goal:           req.Goal,
		LinkToPartner:  req.LinkToPartner,
		DateOfContract: req.DateOfContract,
	}

	createdInstitutionPartnership, err := uc.institutionPartnershipRepo.Create(ctx, institutionPartnership)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionPartnershipDomainToResponseDTO(createdInstitutionPartnership)
	return resp, nil
}

func (uc *institutionPartnershipUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionPartnershipRequest) (*dtos.InstitutionPartnershipResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution achievement: %w", err))
	}

	institutionPartnership := &domain.InstitutionPartnership{
		ID: req.ID,
	}

	if req.PartnerName != nil {
		institutionPartnership.PartnerName = *req.PartnerName
	}

	if req.PartnerType != nil {
		institutionPartnership.PartnerType = *req.PartnerType
	}

	if req.Goal != nil {
		institutionPartnership.Goal = *req.Goal
	}

	if req.LinkToPartner != nil {
		institutionPartnership.LinkToPartner = *req.LinkToPartner
	}

	updatedInstitutionPartnership, err := uc.institutionPartnershipRepo.Update(ctx, institutionPartnership)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionPartnershipDomainToResponseDTO(updatedInstitutionPartnership)
	return resp, nil
}

func (uc *institutionPartnershipUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution achievement"))
	}

	if err := uc.institutionPartnershipRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionPartnershipUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionPartnershipResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution achievement", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution achievement", langCode))
	}

	institutionPartnerships, err := uc.institutionPartnershipRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionPartnershipResponse, len(institutionPartnerships))
	for index, degree := range institutionPartnerships {
		resp[index] = MapInstitutionPartnershipDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionPartnershipDomainToResponseDTO(institutionPartnership *domain.InstitutionPartnership) *dtos.InstitutionPartnershipResponse {
	if institutionPartnership == nil {
		return nil
	}

	return &dtos.InstitutionPartnershipResponse{
		ID:             institutionPartnership.ID,
		PartnerType:    institutionPartnership.PartnerType,
		PartnerName:    institutionPartnership.PartnerName,
		Goal:           institutionPartnership.Goal,
		LinkToPartner:  institutionPartnership.PartnerName,
		DateOfContract: institutionPartnership.DateOfContract,
		CreatedAt:      institutionPartnership.CreatedAt,
		UpdatedAt:      institutionPartnership.UpdatedAt,
	}
}
