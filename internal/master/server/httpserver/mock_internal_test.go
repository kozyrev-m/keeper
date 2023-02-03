package httpserver

import (
	"mime/multipart"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/sqlstore"
)

type mockStore struct {
	createUser      func(m *usermodel.User) error
	findUserByLogin func(login string) (*usermodel.User, error)
	findUserByID    func(id int) (*usermodel.User, error)

	createDataRecord func(c sqlstore.Content) error
	findTextsByOwner func(int) ([]datamodel.Text, error)

	findPairsByOwner func(int) ([]datamodel.LoginPassword, error)

	createFile func(int, string, string, multipart.File) error
	getFileList func(int) ([]datamodel.File, error)
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

func (ms *mockStore) CreateDataRecord(c sqlstore.Content) error {
	return ms.createDataRecord(c)
}

func (ms *mockStore) FindTextsByOwner(userID int) ([]datamodel.Text, error) {
	return ms.findTextsByOwner(userID)
}

func (ms *mockStore) FindPairsByOwner(userID int) ([]datamodel.LoginPassword, error) {
	return ms.findPairsByOwner(userID)
}

func (ms *mockStore) CreateFile(ownerID int, metadata string, filename string, file multipart.File) error {
	return ms.createFile(ownerID, metadata, filename, file)
}

func (ms *mockStore) GetFileList(ownerID int) ([]datamodel.File, error) {
	return ms.getFileList(ownerID)
}