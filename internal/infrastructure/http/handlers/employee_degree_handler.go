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

type EmployeeDegreeHandler struct {
	employeeDegreeUC usecases.EmployeeDegreeUsecase
}

func NewEmployeeDegreeHandler(employeeDegreeUC usecases.EmployeeDegreeUsecase) *EmployeeDegreeHandler {
	return &EmployeeDegreeHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee/degree
// Request body - dto.CreateEmployeeDegreeRequest
// Response body - dto.EmployeeDegreeResponse
func (h *EmployeeDegreeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeDegreeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee degree: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	langCode := utils.GetLanguageFromContext(r.Context())
	req.LanguageCode = langCode
	resp, err := h.employeeDegreeUC.Create(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, resp)
}

// PUT /employee/degree
// Request body - dto.UpdateEmployeeDegreeRequest
// Response body - dto.EmployeeDegreeResponse
func (h *EmployeeDegreeHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeDegreeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee degree: %w", err))
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

// DELETE /employee/degree/{id}
// Request body - None
// Response body - None
func (h *EmployeeDegreeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee degree by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee/degree/{employeeID}
// Request body - none
// Response body - []dtos.EmployeeDegreeResponse
func (h *EmployeeDegreeHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee degree by employeeID: %w", err))
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
