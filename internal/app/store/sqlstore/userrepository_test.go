package sqlstore_test

import (
	"testing"

	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/sotanodroid/gophiato/internal/app/store"
	"github.com/sotanodroid/gophiato/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserReposytory_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("restapi.users")

	s := sqlstore.New(db)
	testUser := model.TestUser(t)
	err := s.User().Create(testUser)

	assert.NoError(t, err)
}

func TestUserReposytory_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("restapi.users")

	s := sqlstore.New(db)
	testUser := model.TestUser(t)
	_, err := s.User().FindByEmail(testUser.Email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(testUser)
	u, err := s.User().FindByEmail(testUser.Email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserReposytory_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("restapi.users")

	s := sqlstore.New(db)
	testUser := model.TestUser(t)
	s.User().Create(testUser)
	u, err := s.User().Find(testUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
