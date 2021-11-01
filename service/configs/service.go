package configs

import "gitlab.com/ulexxander/remoconf/storage"

type Service struct {
	configs storage.ConfigsStore
}

func NewService(configs storage.ConfigsStore) *Service {
	return &Service{configs}
}

func (s *Service) GetAll() ([]storage.Config, error) {
	return s.configs.GetAll()
}

func (s *Service) Create(p storage.ConfigCreateParams) (*storage.CreatedItem, error) {
	return s.configs.Create(p)
}
