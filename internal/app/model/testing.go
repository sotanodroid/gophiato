package model

import "testing"

// TestUser helper function to initialize user instance for testing purpose
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
