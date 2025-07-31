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

type EmployeePatentHandler struct {
	employeeDegreeUC usecases.EmployeePatentUsecase
}

func NewEmployeePatentHandler(employeeDegreeUC usecases.EmployeePatentUsecase) *EmployeePatentHandler {
	return &EmployeePatentHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-patent
// Request body - dto.CreateEmployeePatentRequest
// Response body - dto.EmployeePatentResponse
func (h *EmployeePatentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeePatentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee patent: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	req.LanguageCode = middleware.GetLanguageFromContext(r.Context())
	resp, err := h.employeeDegreeUC.Create(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, resp)
}

// PUT /employee-patent
// Request body - dto.UpdateEmployeePatentRequest
// Response body - dto.EmployeePatentResponse
func (h *EmployeePatentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeePatentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee patent: %w", err))
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

// DELETE /employee-patent/{id}
// Request body - None
// Response body - None
func (h *EmployeePatentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee patent by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-patent/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeePatentResponse
func (h *EmployeePatentHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee patent by employeeID: %w", err))
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
