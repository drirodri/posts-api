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
type PostHandler struct {
	postService services.PostService
	validator   *validator.Validate
}
func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
		validator:   validator.New(),
	}
}
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
func (h *PostHandler) addAuthorInfoIfOwner(r *http.Request, response *dto.PostResponse) {
	if currentUser, ok := middleware.GetUserDataFromContext(r.Context()); ok {
		if currentUser.ID == response.AuthorID {
			response.Author = h.convertUserDTOToUserData(currentUser)
		}
	}
}
func (h *PostHandler) addAuthorInfoToList(r *http.Request, posts []dto.PostResponse) {
	if currentUser, ok := middleware.GetUserDataFromContext(r.Context()); ok {
		userData := h.convertUserDTOToUserData(currentUser)
		for i := range posts {
			if posts[i].AuthorID == currentUser.ID {
				posts[i].Author = userData
			}
		}
	}
}
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var createReq dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid request body",
			"INVALID_JSON",
			err.Error())
		return
	}
	if err := h.validator.Struct(&createReq); err != nil {
		validationErrors := h.extractValidationErrors(err)
		utils.WriteValidationErrorResponse(w, validationErrors)
		return
	}
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
	
	token := h.extractTokenFromRequest(r)
	if token == "" {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"Authorization token required",
			"MISSING_TOKEN",
			"Bearer token is required")
		return
	}
	
	post, err := h.postService.CreatePost(&createReq, userID, token)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}
	utils.WriteSuccessResponse(w, http.StatusCreated, "Post created successfully", post)
}
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
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
	post, err := h.postService.GetPostByID(id)
	if err != nil {
		if err.Error() == "post not found" {
			utils.WriteNotFoundResponse(w, "Post")
			return
		}
		utils.WriteInternalErrorResponse(w, err)
		return
	}
	h.addAuthorInfoIfOwner(r, post)
	utils.WriteSuccessResponse(w, http.StatusOK, "Post retrieved successfully", post)
}
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
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
	posts, err := h.postService.GetAllPosts(page, pageSize)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}
	h.addAuthorInfoToList(r, posts.Posts)
	utils.WriteSuccessResponse(w, http.StatusOK, "Posts retrieved successfully", posts)
}
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
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
	var updateReq dto.UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Invalid request body",
			"INVALID_JSON",
			err.Error())
		return
	}
	if err := h.validator.Struct(&updateReq); err != nil {
		validationErrors := h.extractValidationErrors(err)
		utils.WriteValidationErrorResponse(w, validationErrors)
		return
	}
	if !updateReq.HasChanges() {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"No changes provided",
			"NO_CHANGES",
			"At least one field must be provided for update")
		return
	}
	if err := updateReq.Validate(); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest,
			"Validation failed",
			"BUSINESS_VALIDATION_ERROR",
			err.Error())
		return
	}
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"User not authenticated",
			"AUTHENTICATION_ERROR",
			"User ID not found in request context")
		return
	}
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
	h.addAuthorInfoIfOwner(r, post)
	utils.WriteSuccessResponse(w, http.StatusOK, "Post updated successfully", post)
}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
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
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteErrorResponse(w, http.StatusUnauthorized,
			"User not authenticated",
			"AUTHENTICATION_ERROR",
			"User ID not found in request context")
		return
	}
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
func (h *PostHandler) GetPostsByAuthor(w http.ResponseWriter, r *http.Request) {
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
	posts, err := h.postService.GetPostsByAuthor(authorID, page, pageSize)
	if err != nil {
		utils.WriteInternalErrorResponse(w, err)
		return
	}
	h.addAuthorInfoToList(r, posts.Posts)
	utils.WriteSuccessResponse(w, http.StatusOK, "Author posts retrieved successfully", posts)
}
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
