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

type EmployeeParticipationInProfessionalCommunityHandler struct {
	employeeDegreeUC usecases.EmployeeParticipationInProfessionalCommunityUsecase
}

func NewEmployeeParticipationInProfessionalCommunityHandler(employeeDegreeUC usecases.EmployeeParticipationInProfessionalCommunityUsecase) *EmployeeParticipationInProfessionalCommunityHandler {
	return &EmployeeParticipationInProfessionalCommunityHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-participation-in-professional-community
// Request body - dto.CreateEmployeeParticipationInProfessionalCommunityRequest
// Response body - dto.EmployeeParticipationInProfessionalCommunityResponse
func (h *EmployeeParticipationInProfessionalCommunityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeParticipationInProfessionalCommunityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee participation in professional community: %w", err))
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

// PUT /employee-participation-in-professional-community
// Request body - dto.UpdateEmployeeParticipationInProfessionalCommunityRequest
// Response body - dto.EmployeeParticipationInProfessionalCommunityResponse
func (h *EmployeeParticipationInProfessionalCommunityHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeParticipationInProfessionalCommunityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee participation in professional community: %w", err))
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

// DELETE /employee-participation-in-professional-community/{id}
// Request body - None
// Response body - None
func (h *EmployeeParticipationInProfessionalCommunityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee participation in professional community by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-participation-in-professional-community/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeParticipationInProfessionalCommunityResponse
func (h *EmployeeParticipationInProfessionalCommunityHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee participation in professional community by employeeID: %w", err))
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
