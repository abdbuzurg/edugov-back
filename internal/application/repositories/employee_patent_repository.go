package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeePatentRepository interface {
  //Create - inserts an entry of domain.EmployeePatent into DB
	Create(ctx context.Context, employeePatent *domain.EmployeePatent) (*domain.EmployeePatent, error)

  //Update - modifies an entry of domain.EmployeePatent in DB
  Update(ctx context.Context, employeePatent *domain.EmployeePatent) (*domain.EmployeePatent, error)

  //Delete - removes an entry of domain.EmployeePatent from DB by ID
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.EmployeePatent from DB By ID
  GetByID(ctx context.Context, id int64) (*domain.EmployeePatent, error)

  //GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeePatent from DB By EmployeeID and specified langauge code
  GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeePatent, error)
}
