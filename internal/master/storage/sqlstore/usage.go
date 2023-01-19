package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/integralist/go-findroot/find"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

// NewDB creates connect with postgres db.
func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = goose.SetDialect("pgx"); err != nil {
		return nil, err
	}

	rep, err := find.Repo()
	if err != nil {
		return nil, err
	}

	if err := goose.Up(db, fmt.Sprintf("%s/migrations", rep.Path)); err != nil {
		return nil, err
	}

	return db, nil
}
