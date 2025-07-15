package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionMagazineRepository interface {
	//Create - inserts an entry of domain.InstitutionMagazine into DB
	Create(ctx context.Context, institutionMagazine *domain.InstitutionMagazine) (*domain.InstitutionMagazine, error)

	//Update - modifies an entry of domain.InstitutionMagazine in DB
	Update(ctx context.Context, institutionMagazine *domain.InstitutionMagazine) (*domain.InstitutionMagazine, error)

	//Delete - removes an entry of domain.InstitutionMagazine from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionMagazine from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionMagazine, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionMagazine from DB by ID and specified langauge code.
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionMagazine, error)
}
