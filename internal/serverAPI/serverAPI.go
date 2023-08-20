package serverAPI

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/Storage/sqlStorage"
	"ShopRestAPI/internal/routers"
	"database/sql"
	"log"
	"net/http"
)

type ServerApi struct {
	config *Config
	router *http.ServeMux
	store  Storage.Store
}

func New(config *Config) *ServerApi {
	return &ServerApi{
		config: config,
		router: http.NewServeMux(),
	}
}

func (s *ServerApi) Start() error {

	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}
	routers.ConfigureRoutes(s.router, s.store)

	return http.ListenAndServe(s.config.serverPort, s.router)
}

func (s *ServerApi) configureStore() error {
	db, err := sql.Open(s.config.dbDriver, s.config.dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.store = sqlStorage.New(db)
	return nil
}
