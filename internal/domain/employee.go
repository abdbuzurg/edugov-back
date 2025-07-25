package domain

import "time"

type Employee struct {
	ID             int64
	UniqueID       string
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Degrees                                []EmployeeDegree
	Details                                []*EmployeeDetails
	MainResearchAreas                      []EmployeeMainResearchArea
	ParticipationInEvents                  []EmployeeParticipationInEvent
	ParticipationInProfessionalCommunities []EmployeeParticipationInProfessionalCommunity
	Patents                                []EmployeePatent
	WorkExperiences                        []EmployeeWorkExperience
	Publications                           []EmployeePublication
	RefresherCourses                       []EmployeeRefresherCourse
	ResearchActivity                       []EmployeeResearchActivity
	ScientificAwards                       []EmployeeScientificAward
	Socials                                []EmployeeSocial
}
