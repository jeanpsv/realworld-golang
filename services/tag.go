package services

import "github.com/jeanpsv/realworld-golang/models"

type TagRepository interface {
	Select() ([]*models.Tag, error)
}

type TagService struct {
	tags TagRepository
}

func NewTagService(tagRepository TagRepository) TagService {
	return TagService{
		tags: tagRepository,
	}
}

func (s *TagService) List() ([]*models.Tag, error) {
	tags, err := s.tags.Select()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
