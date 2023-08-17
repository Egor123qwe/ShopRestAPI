package ServerAPI

import (
	"ShopRestAPI/internal/ServerAPI/routes"
	"ShopRestAPI/internal/Storage"
	"log"
	"net/http"
)

type ServerApi struct {
	config *Config
	mux    *http.ServeMux
	store  *Storage.Store
}

func New(config *Config) *ServerApi {
	return &ServerApi{
		config: config,
		mux:    http.NewServeMux(),
		store:  Storage.New(config.store),
	}
}

func (s *ServerApi) Start() error {

	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}

	routes.ConfigureRoutes(s.mux, s.store)

	return http.ListenAndServe(s.config.serverPort, s.mux)
}

func (s *ServerApi) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}
	return nil
}
