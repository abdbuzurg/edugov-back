package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeResearchActivityRepository interface {
	//Create - inserts an entry of domain.EmployeeResearchActivity into DB
	Create(ctx context.Context, employeeRA *domain.EmployeeResearchActivity) (*domain.EmployeeResearchActivity, error)

	//Update - modifies an entry of domain.EmployeeResearchActivity in DB
	Update(ctx context.Context, employeeRA *domain.EmployeeResearchActivity) (*domain.EmployeeResearchActivity, error)

	//Delete - removes an entry of domain.EmployeeResearchActivity from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeeResearchActivity from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeResearchActivity, error)

	//GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeeResearchActivity from DB by EmployeeID and specified language code
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeResearchActivity, error)
}
