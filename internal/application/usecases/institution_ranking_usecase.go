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

type InstitutionRankingUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionRankingRequest) (*dtos.InstitutionRankingResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionRankingRequest) (*dtos.InstitutionRankingResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionRankingResponse, error)
}

type institutionRankingUsecase struct {
	institutionRankingRepo repositories.InstitutionRankingRepository
	validator              *validator.Validate
}

func NewInstitutionRankingUsecase(
	institutionRankingRepo repositories.InstitutionRankingRepository,
	validator *validator.Validate,
) InstitutionRankingUsecase {
	return &institutionRankingUsecase{
		institutionRankingRepo: institutionRankingRepo,
		validator:              validator,
	}
}

func (uc *institutionRankingUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionRankingRequest) (*dtos.InstitutionRankingResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution ranking: %w", err))
	}

	institutionRanking := &domain.InstitutionRanking{
		InstitutionID:       req.InstitutionID,
		LanguageCode:        req.LanguageCode,
		RankingTitle:        req.RankingTitle,
		RankingType:         req.RankingType,
		DateReceived:        req.DateReceived,
		RankingAgency:       req.RankingAgency,
		Description:         req.Description,
		LinkToRankingFile:   req.LinkToRankingFile,
		LinkToRankingAgency: req.LinkToRankingAgencyFile,
	}

	createdInstitutionRanking, err := uc.institutionRankingRepo.Create(ctx, institutionRanking)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionRankingDomainToResponseDTO(createdInstitutionRanking)
	return resp, nil
}

func (uc *institutionRankingUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionRankingRequest) (*dtos.InstitutionRankingResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution ranking: %w", err))
	}

	institutionRanking := &domain.InstitutionRanking{
		ID: req.ID,
	}

	if req.RankingTitle != nil {
		institutionRanking.RankingTitle = *req.RankingTitle
	}

	if req.RankingType != nil {
		institutionRanking.RankingType = *req.RankingType
	}

	if req.DateReceived != nil {
		institutionRanking.DateReceived = *req.DateReceived
	}

	if req.RankingAgency != nil {
		institutionRanking.RankingAgency = *req.RankingAgency
	}

	if req.Description != nil {
		institutionRanking.Description = *req.Description
	}

	if req.LinkToRankingFile != nil {
		institutionRanking.LinkToRankingFile = *req.LinkToRankingFile
	}

	if req.LinkToRankingAgencyFile != nil {
		institutionRanking.LinkToRankingAgency = *req.LinkToRankingAgencyFile
	}

	updatedInstitutionRanking, err := uc.institutionRankingRepo.Update(ctx, institutionRanking)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionRankingDomainToResponseDTO(updatedInstitutionRanking)
	return resp, nil
}

func (uc *institutionRankingUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution ranking"))
	}

	if err := uc.institutionRankingRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionRankingUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionRankingResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution ranking", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution ranking", langCode))
	}

	institutionRankings, err := uc.institutionRankingRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionRankingResponse, len(institutionRankings))
	for index, degree := range institutionRankings {
		resp[index] = MapInstitutionRankingDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionRankingDomainToResponseDTO(institutionRanking *domain.InstitutionRanking) *dtos.InstitutionRankingResponse {
	if institutionRanking == nil {
		return nil
	}

	return &dtos.InstitutionRankingResponse{
		ID:                      institutionRanking.ID,
		RankingTitle:            institutionRanking.RankingTitle,
		RankingType:             institutionRanking.RankingType,
		DateReceived:            institutionRanking.DateReceived,
		RankingAgency:           institutionRanking.RankingAgency,
		Description:             institutionRanking.Description,
		LinkToRankingFile:       institutionRanking.LinkToRankingFile,
		LinkToRankingAgencyFile: institutionRanking.LinkToRankingAgency,
		CreatedAt:               institutionRanking.CreatedAt,
		UpdatedAt:               institutionRanking.UpdatedAt,
	}
}
