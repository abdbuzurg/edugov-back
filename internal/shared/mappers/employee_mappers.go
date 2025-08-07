package mappers

import (
	"backend/internal/application/dtos"
	"backend/internal/domain"
)

func MapEmployeeDomainToResponseDTO(employee *domain.Employee) *dtos.EmployeeResponse {
	if employee == nil {
		return nil
	}

	return &dtos.EmployeeResponse{
		ID:             employee.ID,
		UniqueID:       employee.UniqueID,
		ProfilePicture: employee.ProfilePicture,
		CreatedAt:      employee.CreatedAt,
		UpdatedAt:      employee.UpdatedAt,
	}
}

func MapEmployeeDetailsDomainIntoResponseDTO(employeeDetails *domain.EmployeeDetails) *dtos.EmployeeDetailsResponse {
	if employeeDetails == nil {
		return nil
	}

	return &dtos.EmployeeDetailsResponse{
		ID:                   employeeDetails.ID,
		LanguageCode:         employeeDetails.LanguageCode,
		Surname:              employeeDetails.Surname,
		Name:                 employeeDetails.Name,
		Middlename:           employeeDetails.Middlename,
		IsEmployeeDetailsNew: employeeDetails.IsEmployeeDetailsNew,
		CreatedAt:            employeeDetails.CreatedAt,
		UpdatedAt:            employeeDetails.UpdatedAt,
	}
}

func MapEmployeeDegreeDomainToResponseDTO(employeeDegree *domain.EmployeeDegree) *dtos.EmployeeDegreeResponse {
	if employeeDegree == nil {
		return nil
	}

	return &dtos.EmployeeDegreeResponse{
		ID:                 employeeDegree.ID,
		DegreeLevel:        employeeDegree.DegreeLevel,
		UniversityName:     employeeDegree.UniversityName,
		Speciality:         employeeDegree.Speciality,
		DateStart:          employeeDegree.DateStart,
		DateEnd:            employeeDegree.DateEnd,
		GivenBy:            employeeDegree.GivenBy,
		DateDegreeRecieved: employeeDegree.DateDegreeRecieved,
		CreatedAt:          employeeDegree.CreatedAt,
		UpdatedAt:          employeeDegree.UpdatedAt,
	}
}

func MapEmployeeWorkExperienceDomainToResponseDTO(employeeWorkExperience *domain.EmployeeWorkExperience) *dtos.EmployeeWorkExperienceResponse {
	if employeeWorkExperience == nil {
		return nil
	}

	return &dtos.EmployeeWorkExperienceResponse{
		ID:          employeeWorkExperience.ID,
		Workplace:   employeeWorkExperience.Workplace,
		JobTitle:    employeeWorkExperience.JobTitle,
		Description: employeeWorkExperience.Description,
		DateStart:   employeeWorkExperience.DateStart,
		DateEnd:     employeeWorkExperience.DateEnd,
		CreatedAt:   employeeWorkExperience.CreatedAt,
		UpdatedAt:   employeeWorkExperience.UpdatedAt,
	}
}

func MapEmployeeMainResearchAreaDomainToResponseDTO(employeeMRA *domain.EmployeeMainResearchArea) *dtos.EmployeeMainResearchAreaResponse {
	if employeeMRA == nil {
		return nil
	}

	keyTopics := make([]*dtos.ResearchAreaKeyTopicResponse, len(employeeMRA.KeyTopics))
	for index, kt := range employeeMRA.KeyTopics {
		keyTopics[index] = &dtos.ResearchAreaKeyTopicResponse{
			ID:            kt.ID,
			KeyTopicTitle: kt.KeyTopicTitle,
			CreatedAt:     kt.CreatedAt,
			UpdatedAt:     kt.UpdatedAt,
		}
	}

	return &dtos.EmployeeMainResearchAreaResponse{
		ID:         employeeMRA.ID,
		Discipline: employeeMRA.Discipline,
		Area:       employeeMRA.Area,
		KeyTopics:  keyTopics,
		CreatedAt:  employeeMRA.CreatedAt,
		UpdatedAt:  employeeMRA.UpdatedAt,
	}
}

func MapEmployeePublicationDomainToResponseDTO(employeePublication *domain.EmployeePublication) *dtos.EmployeePublicationResponse {
	if employeePublication == nil {
		return nil
	}

	return &dtos.EmployeePublicationResponse{
		ID:                employeePublication.ID,
		PublicationTitle:  employeePublication.PublicationTitle,
		LinkToPublication: employeePublication.LinkToPublication,
		CreatedAt:         employeePublication.CreatedAt,
		UpdatedAt:         employeePublication.UpdatedAt,
	}
}

