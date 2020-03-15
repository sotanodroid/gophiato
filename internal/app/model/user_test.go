package model_test

import (
	"testing"

	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid user instance",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = "notEmail"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid password length",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)
	assert.NoError(t, user.BeforeCreate())
	assert.NotEmpty(t, user.EncryptedPassword)
}
