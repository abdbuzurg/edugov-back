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

type InstitutionSocialUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionSocialRequest) (*dtos.InstitutionSocialResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionSocialRequest) (*dtos.InstitutionSocialResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByInstitutionID(ctx context.Context, institutionID int64) ([]*dtos.InstitutionSocialResponse, error)
}

type institutionSocialUsecase struct {
	institutionSocialRepo repositories.InstitutionSocialRepository
	validator             *validator.Validate
}

func NewInstitutionSocialUsecase(
	institutionSocialRepo repositories.InstitutionSocialRepository,
	validator *validator.Validate,
) InstitutionSocialUsecase {
	return &institutionSocialUsecase{
		institutionSocialRepo: institutionSocialRepo,
		validator:             validator,
	}
}

func (uc *institutionSocialUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionSocialRequest) (*dtos.InstitutionSocialResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution social: %w", err))
	}

	institutionSocial := &domain.InstitutionSocial{
		InstitutionID: req.InstitutionID,
		LinkToSocial:  req.LinkToSocial,
		SocialName:    req.SocialName,
	}

	createdInstitutionSocial, err := uc.institutionSocialRepo.Create(ctx, institutionSocial)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionSocialDomainToResponseDTO(createdInstitutionSocial)
	return resp, nil
}

func (uc *institutionSocialUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionSocialRequest) (*dtos.InstitutionSocialResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution social: %w", err))
	}

	institutionSocial := &domain.InstitutionSocial{
		ID: req.ID,
	}

	if req.LinkToSocial != nil {
		institutionSocial.LinkToSocial = *req.LinkToSocial
	}

	if req.SocialName != nil {
		institutionSocial.SocialName = *req.SocialName
	}

	updatedInstitutionSocial, err := uc.institutionSocialRepo.Update(ctx, institutionSocial)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionSocialDomainToResponseDTO(updatedInstitutionSocial)
	return resp, nil
}

func (uc *institutionSocialUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution social"))
	}

	if err := uc.institutionSocialRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionSocialUsecase) GetByInstitutionID(ctx context.Context, institutionID int64) ([]*dtos.InstitutionSocialResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution social", institutionID))
	}

	institutionSocials, err := uc.institutionSocialRepo.GetByInstitutionID(ctx, institutionID)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionSocialResponse, len(institutionSocials))
	for index, degree := range institutionSocials {
		resp[index] = MapInstitutionSocialDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionSocialDomainToResponseDTO(institutionSocial *domain.InstitutionSocial) *dtos.InstitutionSocialResponse {
	if institutionSocial == nil {
		return nil
	}

	return &dtos.InstitutionSocialResponse{
		ID:           institutionSocial.ID,
		LinkToSocial: institutionSocial.LinkToSocial,
		SocialName:   institutionSocial.SocialName,
		CreatedAt:    institutionSocial.CreatedAt,
		UpdatedAt:    institutionSocial.UpdatedAt,
	}
}
