package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/mappers"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type InstitutionUsecase interface {
	Create(ctx context.Context, req *dtos.CreateInstitutionRequest) (*dtos.InstitutionResponse, error)
	Update(ctx context.Context, req *dtos.UpdateInstitutionRequest) (*dtos.InstitutionResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, employeeID int64, langCode string) (*dtos.InstitutionResponse, error)
	GetAllInstitutions(ctx context.Context) ([]*dtos.AllInstitutionResponse, error)
}

type institutionUsecase struct {
	institutionRepo repositories.InstitutionRepository
	validator       *validator.Validate
	store           *postgres.Store
}

func NewInstitutionUsecase(
	institutionRepo repositories.InstitutionRepository,
	validator *validator.Validate,
	store *postgres.Store,
) InstitutionUsecase {
	return &institutionUsecase{
		institutionRepo: institutionRepo,
		validator:       validator,
		store:           store,
	}
}

func (uc *institutionUsecase) Create(ctx context.Context, req *dtos.CreateInstitutionRequest) (*dtos.InstitutionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create institutin: %w", err))
	}

	institution := &domain.Institution{
		YearOfEstablishment: req.YearOfEstablishment,
		Email:               req.Email,
		Fax:                 req.Fax,
		OfficialWebsite:     req.OfficialWebsite,
		PhoneNumber:         req.PhoneNumber,
		MailIndex:           req.MailIndex,
	}

	createdInstitution, err := uc.institutionRepo.Create(ctx, institution)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapInstitutionDomainToResponseDTO(createdInstitution)
	return resp, nil
}

func (uc *institutionUsecase) Update(ctx context.Context, req *dtos.UpdateInstitutionRequest) (*dtos.InstitutionResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update institutin: %w", err))
	}

	institution := &domain.Institution{
		ID: req.ID,
	}

	if req.YearOfEstablishment != nil {
		institution.YearOfEstablishment = *req.YearOfEstablishment
	}

	if req.Email != nil {
		institution.Email = *req.Email
	}

	if req.Fax != nil {
		institution.Fax = *req.Fax
	}

	if req.OfficialWebsite != nil {
		institution.OfficialWebsite = *req.OfficialWebsite
	}

	if req.PhoneNumber != nil {
		institution.PhoneNumber = *req.PhoneNumber
	}

	if req.PhoneNumber != nil {
		institution.MailIndex = *req.MailIndex
	}

	updatedInstitution, err := uc.institutionRepo.Update(ctx, institution)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapInstitutionDomainToResponseDTO(updatedInstitution)
	return resp, nil
}

func (uc *institutionUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete institutin"))
	}

	if err := uc.institutionRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *institutionUsecase) GetByID(ctx context.Context, id int64, langCode string) (*dtos.InstitutionResponse, error) {
	institution, err := uc.institutionRepo.GetByID(ctx, id, langCode)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapInstitutionDomainToResponseDTO(institution)
	return resp, nil
}

func (uc *institutionUsecase) GetAllInstitutions(ctx context.Context) ([]*dtos.AllInstitutionResponse, error) {
	var result []*dtos.AllInstitutionResponse
	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txInstitutionRepo := postgres.NewPgInstitutionRepositoryWithQueries(q)
		txInstitutionDetailsRepo := postgres.NewPGInstitutionDetailsRepositoryWithQuery(q)

		institutionsResult, err := txInstitutionRepo.GetAllInstitutions(ctx)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}

		result = make([]*dtos.AllInstitutionResponse, len(institutionsResult))

		langCode := middleware.GetLanguageFromContext(ctx)
		for index, institution := range institutionsResult {
			institutionDetails, err := txInstitutionDetailsRepo.GetByInstitutionIDAndLanguageCode(ctx, institution.ID, langCode)
			if err != nil {
				if custom_errors.IsNotFound(err) {
					continue
				}

				return err
			}

			result[index] = &dtos.AllInstitutionResponse{
				InstitutitonTitleShort: institutionDetails.InstitutionTitleShort,
				InstitutitonTitleLong:  institutionDetails.InstitutionTitleLong,
				MailIndex:              institution.MailIndex,
				City:                   institutionDetails.City,
				Address:                institutionDetails.LegalAddress,
				OfficialWebsite:        institution.OfficialWebsite,
				Email:                  institution.Email,
				PhoneNumber:            institution.PhoneNumber,
			}
		}

		return nil
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("transaction failed to retrive information about all institutions: %w", err))
	}

	return result, nil
}
