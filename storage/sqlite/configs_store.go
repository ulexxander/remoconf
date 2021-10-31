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

const configGetAllQuery = `SELECT * FROM configs`

func (us *ConfigsStore) GetAll() ([]storage.Config, error) {
	var items []storage.Config
	if err := us.db.Select(&items, configGetAllQuery); err != nil {
		return nil, err
	}
	return items, nil
}

const configCreateQuery = `INSERT INTO configs (project_id, version, content, created_by)
VALUES ($1, $2, $3, $4)
RETURNING id`

func (us *ConfigsStore) Create(p storage.ConfigCreateParams) (int, error) {
	var id int
	if err := us.db.Get(
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
