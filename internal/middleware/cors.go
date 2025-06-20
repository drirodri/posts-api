package middleware
import (
	"net/http"
	"github.com/gorilla/handlers"
)
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	MaxAge         int
}
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins: []string{"*"}, // Allow all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin", "X-CSRF-Token"},
		MaxAge:         86400, // 24 hours
	}
}
func NewCORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins(config.AllowedOrigins),
		handlers.AllowedMethods(config.AllowedMethods),
		handlers.AllowedHeaders(config.AllowedHeaders),
		handlers.MaxAge(config.MaxAge),
	)
}
