package domain

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
	Create(post *Post) error
	GetByID(id int) (*Post, error)
	GetAll(searchTerm string) ([]*Post, error)
	Update(post *Post) error
	Delete(id int) error
}
