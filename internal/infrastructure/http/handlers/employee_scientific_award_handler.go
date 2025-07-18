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

type EmployeeScientificAwardHandler struct {
	employeeDegreeUC usecases.EmployeeScientificAwardUsecase
}

func NewEmployeeScientificAwardHandler(employeeDegreeUC usecases.EmployeeScientificAwardUsecase) *EmployeeScientificAwardHandler {
	return &EmployeeScientificAwardHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-scientific-award
// Request body - dto.CreateEmployeeScientificAwardRequest
// Response body - dto.EmployeeScientificAwardResponse
func (h *EmployeeScientificAwardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeScientificAwardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee scientific award: %w", err))
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

// PUT /employee-scientific-award
// Request body - dto.UpdateEmployeeScientificAwardRequest
// Response body - dto.EmployeeScientificAwardResponse
func (h *EmployeeScientificAwardHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeScientificAwardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee scientific award: %w", err))
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

// DELETE /employee-scientific-award/{id}
// Request body - None
// Response body - None
func (h *EmployeeScientificAwardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee scientific award by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-scientific-award/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeScientificAwardResponse
func (h *EmployeeScientificAwardHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee scientific award by employeeID: %w", err))
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
