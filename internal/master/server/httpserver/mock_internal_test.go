package httpserver

import (
	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
)


type mockStore struct {
	createUser func(m *usermodel.User) error
	findUserByLogin func(login string) (*usermodel.User, error)
	findUserByID func(id int) (*usermodel.User, error)
}

func (ms *mockStore) CreateUser(m *usermodel.User) error {
	return ms.createUser(m)
}

func (ms *mockStore) FindUserByLogin(login string) (*usermodel.User, error) {
	return ms.findUserByLogin(login)
}

func (ms *mockStore) FindUserByID(id int) (*usermodel.User, error) {
	return ms.findUserByID(id)
}