func MapEmployeeScientificAwardDomainToResponseDTO(employeeScientificAward *domain.EmployeeScientificAward) *dtos.EmployeeScientificAwardResponse {
	if employeeScientificAward == nil {
		return nil
	}

	return &dtos.EmployeeScientificAwardResponse{
		ID:                   employeeScientificAward.ID,
		ScientificAwardTitle: employeeScientificAward.ScientificAwardTitle,
		GivenBy:              employeeScientificAward.GivenBy,
		CreatedAt:            employeeScientificAward.CreatedAt,
		UpdatedAt:            employeeScientificAward.UpdatedAt,
	}
}

func MapEmployeePatentDomainToResponseDTO(employeePatent *domain.EmployeePatent) *dtos.EmployeePatentResponse {
	if employeePatent == nil {
		return nil
	}

	return &dtos.EmployeePatentResponse{
		ID:          employeePatent.ID,
		PatentTitle: employeePatent.PatentTitle,
		Description: employeePatent.Description,
		CreatedAt:   employeePatent.CreatedAt,
		UpdatedAt:   employeePatent.UpdatedAt,
	}
}

func MapEmployeeParticipationInProfessionalCommunityDomainToResponseDTO(employeeParticipationInProfessionalCommunity *domain.EmployeeParticipationInProfessionalCommunity) *dtos.EmployeeParticipationInProfessionalCommunityResponse {
	if employeeParticipationInProfessionalCommunity == nil {
		return nil
	}

	return &dtos.EmployeeParticipationInProfessionalCommunityResponse{
		ID:                          employeeParticipationInProfessionalCommunity.ID,
		ProfessionalCommunityTitle:  employeeParticipationInProfessionalCommunity.ProfessionalCommunityTitle,
		RoleInProfessionalCommunity: employeeParticipationInProfessionalCommunity.RoleInProfessionalCommunity,
		CreatedAt:                   employeeParticipationInProfessionalCommunity.CreatedAt,
		UpdatedAt:                   employeeParticipationInProfessionalCommunity.UpdatedAt,
	}
}

func MapEmployeeRefresherCourseDomainToResponseDTO(employeeRefresherCourse *domain.EmployeeRefresherCourse) *dtos.EmployeeRefresherCourseResponse {
	if employeeRefresherCourse == nil {
		return nil
	}

	return &dtos.EmployeeRefresherCourseResponse{
		ID:          employeeRefresherCourse.ID,
		CourseTitle: employeeRefresherCourse.CourseTitle,
		DateStart:   employeeRefresherCourse.DateStart,
		DateEnd:     employeeRefresherCourse.DateEnd,
		CreatedAt:   employeeRefresherCourse.CreatedAt,
		UpdatedAt:   employeeRefresherCourse.UpdatedAt,
	}
}

func MapEmployeeParticipationInEventDomainToResponseDTO(employeeParticipationInEvent *domain.EmployeeParticipationInEvent) *dtos.EmployeeParticipationInEventResponse {
	if employeeParticipationInEvent == nil {
		return nil
	}

	return &dtos.EmployeeParticipationInEventResponse{
		ID:         employeeParticipationInEvent.ID,
		EventTitle: employeeParticipationInEvent.EventTitle,
		EventDate:  employeeParticipationInEvent.EventDate,
		CreatedAt:  employeeParticipationInEvent.CreatedAt,
		UpdatedAt:  employeeParticipationInEvent.UpdatedAt,
	}
}

func MapEmployeeResearchActivityDomainToResponseDTO(employeeResearchActivity *domain.EmployeeResearchActivity) *dtos.EmployeeResearchActivityResponse {
	if employeeResearchActivity == nil {
		return nil
	}

	return &dtos.EmployeeResearchActivityResponse{
		ID:                    employeeResearchActivity.ID,
		ResearchActivityTitle: employeeResearchActivity.ResearchActivityTitle,
		EmployeeRole:          employeeResearchActivity.EmployeeRole,
		CreatedAt:             employeeResearchActivity.CreatedAt,
		UpdatedAt:             employeeResearchActivity.UpdatedAt,
	}
}

func MapEmployeeSocialDomainToResponseDTO(employeeSocial *domain.EmployeeSocial) *dtos.EmployeeSocialResponse {
	if employeeSocial == nil {
		return nil
	}

	return &dtos.EmployeeSocialResponse{
		ID:           employeeSocial.ID,
		SocialName:   employeeSocial.SocialName,
		LinkToSocial: employeeSocial.LinkToSocial,
		CreatedAt:    employeeSocial.CreatedAt,
		UpdatedAt:    employeeSocial.UpdatedAt,
	}
}
