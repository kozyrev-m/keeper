// Package storage provides common interface of store.
package storage

import (
	"github.com/kozyrev-m/keeper/internal/master/model"
)

// Store is store iterface.
type Storage interface {
	CreateUser(*model.User) error
	FindUserByLogin(string) (*model.User, error)
}
