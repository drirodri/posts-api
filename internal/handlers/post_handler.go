// post_handler.go - HTTP handlers for post endpoints
package handlers

import (
	"encoding/json"
	"net/http"
	"posts-api/internal/dto"
	"posts-api/internal/middleware"
	"posts-api/internal/services"
	"posts-api/pkg/utils"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// PostHandler handles HTTP requests for posts
type PostHandler struct {
	postService services.PostService
	validator   *validator.Validate
}

// NewPostHandler creates a new PostHandler
func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
		validator:   validator.New(),
	}
}

// convertUserDTOToUserData converts UserDTO to dto.UserData
func (h *PostHandler) convertUserDTOToUserData(user *services.UserDTO) *dto.UserData {
	if user == nil {
		return nil
	}
	return &dto.UserData{
		ID:       user.ID,
		Username: user.Name,
		Email:    user.Email,
	}
}

// addAuthorInfoIfOwner adds author information to the post response if the current user is the author
func (h *PostHandler) addAuthorInfoIfOwner(r *http.Request, response *dto.PostResponse) {
	// Get current user from context (if authenticated)
	if currentUser, ok := middleware.GetUserDataFromContext(r.Context()); ok {
		// If the current user is the author of this post, add author info
		if currentUser.ID == response.AuthorID {
			response.Author = h.convertUserDTOToUserData(currentUser)
		}
	}
}

// addAuthorInfoToList adds author information to posts authored by the current user
func (h *PostHandler) addAuthorInfoToList(r *http.Request, posts []dto.PostResponse) {
	// Get current user from context (if authenticated)
	if currentUser, ok := middleware.GetUserDataFromContext(r.Context()); ok {
		userData := h.convertUserDTOToUserData(currentUser)
		// Add author info to posts authored by the current user
		for i := range posts {
			if posts[i].AuthorID == currentUser.ID {
				posts[i].Author = userData
			}
		}
	}
}

// CreatePost handles POST /posts
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var createReq dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid request body",
			"INVALID_JSON",
			err.Error())
		return
	}

	// Validate request
	if err := h.validator.Struct(&createReq); err != nil {
		validationErrors := h.extractValidationErrors(err)
		utils.WriteValidationErrorResponse(w, validationErrors)
		return
	}

	// Business validation
	if err := createReq.Validate(); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Validation failed",
			"BUSINESS_VALIDATION_ERROR",
			err.Error())
		return
	}	// Extract user ID from JWT token (provided by auth middleware)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"User not authenticated",
			"AUTHENTICATION_ERROR",
			"User ID not found in request context")
		return
	}
	
	// Extract token for user API calls
	token := h.extractTokenFromRequest(r)
	if token == "" {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"Authorization token required",
			"MISSING_TOKEN",
			"Bearer token is required")
		return
	}
	
	// Create post
	post, err := h.postService.CreatePost(&createReq, userID, token)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, "Post created successfully", post)
}

// GetPost handles GET /posts/{id}
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Post ID is required",
			"MISSING_PARAMETER",
			"Post ID must be provided in the URL")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid post ID",
			"INVALID_PARAMETER",
			"Post ID must be a valid number")
		return
	}
	// Get post
	post, err := h.postService.GetPostByID(id)
	if err != nil {
		if err.Error() == "post not found" {
			utils.WriteNotFoundResponse(w, "Post")
			return
		}
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	// Add author information if the current user is the author
	h.addAuthorInfoIfOwner(r, post)

	utils.WriteSuccessResponse(w, http.StatusOK, "Post retrieved successfully", post)
}

// GetAllPosts handles GET /posts
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page := 1
	pageSize := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}
	// Get posts
	posts, err := h.postService.GetAllPosts(page, pageSize)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	// Add author information to posts authored by the current user
	h.addAuthorInfoToList(r, posts.Posts)

	utils.WriteSuccessResponse(w, http.StatusOK, "Posts retrieved successfully", posts)
}

