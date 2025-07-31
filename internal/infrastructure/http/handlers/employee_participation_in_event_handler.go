package handlers

import (
	"backend/internal/application/dtos"
	"backend/internal/application/usecases"
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EmployeeParticipationInEventHandler struct {
	employeeDegreeUC usecases.EmployeeParticipationInEventUsecase
}

func NewEmployeeParticipationInEventHandler(employeeDegreeUC usecases.EmployeeParticipationInEventUsecase) *EmployeeParticipationInEventHandler {
	return &EmployeeParticipationInEventHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-participation-in-event
// Request body - dto.CreateEmployeeParticipationInEventRequest
// Response body - dto.EmployeeParticipationInEventResponse
func (h *EmployeeParticipationInEventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeParticipationInEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee participation in event: %w", err))
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

// PUT /employee-participation-in-event
// Request body - dto.UpdateEmployeeParticipationInEventRequest
// Response body - dto.EmployeeParticipationInEventResponse
func (h *EmployeeParticipationInEventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeParticipationInEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee participation in event: %w", err))
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

// DELETE /employee-participation-in-event/{id}
// Request body - None
// Response body - None
func (h *EmployeeParticipationInEventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee participation in event by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-participation-in-event/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeParticipationInEventResponse
func (h *EmployeeParticipationInEventHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee participation in event by employeeID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	langCode := middleware.GetLanguageFromContext(r.Context())
	resp, err := h.employeeDegreeUC.GetByEmployeeIDAndLanguageCode(r.Context(), int64(employeeID), langCode)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
