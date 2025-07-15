package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionAchievementRepository interface {
	//Create - inserts an entry of domain.InstitutionAchievement into DB
	Create(ctx context.Context, institutionAchievement *domain.InstitutionAchievement) (*domain.InstitutionAchievement, error)

  //Update - modifies an entry of domain.InstitutionAchievement in DB
  Update(ctx context.Context, institutionAchievement *domain.InstitutionAchievement) (*domain.InstitutionAchievement, error)

  //Delete - remove an entry of domain.InstitutionAchievement from DB
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.InstitutionAchievement from DB by ID 
  GetByID(ctx context.Context, id int64) (*domain.InstitutionAchievement, error)

  //GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionAchievement from DB by ID and specified language code 
  GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionAchievement, error)
}
