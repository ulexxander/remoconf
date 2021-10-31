package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const migrationUp = `
PRAGMA foreign_keys = ON;

CREATE TABLE users (
	id integer PRIMARY KEY AUTOINCREMENT,
  login text NOT NULL,
  password text NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp
);

CREATE TABLE projects (
	id integer PRIMARY KEY AUTOINCREMENT,
	title text NOT NULL,
	description text NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by integer NOT NULL
		REFERENCES users(id),
	updated_at timestamp,
	updated_by integer
		REFERENCES users(id)
);
`

const migrationDown = `
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS projects;
`

func Migrate(db *sqlx.DB) error {
	_, err := db.Exec(migrationDown)
	if err != nil {
		return errors.Wrap(err, "migration down")
	}
	_, err = db.Exec(migrationUp)
	if err != nil {
		return errors.Wrap(err, "migration up")
	}
	return nil
}
