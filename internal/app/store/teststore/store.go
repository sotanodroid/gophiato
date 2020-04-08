package teststore

import (
	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/sotanodroid/gophiato/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New creates new database connection
func New() *Store {
	return &Store{}
}

// User implements base permission for users to use repository
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
