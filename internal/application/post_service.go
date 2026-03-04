package application

import (
	"fmt"
	"time"

	"github.com/ArtoIi/Blogging-Platform-API/internal/domain"
)

type PostService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *domain.Post) error {
	if len(post.Title) < 5 {
		return fmt.Errorf("título muito curto")
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	return s.repo.Create(post)
}

func (s *PostService) GetAllPosts(search string) ([]*domain.Post, error) {
	return s.repo.GetAll(search)
}

func (s *PostService) GetPostByID(id int) (*domain.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) UpdatePost(post *domain.Post) error {
	existingPost, err := s.repo.GetByID(post.ID)
	if err != nil {
		return err
	}
	post.CreatedAt = existingPost.CreatedAt
	post.UpdatedAt = time.Now()
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}
