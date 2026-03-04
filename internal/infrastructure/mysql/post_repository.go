package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ArtoIi/Blogging-Platform-API/internal/domain"
)

type mysqlPostRepository struct {
	db *sql.DB
}

func NewMySQLPostRepository(db *sql.DB) domain.PostRepository {
	return &mysqlPostRepository{db: db}
}

func (r *mysqlPostRepository) Create(post *domain.Post) error {
	tagsString := strings.Join(post.Tags, ",")

	query := `INSERT INTO posts (title, content, category, tags) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(query, post.Title, post.Content, post.Category, tagsString)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		post.ID = int(id)
	}

	return nil
}

func (r *mysqlPostRepository) GetByID(id int) (*domain.Post, error) {
	query := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = ?`

	var post domain.Post
	var tagsString string

	err := r.db.QueryRow(query, id).Scan(
		&post.ID, &post.Title, &post.Content, &post.Category, &tagsString, &post.CreatedAt, &post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if tagsString != "" {
		post.Tags = strings.Split(tagsString, ",")
	}

	return &post, nil
}

func (r *mysqlPostRepository) GetAll(searchTerm string) ([]*domain.Post, error) {
	query := `SELECT id, title, content, category, tags, created_at, updated_at 
              FROM posts 
              WHERE title LIKE ? OR category LIKE ? OR content LIKE ? OR tags LIKE ?`

	filter := "%" + searchTerm + "%"

	rows, err := r.db.Query(query, filter, filter, filter, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*domain.Post{}

	for rows.Next() {
		var p domain.Post
		var tagsString string

		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &tagsString, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if tagsString != "" {
			p.Tags = strings.Split(tagsString, ",")
		} else {
			p.Tags = []string{}
		}

		posts = append(posts, &p)
	}

	return posts, nil
}

func (r *mysqlPostRepository) Update(post *domain.Post) error {
	tagsString := strings.Join(post.Tags, ",")

	query := `UPDATE posts SET title = ?, content = ?, category = ?, tags = ? WHERE id = ?`

	_, err := r.db.Exec(query, post.Title, post.Content, post.Category, tagsString, post.ID)
	return err
}

func (r *mysqlPostRepository) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("nenhum post encontrado com o ID %d", id)
	}
	return nil
}
