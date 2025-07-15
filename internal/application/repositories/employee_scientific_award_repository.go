package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeScientificAwardRepository interface {
	//Create - inserts an entry of domain.EmployeeScientificAward into DB
	Create(ctx context.Context, employeeSA *domain.EmployeeScientificAward) (*domain.EmployeeScientificAward, error)

	//Update - modifies an entry of domain.EmployeeScientificAward in DB
	Update(ctx context.Context, employeeSA *domain.EmployeeScientificAward) (*domain.EmployeeScientificAward, error)

	//Delete - removes an entry of domain.EmployeeScientificAward from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeeScientificAward from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeScientificAward, error)

	//GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeeScientificAward from DB by EmployeeID and specified language code
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeScientificAward, error)
}
