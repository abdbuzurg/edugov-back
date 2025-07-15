package domain

import "time"

type InstitutionRanking struct {
	ID                  int64
	InstitutionID       int64
	LanguageCode        string
	RankingTitle        string
	RankingType         string
	DateReceived        time.Time
	RankingAgency       string
	Description         string
	LinkToRankingFile   string
	LinkToRankingAgency string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
