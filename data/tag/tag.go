package tag

import "time"

type Tag struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	List() ([]*Tag, error)
}

type Repository interface {
	Select() ([]*Tag, error)
}
