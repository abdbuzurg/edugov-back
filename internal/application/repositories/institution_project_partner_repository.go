package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionProjectPartnerRepository interface {
	//Create - insert an entry of domain.InstitutionProjectPartner into DB
  //Note - the field project_
	Create(ctx context.Context, institutionPP *domain.InstitutionProjectPartner) (*domain.InstitutionProjectPartner, error)

	//Update - modifies an entry of domain.InstitutionProjectPartner in DB
	Update(ctx context.Context, institutionPP *domain.InstitutionProjectPartner) (*domain.InstitutionProjectPartner, error)

	//Delete - removes an entry of domain.InstitutionProjectPartner from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionProjectPartner from DB by ID 
	GetByID(ctx context.Context, id int64) (*domain.InstitutionProjectPartner, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionProjectPartner from DB by InstitutionProjectID and specified language code
	GetByInstitutionProjectIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionProjectPartner, error)
}
