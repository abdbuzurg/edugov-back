package domain

import "time"

type EmployeeScientificAward struct {
	ID                        int64
	EmployeeID                int64
	LanguageCode              string
	ScientificAwardTitle      string
	GivenBy                   string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}
