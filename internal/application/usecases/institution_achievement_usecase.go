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

type InstitutionAchievementUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionAchievementRequest) (*dtos.InstitutionAchievementResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionAchievementRequest) (*dtos.InstitutionAchievementResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.InstitutionAchievementResponse, error)
}

type institutionAchievementUsecase struct {
	institutionAchievementRepo repositories.InstitutionAchievementRepository
	validator                  *validator.Validate
}

func NewInstitutionAchievementUsecase(
	institutionAchievementRepo repositories.InstitutionAchievementRepository,
	validator *validator.Validate,
) InstitutionAchievementUsecase {
	return &institutionAchievementUsecase{
		institutionAchievementRepo: institutionAchievementRepo,
		validator:                  validator,
	}
}

func (uc *institutionAchievementUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionAchievementRequest) (*dtos.InstitutionAchievementResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institution achievement: %w", err))
	}

	institutionAchievement := &domain.InstitutionAchievement{
		InstitutionID:    req.InstitutionID,
		LanguageCode:     req.LanguageCode,
		AchievementType:  req.AchievementType,
		AchievementTitle: req.AchievementTitle,
		DateReceived:     req.DateReceived,
		GivenBy:          req.GivenBy,
		LinkToFile:       req.LinkToFile,
		Description:      req.Description,
	}

	createdInstitutionAchievement, err := uc.institutionAchievementRepo.Create(ctx, institutionAchievement)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionAchievementDomainToResponseDTO(createdInstitutionAchievement)
	return resp, nil
}

func (uc *institutionAchievementUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionAchievementRequest) (*dtos.InstitutionAchievementResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institution achievement: %w", err))
	}

	institutionAchievement := &domain.InstitutionAchievement{
		ID: req.ID,
	}

	if req.AchievementType != nil {
		institutionAchievement.AchievementType = *req.AchievementType
	}

	if req.AchievementType != nil {
		institutionAchievement.AchievementType = *req.AchievementType
	}

	if req.DateReceived != nil {
		institutionAchievement.DateReceived = *req.DateReceived
	}

	if req.GivenBy != nil {
		institutionAchievement.GivenBy = *req.GivenBy
	}

	if req.Description != nil {
		institutionAchievement.Description = *req.Description
	}

	if req.LinkToFile != nil {
		institutionAchievement.LinkToFile = *req.LinkToFile
	}

	updatedInstitutionAchievement, err := uc.institutionAchievementRepo.Update(ctx, institutionAchievement)
	if err != nil {
		return nil, err
	}

	resp := MapInstitutionAchievementDomainToResponseDTO(updatedInstitutionAchievement)
	return resp, nil
}

func (uc *institutionAchievementUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institution achievement"))
	}

	if err := uc.institutionAchievementRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionAchievementUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*dtos.InstitutionAchievementResponse, error) {
	if institutionID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - InstitutionID(%d) to retrive institution achievement", institutionID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive institution achievement", langCode))
	}

	institutionAchievements, err := uc.institutionAchievementRepo.GetByInstitutionIDAndLanguageCode(ctx, institutionID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.InstitutionAchievementResponse, len(institutionAchievements))
	for index, degree := range institutionAchievements {
		resp[index] = MapInstitutionAchievementDomainToResponseDTO(degree)
	}

	return resp, nil
}

func MapInstitutionAchievementDomainToResponseDTO(institutionAchievement *domain.InstitutionAchievement) *dtos.InstitutionAchievementResponse {
	if institutionAchievement == nil {
		return nil
	}

	return &dtos.InstitutionAchievementResponse{
		ID:               institutionAchievement.ID,
		AchievementType:  institutionAchievement.AchievementType,
		AchievementTitle: institutionAchievement.AchievementType,
		DateReceived:     institutionAchievement.DateReceived,
		GivenBy:          institutionAchievement.GivenBy,
		Description:      institutionAchievement.Description,
		LinkToFile:       institutionAchievement.LinkToFile,
		CreatedAt:        institutionAchievement.CreatedAt,
		UpdatedAt:        institutionAchievement.UpdatedAt,
	}
}
