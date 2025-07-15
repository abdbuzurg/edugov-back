package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionPartnershipRepository interface {
	//Create - inserts an entry of domain.InstitutionPartnership into DB.
	Create(ctx context.Context, institutionPartnership *domain.InstitutionPartnership) (*domain.InstitutionPartnership, error)

	//Update - modifies an entry of domain.InstitutionPartnership in DB
	Update(ctx context.Context, institutionPartnership *domain.InstitutionPartnership) (*domain.InstitutionPartnership, error)

	//Delete - removes an entry of domain.InstitutionPartnership from DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionPartnership from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionPartnership, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionPartnership from DB by InstituionID and specified language code
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionPartnership, error)
}
