package domain

import "time"

type EmployeeResearchActivity struct {
	ID                    int64
	EmployeeID            int64
	LanguageCode          string
	ResearchActivityTitle string
	EmployeeRole          string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
