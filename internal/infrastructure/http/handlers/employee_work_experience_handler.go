package handlers

import (
	"backend/internal/application/dtos"
	"backend/internal/application/usecases"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EmployeeWorkExperienceHandler struct {
	employeeDegreeUC usecases.EmployeeWorkExperienceUsecase
}

func NewEmployeeWorkExperienceHandler(employeeDegreeUC usecases.EmployeeWorkExperienceUsecase) *EmployeeWorkExperienceHandler {
	return &EmployeeWorkExperienceHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-professional-activity-in-education
// Request body - dto.CreateEmployeeWorkExperienceRequest
// Response body - dto.EmployeeWorkExperienceResponse
func (h *EmployeeWorkExperienceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeWorkExperienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee work experience: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	resp, err := h.employeeDegreeUC.Create(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, resp)
}

// PUT /employee-professional-activity-in-education
// Request body - dto.UpdateEmployeeWorkExperienceRequest
// Response body - dto.EmployeeWorkExperienceResponse
func (h *EmployeeWorkExperienceHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeWorkExperienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee work experience: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	resp, err := h.employeeDegreeUC.Update(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// DELETE /employee-professional-activity-in-education/{id}
// Request body - None
// Response body - None
func (h *EmployeeWorkExperienceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee work experience by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-professional-activity-in-education/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeWorkExperienceResponse
func (h *EmployeeWorkExperienceHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee work experience by employeeID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	langCode := utils.GetLanguageFromContext(r.Context())
	resp, err := h.employeeDegreeUC.GetByEmployeeIDAndLanguageCode(r.Context(), int64(employeeID), langCode)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
