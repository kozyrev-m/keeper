package storage

import (
	"github.com/kozyrev-m/keeper/internal/master/config"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/sqlstore"
)

func CreateSqlStore(config *config.Config) (*sqlstore.Store, error) {

	db, err := sqlstore.NewDB(config.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	store := sqlstore.New(db)

	return store, err
}
