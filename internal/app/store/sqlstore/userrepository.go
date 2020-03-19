package sqlstore

import (
	"database/sql"

	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/sotanodroid/gophiato/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create creates new user instance
func (r *UserRepository) Create(u *model.User) error {
	query := `
		INSERT INTO restapi.users (
			email,
			encrypted_password
		) VALUES (
			$1,
			$2
		) RETURNING id
	`

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(
		query,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

// FindByEmail finds user in DB by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, encrypted_password
		FROM restapi.users
		WHERE email = $1
	`

	u := &model.User{}

	if err := r.store.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
