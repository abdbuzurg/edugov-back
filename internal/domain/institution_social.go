package domain

import "time"

type InstitutionSocial struct {
	ID            int64
	InstitutionID int64
	LinkToSocial  string
	SocialName    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
