package handlers

import (
	"backend/internal/application/dtos"
	"backend/internal/application/usecases"
	"backend/internal/shared/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func (h *AuthHandler) setAuthCookie(w http.ResponseWriter, resp *dtos.AuthResponse) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    resp.RefreshToken,
		Expires:  time.Now().Add(time.Duration(h.authUsecase.GetRefreshTokenDuration())),
		HttpOnly: true,
		Secure:   h.cookieSecure,
		SameSite: http.SameSiteLaxMode,
		Domain:   h.cookieDomain,
		Path:     "/auth/refresh-token",
	})
}

func (h *AuthHandler) clearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   h.cookieSecure,
		SameSite: http.SameSiteLaxMode,
		Domain:   h.cookieDomain,
		Path:     "/auth/refresh-token",
	})
}

// POST /auth/register
// Request body - dtos.AuthRequest
// Response body - none
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.AuthRequest
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

	h.setAuthCookie(w, resp)

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// POST /auth/refresh-token
// Request body - dtos.RefreshRequest
// Response body - dtos.AuthResponse
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	resfreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		h.clearAuthCookie(w)
		utils.RespondWithJSON(w, r, http.StatusNoContent, nil)
		return
	}

	resp, err := h.authUsecase.RefreshToken(r.Context(), resfreshToken.Value)
	if err != nil {
		h.clearAuthCookie(w)
		utils.RespondWithError(w, r, err)
		return
	}

	h.setAuthCookie(w, resp)

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// POST /auth/logout
// Request body - dtos.RefreshRequest
// Response body - none
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
  fmt.Println("kek_logout")
	var req dtos.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, r, fmt.Errorf("invalid request body to login: %w", err))
		return
	}

  err := h.authUsecase.Logout(r.Context(), req.RefreshToken)
	if err != nil {
		h.clearAuthCookie(w)
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusNoContent, nil)
}
