package sqlStorage

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/Storage/repositories/products"
	productRep "ShopRestAPI/internal/Storage/sqlStorage/repositories/products"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sql.DB
	product products.ProductRepository
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