// UpdatePost handles PUT /posts/{id}
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Post ID is required",
			"MISSING_PARAMETER",
			"Post ID must be provided in the URL")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid post ID",
			"INVALID_PARAMETER",
			"Post ID must be a valid number")
		return
	}

	// Parse request body
	var updateReq dto.UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid request body",
			"INVALID_JSON",
			err.Error())
		return
	}

	// Validate request
	if err := h.validator.Struct(&updateReq); err != nil {
		validationErrors := h.extractValidationErrors(err)
		utils.WriteValidationErrorResponse(w, validationErrors)
		return
	}

	// Check if there are any changes
	if !updateReq.HasChanges() {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"No changes provided",
			"NO_CHANGES",
			"At least one field must be provided for update")
		return
	}

	// Business validation
	if err := updateReq.Validate(); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Validation failed",
			"BUSINESS_VALIDATION_ERROR",
			err.Error())
		return
	}

	// Extract user ID from JWT token (provided by auth middleware)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"User not authenticated",
			"AUTHENTICATION_ERROR",
			"User ID not found in request context")
		return
	}
	// Update post
	post, err := h.postService.UpdatePost(id, &updateReq, userID)
	if err != nil {
		if err.Error() == "post not found" {
			utils.WriteNotFoundResponse(w, "Post")
			return
		}
		if err.Error() == "unauthorized: only the author can update this post" {
			utils.WriteErrorResponse(w, http.StatusForbidden,
				"Access denied",
				"FORBIDDEN",
				"You can only update your own posts")
			return
		}
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	// Add author information since the current user is the author
	h.addAuthorInfoIfOwner(r, post)

	utils.WriteSuccessResponse(w, http.StatusOK, "Post updated successfully", post)
}

// DeletePost handles DELETE /posts/{id}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Post ID is required",
			"MISSING_PARAMETER",
			"Post ID must be provided in the URL")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid post ID",
			"INVALID_PARAMETER",
			"Post ID must be a valid number")
		return
	}

	// Extract user ID from JWT token (provided by auth middleware)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"User not authenticated",
			"AUTHENTICATION_ERROR",
			"User ID not found in request context")
		return
	}

	// Delete post
	err = h.postService.DeletePost(id, userID)
	if err != nil {		if err.Error() == "post not found" {
			utils.WriteNotFoundResponse(w, "Post")
			return
		}
		if err.Error() == "unauthorized: only the author can delete this post" {
			utils.WriteErrorResponse(w, http.StatusForbidden,
				"Access denied",
				"FORBIDDEN",
				"You can only delete your own posts")
			return
		}
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Post deleted successfully", nil)
}

// GetPostsByAuthor handles GET /posts/author/{authorId}
func (h *PostHandler) GetPostsByAuthor(w http.ResponseWriter, r *http.Request) {
	// Extract author ID from URL
	vars := mux.Vars(r)
	authorIDStr, exists := vars["authorId"]
	if !exists {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Author ID is required",
			"MISSING_PARAMETER",
			"Author ID must be provided in the URL")
		return
	}

	authorID, err := strconv.ParseInt(authorIDStr, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid author ID",
			"INVALID_PARAMETER",
			"Author ID must be a valid number")
		return
	}

	// Parse query parameters for pagination
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page := 1
	pageSize := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}
	// Get posts by author
	posts, err := h.postService.GetPostsByAuthor(authorID, page, pageSize)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}

	// Add author information to posts authored by the current user
	h.addAuthorInfoToList(r, posts.Posts)

	utils.WriteSuccessResponse(w, http.StatusOK, "Author posts retrieved successfully", posts)
}

// extractValidationErrors converts validator errors to our custom format
func (h *PostHandler) extractValidationErrors(err error) []utils.ValidationError {
	var validationErrors []utils.ValidationError

	if validatorErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validatorErrs {
			validationErrors = append(validationErrors, utils.ValidationError{
				Field:   e.Field(),
				Message: h.getValidationMessage(e),
			})
		}
	}

	return validationErrors
}

// getValidationMessage returns a user-friendly validation message
func (h *PostHandler) getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must be at most " + e.Param() + " characters"
	default:
		return e.Field() + " is invalid"
	}
}

// extractTokenFromRequest extracts the Bearer token from Authorization header
func (h *PostHandler) extractTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	
	return parts[1]
}