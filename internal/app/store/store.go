package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New creates new database connection
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open initialize connection
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close closes connection
func (s *Store) Close() {
	s.db.Close()
}
