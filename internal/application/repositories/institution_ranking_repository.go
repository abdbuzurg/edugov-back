package repositories

import (
	"backend/internal/domain"
	"context"
)

type InstitutionRankingRepository interface {
	//Create - inserts an entry of domain.InstitutionRanking into DB
	Create(ctx context.Context, institutionRankings *domain.InstitutionRanking) (*domain.InstitutionRanking, error)

	//Update - modifies an netry of domain.InstitutionRanking in DB
	Update(ctx context.Context, institutionRankings *domain.InstitutionRanking) (*domain.InstitutionRanking, error)

	//Delete - removes an entry of domain.InstitutionRanking from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.InstitutionRanking from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.InstitutionRanking, error)

	//GetByInstitutionIDAndLanguageCode - retrives an entry of domain.InstitutionRanking from DB by InstitutionID with specific language
	GetByInstitutionIDAndLanguageCode(ctx context.Context, institutionID int64, langCode string) ([]*domain.InstitutionRanking, error)
}
