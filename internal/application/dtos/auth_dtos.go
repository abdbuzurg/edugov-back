package dtos

// ---- REQUEST DTOs ----
type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
  RefreshToken string `json:"refreshToken" validate:"required"`
}

// ---- RESPONSE DTOs ----
type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType,omitempty"`
	UID          string `json:"uid"`
	UserRole     string `json:"userRole"`
}
