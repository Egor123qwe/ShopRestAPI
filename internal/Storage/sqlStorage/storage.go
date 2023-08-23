package sqlStorage

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/Storage/repositories/products"
	"ShopRestAPI/internal/Storage/repositories/users"
	productRep "ShopRestAPI/internal/Storage/sqlStorage/repositories/products"
	userRep "ShopRestAPI/internal/Storage/sqlStorage/repositories/users"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sql.DB
	product products.ProductRepository
	user    users.UserRepository
}

func New(db *sql.DB) Storage.Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Product() products.ProductRepository {
	if s.product != nil {
		return s.product
	}

	s.product = productRep.New(s.db)
	return s.product
}

func (s *Store) User() users.UserRepository {
	if s.user != nil {
		return s.user
	}

	s.user = userRep.New(s.db)
	return s.user
}
