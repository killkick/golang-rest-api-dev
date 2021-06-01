package store

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/store"
)

type UserRepository struct {
	Store *store.Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO users(email, ecnrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) findByEmail(email string) (*model.User, error) {
	if err := r.Store.Db.QueryRow("SELECT ALL FROM users WHERE users.id = $1 RETURNING ID", email).Scan(); err != nil {
		return nil, err
	}
	return
}
