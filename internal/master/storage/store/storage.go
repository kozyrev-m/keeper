// Package store provides common interface of store.
package store

import (
	"github.com/kozyrev-m/keeper/internal/master/model"
)

// Store is store iterface.
type Store interface {
	CreateUser(*model.User) error
	FindUserByLogin(string) (*model.User, error)
}
