package domain

import "time"

type EmployeeDetails struct {
	ID                   int64
	EmployeeID           int64
	LanguageCode         string
	Surname              string
	Name                 string
	Middlename           string
	IsEmployeeDetailsNew bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (d *EmployeeDetails) GetID() int64 {
  return d.ID
}

func (d *EmployeeDetails) IsNew() bool {
  return d.ID == 0
}
