package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeWorkExperienceRepository interface {
	//Create - inserts an entry of domain.EmployeeWorkExperience into DB
	Create(ctx context.Context, employeeWorkExperience *domain.EmployeeWorkExperience) (*domain.EmployeeWorkExperience, error)

	//Update - modifies an entry of domain.EmployeeWorkExperience in DB
	Update(ctx context.Context, employeeWorkExperience *domain.EmployeeWorkExperience) (*domain.EmployeeWorkExperience, error)

	//Delete - removes an entry of domain.EmployeeWorkExperience from DB
	Delete(ctx context.Context, id int64) error

	//GetByID - retrives an entry of domain.EmployeeWorkExperience from DB by ID
	GetByID(ctx context.Context, id int64) (*domain.EmployeeWorkExperience, error)

	//GetByEmploployeeIDAndLanguageCode - retrives an entry of domain.EmployeeWorkExperience from DB by EmployeeID and specified language code
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeWorkExperience, error)

	ListUniqueOngoingWorkplaces(ctx context.Context, langCode string) ([]string, error)
}
