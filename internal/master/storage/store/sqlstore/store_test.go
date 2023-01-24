package sqlstore_test

import (
	"github.com/kozyrev-m/keeper/internal/master/model"
	"github.com/kozyrev-m/keeper/internal/master/storage/store"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/sqlstore"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	databaseURL = "host=localhost port=5432 dbname=keeper_test password=12345 sslmode=disable"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()
	
	s := sqlstore.New(db)
	u := model.TestUser(t)
	err := s.CreateUser(u)

	assert.NoError(t, err)
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindUserByLogin(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()

	s := sqlstore.New(db)

	login := "someuser"
	_, err := s.FindUserByLogin(login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Login = login

	err = s.CreateUser(u)
	require.NoError(t, err)

	u, err = s.FindUserByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}