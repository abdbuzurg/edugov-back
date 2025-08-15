package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/mappers"
	"backend/internal/shared/utils"
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type EmployeeUsecase interface {
	Delete(ctx context.Context, id int64) error
	GetByUniqueID(ctx context.Context, uniqueID string) (*dtos.EmployeeResponse, error)
	GetPersonnelPaginated(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) (*dtos.PersonnelPaginatedResponse, error)
}

type employeeUsecase struct {
	employeeRepo repositories.EmployeeRepository
	store        *postgres.Store
	validator    *validator.Validate
}

func NewEmployeeUsecase(
	employeeRepo repositories.EmployeeRepository,
	store *postgres.Store,
	validator *validator.Validate,
) EmployeeUsecase {
	return &employeeUsecase{
		employeeRepo: employeeRepo,
		store:        store,
		validator:    validator,
	}
}

func (uc *employeeUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return custom_errors.BadRequest(fmt.Errorf("invalid input provided for employee deletion"))
	}

	return uc.employeeRepo.Delete(ctx, id)
}

func (uc *employeeUsecase) GetByUniqueID(ctx context.Context, uniqueID string) (*dtos.EmployeeResponse, error) {
	if uniqueID == "" {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid input - uniqueID(%s) is provided for retrival", uniqueID))
	}

	var resp *dtos.EmployeeResponse
	langCode := middleware.GetLanguageFromContext(ctx)
	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeRepo := postgres.NewPgEmployeeRepositoryWithQuery(q)
		txEmployeeDetailsRepo := postgres.NewPGEmployeeDetailsRepositoryWithQueries(q)
		txEmployeeDegreeRepo := postgres.NewPgEmployeeDegreeRepositoryWithQuery(q)
		txEmployeeWorkExperienceRepo := postgres.NewPgEmployeeWorkExperienceRepositoryWithQuery(q)
		txEmployeeMainResearchAreaRepo := postgres.NewPgEmployeeMainResearchAreaRepositoryWithQueries(q)
		txEmployeePublicationRepo := postgres.NewPgEmployeePublicationRepositoryWithQuery(q)
		txEmployeeScientificAwardRepo := postgres.NewPgEmployeeScientificAwardRepositoryWithQuery(q)
		txEmployeePatentRepo := postgres.NewPgEmployeePatentRepositoryWithQuery(q)
		txEmployeeParticipationInProfessionalCommunityRepo := postgres.NewPgEmployeeParticipationInProfessionalCommunityRepositoryWithQuery(q)
		txEmployeeRefresherCourseRepo := postgres.NewPgEmployeeRefresherCourseRepositoryWithQuery(q)
		txEmployeeParticipationInEvenRepo := postgres.NewPgEmployeeParticipationInEventRepositoryWithQuery(q)
		txEmployeeResearchActivityRepo := postgres.NewPgEmployeeResearchActivityRepositoryWithQueries(q)
		txEmployeeSocialRepo := postgres.NewPgEmployeeSocialRepositoryWithQueries(q)

		employee, err := txEmployeeRepo.GetByUniqueID(ctx, uniqueID)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		} else if custom_errors.IsNotFound(err) {
			return custom_errors.BadRequest(fmt.Errorf("no user with given unique id"))
		}
		resp = mappers.MapEmployeeDomainToResponseDTO(employee)

		//Employee Details
		employeeDetails, err := txEmployeeDetailsRepo.GetByEmployeeID(ctx, employee.ID)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.Details = make([]*dtos.EmployeeDetailsResponse, len(employeeDetails))
		for index, details := range employeeDetails {
			resp.Details[index] = mappers.MapEmployeeDetailsDomainIntoResponseDTO(details)
		}

		//Employee Degress
		employeeDegrees, err := txEmployeeDegreeRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.Degrees = make([]*dtos.EmployeeDegreeResponse, len(employeeDegrees))
		for index, degree := range employeeDegrees {
			resp.Degrees[index] = mappers.MapEmployeeDegreeDomainToResponseDTO(degree)
		}

		//Employee Work Experience
		employeeWorkExperiences, err := txEmployeeWorkExperienceRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.WorkExperiences = make([]*dtos.EmployeeWorkExperienceResponse, len(employeeWorkExperiences))
		for index, workExperience := range employeeWorkExperiences {
			resp.WorkExperiences[index] = mappers.MapEmployeeWorkExperienceDomainToResponseDTO(workExperience)
		}

		//Employee Main Research Area
		employeeMRAs, err := txEmployeeMainResearchAreaRepo.GetMRAByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		for index, mra := range employeeMRAs {
			rakts, err := txEmployeeMainResearchAreaRepo.GetRAKTByMRAIDAndLanguageCode(ctx, mra.ID)
			if err != nil && !custom_errors.IsNotFound(err) {
				return err
			}

			employeeMRAs[index].KeyTopics = rakts
		}
		resp.MainResearchAreas = make([]*dtos.EmployeeMainResearchAreaResponse, len(employeeMRAs))
		for index, mra := range employeeMRAs {
			resp.MainResearchAreas[index] = mappers.MapEmployeeMainResearchAreaDomainToResponseDTO(mra)
		}

		//Employee Publications
		employeePublications, err := txEmployeePublicationRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.Publications = make([]*dtos.EmployeePublicationResponse, len(employeePublications))
		for index, publication := range employeePublications {
			resp.Publications[index] = mappers.MapEmployeePublicationDomainToResponseDTO(publication)
		}

		//Employee Scientific Awards
		employeeScientficAwards, err := txEmployeeScientificAwardRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.ScientificAwards = make([]*dtos.EmployeeScientificAwardResponse, len(employeeScientficAwards))
		for index, award := range employeeScientficAwards {
			resp.ScientificAwards[index] = mappers.MapEmployeeScientificAwardDomainToResponseDTO(award)
		}

		//Employee Patents
		employeePatents, err := txEmployeePatentRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.Patents = make([]*dtos.EmployeePatentResponse, len(employeePatents))
		for index, patent := range employeePatents {
			resp.Patents[index] = mappers.MapEmployeePatentDomainToResponseDTO(patent)
		}

		//Employee Participation In Professional Communities
		employeePIPCs, err := txEmployeeParticipationInProfessionalCommunityRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.ParticipationInProfessionalCommunities = make([]*dtos.EmployeeParticipationInProfessionalCommunityResponse, len(employeePIPCs))
		for index, pipc := range employeePIPCs {
			resp.ParticipationInProfessionalCommunities[index] = mappers.MapEmployeeParticipationInProfessionalCommunityDomainToResponseDTO(pipc)
		}

		//Employee Refresher Courses
		employeeRCs, err := txEmployeeRefresherCourseRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.RefresherCourses = make([]*dtos.EmployeeRefresherCourseResponse, len(employeeRCs))
		for index, rc := range employeeRCs {
			resp.RefresherCourses[index] = mappers.MapEmployeeRefresherCourseDomainToResponseDTO(rc)
		}

		//Employee Participation In Events
		employeePIE, err := txEmployeeParticipationInEvenRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.ParticipationInEvents = make([]*dtos.EmployeeParticipationInEventResponse, len(employeePIE))
		for index, pie := range employeePIE {
			resp.ParticipationInEvents[index] = mappers.MapEmployeeParticipationInEventDomainToResponseDTO(pie)
		}

		//Employee Research Activities
		employeeRAs, err := txEmployeeResearchActivityRepo.GetByEmployeeIDAndLanguageCode(ctx, employee.ID, langCode)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.ResearchActivities = make([]*dtos.EmployeeResearchActivityResponse, len(employeeRAs))
		for index, ra := range employeeRAs {
			resp.ResearchActivities[index] = mappers.MapEmployeeResearchActivityDomainToResponseDTO(ra)
		}

		//Employee Socials
		employeeSocials, err := txEmployeeSocialRepo.GetByEmployeeID(ctx, employee.ID)
		if err != nil && !custom_errors.IsNotFound(err) {
			return err
		}
		resp.Socials = make([]*dtos.EmployeeSocialResponse, len(employeeSocials))
		for index, social := range employeeSocials {
			resp.Socials[index] = mappers.MapEmployeeSocialDomainToResponseDTO(social)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (uc *employeeUsecase) GetPersonnelPaginated(ctx context.Context, filter *dtos.PersonnelPaginatedQueryParameters) (*dtos.PersonnelPaginatedResponse, error) {
	result := &dtos.PersonnelPaginatedResponse{}
	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txEmployeeRepo := postgres.NewPgEmployeeRepositoryWithQuery(q)
		txEmployeeDetailsRepo := postgres.NewPGEmployeeDetailsRepositoryWithQueries(q)
		txEmployeeDegreeRepo := postgres.NewPgEmployeeDegreeRepositoryWithQuery(q)
		txEmployeeWorkExperienceRepo := postgres.NewPgEmployeeWorkExperienceRepositoryWithQuery(q)
		txEmployeeSocialRepo := postgres.NewPgEmployeeSocialRepositoryWithQueries(q)

		employeeIDsAndUIDs, err := txEmployeeRepo.GetPersonnelIDsPaginated(ctx, filter)
		if err != nil {
			if custom_errors.IsNotFound(err) {
				result.Data = []dtos.PersonnelProfileData{}
				result.NextPage = 0
				return nil
			}

			return err
		}
		totalPersonnelByFilter, err := txEmployeeRepo.CountPersonnel(ctx, filter)
		if err != nil && custom_errors.IsNotFound(err) {
			return err
		}

		if totalPersonnelByFilter-filter.Page*filter.Limit > 0 {
			result.NextPage = filter.Page + 1
		} else {
			result.NextPage = 0
		}

		personnelData := make([]dtos.PersonnelProfileData, len(employeeIDsAndUIDs))
		for index := range personnelData {
			//personnel UID
			currentPersonnel := dtos.PersonnelProfileData{
				UID: employeeIDsAndUIDs[index].UniqueID,
			}
			// personnel fullname
			currentEmployeeDetails, err := txEmployeeDetailsRepo.GetCurrentDetailsByEmployeeIDAndLanguageCode(ctx, employeeIDsAndUIDs[index].ID, filter.LanguageCode)
			if err != nil {
				if custom_errors.IsNotFound(err) {
					continue
				}
				return err
			}
			currentPersonnel.Fullname = fmt.Sprintf("%s %s", currentEmployeeDetails.Surname, currentEmployeeDetails.Name)
			if currentEmployeeDetails.Middlename != "" {
				currentPersonnel.Fullname += " " + currentEmployeeDetails.Middlename
			}

			//personnel degree (HighestAcademicDegree, Speciality)
			currentEmployeeDegrees, err := txEmployeeDegreeRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeIDsAndUIDs[index].ID, filter.LanguageCode)
			if err != nil {
				if custom_errors.IsNotFound(err) {
					continue
				}

				return err
			}
			currentPersonnel.HighestAcademicDegree = currentEmployeeDegrees[0].DegreeLevel
			currentPersonnel.Speciality = currentEmployeeDegrees[0].Speciality

			// personnel work experience (CurrentWorkplace, WorkExperienceYears, WorkExperienceMonths)
			currentEmployeeWorkExperiences, err := txEmployeeWorkExperienceRepo.GetByEmployeeIDAndLanguageCode(ctx, employeeIDsAndUIDs[index].ID, filter.LanguageCode)
			if err != nil {
				if custom_errors.IsNotFound(err) {
					continue
				}
				return err
			}
			currentPersonnel.CurrentWorkplace = currentEmployeeWorkExperiences[0].Workplace

			workExperienceCountYears := 0
			workExperienceCountMonths := 0
			for _, workExperience := range currentEmployeeWorkExperiences {
				if workExperience.Ongoing {
					workExperience.DateEnd = time.Now()
				}
				yearDiff, monthDiff := utils.DateDifference(workExperience.DateStart, workExperience.DateEnd)
				workExperienceCountYears += yearDiff
				workExperienceCountMonths += monthDiff
			}

			workExperienceCountYears += workExperienceCountMonths / 12

			currentPersonnel.WorkExperience = int64(workExperienceCountYears)

			// personnel socials
			currentEmployeeSocials, err := txEmployeeSocialRepo.GetByEmployeeID(ctx, employeeIDsAndUIDs[index].ID)
			if err != nil && !custom_errors.IsNotFound(err) {
				return err
			}

			for _, social := range currentEmployeeSocials {
				if len(currentPersonnel.Socials) == 5 {
					break
				}

				currentPersonnel.Socials = append(currentPersonnel.Socials, *mappers.MapEmployeeSocialDomainToResponseDTO(social))
			}

			personnelData[index] = currentPersonnel
		}

		result.Data = personnelData
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
