package Storage

import (
	"ShopRestAPI/internal/Storage/migrations"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db     *sql.DB
	config *Config
	Reps   *Reps
}

func New(config *Config) *Store {
	return &Store{
		config: config,
		Reps:   new(Reps),
	}
}

func (s *Store) Open() error {
	db, err := sql.Open(s.config.dbDriver, s.config.dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	migrations.CreateTables(s.db) //////////DELETE

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
