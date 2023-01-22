package httpserver_test

import "github.com/kozyrev-m/keeper/internal/master/model"


type mockStore struct {
	createUser func(m *model.User) error
	findUserByLogin func(login string) (*model.User, error)
}

func (ms *mockStore) CreateUser(m *model.User) error {
	return ms.createUser(m)
}

func (ms *mockStore) FindUserByLogin(login string) (*model.User, error) {
	return ms.findUserByLogin(login)
}
