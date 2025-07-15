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

type EmployeeRefresherCourseHandler struct {
	employeeDegreeUC usecases.EmployeeRefresherCourseUsecase
}

func NewEmployeeRefresherCourseHandler(employeeDegreeUC usecases.EmployeeRefresherCourseUsecase) *EmployeeRefresherCourseHandler {
	return &EmployeeRefresherCourseHandler{
		employeeDegreeUC: employeeDegreeUC,
	}
}

// POST /employee-refresher-course
// Request body - dto.CreateEmployeeRefresherCourseRequest
// Response body - dto.EmployeeRefresherCourseResponse
func (h *EmployeeRefresherCourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateEmployeeRefresherCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee refresher course: %w", err))
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

// PUT /employee-refresher-course
// Request body - dto.UpdateEmployeeRefresherCourseRequest
// Response body - dto.EmployeeRefresherCourseResponse
func (h *EmployeeRefresherCourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dtos.UpdateEmployeeRefresherCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to update employee refresher course: %w", err))
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

// DELETE /employee-refresher-course/{id}
// Request body - None
// Response body - None
func (h *EmployeeRefresherCourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee refresher course by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeDegreeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee-refresher-course/employee-{employeeID}/
// Request body - none
// Response body - []dtos.EmployeeRefresherCourseResponse
func (h *EmployeeRefresherCourseHandler) GetByEmployeeIDAndLanguageCode(w http.ResponseWriter, r *http.Request) {
	employeeID, err := strconv.Atoi(r.PathValue("employeeID"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to retrive employee refresher course by employeeID: %w", err))
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
