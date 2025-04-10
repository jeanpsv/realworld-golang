package tag

import "github.com/jeanpsv/realworld-golang/domain"

type Repository interface {
	Select() ([]*domain.Tag, error)
}

// type Service interface {
// 	List() ([]*domain.Tag, error)
// }

// type TagService struct {
// 	repo Repository
// }

type Service struct {
	repo Repository
}

func NewService(t Repository) Service {
	return Service{
		repo: t,
	}
}

func (s *Service) List() ([]*domain.Tag, error) {
	tags, err := s.repo.Select()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
