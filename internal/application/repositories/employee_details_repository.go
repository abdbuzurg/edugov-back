package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeDetailsRepository interface {
	//Create - insterts employee degree data into DB
	Create(ctx context.Context, employeeDegree *domain.EmployeeDetails) (*domain.EmployeeDetails, error)

	//Update - modifies existing employee in DB
	Update(ctx context.Context, employeeDegree *domain.EmployeeDetails) (*domain.EmployeeDetails, error)

	//Delete - deletes employee degree data from DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives a single employee degree data by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeDetails, error)

	//GetByEmployeeIDAndLanguageCode - retrives a single employee degree data by EmployeeID and specified language code
	GetByEmployeeID(ctx context.Context, employeeID int64) ([]*domain.EmployeeDetails, error)
}
