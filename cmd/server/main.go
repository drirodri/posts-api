// main.go - Entry point for the Posts API server
package main

import (
	"fmt"
	"log"
	"net/http"
	"posts-api/internal/config"
	"posts-api/internal/models"
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

	port := ":8080"

	if portEnv := appConfig.Server.Port; portEnv != "" {
		port = ":" + portEnv
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Posts API is running!")
    })

	fmt.Printf("Server starting on port %s\n", port)
    fmt.Println("Visit http://localhost" + port + " to test the API")

    log.Fatal(http.ListenAndServe(port, nil))
}
