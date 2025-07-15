package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionLicenceRepository interface {
	//Create - inserts an entry of domain.InstitutionLicence into DB.
	Create(ctx context.Context, institutionLicence *domain.InstitutionLicence) (*domain.InstitutionLicence, error)

  //Update - modifies an entry of domain.InstitutionLicence in DB
  Update(ctx context.Context, institutionLicence *domain.InstitutionLicence) (*domain.InstitutionLicence, error)

  //Delete - removes an entry of domain.InstitutionLicence from DB
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.InstitutionLicence from DB by ID 
  GetByID(ctx context.Context, id int64) (*domain.InstitutionLicence, error)

  //GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionLicence from DB by InstititonID and specified language code.
  GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionLicence, error)
}
