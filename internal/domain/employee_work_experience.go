package domain

import "time"

type EmployeeWorkExperience struct {
	ID           int64
	EmployeeID   int64
	LanguageCode string
	Workplace    string
	JobTitle     string
	Description  string
	DateStart    time.Time
	DateEnd      time.Time
	Ongoing      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
