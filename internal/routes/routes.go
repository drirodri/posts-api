// routes.go - HTTP routes configuration
package routes

import (
	"net/http"
	"posts-api/internal/handlers"
	"posts-api/internal/middleware"
	"posts-api/internal/services"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all HTTP routes for the application
func SetupRoutes(postHandler *handlers.PostHandler, userService services.UserService) *mux.Router {
	router := mux.NewRouter()

	// Apply CORS middleware to all routes
	router.Use(middleware.CORS)

	// Health check endpoint (no auth required)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "posts-api"}`))
	}).Methods("GET")

	// Root endpoint (no auth required)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Posts API is running!", "version": "1.0.0"}`))
	}).Methods("GET")

	// Public routes (no authentication required)
	public := router.PathPrefix("/api/v1").Subrouter()
	
	// GET /api/v1/posts - Get all posts (public)
	public.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	
	// GET /api/v1/posts/{id} - Get single post (public)
	public.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPost).Methods("GET")
	
	// GET /api/v1/posts/author/{authorId} - Get posts by author (public)
	public.HandleFunc("/posts/author/{authorId:[0-9]+}", postHandler.GetPostsByAuthor).Methods("GET")
	// Protected routes (authentication required)
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.JWTMiddleware(userService))
	
	// POST /api/v1/posts - Create new post (auth required)
	protected.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	
	// PUT /api/v1/posts/{id} - Update post (auth required, only by author)
	protected.HandleFunc("/posts/{id:[0-9]+}", postHandler.UpdatePost).Methods("PUT")
	
	// DELETE /api/v1/posts/{id} - Delete post (auth required, only by author)
	protected.HandleFunc("/posts/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")

	return router
}
