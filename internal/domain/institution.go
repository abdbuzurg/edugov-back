package domain

import "time"

type Institution struct {
	ID                  int64
	YearOfEstablishment int32
	Email               string
	Fax                 string
	OfficialWebsite     string
	PhoneNumber         string
	MailIndex           string
	CreatedAt           time.Time
	UpdatedAt           time.Time

	Accreditations                 []InstitutionAccreditation
	Achievements                   []InstitutionAchievement
	Conferences                    []InstitutionConference
	Details                        *InstitutionDetails
	Licences                       []InstitutionLicence
	Magazine                       []InstitutionMagazine
	MainResearchDirections         []InstitutionMainResearchDirection
	Patents                        []InstitutionPatent
	Projects                       []InstitutionProject
	Rankings                       []InstitutionRanking
	ResearchSupportInfrastructures []InstitutionResearchSupportInfrastructure
	Socials                        []InstitutionSocial
}
