package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jeanpsv/realworld-golang/models"
	"github.com/jeanpsv/realworld-golang/services"
)

type TagRepository struct {
	conn *sql.DB
}

func NewTagRepository(c *sql.DB) services.TagRepository {
	return &TagRepository{
		conn: c,
	}
}

func (r *TagRepository) Select() ([]*models.Tag, error) {
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

	tags := []*models.Tag{}

	for rows.Next() {
		var tag models.Tag

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
