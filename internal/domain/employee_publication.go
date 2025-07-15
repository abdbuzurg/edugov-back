package domain

import "time"

type EmployeePublication struct {
	ID                int64
	EmployeeID        int64
	LanguageCode      string
	PublicationTitle  string
	LinkToPublication string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
