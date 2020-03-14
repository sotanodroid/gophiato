package store_test

import (
	"testing"

	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/sotanodroid/gophiato/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserReposytory_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("restapi.users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserReposytory_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("restapi.users")

	email := "user@example.com"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.com",
	})

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
