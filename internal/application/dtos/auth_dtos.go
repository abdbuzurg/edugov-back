package dtos

// ---- REQUEST DTOs ----
type RegisterRequest struct {
	Tin      string `json:"tin" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenRequest struct {
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

type MeResponse struct {
	UniqueID string `json:"uniqueID"`
}
