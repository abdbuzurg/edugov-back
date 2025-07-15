package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeSocialRepository interface {
	//Create - inserts an entry of domain.EmployeeSocials into DB.
	Create(ctx context.Context, employeeSocials *domain.EmployeeSocial) (*domain.EmployeeSocial, error)

	//Update - modifies an entry of domain.EmployeeSocials in DB
	Update(ctx context.Context, employeeSocials *domain.EmployeeSocial) (*domain.EmployeeSocial, error)

	//Delete - removes an entry of domain.EmployeeSocials from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeeSocials from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeSocial, error)

	//GetByEmployeeID - retrives an entry of domain.EmployeeSocials from DB by EmployeeID
	GetByEmployeeID(ctx context.Context, id int64) ([]*domain.EmployeeSocial, error)
}
