package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeMainResearchAreaRequest struct {
	EmployeeID   int64                                `json:"employeeID" validate:"required,min=1"`
	LanguageCode string                               `json:"languageCode" validate:"required,len=2"`
	Area         string                               `json:"area" validate:"required"`
	Discipline   string                               `json:"discipline" validate:"required"`
	KeyTopics    []*CreateResearchAreaKeyTopicRequest `json:"keyTopics" validate:"required,dive"`
}

type UpdateEmployeeMainResearchAreaRequest struct {
	ID         int64                                `json:"id" validate:"required,min=1"`
	Discipline *string                              `json:"discipline" validate:"omitempty"`
	Area       *string                              `json:"area" validate:"omitempty"`
	KeyTopics  []*UpdateResearchAreaKeyTopicRequest `json:"keyTopics" required:"omitempty,dive"`
}

type CreateResearchAreaKeyTopicRequest struct {
	LanguageCode  string `json:"languageCode" validate:"required,len=2"`
	KeyTopicTitle string `json:"keyTopicTitle"`
}

type UpdateResearchAreaKeyTopicRequest struct {
	ID            int64   `json:"id" validate:"required,min=1"`
	KeyTopicTitle *string `json:"keyTopicTitle" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeMainResearchAreaResponse struct {
	ID         int64                           `json:"id"`
	Discipline string                          `json:"discipline"`
	Area       string                          `json:"area"`
	KeyTopics  []*ResearchAreaKeyTopicResponse `json:"keyTopics,omitempty"`
	CreatedAt  time.Time                       `json:"createdAt"`
	UpdatedAt  time.Time                       `json:"updatedAt"`
}

type ResearchAreaKeyTopicResponse struct {
	ID            int64     `json:"id"`
	KeyTopicTitle string    `json:"keyTopicTitle"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
