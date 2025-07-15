package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeParticipationInEventRepository interface {
	//Create - insterts  domain.EmployeeParticipationInEvent data into DB.
	Create(ctx context.Context, employeePIE *domain.EmployeeParticipationInEvent) (*domain.EmployeeParticipationInEvent, error)
  
  //Update - modifies an entry of domain.EmployeeParticipationInEvent in DB.
  Update(ctx context.Context, employeePIE *domain.EmployeeParticipationInEvent) (*domain.EmployeeParticipationInEvent, error)

  //Delete - removes an entry of domain.EmployeeParticipationInEvent by ID in DB.
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives a single domain.EmployeeParticipationInEvent from DB by ID 
  GetByID(ctx context.Context, id int64) (*domain.EmployeeParticipationInEvent, error)

  //GetByEmployeeIDAndLanguageCode - retrives a single domain.EmployeeParticipationInEvent from DB by EmployeeID and specified language code
  GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeParticipationInEvent, error)
}
