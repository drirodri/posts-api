// post_service.go - Post business logic
package services

import (
	"errors"
	"fmt"
	"math"
	"posts-api/internal/dto"
	"posts-api/internal/repository"

	"gorm.io/gorm"
)

// PostService defines the business logic interface for posts
type PostService interface {
	CreatePost(req *dto.CreatePostRequest, authorID int64, token string) (*dto.PostResponse, error)
	GetPostByID(id int64) (*dto.PostResponse, error)
	GetAllPosts(page, pageSize int) (*dto.PostListResponse, error)
	UpdatePost(id int64, req *dto.UpdatePostRequest, authorID int64) (*dto.PostResponse, error)
	DeletePost(id int64, authorID int64) error
	GetPostsByAuthor(authorID int64, page, pageSize int) (*dto.PostListResponse, error)
}

type postService struct {
	postRepo    repository.PostRepository
	userService UserService
}

// NewPostService creates a new post service instance
func NewPostService(postRepo repository.PostRepository, userService UserService) PostService {
	return &postService{
		postRepo:    postRepo,
		userService: userService,
	}
}

// CreatePost creates a new post after validating the author
func (s *postService) CreatePost(req *dto.CreatePostRequest, authorID int64, token string) (*dto.PostResponse, error) {
	// Get user information from token
	userData, err := s.userService.GetUserFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %w", err)
	}
	
	// Verify that the token user ID matches the provided authorID
	if userData.ID != authorID {
		return nil, errors.New("token user ID does not match provided author ID")
	}
	
	// Convert DTO to model with author information
	post := req.ToModelWithAuthor(authorID, userData.Name, userData.Email)

	// Save to database
	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	// Create response
	response := &dto.PostResponse{}
	response.FromModel(post)

	return response, nil
}

// GetPostByID retrieves a post by its ID without author information
func (s *postService) GetPostByID(id int64) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	// Create response
	response := &dto.PostResponse{}
	response.FromModel(post)

	// Note: Author information cannot be fetched here because GET /users/:id 
	// requires admin access. Author info should be added by handler if the 
	// current user is the author (from JWT context)

	return response, nil
}

// GetAllPosts retrieves all posts with pagination without author information
func (s *postService) GetAllPosts(page, pageSize int) (*dto.PostListResponse, error) {
	// Set default pagination values
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	posts, err := s.postRepo.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	total, err := s.postRepo.GetTotalPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to get total posts count: %w", err)
	}

	// Calculate pagination
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	offset := (page - 1) * pageSize

	// Apply pagination
	var paginatedPosts []*dto.PostResponse
	if offset < len(posts) {
		end := offset + pageSize
		if end > len(posts) {
			end = len(posts)
		}
		
		postSlice := posts[offset:end]
		paginatedPosts = make([]*dto.PostResponse, len(postSlice))

		// Convert posts to responses without author information
		for i, post := range postSlice {
			response := &dto.PostResponse{}
			response.FromModel(post)
			
			// Note: Author information cannot be fetched here because GET /users/:id 
			// requires admin access. Author info should be added by handler for 
			// posts authored by the current user (from JWT context)
			
			paginatedPosts[i] = response
		}
	} else {
		paginatedPosts = []*dto.PostResponse{}
	}

	// Convert to []PostResponse (not []*PostResponse)
	posts_slice := make([]dto.PostResponse, len(paginatedPosts))
	for i, post := range paginatedPosts {
		posts_slice[i] = *post
	}

	return &dto.PostListResponse{
		Posts:      posts_slice,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// UpdatePost updates an existing post (only by the author)
func (s *postService) UpdatePost(id int64, req *dto.UpdatePostRequest, authorID int64) (*dto.PostResponse, error) {
	// Get existing post
	existingPost, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	// Check if the user is the author
	if existingPost.AuthorID != authorID {
		return nil, errors.New("unauthorized: only the author can update this post")
	}

	// Apply updates using the correct method name
	req.UpdateModel(existingPost)

	// Save updates
	if err := s.postRepo.UpdatePost(existingPost); err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	// Create response
	response := &dto.PostResponse{}
	response.FromModel(existingPost)

	// Note: Author information should be added by the handler using context data
	// We don't include author details here to avoid admin-only API calls

	return response, nil
}

// DeletePost deletes a post (only by the author)
func (s *postService) DeletePost(id int64, authorID int64) error {
	// Get existing post
	existingPost, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return fmt.Errorf("failed to get post: %w", err)
	}

	// Check if the user is the author
	if existingPost.AuthorID != authorID {
		return errors.New("unauthorized: only the author can delete this post")
	}

	// Delete the post
	if err := s.postRepo.DeletePost(id); err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

// GetPostsByAuthor retrieves all posts by a specific author with pagination
func (s *postService) GetPostsByAuthor(authorID int64, page, pageSize int) (*dto.PostListResponse, error) {
	// Set default pagination values
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// Note: We cannot validate if the author exists because GET /users/:id 
	// requires admin access. We'll just query posts and let the empty result 
	// speak for itself if the author doesn't exist.

	posts, err := s.postRepo.GetByAuthorID(authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts by author: %w", err)
	}

	// Calculate pagination
	total := int64(len(posts))
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	offset := (page - 1) * pageSize

	// Apply pagination
	var paginatedPosts []*dto.PostResponse
	if offset < len(posts) {
		end := offset + pageSize
		if end > len(posts) {
			end = len(posts)
		}

		postSlice := posts[offset:end]
		paginatedPosts = make([]*dto.PostResponse, len(postSlice))

		// Convert posts to responses without author information
		for i, post := range postSlice {
			response := &dto.PostResponse{}
			response.FromModel(post)
			
			// Note: Author information cannot be fetched here because GET /users/:id 
			// requires admin access. Author info should be added by handler if the 
			// current user is the author (from JWT context)
			
			paginatedPosts[i] = response
		}
	} else {
		paginatedPosts = []*dto.PostResponse{}
	}

	// Convert to []PostResponse (not []*PostResponse)
	posts_slice := make([]dto.PostResponse, len(paginatedPosts))
	for i, post := range paginatedPosts {
		posts_slice[i] = *post
	}

	return &dto.PostListResponse{
		Posts:      posts_slice,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}