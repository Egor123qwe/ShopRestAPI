package Storage

import (
	"ShopRestAPI/internal/Storage/Repositories"
	"ShopRestAPI/internal/Storage/migrations"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config     *Config
	db         *sql.DB
	ProductRep *Repositories.ProductRep
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	migrations.CreateTables(s.db) //////////

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Product() *Repositories.ProductRep {
	if s.ProductRep != nil {
		return s.ProductRep
	}

	s.ProductRep = Repositories.NewProductRep(s.db)
	return s.ProductRep
}
