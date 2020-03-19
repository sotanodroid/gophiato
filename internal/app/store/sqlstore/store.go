package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/sotanodroid/gophiato/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New creates new database connection
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User implements base permission for users to use repository
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
