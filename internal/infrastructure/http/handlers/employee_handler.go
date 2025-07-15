package handlers

import (
	"backend/internal/application/usecases"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/utils"
	"fmt"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	employeeUC usecases.EmployeeUsecase
}

func NewEmployeeHandler(employeeUC usecases.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUC: employeeUC,
	}
}

// DELETE /employee/{id}
// Request body - None
// Response body - None
func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee/{uiq}
// Request body - none
// Response body - dtos.EmployeeResponse
func (h *EmployeeHandler) GetByUID(w http.ResponseWriter, r *http.Request) {
  fmt.Println(utils.GetLanguageFromContext(r.Context()))
	resp, err := h.employeeUC.GetByUniqueID(r.Context(), r.PathValue("uid"))
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}
  fmt.Println(resp)

  utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
