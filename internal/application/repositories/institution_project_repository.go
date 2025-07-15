package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionProjectRepository interface {
	//Create - inserts an entry of domain.InstitutionProject into DB
	Create(ctx context.Context, institutionProject *domain.InstitutionProject) (*domain.InstitutionProject, error)

  //Update - modifies an entry of domain.InstitutionProject in DB
  Update(ctx context.Context, institutionProject *domain.InstitutionProject) (*domain.InstitutionProject, error)

  //Delete - removes an entry of domain.InsititutionProject from DB by ID
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.InstitutionProject from DB by ID  
  GetByID(ctx context.Context, id int64) (*domain.InstitutionProject, error)

  //GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionProject from DB by InstitutionID and specified language code 
  GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionProject, error)
}
