package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionConferenceRepository interface {
	//Create - inserts an entry of domain.InstitutionConference into DB
	Create(ctx context.Context, institutionConference *domain.InstitutionConference) (*domain.InstitutionConference, error)

  //Update - modifies an entry of domain.InstitutionConference in DB
  Update(ctx context.Context, institutionConference *domain.InstitutionConference) (*domain.InstitutionConference, error)

  //Delete - removes an entry of domain.InstitutionConference from DB
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.InstitutionConference from DB by ID
  GetByID(ctx context.Context, id int64) (*domain.InstitutionConference, error)

  //GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionConference from DB by Institution ID and specified language code
  GetByInstitutionIDAndLanguageCode(ctx context.Context, id int64, langCode string) ([]*domain.InstitutionConference, error)
}
