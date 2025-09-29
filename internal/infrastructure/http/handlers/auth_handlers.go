package handlers

import (
	"backend/internal/application/dtos"
	"backend/internal/application/usecases"
	"backend/internal/shared/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	authUsecase  usecases.AuthUsecase
	cookieDomain string
	cookieSecure bool
}

func NewAuthHandler(
	authUsecase usecases.AuthUsecase,
	cookieDomain string,
	cookieSecure bool,
) *AuthHandler {
	return &AuthHandler{
		authUsecase:  authUsecase,
		cookieDomain: cookieDomain,
		cookieSecure: cookieSecure,
	}
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	resp, err := h.authUsecase.Me(r.Context())
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// POST /auth/register
// Request body - dtos.AuthRequest
// Response body - none
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, r, fmt.Errorf("invalid request body to register: %w", err))
		return
	}

	err := h.authUsecase.Register(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusCreated, nil)
}

// POST /auth/login
// Request body - dtos.AuthRequest
// Response body - dtos.AuthResponse
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, r, fmt.Errorf("invalid request body to login: %w", err))
		return
	}

	resp, err := h.authUsecase.Login(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// POST /auth/refresh-token
// Request body - dtos.RefreshRequest
// Response body - dtos.AuthResponse
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req dtos.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, r, fmt.Errorf("invalid request body to refresh token: %w", err))
		return
	}

	resp, err := h.authUsecase.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// POST /auth/logout
// Request body - dtos.RefreshRequest
// Response body - none
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req dtos.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, r, fmt.Errorf("invalid request body to login: %w", err))
		return
	}

	err := h.authUsecase.Logout(r.Context(), req.RefreshToken)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusNoContent, nil)
}
