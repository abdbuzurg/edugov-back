package repositories

import (
	"backend/internal/domain"
	"context"
)

type EmployeeMainResearchArea interface {
	// CreateMRA - insterts employee main research area data into DB.
	CreateMRA(ctx context.Context, employeeMRA *domain.EmployeeMainResearchArea) (*domain.EmployeeMainResearchArea, error)

	// CreateRAKT - insterts  research area key topic data into DB.
	CreateRAKT(ctx context.Context, rakt *domain.ResearchAreaKeyTopic) (*domain.ResearchAreaKeyTopic, error)

	// UpdateMRA - modifies employee main research area in DB.
	UpdateMRA(ctx context.Context, employeeMRA *domain.EmployeeMainResearchArea) (*domain.EmployeeMainResearchArea, error)

	// UpdateRAKT - modifies research area key topic in DB.
	UpdateRAKT(ctx context.Context, rakt *domain.ResearchAreaKeyTopic) (*domain.ResearchAreaKeyTopic, error)

	//Delete - removes employee main research entry from DB.
	DeleteMRA(ctx context.Context, id int64) error

	//Delete - removes main research key topic entry from DB.
	DeleteRAKT(ctx context.Context, id int64) error

	//GetMRAByID - retrives employee main research area entry from DB by ID
	GetMRAByID(ctx context.Context, id int64) (*domain.EmployeeMainResearchArea, error)

	//GetRAKTByID - retrives  main research area entry key topic from DB by ID
	GetRAKTByID(ctx context.Context, id int64) (*domain.ResearchAreaKeyTopic, error)

	//GetByEmployeeIDAndLanguageCode - retrives employee main research area entries from DB by EmployeeID and specified language code.
	GetMRAByEmployeeIDAndLanguageCode(ctx context.Context, employeeID int64, langCode string) ([]*domain.EmployeeMainResearchArea, error)

	//GetRAKTByMRAIDAndLanguageCode - retrives main research area key topic entries from DB by EmployeeMainResearchAreaID and specified language code.
	GetRAKTByMRAIDAndLanguageCode(ctx context.Context, employeeMRAID int64) ([]*domain.ResearchAreaKeyTopic, error)
}
