package domain

import "time"

type EmployeeParticipationInProfessionalCommunity struct {
	ID                          int64
	EmployeeID                  int64
	LanguageCode                string
	ProfessionalCommunityTitle  string
	RoleInProfessionalCommunity string
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}
