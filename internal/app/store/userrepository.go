package store

import "github.com/sotanodroid/gophiato/internal/app/model"

// UserRepository ...
type UserRepository struct {
	stroe *Store
}

// Create creates new user instance
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	query := `
		INSERT INTO restapi.users (
			email,
			encrypted_password
		) VALUES (
			$1,
			$2
		) RETURNING id
	`

	if err := r.stroe.db.QueryRow(query, u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail finds user in DB by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, encrypted_password
		FROM restapi.users
		WHERE email = $1
	`

	u := &model.User{}

	if err := r.stroe.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}