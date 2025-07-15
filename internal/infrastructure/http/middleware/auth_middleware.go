package middleware

import (
	"backend/internal/infrastructure/security"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	UserIDContextKey   string = "userID"
	UserRoleContextKey string = "userRole"
)

// AuthMiddleware creates a middleware that authenticates requests using JWT.

// CreateAuthMiddleware is a factory that creates a new authentication middleware function.
// It takes the necessary dependencies (like a token manager and an error responder)
// and returns a middleware decorator for http.HandlerFunc.
func CreateAuthMiddleware(
	tokenManager *security.TokenManager,
	respondWithError func(w http.ResponseWriter, r *http.Request, err error),
) func(next http.HandlerFunc) http.HandlerFunc {
	// The returned function is the middleware decorator.
	return func(next http.HandlerFunc) http.HandlerFunc {
		// This is the actual http.HandlerFunc that will wrap your handler.
		return func(w http.ResponseWriter, r *http.Request) {
			// 1. Extract JWT from the Authorization header.
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondWithError(w, r, custom_errors.Unauthorized(fmt.Errorf("authorization header missing")))
				return
			}

			// The expected format is "Bearer <token>".
			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
				respondWithError(w, r, custom_errors.Unauthorized(fmt.Errorf("invalid authorization header format")))
				return
			}
			accessToken := headerParts[1]

			// 2. Validate the Access Token.
			claims, err := tokenManager.ValidateAccessToken(accessToken)
			if err != nil {
				// It's assumed that tokenManager.ValidateAccessToken returns a custom_errors.Unauthorized
				// for tokens that are invalid or expired.
				respondWithError(w, r, err)
				return
			}

			// 3. Inject UserID and UserRole into the request context.
			ctx := context.WithValue(r.Context(), UserIDContextKey, claims.UserID)
			ctx = context.WithValue(ctx, UserRoleContextKey, claims.Role)

			// 4. Call the next handler in the chain with the updated context.
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDContextKey).(int64)
	return userID, ok
}

// GetUserRoleFromContext is a helper function to retrieve the UserRole from the request context.
func GetUserRoleFromContext(ctx context.Context) (string, bool) {
	userRole, ok := ctx.Value(UserRoleContextKey).(string)
	return userRole, ok
}
