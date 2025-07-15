package usecases

import (
	"backend/internal/application/dtos"
	"backend/internal/application/repositories"
	"backend/internal/domain"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/infrastructure/security"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/utils"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *dtos.AuthRequest) error
	Login(ctx context.Context, req *dtos.AuthRequest) (*dtos.AuthResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*dtos.AuthResponse, error)
	Logout(ctx context.Context, refreshToken string) error
	GetRefreshTokenDuration() time.Duration
}

type authUsecase struct {
	userRepo        repositories.UserRepository
	userSessionRepo repositories.UserSessionRepository
	employeeRepo    repositories.EmployeeRepository
	store           *postgres.Store
	tokenManager    *security.TokenManager
	validator       *validator.Validate
}

func NewAuthUsecase(
	userRepo repositories.UserRepository,
	userSessionRepo repositories.UserSessionRepository,
	employeeRepo repositories.EmployeeRepository,
	store *postgres.Store,
	tokenManager *security.TokenManager,
	validator *validator.Validate,
) AuthUsecase {
	return &authUsecase{
		userRepo:        userRepo,
		userSessionRepo: userSessionRepo,
		employeeRepo:    employeeRepo,
		store:           store,
		tokenManager:    tokenManager,
		validator:       validator,
	}
}

var clientErrorMessages = map[string]map[string]string{
	"registerSameEmail": {
		"en": "User with the same email '%s' already exists.",
		"ru": "Пользователь с таким же адресом электронной почты '%s' уже существует.",
		"tg": "Истифодабаранда бо ҳамин почтаи электронӣ '%s' аллакай вуҷуд дорад.",
	},
	"invalidCredentials": {
		"en": "Invalid credentials",
		"ru": "Неверные учётные данные",
		"tg": "Маълумоти воридшавӣ нодуруст аст",
	},
}

func (uc *authUsecase) GetRefreshTokenDuration() time.Duration {
	return uc.tokenManager.GetRefreshTokenDuration()
}

func (uc *authUsecase) generateAndStoreTokens(ctx context.Context, user *domain.User) (*dtos.AuthResponse, error) {
	accessToken, _, err := uc.tokenManager.GenerateAccessToken(user.ID, user.UserType)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to generate access token: %w", err))
	}

	refreshToken, refreshTokenExp, err := uc.tokenManager.GenereateRefreshToken(user.ID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to generate refresh token: %w", err))
	}

	session := &domain.UserSession{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenExp,
	}

	_, err = uc.userSessionRepo.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	var uid string
	if user.UserType == "employee" {
		employee, err := uc.employeeRepo.GetByID(ctx, user.EntityID)
		if err != nil {
			return nil, err
		}

		uid = employee.UniqueID
	}

	return &dtos.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		UID:          uid,
		UserRole:     user.UserType,
	}, nil
}

func (uc *authUsecase) generateNumericUniqueID() (string, error) {
	// Initialize a strings.Builder for efficient string concatenation.
	// We know the final string length will be 9 (8 digits + 1 hyphen).
	var builder strings.Builder
	builder.Grow(9)

	// Loop 8 times to generate each of the 8 required digits.
	for i := 0; i < 8; i++ {
		// After the first 4 digits, append a hyphen to match the XXXX-XXXX format.
		if i == 4 {
			builder.WriteByte('-')
		}

		// Generate a random digit (0-9).
		// rand.Int(rand.Reader, big.NewInt(10)) generates a cryptographically
		// secure random integer in the range [0, 10), which means 0 through 9.
		// This method ensures a uniform distribution and avoids any statistical bias.
		digitBig, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			// If there's an error during random number generation, return an error.
			return "", fmt.Errorf("failed to generate random digit: %w", err)
		}

		// Convert the *big.Int result to an int64 and then to a string,
		// appending it to our string builder.
		builder.WriteString(fmt.Sprintf("%d", digitBig.Int64()))
	}

	// Return the final generated string and no error (nil).
	return builder.String(), nil
}

