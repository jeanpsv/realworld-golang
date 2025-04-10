package tag

type Service struct {
	repo Repository
}

func NewService(r Repository) UseCase {
	return &Service{
		repo: r,
	}
}

func (s *Service) List() ([]*Tag, error) {
	return s.repo.Select()
}
