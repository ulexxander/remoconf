package configs

import (
	"github.com/pkg/errors"
	"gitlab.com/ulexxander/remoconf/storage"
)

type Service struct {
	configs storage.ConfigsStore
}

func NewService(configs storage.ConfigsStore) *Service {
	return &Service{configs}
}

func (s *Service) GetByProject(id int) ([]storage.Config, error) {
	return s.configs.GetByProject(id)
}

type ConfigCreateParams struct {
	ProjectID int
	Content   string
	CreatedBy int
}

func (s *Service) Create(p ConfigCreateParams) (*storage.CreatedItem, error) {
	projectConfigs, err := s.configs.GetByProject(p.ProjectID)
	if err != nil {
		return nil, errors.Wrap(err, "getting project configs")
	}
	version := 1
	if len(projectConfigs) > 0 {
		lastConfig := projectConfigs[len(projectConfigs)-1]
		version = lastConfig.Version + 1
	}
	created, err := s.configs.Create(storage.ConfigCreateParams{
		ProjectID: p.ProjectID,
		Version:   version,
		Content:   p.Content,
		CreatedBy: p.CreatedBy,
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating config in db")
	}
	return created, nil
}
