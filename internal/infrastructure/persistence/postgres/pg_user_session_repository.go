package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type pgUserSessionRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPGUserSessionRepository(store *Store) repositories.UserSessionRepository {
	return &pgUserSessionRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgUserSessionWithQuery(q *sqlc.Queries) repositories.UserSessionRepository {
	return &pgUserSessionRepository{
		queries: q,
	}
}

func (r *pgUserSessionRepository) CreateSession(ctx context.Context, session *domain.UserSession) (*domain.UserSession, error) {
	createdUserSession, err := r.queries.CreateUserSession(ctx, sqlc.CreateUserSessionParams{
		UserID:       session.UserID,
		RefreshToken: session.RefreshToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  session.ExpiresAt,
			Valid: !session.ExpiresAt.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create user session: %w", err))
	}

	session.ID = createdUserSession.ID
	session.CreatedAt = createdUserSession.CreatedAt.Time
	session.UpdatedAt = createdUserSession.UpdatedAt.Time

	return session, nil
}

func (r *pgUserSessionRepository) UpdateSession(ctx context.Context, session *domain.UserSession) (*domain.UserSession, error) {
	updatedUserSession, err := r.queries.UpdateUserSession(ctx, sqlc.UpdateUserSessionParams{
		ID:           session.ID,
		RefreshToken: session.RefreshToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  session.ExpiresAt,
			Valid: !session.ExpiresAt.IsZero(),
		},
	})
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to update user session: %w", err))
	}

	session.CreatedAt = updatedUserSession.CreatedAt.Time
	session.UpdatedAt = updatedUserSession.UpdatedAt.Time

	return session, nil
}

func (r *pgUserSessionRepository) DeleteSession(ctx context.Context, id int64) error {
	err := r.queries.DeleteUserSessionByID(ctx, id)
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete user session by ID(%d): %w", id, err))
	}

	return nil
}

func (r *pgUserSessionRepository) DeleteSessionsByUserID(ctx context.Context, userID int64) error {
	err := r.queries.DeleteUserSessionByUserID(ctx, userID)
	if err != nil {
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete user session by user_id (%d): %w", userID, err))
	}

	return nil
}

func (r *pgUserSessionRepository) GetSessionByToken(ctx context.Context, token string) (*domain.UserSession, error) {
	userSessionResult, err := r.queries.GetUserSessionByToken(ctx, token)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive user session by refresh_token (%s): %w", token, err))
	}

	return &domain.UserSession{
		ID:           userSessionResult.ID,
		UserID:       userSessionResult.UserID,
		RefreshToken: userSessionResult.RefreshToken,
		ExpiresAt:    userSessionResult.ExpiresAt.Time,
		CreatedAt:    userSessionResult.CreatedAt.Time,
		UpdatedAt:    userSessionResult.UpdatedAt.Time,
	}, nil
}
