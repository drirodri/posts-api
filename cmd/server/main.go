// main.go - Entry point for the Posts API server
package main

import (
	"fmt"
	"log"
	"net/http"
	"posts-api/internal/config"
	"posts-api/internal/handlers"
	"posts-api/internal/models"
	"posts-api/internal/repository"
	"posts-api/internal/routes"
	"posts-api/internal/services"
)

func main() {
    fmt.Println("Server starting...")

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	fmt.Println("Initializing database connection...")
	config.InitDatabase(appConfig.Database)
	defer config.CloseDatabase()

	fmt.Println("Database connection established successfully\nMigrating database...")

	err = config.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("Database migration completed successfully")

	// Initialize dependencies
	fmt.Println("Initializing services...")
	
	// Create repository
	postRepo := repository.NewPostRepository(config.DB)
		// Create user service (for external Users API integration)
	usersAPIURL := appConfig.Server.UsersAPIURL
	if usersAPIURL == "" {
		log.Fatal("USERS_API_URL environment variable is required")
	}
	userService := services.NewUserService(usersAPIURL)
	
	// Create post service
	postService := services.NewPostService(postRepo, userService)
	
	// Create handlers
	postHandler := handlers.NewPostHandler(postService)
		// Setup routes
	router := routes.SetupRoutes(postHandler, userService)

	// Configure server
	port := ":8080"
	if portEnv := appConfig.Server.Port; portEnv != "" {
		port = ":" + portEnv
	}

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Visit http://localhost" + port + " to test the API")
	fmt.Println("API endpoints:")
	fmt.Println("  GET    / - Root endpoint")
	fmt.Println("  GET    /health - Health check")
	fmt.Println("  GET    /api/v1/posts - Get all posts")
	fmt.Println("  GET    /api/v1/posts/{id} - Get single post")
	fmt.Println("  GET    /api/v1/posts/author/{authorId} - Get posts by author")
	fmt.Println("  POST   /api/v1/posts - Create post (auth required)")
	fmt.Println("  PUT    /api/v1/posts/{id} - Update post (auth required)")
	fmt.Println("  DELETE /api/v1/posts/{id} - Delete post (auth required)")

    log.Fatal(http.ListenAndServe(port, router))
}
