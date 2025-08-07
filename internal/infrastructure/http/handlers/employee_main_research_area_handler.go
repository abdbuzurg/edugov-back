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

type EmployeeMainResearchAreaHandler struct {
	employeeMRAUC usecases.EmployeeMainResearchAreaUsecase
}

func NewEmployeeMainResearchAreaHandler(employeeMRAUC usecases.EmployeeMainResearchAreaUsecase) *EmployeeMainResearchAreaHandler {
	return &EmployeeMainResearchAreaHandler{
		employeeMRAUC: employeeMRAUC,
	}
}

// POST /employee-main-research-area
// Request body - dto.CreateEmployeeMainResearchAreaRequest
// Response body - dto.EmployeeMainResearchAreaResponse
func (h *EmployeeMainResearchAreaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeMainResearchAreaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee main research area: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	lang := middleware.GetLanguageFromContext(r.Context())
	req.LanguageCode = lang
	for index := range req.KeyTopics {
		req.KeyTopics[index].LanguageCode = lang
	}
	resp, err := h.employeeMRAUC.Create(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, resp)
}

// PUT /employee-main-research-area
// Request body - dto.UpdateEmployeeMainResearchAreaRequest
// Response body - dto.EmployeeMainResearchAreaResponse
func (h *EmployeeMainResearchAreaHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeMainResearchAreaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee main research area: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	resp, err := h.employeeMRAUC.Update(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// DELETE /employee-main-research-area/{id}
// Request body - None
// Response body - None
func (h *EmployeeMainResearchAreaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee main research area by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeMRAUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-main-research-area/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeMainResearchAreaResponse
func (h *EmployeeMainResearchAreaHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee main research area by employeeID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	langCode := middleware.GetLanguageFromContext(r.Context())
	resp, err := h.employeeMRAUC.GetByEmployeeIDAndLanguageCode(r.Context(), int64(employeeID), langCode)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
