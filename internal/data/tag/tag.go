package tag

import "time"

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UseCase interface {
	List() ([]*Tag, error)
}

type Repository interface {
	Select() ([]*Tag, error)
}
