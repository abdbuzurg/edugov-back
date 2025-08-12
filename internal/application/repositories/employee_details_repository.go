package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeDetailsRepository interface {
	//Create - insterts employee Details data into DB
	Create(ctx context.Context, employeeDetails *domain.EmployeeDetails) (*domain.EmployeeDetails, error)

	//Update - modifies existing employee in DB
	Update(ctx context.Context, employeeDetails *domain.EmployeeDetails) (*domain.EmployeeDetails, error)

	//Delete - deletes employee Details data from DB by ID
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives a single employee Details data by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeDetails, error)

	//GetByEmployeeIDAndLanguageCode - retrives a employee details data by EmployeeID
	GetByEmployeeID(ctx context.Context, employeeID int64) ([]*domain.EmployeeDetails, error)

	//GetCurrentDetailsByEmployeeIDAndLanguageCode - retrives a single employee details by employeeID and language code where is_new_employee_details = true (current employee information)
	GetCurrentDetailsByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) (*domain.EmployeeDetails, error)
}
