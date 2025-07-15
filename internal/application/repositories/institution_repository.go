package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionRepository interface {
	//Create - inserts an entry of domain.Institution into DB.
	//Note - this will also insert details(*domain.InstitutionDetails) into DB.
	Create(ctx context.Context, institution *domain.Institution) (*domain.Institution, error)
  
  //Update - modifies an entry of domain.Institution in DB.
  //Note - this will also modify details information if changed.
  Update(ctx context.Context, institution *domain.Institution) (*domain.Institution, error)

  //Delete - removes an entry of domain.Institution from DB
  //Note - this will automatically remove all related institution tables.
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.Institution and its details(*domain.InstitutionDetails) by ID and specified langauge code.
  GetByID(ctx context.Context, id int64, langCode string) (*domain.Institution, error)
}
