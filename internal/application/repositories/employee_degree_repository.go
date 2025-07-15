package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeDegreeRepository interface {
  //Create - insterts employee degree data into DB
	Create(ctx context.Context, employeeDegree *domain.EmployeeDegree) (*domain.EmployeeDegree, error)

  //Update - modifies existing employee in DB
  Update(ctx context.Context, employeeDegree *domain.EmployeeDegree) (*domain.EmployeeDegree, error)

  //Delete - deletes employee degree data from DB by ID
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives a single employee degree data by ID 
  GetByID(ctx context.Context, id int64) (*domain.EmployeeDegree, error)

  //GetByEmployeeIDAndLanguageCode - retrives a single employee degree data by EmployeeID and specified language code
  GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeDegree, error)
}
