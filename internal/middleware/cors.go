// cors.go - CORS middleware using Gorilla Handlers
package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	MaxAge         int
}

// DefaultCORSConfig returns a default CORS configuration for development
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins: []string{"*"}, // Allow all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin", "X-CSRF-Token"},
		MaxAge:         86400, // 24 hours
	}
}

// NewCORSMiddleware creates a new CORS middleware using Gorilla Handlers
func NewCORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins(config.AllowedOrigins),
		handlers.AllowedMethods(config.AllowedMethods),
		handlers.AllowedHeaders(config.AllowedHeaders),
		handlers.MaxAge(config.MaxAge),
	)
}