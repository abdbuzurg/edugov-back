package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionAccreditationRepository interface {
	//Create - inserts an entry of domain.InstitutionAccreditation into DB
	Create(ctx context.Context, institutionAccreditation *domain.InstitutionAccreditation) (*domain.InstitutionAccreditation, error)

	//Update - modifies an entry of domain.InstitutionAccreditation in DB
	Update(ctx context.Context, institutionAccreditation *domain.InstitutionAccreditation) (*domain.InstitutionAccreditation, error)

	//Delete - removes an entry of domain.InstitutionAccreditation from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionAccreditation from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionAccreditation, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domian.InstitutionAccrediation
	//by institutionID and specified language code
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionAccreditation, error)
}
