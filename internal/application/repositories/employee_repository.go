package repositories

import (
	"backend/internal/application/dtos"
	"backend/internal/domain"
	"context"
)

type GetPersonnelPaginatedQueryResult struct {
	UniqueID              string `json:"unique_id"`
	EmployeeID            int64  `json:"employee_id"`
	Surname               string `json:"surname"`
	Name                  string `json:"name"`
	Middlename            string `json:"middlename"`
	Currentworkplace      string `json:"currentworkplace"`
	Highestacademicdegree string `json:"highestacademicdegree"`
	Speciality            string `json:"speciality"`
}

type EmployeeRepository interface {
	//Create - inserts employee into the database.
	//it takes domain.Employee with details included(*domain.Details) and return it with ID and timestamps
	Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error)

	//GetByID - retrives a single employee by their ID with the specified language.
	//Note - this retrival will only give the main employee fields and the details information. Other fields will be omitted.
	GetByID(ctx context.Context, id int64) (*domain.Employee, error)

	//Delete - removes an employee from database by ID
	//Note - all related employee data in other tables will be deleted, becuase of on cascade delete constraint
	Delete(ctx context.Context, id int64) error

	//GetByUniqueID - retrives a single employee by their uniqueIdentifer.
	GetByUniqueID(ctx context.Context, uniqueID string) (*domain.Employee, error)

	GetByUserID(ctx context.Context, userID int64) (*domain.Employee, error)

	//GetPersonnelIDsPaginated - retrives employee IDs from db that satisfy the filter parameters in paginated form
	GetPersonnelIDsPaginated(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) ([]*GetPersonnelPaginatedQueryResult, error)

	//CountPersonnel - count total number of personnel (by unique employee_id) from db that satisfy the filter paramenter
	CountPersonnel(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) (int64, error)
}
