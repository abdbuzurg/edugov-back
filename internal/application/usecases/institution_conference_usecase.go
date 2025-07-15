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

type InstitutionConferenceUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionConferenceRequest) (*dtos.InstitutionConferenceResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionConferenceRequest) (*dtos.InstitutionConferenceResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionConferenceResponse, error)
}

type institutionConferenceUsecase struct {
	institutionConferenceRepo repositories.InstitutionConferenceRepository
	validator                 *validator.Validate
}

func NewInstitutionConferenceUsecase(
	institutionConferenceRepo repositories.InstitutionConferenceRepository,
	validator *validator.Validate,
) InstitutionConferenceUsecase {
	return &institutionConferenceUsecase{
		institutionConferenceRepo: institutionConferenceRepo,
		validator:                 validator,
	}
}

func (uc *institutionConferenceUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionConferenceRequest) (*dtos.InstitutionConferenceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution conference: %w", err))
	}

	institutionConference := &domain.InstitutionConference{
		InstitutionID:    req.InstitutionID,
		LanguageCode:     req.LanguageCode,
		ConferenceTitle:  req.ConferenceTitle,
		Link:             req.Link,
		LinkToRINC:       req.LinkToRINC,
		DateOfConference: req.DateOfConference,
	}

	createdInstitutionConference, err := uc.institutionConferenceRepo.Create(ctx, institutionConference)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionConferenceDomainToResponseDTO(createdInstitutionConference)
	return resp, nil
}

func (uc *institutionConferenceUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionConferenceRequest) (*dtos.InstitutionConferenceResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution conference: %w", err))
	}

	institutionConference := &domain.InstitutionConference{
		ID: req.ID,
	}

	if req.ConferenceTitle != nil {
		institutionConference.ConferenceTitle = *req.ConferenceTitle
	}

	if req.Link != nil {
		institutionConference.Link = *req.Link
	}

	if req.LinkToRINC != nil {
		institutionConference.LinkToRINC = *req.LinkToRINC
	}

	if req.DateOfConference != nil {
		institutionConference.DateOfConference = *req.DateOfConference
	}

	updatedInstitutionConference, err := uc.institutionConferenceRepo.Update(ctx, institutionConference)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionConferenceDomainToResponseDTO(updatedInstitutionConference)
	return resp, nil
}

func (uc *institutionConferenceUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution conference"))
	}

	if err := uc.institutionConferenceRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionConferenceUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionConferenceResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution conference", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution conference", langCode))
	}

	institutionConferences, err := uc.institutionConferenceRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionConferenceResponse, len(institutionConferences))
	for index, degree := range institutionConferences {
		resp[index] = MapInstitutionConferenceDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionConferenceDomainToResponseDTO(institutionConference *domain.InstitutionConference) *dtos.InstitutionConferenceResponse {
	if institutionConference == nil {
		return nil
	}

	return &dtos.InstitutionConferenceResponse{
		ID:               institutionConference.ID,
		ConferenceTitle:  institutionConference.ConferenceTitle,
		Link:             institutionConference.Link,
		LinkToRINC:       institutionConference.LinkToRINC,
		DateOfConference: institutionConference.DateOfConference,
		CreatedAt:        institutionConference.CreatedAt,
		UpdatedAt:        institutionConference.UpdatedAt,
	}
}
