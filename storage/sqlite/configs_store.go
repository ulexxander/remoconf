package sqlite

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/ulexxander/remoconf/storage"
)

type ConfigsStore struct {
	db *sqlx.DB
}

func NewConfigsStore(db *sqlx.DB) *ConfigsStore {
	return &ConfigsStore{db}
}

const configsGetAllQuery = `SELECT * FROM configs`

func (s *ConfigsStore) GetAll() ([]storage.Config, error) {
	var items []storage.Config
	if err := s.db.Select(&items, configsGetAllQuery); err != nil {
		return nil, err
	}
	return items, nil
}

const configCreateQuery = `INSERT INTO configs (project_id, version, content, created_by)
VALUES ($1, $2, $3, $4)
RETURNING id`

func (s *ConfigsStore) Create(p storage.ConfigCreateParams) (int, error) {
	var id int
	if err := s.db.Get(
		&id,
		configCreateQuery,
		p.ProjectID,
		p.Version,
		p.Content,
		p.CreatedBy,
	); err != nil {
		return 0, err
	}
	return id, nil
}
