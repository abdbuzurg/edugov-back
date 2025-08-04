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

type EmployeeSocialHandler struct {
	employeeDegreeUC usecases.EmployeeSocialUsecase
}

func NewEmployeeSocialHandler(employeeDegreeUC usecases.EmployeeSocialUsecase) *EmployeeSocialHandler {
	return &EmployeeSocialHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee/social
// Request body - dto.CreateEmployeeSocialRequest
// Response body - dto.EmployeeSocialResponse
func (h *EmployeeSocialHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeSocialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee social: %w", err))
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

// PUT /employee/social
// Request body - dto.UpdateEmployeeSocialRequest
// Response body - dto.EmployeeSocialResponse
func (h *EmployeeSocialHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeSocialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee social: %w", err))
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

// DELETE /employee/social/{id}
// Request body - None
// Response body - None
func (h *EmployeeSocialHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee social by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee/social/{employeeID}
// Request body - none
// Response body - []dtos.EmployeeSocialResponse
func (h *EmployeeSocialHandler) GetByEmployeeID(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee social by employeeID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	resp, err := h.employeeDegreeUC.GetByEmployeeID(r.Context(), int64(employeeID))
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
