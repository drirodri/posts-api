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
type AuthContextKey string
const (
	UserIDKey   AuthContextKey = "userID"
	UserDataKey AuthContextKey = "userData"
)
func JWTMiddleware(userService services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized,
					"Authorization header required",
					"MISSING_TOKEN",
					"Authorization header with Bearer token is required")
				return
			}
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
			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
			ctx = context.WithValue(ctx, UserDataKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Authorization header required",
				"MISSING_TOKEN",
				"Authorization header with Bearer token is required")
			return
		}
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
		userID, err := validateTokenPlaceholder(token)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized,
				"Invalid token",
				"INVALID_TOKEN",
				err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func validateTokenPlaceholder(token string) (int64, error) {
	if token == "test-token" {
		return 1, nil
	}
	if userID, err := strconv.ParseInt(token, 10, 64); err == nil && userID > 0 {
		return userID, nil
	}
	return 0, errors.New("invalid token")
}
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}
func GetUserDataFromContext(ctx context.Context) (*services.UserDTO, bool) {
	userData, ok := ctx.Value(UserDataKey).(*services.UserDTO)
	return userData, ok
}
