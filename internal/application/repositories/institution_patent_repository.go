package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionPatentRepository interface {
	//Create - inserts an entry of domain.InstitutionPatent into DB
	Create(ctx context.Context, institutionPatent *domain.InstitutionPatent) (*domain.InstitutionPatent, error)

	//Update - modifies an entry of domain.InstitutionPatent in DB
	Update(ctx context.Context, institutionPatent *domain.InstitutionPatent) (*domain.InstitutionPatent, error)

	//Delete - remove an entry of domain.InstitutionPatent from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionPatent from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionPatent, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionPatent from DB by InstitutionID with specified language code
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionPatent, error)
}
