package domain

import "time"

type EmployeeDegree struct {
	ID                 int64
	EmployeeID         int64
	LanguageCode       string
	DegreeLevel        string
	UniversityName     string
	Speciality         string
	DateStart          time.Time
	DateEnd            time.Time
	GivenBy            string
	DateDegreeRecieved time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
