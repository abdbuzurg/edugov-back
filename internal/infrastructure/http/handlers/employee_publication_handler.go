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

type EmployeePublicationHandler struct {
	employeeDegreeUC usecases.EmployeePublicationUsecase
}

func NewEmployeePublicationHandler(employeeDegreeUC usecases.EmployeePublicationUsecase) *EmployeePublicationHandler {
	return &EmployeePublicationHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-publication
// Request body - dto.CreateEmployeePublicationRequest
// Response body - dto.EmployeePublicationResponse
func (h *EmployeePublicationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeePublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee publication: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	req.LanguageCode = utils.GetLanguageFromContext(r.Context())
	resp, err := h.employeeDegreeUC.Create(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, resp)
}

// PUT /employee-publication
// Request body - dto.UpdateEmployeePublicationRequest
// Response body - dto.EmployeePublicationResponse
func (h *EmployeePublicationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeePublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee publication: %w", err))
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

// DELETE /employee-publication/{id}
// Request body - None
// Response body - None
func (h *EmployeePublicationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee publication by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-publication/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeePublicationResponse
func (h *EmployeePublicationHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee publication by employeeID: %w", err))
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
