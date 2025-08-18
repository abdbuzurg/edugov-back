package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionDetailsRepository interface {

	//Create - inserts an entry of domain.Institution into DB.
	Create(ctx context.Context, institutionDetails *domain.InstitutionDetails) (*domain.InstitutionDetails, error)

	//Update - modifies an entry of domain.Institution in DB.
	Update(ctx context.Context, institutionDetails *domain.InstitutionDetails) (*domain.InstitutionDetails, error)

	//Delete - removes an entry of domain.Institution from DB
	Delete(ctx context.Context, id int64) error

	//GetByInstitutionIDAndLanguageCode - retrives entries institution_details(*domain.InstitutionDetails) by ID and specified langauge code.
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) (*domain.InstitutionDetails, error)
}
