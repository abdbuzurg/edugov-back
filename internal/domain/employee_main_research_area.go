package domain

import "time"

type EmployeeMainResearchArea struct {
	ID           int64
	EmployeeID   int64
	LanguageCode string
	Area         string
	Discipline   string
	KeyTopics    []*ResearchAreaKeyTopic
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ResearchAreaKeyTopic struct {
	ID                         int64
	EmployeeMainResearchAreaID int64
	KeyTopicTitle              string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

func(d *ResearchAreaKeyTopic) GetID() int64 {
  return d.ID
}

func(d *ResearchAreaKeyTopic) IsNew() bool {
  return d.ID == 0
}
