package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/mappers"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type EmployeeRefresherCourseUsecase interface {
	Create(ctx context.Context, req *dtos.CreateEmployeeRefresherCourseRequest) (*dtos.EmployeeRefresherCourseResponse, error)
	Update(ctx context.Context, req *dtos.UpdateEmployeeRefresherCourseRequest) (*dtos.EmployeeRefresherCourseResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeRefresherCourseResponse, error)
}

type employeeRefresherCourseUsecase struct {
	employeeRefresherCourseRepo repositories.EmployeeRefresherCourseRepository
	validator                   *validator.Validate
}

func NewEmployeeRefresherCourseUsecase(
	employeeRefresherCourseRepo repositories.EmployeeRefresherCourseRepository,
	validator *validator.Validate,
) EmployeeRefresherCourseUsecase {
	return &employeeRefresherCourseUsecase{
		employeeRefresherCourseRepo: employeeRefresherCourseRepo,
		validator:                   validator,
	}
}

func (uc *employeeRefresherCourseUsecase) Create(ctx context.Context, req *dtos.CreateEmployeeRefresherCourseRequest) (*dtos.EmployeeRefresherCourseResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to create employee refresher course: %w", err))
	}

	employeeRefresherCourse := &domain.EmployeeRefresherCourse{
		EmployeeID:   req.EmployeeID,
		LanguageCode: req.LanguageCode,
		CourseTitle:  req.CourseTitle,
		DateStart:    req.DateStart,
		DateEnd:      req.DateEnd,
	}

	createdEmployeeRefresherCourse, err := uc.employeeRefresherCourseRepo.Create(ctx, employeeRefresherCourse)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeRefresherCourseDomainToResponseDTO(createdEmployeeRefresherCourse)
	return resp, nil
}

func (uc *employeeRefresherCourseUsecase) Update(ctx context.Context, req *dtos.UpdateEmployeeRefresherCourseRequest) (*dtos.EmployeeRefresherCourseResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input to update employee refresher course: %w", err))
	}

	employeeRefresherCourse := &domain.EmployeeRefresherCourse{
		ID: req.ID,
	}

	if req.CourseTitle != nil {
		employeeRefresherCourse.CourseTitle = *req.CourseTitle
	}

	if req.DateStart != nil {
		employeeRefresherCourse.DateStart = *req.DateStart
	}

	if req.DateEnd != nil {
		employeeRefresherCourse.DateEnd = *req.DateEnd
	}

	updatedEmployeeRefresherCourse, err := uc.employeeRefresherCourseRepo.Update(ctx, employeeRefresherCourse)
	if err != nil {
		return nil, err
	}

	resp := mappers.MapEmployeeRefresherCourseDomainToResponseDTO(updatedEmployeeRefresherCourse)
	return resp, nil
}

func (uc *employeeRefresherCourseUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input to delete employee refresher course"))
	}

	if err := uc.employeeRefresherCourseRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *employeeRefresherCourseUsecase) GetByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*dtos.EmployeeRefresherCourseResponse, error) {
	if employeeID <= 0 {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - EmployeeID(%d) to retrive employee refresher course", employeeID))
	}

	if langCode != "en" && langCode != "ru" && langCode != "tg" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - LanguageCode(%s) to retrive employee refresher course", langCode))
	}

	employeeRefresherCourses, err := uc.employeeRefresherCourseRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeID, langCode)
	if err != nil {
		return nil, err
	}

	resp := make([]*dtos.EmployeeRefresherCourseResponse, len(employeeRefresherCourses))
	for index, degree := range employeeRefresherCourses {
		resp[index] = mappers.MapEmployeeRefresherCourseDomainToResponseDTO(degree)
	}

	return resp, nil
}


