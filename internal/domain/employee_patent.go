package domain

import "time"

type EmployeePatent struct {
	ID               int64
	EmployeeID       int64
	LanguageCode     string
	PatentTitle      string
	Description      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
