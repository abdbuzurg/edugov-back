package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionSocialRepository interface {
	//Create - inserts an entry domain.InstitutionSocial into DB
	Create(ctx context.Context, institutionSocial *domain.InstitutionSocial) (*domain.InstitutionSocial, error)

  //Update - modifies an entry of domain.InstituitonSocial in DB
  Update(ctx context.Context, institutionSocial *domain.InstitutionSocial) (*domain.InstitutionSocial, error)
  
  //Delete - removes an entry of domain.InstitutionSocial from DB by ID
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.InstitutionSocial from DB by ID 
  GetByID(ctx context.Context, id int64) (*domain.InstitutionSocial, error)

  //GetByInstitutionID - retrives an entry of domain.InstitutionSocial from DB by InstitutionID 
  GetByInstitutionID(ctx context.Context, id int64) ([]*domain.InstitutionSocial, error)
}
