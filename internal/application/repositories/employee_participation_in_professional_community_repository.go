package repositories

import (
	"backend/internal/domain"
	"context"
)

type ProfessionalCommunitiesTitlesLookUp struct {
	ID                         int64  `json:"ID"`
	ProfessionalCommunityTitle string `json:"professionalCommunityTitle"`
}

type EmployeeParticipationInProfessionalCommunityRepository interface {
	//Create - inserts an entry of domain.EmployeeParticipationInProfessionalCommunity into DB
	Create(ctx context.Context, employeePIPC *domain.EmployeeParticipationInProfessionalCommunity) (*domain.EmployeeParticipationInProfessionalCommunity, error)

	//Update - modifies an entry of domain.EmployeeParticipationInProfessionalCommunity in DB
	Update(ctx context.Context, employeePIPC *domain.EmployeeParticipationInProfessionalCommunity) (*domain.EmployeeParticipationInProfessionalCommunity, error)

	//Delete - removes an entry of domain.EmployeeParticipationInProfessionalCommunity in DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeeParticipationInProfessionalCommunity from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeParticipationInProfessionalCommunity, error)

	//GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeeParticipationInProfessionalCommunity from DB by EmployeeID and specified language code
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeParticipationInProfessionalCommunity, error)
}
