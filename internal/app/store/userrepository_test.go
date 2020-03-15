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

	testUser := model.TestUser(t)
	u, err := s.User().Create(testUser)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserReposytory_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("restapi.users")

	testUser := model.TestUser(t)
	_, err := s.User().FindByEmail(testUser.Email)

	assert.Error(t, err)

	s.User().Create(testUser)
	u, err := s.User().FindByEmail(testUser.Email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
