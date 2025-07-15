package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionResearchSupportInfrastructureRepository interface {
	//Create - inserts an entry of domain.InstitutionResearchSupportInfrastructure into DB
	Create(ctx context.Context, institutionRSI *domain.InstitutionResearchSupportInfrastructure) (*domain.InstitutionResearchSupportInfrastructure, error)

	//Update - modifies an entry of domain.InstitutionResearchSupportInfrastructure in DB
	Update(ctx context.Context, institutionRSI *domain.InstitutionResearchSupportInfrastructure) (*domain.InstitutionResearchSupportInfrastructure, error)

	//Delete - removes an entry of domain.InstitutionResearchSupportInfrastructure from DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionResearchSupportInfrastructure from DB
	GetByID(ctx context.Context, id int64) (*domain.InstitutionResearchSupportInfrastructure, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionResearchSupportInfrastructure from DB by InstitutionID with specified language code
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionResearchSupportInfrastructure, error)
}
