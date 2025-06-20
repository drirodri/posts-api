package routes
import (
	"net/http"
	"posts-api/internal/handlers"
	"posts-api/internal/middleware"
	"posts-api/internal/services"
	"github.com/gorilla/mux"
)
func SetupRoutes(postHandler *handlers.PostHandler, userService services.UserService) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "posts-api"}`))
	}).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Posts API is running!", "version": "1.0.0"}`))
	}).Methods("GET")
	public := router.PathPrefix("/api/v1").Subrouter()
	
	public.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	
	public.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPost).Methods("GET")
	
	public.HandleFunc("/posts/author/{authorId:[0-9]+}", postHandler.GetPostsByAuthor).Methods("GET")
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.JWTMiddleware(userService))
	
	protected.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	
	protected.HandleFunc("/posts/{id:[0-9]+}", postHandler.UpdatePost).Methods("PUT")
	protected.HandleFunc("/posts/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")
	corsConfig := middleware.DefaultCORSConfig()
	corsMiddleware := middleware.NewCORSMiddleware(corsConfig)
	
	return corsMiddleware(router)
}
