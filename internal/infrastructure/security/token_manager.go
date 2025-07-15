package security

import (
	"backend/internal/shared/custom_errors"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	UserID int64  `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
  UserID int64 `json:"userID"`
  jwt.RegisteredClaims
}

type TokenManager struct {
  accessTokenSecret []byte
  refreshTokenSecret []byte
  accessTokenDuration time.Duration
  refreshTokenDuration time.Duration
}

func NewTokenManager(
  accessTokenSecret []byte,
  refreshTokenSecret []byte,
  accessTokenDuration time.Duration,
  refreshTokenDuration time.Duration,
) *TokenManager {
  return &TokenManager{
    accessTokenSecret: accessTokenSecret,
    refreshTokenSecret: refreshTokenSecret,
    accessTokenDuration: accessTokenDuration,
    refreshTokenDuration: refreshTokenDuration,
  }
}

func (tm *TokenManager) GetRefreshTokenDuration() time.Duration {
  return tm.refreshTokenDuration
}

func(tm *TokenManager) GenerateAccessToken(userID int64, role string) (string, time.Time, error) {
  expirationTime := time.Now().Add(tm.accessTokenDuration)
  claims := &AccessClaims{
    UserID: userID,
    Role: role,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(expirationTime),
      IssuedAt: jwt.NewNumericDate(time.Now()),
      NotBefore: jwt.NewNumericDate(time.Now()),
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(tm.accessTokenSecret)
  return tokenString, expirationTime, err
}

func(tm *TokenManager) GenereateRefreshToken(userID int64) (string, time.Time, error) {
  expirationTime := time.Now().Add(tm.refreshTokenDuration)
  claims := &RefreshClaims{
    UserID: userID,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(expirationTime),
      IssuedAt: jwt.NewNumericDate(time.Now()),
      NotBefore: jwt.NewNumericDate(time.Now()),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(tm.accessTokenSecret)
  return tokenString, expirationTime, err
}

func(tm *TokenManager) ValidateAccessToken(tokenString string) (*AccessClaims, error) {
  claims := &AccessClaims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
    } 

    return tm.accessTokenSecret, nil
  })
  if err != nil {
    if errors.Is(err, jwt.ErrTokenExpired) {
      return nil, custom_errors.Unauthorized(fmt.Errorf("access token expired: %w", err))
    }

    return nil, custom_errors.Unauthorized(fmt.Errorf("invalid access token: %w", err))
  }

  if !token.Valid {
    return nil, custom_errors.Unauthorized(fmt.Errorf("invalid access token"))
  }

  return claims, nil
}

func (tm *TokenManager) ValidateRefreshToken(tokenString string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return tm.refreshTokenSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, custom_errors.Unauthorized(fmt.Errorf("refresh token is expired: %w", err))
		}
		return nil, custom_errors.Unauthorized(fmt.Errorf("invalid refresh token: %w", err))
	}

	if !token.Valid {
		return nil, custom_errors.Unauthorized(fmt.Errorf("invalid refresh token"))
	}

	return claims, nil
}
