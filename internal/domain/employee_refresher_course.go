package domain

import "time"

type EmployeeRefresherCourse struct {
	ID           int64
	EmployeeID   int64
	LanguageCode string
	CourseTitle  string
	DateStart    time.Time
	DateEnd      time.Time
	CreatedAt    time.Time
	UpdatedAt     time.Time
}
