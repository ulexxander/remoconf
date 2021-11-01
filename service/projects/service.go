package projects

import "gitlab.com/ulexxander/remoconf/storage"

type Service struct {
	projects storage.ProjectsStore
}

func NewService(ps storage.ProjectsStore) *Service {
	return &Service{ps}
}

func (s *Service) GetAll() ([]storage.Project, error) {
	return s.projects.GetAll()
}

func (s *Service) Create(p storage.ProjectCreateParams) (*storage.CreatedItem, error) {
	return s.projects.Create(p)
}
