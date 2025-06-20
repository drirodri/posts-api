package services
import (
	"errors"
	"fmt"
	"math"
	"posts-api/internal/dto"
	"posts-api/internal/repository"
	"gorm.io/gorm"
)
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
func NewPostService(postRepo repository.PostRepository, userService UserService) PostService {
	return &postService{
		postRepo:    postRepo,
		userService: userService,
	}
}
func (s *postService) CreatePost(req *dto.CreatePostRequest, authorID int64, token string) (*dto.PostResponse, error) {
	userData, err := s.userService.GetUserFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %w", err)
	}
	
	if userData.ID != authorID {
		return nil, errors.New("token user ID does not match provided author ID")
	}
	
	post := req.ToModelWithAuthor(authorID, userData.Name, userData.Email)
	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	response := &dto.PostResponse{}
	response.FromModel(post)
	return response, nil
}
func (s *postService) GetPostByID(id int64) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	response := &dto.PostResponse{}
	response.FromModel(post)
	return response, nil
}
func (s *postService) GetAllPosts(page, pageSize int) (*dto.PostListResponse, error) {
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
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	offset := (page - 1) * pageSize
	var paginatedPosts []*dto.PostResponse
	if offset < len(posts) {
		end := offset + pageSize
		if end > len(posts) {
			end = len(posts)
		}
		
		postSlice := posts[offset:end]
		paginatedPosts = make([]*dto.PostResponse, len(postSlice))
		for i, post := range postSlice {
			response := &dto.PostResponse{}
			response.FromModel(post)
			
			
			paginatedPosts[i] = response
		}
	} else {
		paginatedPosts = []*dto.PostResponse{}
	}
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
func (s *postService) UpdatePost(id int64, req *dto.UpdatePostRequest, authorID int64) (*dto.PostResponse, error) {
	existingPost, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	if existingPost.AuthorID != authorID {
		return nil, errors.New("unauthorized: only the author can update this post")
	}
	req.UpdateModel(existingPost)
	if err := s.postRepo.UpdatePost(existingPost); err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}
	response := &dto.PostResponse{}
	response.FromModel(existingPost)
	return response, nil
}
func (s *postService) DeletePost(id int64, authorID int64) error {
	existingPost, err := s.postRepo.GetPostByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return fmt.Errorf("failed to get post: %w", err)
	}
	if existingPost.AuthorID != authorID {
		return errors.New("unauthorized: only the author can delete this post")
	}
	if err := s.postRepo.DeletePost(id); err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}
func (s *postService) GetPostsByAuthor(authorID int64, page, pageSize int) (*dto.PostListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	posts, err := s.postRepo.GetByAuthorID(authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts by author: %w", err)
	}
	total := int64(len(posts))
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	offset := (page - 1) * pageSize
	var paginatedPosts []*dto.PostResponse
	if offset < len(posts) {
		end := offset + pageSize
		if end > len(posts) {
			end = len(posts)
		}
		postSlice := posts[offset:end]
		paginatedPosts = make([]*dto.PostResponse, len(postSlice))
		for i, post := range postSlice {
			response := &dto.PostResponse{}
			response.FromModel(post)
			
			
			paginatedPosts[i] = response
		}
	} else {
		paginatedPosts = []*dto.PostResponse{}
	}
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
