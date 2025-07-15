package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionMainResearchDirectionRepository interface {
	//Create - inserts an entry of domain.InstitutionMainResearchDirection into DB
	Create(ctx context.Context, institutionMRD *domain.InstitutionMainResearchDirection) (*domain.InstitutionMainResearchDirection, error)

	//Update - modifies an entry of domain.InstitutionMainResearchDirection in DB
	Update(ctx context.Context, institutionMRD *domain.InstitutionMainResearchDirection) (*domain.InstitutionMainResearchDirection, error)

	//Delete - remove an entry of domain.InstitutionMainResearchDirection from DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionMainResearchDirection from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionMainResearchDirection, error)

	//GetByInstituionIDAndLanguageCode - retrives an entry of domain.InstitutionMainResearchDirection from DB by institutionID and specified langauge code.
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionMainResearchDirection, error)
}
