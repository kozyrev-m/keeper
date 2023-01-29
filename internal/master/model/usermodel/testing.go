package usermodel

import "testing"

// TestUser creates user for testing.
func TestUser(t *testing.T) *User {
	return &User{
		Login:    "someuser",
		Password: "password",
	}
}