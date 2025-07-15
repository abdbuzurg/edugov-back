package domain

import "time"

type EmployeeParticipationInEvent struct {
	ID                             int64
	EmployeeID                     int64
	LanguageCode                   string
	EventTitle                     string
	EventDate                      time.Time
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}
