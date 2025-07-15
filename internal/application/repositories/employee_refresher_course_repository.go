package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeRefresherCourseRepository interface {
	//Create - inserts an entry of domain.EmployeeRefresherCourse into DB
	Create(ctx context.Context, employeeRC *domain.EmployeeRefresherCourse) (*domain.EmployeeRefresherCourse, error)

  //Update - modifies an entry of domain.EmployeeRefresherCourse in DB
  Update(ctx context.Context, employeeRC *domain.EmployeeRefresherCourse) (*domain.EmployeeRefresherCourse, error)

  //Delete - removes an entry of domain.EmployeeRefresherCourse from DB
  Delete(ctx context.Context, id int64) error

  //GetByID - retrives an entry of domain.EmployeeRefresherCourseReposiotry from DB by ID
  GetByID(ctx context.Context, id int64) (*domain.EmployeeRefresherCourse, error)

  //GetByEmployeeIDAndLanguageCode - retrives an entry of domain.EmployeeRefresherCourseReposiotry from DB by EmployeeID and specified langauge code
  GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeRefresherCourse, error)
} 
