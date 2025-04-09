package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jeanpsv/realworld-golang/data/tag"
)

type DB struct {
	*sql.DB
}

func New(db *sql.DB) tag.Repository {
	return &DB{db}
}

func (r *DB) Select() ([]*tag.Tag, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tags
		ORDER BY id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tags := []*tag.Tag{}

	for rows.Next() {
		var tag tag.Tag

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
