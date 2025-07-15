package repositories

import (
	"backend/internal/domain"
	"context"
)

type UserSessionRepository interface {
	CreateSession(ctx context.Context, session *domain.UserSession) (*domain.UserSession, error)
	GetSessionByToken(ctx context.Context, token string) (*domain.UserSession, error)
	DeleteSession(ctx context.Context, id int64) error
	DeleteSessionsByUserID(ctx context.Context, userID int64) error 
	UpdateSession(ctx context.Context, session *domain.UserSession) (*domain.UserSession, error)
}
