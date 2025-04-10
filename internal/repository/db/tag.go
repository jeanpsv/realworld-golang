package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jeanpsv/realworld-golang/domain"
	"github.com/jeanpsv/realworld-golang/tag"
)

type TagRepository struct {
	conn *sql.DB
}

func NewTagRepository(c *sql.DB) tag.Repository {
	return &TagRepository{
		conn: c,
	}
}

func (r *TagRepository) Select() ([]*domain.Tag, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tags
		ORDER BY id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tags := []*domain.Tag{}

	for rows.Next() {
		var tag domain.Tag

		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
