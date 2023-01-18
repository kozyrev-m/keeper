package model_test

import (
	"github.com/kozyrev-m/keeper/internal/master/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		payload func() *model.User
		isValid bool
	} {
		{
			name: "valid",
			payload: func() *model.User {
				u := model.TestUser(t)
				return u
			},
			isValid: true,
		},
		{
			name: "empty login",
			payload: func() *model.User {
				u := model.TestUser(t)
				u.Login = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			payload: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			payload: func() *model.User {
				u := model.TestUser(t)
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
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}