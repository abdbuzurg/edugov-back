package postgres

import (
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
)

type pgUserRepository struct {
	store   *Store
	queries *sqlc.Queries
}

func NewPGUserRepository(store *Store) repositories.UserRepository {
	return &pgUserRepository{
		store:   store,
		queries: store.Queries,
	}
}

func NewPgUserRepositoryWithQueries(q *sqlc.Queries) repositories.UserRepository {
	return &pgUserRepository{
		queries: q,
	}
}

func (r *pgUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {

		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to create user: %w", err))
	}

	user.ID = createdUser.ID
	user.CreatedAt = createdUser.CreatedAt.Time
	user.UpdatedAt = createdUser.UpdatedAt.Time

	return user, nil
}

func (r *pgUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	userResult, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive user by email(%s): %w", email, err))
	}

	return &domain.User{
		ID:           userResult.ID,
		Email:        userResult.Email,
		PasswordHash: userResult.PasswordHash,
		CreatedAt:    userResult.CreatedAt.Time,
		UpdatedAt:    userResult.UpdatedAt.Time,
	}, nil
}

func (r *pgUserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	userResult, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive user: %w", err))
	}

	return &domain.User{
		ID:           userResult.ID,
		Email:        userResult.Email,
		PasswordHash: userResult.PasswordHash,
		CreatedAt:    userResult.CreatedAt.Time,
		UpdatedAt:    userResult.UpdatedAt.Time,
	}, nil
}
