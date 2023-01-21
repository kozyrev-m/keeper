package sqlstore

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/integralist/go-findroot/find"
	"github.com/pressly/goose/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
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

// TestDB creates connect with postgres db for testing.
func TestDB(t testing.TB, databaseURL string) (*sql.DB, func()) {
	t.Helper()

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)

	}

	if err = goose.SetDialect("pgx"); err != nil {
		t.Fatal(err)
	}

	rep, err := find.Repo()
	if err != nil {
		t.Fatal(err)
	}

	if err := goose.Up(db, fmt.Sprintf("%s/migrations", rep.Path)); err != nil {
		t.Fatal(err)
	}


	return db, func() {
		if err := goose.Down(db, fmt.Sprintf("%s/migrations", rep.Path)); err != nil {
			t.Fatal(err)
		}
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}
}