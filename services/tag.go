package services

import "github.com/jeanpsv/realworld-golang/models"

type TagRepository interface {
	Select() ([]*models.Tag, error)
}

type TagService struct {
	repo TagRepository
}

func NewTagService(t TagRepository) TagService {
	return TagService{
		repo: t,
	}
}

func (s *TagService) List() ([]*models.Tag, error) {
	tags, err := s.repo.Select()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
