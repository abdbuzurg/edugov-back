package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionRepository interface {
	//Create - inserts an entry of domain.Institution into DB.
	Create(ctx context.Context, institution *domain.Institution) (*domain.Institution, error)

	//Update - modifies an entry of domain.Institution in DB.
	Update(ctx context.Context, institution *domain.Institution) (*domain.Institution, error)

	//Delete - removes an entry of domain.Institution from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.Institution and its details(*domain.InstitutionDetails) by ID and specified langauge code.
	GetByID(ctx context.Context, id int64, langCode string) (*domain.Institution, error)

	GetAllInstitutions(ctx context.Context) ([]*domain.Institution, error)
}
