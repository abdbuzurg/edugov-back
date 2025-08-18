package handlers

import (
	"backend/internal/application/usecases"
	"backend/internal/shared/utils"
	"net/http"
)

type institutionHandler struct {
	institutionUC usecases.InstitutionUsecase
}

func NewInstitutionHandler(institutionUsecase usecases.InstitutionUsecase) *institutionHandler {
	return &institutionHandler{
		institutionUC: institutionUsecase,
	}
}

// GET /institution/all
func (h *institutionHandler) GetAllInstitutions(w http.ResponseWriter, r *http.Request) {
	allInstitutions, err := h.institutionUC.GetAllInstitutions(r.Context())
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, allInstitutions)
}
