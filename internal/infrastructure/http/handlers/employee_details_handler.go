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

type employeeDetailsHandler struct {
	employeeDetailsUC usecases.EmployeeDetailsUsecase
}

func NewEmployeeDetailsHandler(employeeDetailsUC usecases.EmployeeDetailsUsecase) *employeeDetailsHandler {
  return &employeeDetailsHandler{
    employeeDetailsUC: employeeDetailsUC,
  }
}

func (h *employeeDetailsHandler) GetByEmployeeID(w http.ResponseWriter, r *http.Request) {
  employeeIDStr := r.PathValue("employeeID")
  employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
  if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to retrive employee details: %w", err))
    utils.RespondWithError(w, r, err)
    return
  }

  resp, err := h.employeeDetailsUC.GetByEmployeeID(r.Context(), employeeID)
  if err != nil {
    utils.RespondWithError(w, r, err)
    return 
  }

  utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

func (h *employeeDetailsHandler) Update(w http.ResponseWriter, r *http.Request) {
  var req dtos.UpdateFullEmployeeData
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request body to create employee details: %w", err))
    utils.RespondWithError(w, r, err)
    return
  }
  
  resp, err := h.employeeDetailsUC.Update(r.Context(), req.Data)
  if err != nil {
		utils.RespondWithError(w, r, err)
    return 
  }

  utils.RespondWithJSON(w, r, http.StatusOK, resp)
}
