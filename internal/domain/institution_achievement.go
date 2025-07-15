package domain

import "time"

type InstitutionAchievement struct {
	ID               int64
	InstitutionID    int64
	LanguageCode     string
	AchievementTitle string
	AchievementType  string
	DateReceived     time.Time
	GivenBy          string
	LinkToFile       string
	Description      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