func (uc *authUsecase) Register(ctx context.Context, req *dtos.AuthRequest) error {
	if err := uc.validator.Struct(req); err != nil {
		return custom_errors.BadRequest(fmt.Errorf("invalid registration input: %w", err))
	}

	err := uc.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		txUserRepo := postgres.NewPgUserRepositoryWithQueries(q)
		txEmployeeRepo := postgres.NewPgEmployeeRepositoryWithQuery(q)

		isUniqueExists := true
		var uniqueID string
		var employeeResult *domain.Employee
		var err error
		for isUniqueExists {
			uniqueID, _ = uc.generateNumericUniqueID()
			employeeResult, err = txEmployeeRepo.Create(ctx, &domain.Employee{
				UniqueID: uniqueID,
			})
			if err != nil && !custom_errors.IsUniqueConstraintError(err) {
				return err
			}

      if employeeResult != nil {
        isUniqueExists = false
      }

		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return custom_errors.InternalServerError(fmt.Errorf("failed to hash password: %w", err))
		}

		_, err = txUserRepo.CreateUser(ctx, &domain.User{
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
			UserType:     "employee",
			EntityID:     employeeResult.ID,
		})
		if err != nil && !custom_errors.IsUniqueConstraintError(err) {
			return err
		} else if custom_errors.IsUniqueConstraintError(err) {
			lang := utils.GetLanguageFromContext(ctx)
			return custom_errors.BadRequest(fmt.Errorf(clientErrorMessages["registerSameEmail"][lang], req.Email))
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *authUsecase) Login(ctx context.Context, req *dtos.AuthRequest) (*dtos.AuthResponse, error) {
	if err := uc.validator.Struct(req); err != nil {
		return nil, custom_errors.BadRequest(fmt.Errorf("invalid credentials: %w", err))
	}

	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if custom_errors.IsNotFound(err) {
			lang := utils.GetLanguageFromContext(ctx)
			return nil, custom_errors.BadRequest(fmt.Errorf(clientErrorMessages["invalidCredentials"][lang]))
		}

		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive user by email(%s): %w", req.Email, err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		lang := utils.GetLanguageFromContext(ctx)
		return nil, custom_errors.BadRequest(fmt.Errorf(clientErrorMessages["invalidCredentials"][lang]))
	}

	return uc.generateAndStoreTokens(ctx, user)
}

func (uc *authUsecase) RefreshToken(ctx context.Context, refreshToken string) (*dtos.AuthResponse, error) {
	refreshClaims, err := uc.tokenManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, custom_errors.Unauthorized(fmt.Errorf("invalid or expired refresh token: %w", err))
	}

	session, err := uc.userSessionRepo.GetSessionByToken(ctx, refreshToken)
	if err != nil {
		if custom_errors.IsNotFound(err) {
			return nil, custom_errors.Unauthorized(fmt.Errorf("refresh token not found or revoked"))
		}

		return nil, err
	}

	if session.UserID != refreshClaims.UserID {
		return nil, custom_errors.Unauthorized(fmt.Errorf("refresh token user mismatch"))
	}

	err = uc.userSessionRepo.DeleteSession(ctx, session.ID)
	if err != nil {
		return nil, custom_errors.InternalServerError(fmt.Errorf("failed to revoke old refresh token: %w", err))
	}

	user, err := uc.userRepo.GetByID(ctx, session.UserID)
	if err != nil {
		if custom_errors.IsNotFound(err) {
			return nil, custom_errors.Unauthorized(fmt.Errorf("user not found for refresh token"))
		}

		return nil, err
	}

	return uc.generateAndStoreTokens(ctx, user)
}

func (uc *authUsecase) removeSessionOnError(ctx context.Context, userID int64) {
  _ = uc.userSessionRepo.DeleteSessionsByUserID(ctx, userID)
}

func (uc *authUsecase) Logout(ctx context.Context, refreshToken string) error {
	refreshClaims, err := uc.tokenManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return custom_errors.Unauthorized(fmt.Errorf("invalid refresh token for logout: %w", err))
	}

	session, err := uc.userSessionRepo.GetSessionByToken(ctx, refreshToken)
	if err != nil {
		if custom_errors.IsNotFound(err) {
			return custom_errors.NotFound(fmt.Errorf("session for token not found")) // Already logged out or never existed
		}

    uc.removeSessionOnError(ctx, refreshClaims.UserID)
		return err
	}

	if session.UserID != refreshClaims.UserID {
    uc.removeSessionOnError(ctx, refreshClaims.UserID)
		return custom_errors.Unauthorized(fmt.Errorf("refresh token user ID mismatch during logout"))
	}

	err = uc.userSessionRepo.DeleteSession(ctx, session.ID)
	if err != nil {
    uc.removeSessionOnError(ctx, refreshClaims.UserID)
		return custom_errors.InternalServerError(fmt.Errorf("failed to delete user session on logout: %w", err))
	}

	return nil
}
