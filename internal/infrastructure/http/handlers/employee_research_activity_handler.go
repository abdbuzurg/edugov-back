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

type EmployeeResearchActivityHandler struct {
	employeeDegreeUC usecases.EmployeeResearchActivityUsecase
}

func NewEmployeeResearchActivityHandler(employeeDegreeUC usecases.EmployeeResearchActivityUsecase) *EmployeeResearchActivityHandler {
	return &EmployeeResearchActivityHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-research-activity
// Request body - dto.CreateEmployeeResearchActivityRequest
// Response body - dto.EmployeeResearchActivityResponse
func (h *EmployeeResearchActivityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeResearchActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee research activity: %w", err))
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

// PUT /employee-research-activity
// Request body - dto.UpdateEmployeeResearchActivityRequest
// Response body - dto.EmployeeResearchActivityResponse
func (h *EmployeeResearchActivityHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeResearchActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee research activity: %w", err))
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

// DELETE /employee-research-activity/{id}
// Request body - None
// Response body - None
func (h *EmployeeResearchActivityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee research activity by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-research-activity/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeResearchActivityResponse
func (h *EmployeeResearchActivityHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee research activity by employeeID: %w", err))
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
