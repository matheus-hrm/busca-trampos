package post

import (
	"database/sql"

	"github.com/matheus-hrm/trampos/types"
)

type Store struct {
	db *sql.DB
}

// CreatePost implements types.PostsStore.
func (s *Store) CreatePost(types.Post) error {
	panic("unimplemented")
}

// DeletePost implements types.PostsStore.
func (s *Store) DeletePost(id int) error {
	panic("unimplemented")
}

// UpdatePost implements types.PostsStore.
func (s *Store) UpdatePost(types.Post) error {
	panic("unimplemented")
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := make([]types.Post, 0)
	for rows.Next() {
		p, err := scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}
	return posts, nil
}

func (s *Store) GetPostByID(id int) (*types.Post, error) {
	row := s.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
	post := new(types.Post)

	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.Address,
		&post.Salary,
		&post.Status,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func scanRowsIntoPost(rows *sql.Rows) (*types.Post, error) {
	post := new(types.Post)
	err := rows.Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.Address,
		&post.Salary,
		&post.Status,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}
