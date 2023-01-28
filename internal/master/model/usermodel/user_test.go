package usermodel_test

import (
	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		payload func() *usermodel.User
		isValid bool
	} {
		{
			name: "valid",
			payload: func() *usermodel.User {
				u := usermodel.TestUser(t)
				return u
			},
			isValid: true,
		},
		{
			name: "empty login",
			payload: func() *usermodel.User {
				u := usermodel.TestUser(t)
				u.Login = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			payload: func() *usermodel.User {
				u := usermodel.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			payload: func() *usermodel.User {
				u := usermodel.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			
			u := tc.payload()
			err := u.Validate()

			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := usermodel.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}