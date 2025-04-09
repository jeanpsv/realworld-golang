package data

import (
	"database/sql"

	"github.com/jeanpsv/realworld-golang/data/tag"
	"github.com/jeanpsv/realworld-golang/data/tag/mysql"
)

type Models struct {
	Tags tag.UseCase
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tags: tag.NewService(mysql.New(db)),
	}
}
