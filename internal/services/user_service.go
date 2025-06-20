// user_service.go - HTTP client to Users API
package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// UserDTO represents user data from the external Users API
type UserDTO struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// UsersAPIResponse represents the direct response from Users API /auth/me endpoint
type UsersAPIResponse struct {
	UserID int64  `json:"userId"`
	Name string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

// UserService handles communication with the external Users API
type UserService interface {
	ValidateToken(token string) (*UserDTO, error)
	GetUserFromToken(token string) (*UserDTO, error) // Alias for ValidateToken for clarity
}

type userService struct {
	usersAPIURL string
	httpClient  *http.Client
}

// NewUserService creates a new user service instance
func NewUserService(usersAPIURL string) UserService {
	return &userService{
		usersAPIURL: usersAPIURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetUserFromToken is an alias for ValidateToken for better clarity in context
func (s *userService) GetUserFromToken(token string) (*UserDTO, error) {
	return s.ValidateToken(token)
}

// ValidateToken validates JWT token with the Users API and returns user info
func (s *userService) ValidateToken(token string) (*UserDTO, error) {
	url := fmt.Sprintf("%s/auth/me", s.usersAPIURL)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to users API: %w", err)
	}
	defer resp.Body.Close()
	
	// Read the response body for debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
		// Log the full response for debugging
	log.Printf("Users API Response - Status: %d, Body: %s", resp.StatusCode, string(body))
	
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid token - users API returned 401")
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("users API returned status %d, body: %s", resp.StatusCode, string(body))
	}
	
	// Parse the actual Users API response format
	var usersAPIResp UsersAPIResponse
	if err := json.Unmarshal(body, &usersAPIResp); err != nil {
		return nil, fmt.Errorf("failed to decode users API response: %w, body: %s", err, string(body))
	}
	
	// Convert UsersAPIResponse to UserDTO
	userDTO := &UserDTO{
		ID:    usersAPIResp.UserID,
		Name:  usersAPIResp.Name, // Use email as name since Users API doesn't provide name
		Email: usersAPIResp.Email,
		Role:  usersAPIResp.Role,
	}
	
	return userDTO, nil
}