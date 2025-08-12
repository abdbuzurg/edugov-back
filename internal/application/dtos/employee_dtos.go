package dtos

import "time"

// ---- REQUEST DTOS ----
type UpdateEmployeeRequest struct {
	ID           int64                           `json:"id" validate:"required"`
	DateOfBirth  *time.Time                      `json:"dateOfBirth" validate:"omitempty,datetime"`
	Gender       *string                         `json:"gender" validate:"omitempty,len=1"`
	Email        *string                         `json:"email" validate:"omitempty,email"`
	MobileNumber *string                         `json:"mobileNumber" validate:"omitempty,e164"`
	PassportCode *string                         `json:"passportCode" validate:"omitempty"`
	LinkToCVFile *string                         `json:"linkToCVFile" validate:"omitempty"`
	Details      []*UpdateEmployeeDetailsRequest `json:"details" validate:"omitempty,dive"`
}

type PersonnelPaginatedQueryParameters struct {
	LanguageCode          string
	UID                   string
	Name                  string
	Surname               string
	Middlename            string
	HighestAcademicDegree string
	Speciality            string
	WorkExperience        int64
	Limit                 int64
	Page                  int64
}

// ---- RESPONSE DTOS ----

type EmployeeResponse struct {
	ID             int64     `json:"id"`
	UniqueID       string    `json:"uniqueID"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`

	Details                                []*EmployeeDetailsResponse                              `json:"details,omitempty"`
	Degrees                                []*EmployeeDegreeResponse                               `json:"degrees,omitempty"`
	WorkExperiences                        []*EmployeeWorkExperienceResponse                       `json:"workExperiences,omitempty"`
	MainResearchAreas                      []*EmployeeMainResearchAreaResponse                     `json:"mainResearchAreas,omitempty"`
	Publications                           []*EmployeePublicationResponse                          `json:"publications,omitempty"`
	ScientificAwards                       []*EmployeeScientificAwardResponse                      `json:"scientificAwards,omitempty"`
	Patents                                []*EmployeePatentResponse                               `json:"patents,omitempty"`
	ParticipationInProfessionalCommunities []*EmployeeParticipationInProfessionalCommunityResponse `json:"participationInProfessionalCommunities,omitempty"`
	RefresherCourses                       []*EmployeeRefresherCourseResponse                      `json:"refresherCourses,omitempty"`
	ParticipationInEvents                  []*EmployeeParticipationInEventResponse                 `json:"participationInEvents,omitempty"`
	ResearchActivities                     []*EmployeeResearchActivityResponse                     `json:"researchActivities,omitempty"`
	Socials                                []*EmployeeSocialResponse                               `json:"socials,omitempty"`
}

type PersonnelProfileData struct {
	Fullname              string                   `json:"fullname"`
	UID                   string                   `json:"uid"`
	HighestAcademicDegree string                   `json:"highestAcademicDegree"`
	Speciality            string                   `json:"speciality"`
	CurrentWorkplace      string                   `json:"currentWorkplace"`
	WorkExperienceYears   int64                    `json:"workExperienceYears"`
	WorkExperienceMonths  int64                    `json:"workExperienceMonths"`
	Socials               []EmployeeSocialResponse `json:"socials"`
}

type PersonnelPaginatedResponse struct {
	Data     []PersonnelProfileData `json:"data"`
	NextPage int64                  `json:"nextPage,omitempty"`
}
