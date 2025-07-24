package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeRepository interface {
	//Create - inserts employee into the database.
	//it takes domain.Employee with details included(*domain.Details) and return it with ID and timestamps
	Create(ctx context.Context, employee *domain.Employee) (*domain.Employee, error)

  //UpdateProfilePicture - update the filed profile picture into database for table employee
  //takes employee ID and profilePicture, and return only error 
  UpdateProfilePicture(ctx context.Context, uid string, profilePicture string) error

	//GetByID - retrives a single employee by their ID with the specified language.
	//Note - this retrival will only give the main employee fields and the details information. Other fields will be omitted.
	GetByID(ctx context.Context, id int64) (*domain.Employee, error)

  GetProfilePictureFileNameByUniqueID(ctx context.Context, uid string) (*string, error)

	//Delete - removes an employee from database by ID
	//Note - all related employee data in other tables will be deleted, becuase of on cascade delete constraint
	Delete(ctx context.Context, id int64) error

	//GetByUniqueID - retrives a single employee by their uniqueIdentifer.
	GetByUniqueID(ctx context.Context, uniqueID string) (*domain.Employee, error)
}
