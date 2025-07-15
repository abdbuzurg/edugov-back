package domain

import "time"

type EmployeeSocial struct {
	ID           int64
	EmployeeID   int64
	SocialName   string
	LinkToSocial string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
