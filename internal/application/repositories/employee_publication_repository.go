package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeePublicationRepository interface {
	//Create - inserts an entry of domain.EmployeePublication into DB
	Create(ctx context.Context, employeePublication *domain.EmployeePublication) (*domain.EmployeePublication, error)

	//Update - modifes an entry of domain.EmployeePublication in DB
	Update(ctx context.Context, employeePublication *domain.EmployeePublication) (*domain.EmployeePublication, error)

	//Delete - removes an entry of domain.EmployeePublication from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeePublication from DB
	GetByID(ctx context.Context, id int64) (*domain.EmployeePublication, error)

	//GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeePublication from DB by EmployeeID and specified language code
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeePublication, error)
}
