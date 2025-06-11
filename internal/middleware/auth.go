// auth.go - JWT validation middleware
package middleware

import (
	"context"
	"errors"
	"net/http"
	"posts-api/internal/services"
	"posts-api/pkg/utils"
	"strconv"
	"strings"
)

// AuthContextKey is the context key for user data
type AuthContextKey string

const (
	UserIDKey   AuthContextKey = "userID"
	UserDataKey AuthContextKey = "userData"
)

// JWTMiddleware validates JWT tokens using the Users API
func JWTMiddleware(userService services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized,
					"Authorization header required",
					"MISSING_TOKEN",
					"Authorization header with Bearer token is required")
				return
			}

			// Check Bearer token format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized,
					"Invalid authorization header format",
					"INVALID_TOKEN_FORMAT",
					"Authorization header must be in format: Bearer <token>")
				return
			}

			token := parts[1]
			if token == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized,
					"Token is required",
					"EMPTY_TOKEN",
					"JWT token cannot be empty")
				return
			}			// Validate token with Users API
			user, err := userService.ValidateToken(token)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusUnauthorized,
					"Invalid token",
					"INVALID_TOKEN",
					err.Error())
				return
			}

			// Add both user ID and full user data to request context
			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
			ctx = context.WithValue(ctx, UserDataKey, user)
			r = r.WithContext(ctx)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// JWT is the legacy middleware that uses placeholder validation
// Keeping this for backward compatibility during development
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Authorization header required",
				"MISSING_TOKEN",
				"Authorization header with Bearer token is required")
			return
		}

		// Check Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Invalid authorization header format",
				"INVALID_TOKEN_FORMAT",
				"Authorization header must be in format: Bearer <token>")
			return
		}

		token := parts[1]
		if token == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Token is required",
				"EMPTY_TOKEN",
				"JWT token cannot be empty")
			return
		}

		// TODO: Validate JWT token with external Users API
		// For now, we'll use a placeholder implementation
		userID, err := validateTokenPlaceholder(token)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Invalid token",
				"INVALID_TOKEN",
				err.Error())
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// validateTokenPlaceholder is a placeholder for token validation
// In production, this should validate the token with the Users API
func validateTokenPlaceholder(token string) (int64, error) {
	// Placeholder: For testing, we'll accept "test-token" and return user ID 1
	if token == "test-token" {
		return 1, nil
	}

	// Try to parse token as user ID for testing
	if userID, err := strconv.ParseInt(token, 10, 64); err == nil && userID > 0 {
		return userID, nil
	}

	return 0, errors.New("invalid token")
}

// GetUserIDFromContext extracts user ID from request context
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

// GetUserDataFromContext extracts full user data from request context
func GetUserDataFromContext(ctx context.Context) (*services.UserDTO, bool) {
	userData, ok := ctx.Value(UserDataKey).(*services.UserDTO)
	return userData, ok
}