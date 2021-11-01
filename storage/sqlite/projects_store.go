package sqlite

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/ulexxander/remoconf/storage"
)

type ProjectsStore struct {
	db *sqlx.DB
}

func NewProjectsStore(db *sqlx.DB) *ProjectsStore {
	return &ProjectsStore{db}
}

const projectsGetAllQuery = `SELECT * FROM projects`

func (s *ProjectsStore) GetAll() ([]storage.Project, error) {
	var items []storage.Project
	if err := s.db.Select(&items, projectsGetAllQuery); err != nil {
		return nil, err
	}
	return items, nil
}

const projectCreateQuery = `INSERT INTO projects (title, description, created_by)
VALUES ($1, $2, $3)
RETURNING id`

func (s *ProjectsStore) Create(p storage.ProjectCreateParams) (int, error) {
	var id int
	if err := s.db.Get(
		&id,
		projectCreateQuery,
		p.Title,
		p.Description,
		p.CreatedBy,
	); err != nil {
		return 0, err
	}
	return id, nil
}
