package users

import (
	"ShopRestAPI/internal/Storage/repositories/users"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) users.UserRepository {
	return &user{db: db}
}
