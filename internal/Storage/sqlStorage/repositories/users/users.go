package users

import (
	"ShopRestAPI/internal/Storage/repositories/users"
	model "ShopRestAPI/internal/models/users"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) users.UserRepository {
	return &user{db: db}
}

func (r *user) Create(u *model.Users) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.db.QueryRow(
		"INSERT INTO users (email, password, role) "+
			"VALUES ($1, $2, $3);",
		u.Email, u.EncryptedPassword, u.Role,
	)
	return nil
}

func (r *user) FindByEmail(email string) (*model.Users, error) {
	u := &model.Users{}
	if err := r.db.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = $1", email,
	).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Role); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *user) Find(id int) (*model.Users, error) {
	u := &model.Users{}
	if err := r.db.QueryRow(
		"SELECT id, email, password, role FROM users WHERE id = $1", id,
	).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Role); err != nil {
		return nil, err
	}
	return u, nil
}
