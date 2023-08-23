package ServerApi

import (
	"ShopRestAPI/internal/Storage/sqlStorage"
	"ShopRestAPI/internal/routers"
	"ShopRestAPI/internal/server"
	"database/sql"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func New(config *server.Config) *server.ServerApi {
	return &server.ServerApi{
		config:       config,
		Router:       http.NewServeMux(),
		SessionStore: sessions.NewCookieStore([]byte(config.sessionKey)),
	}
}

func (s *server.ServerApi) Start() error {

	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}
	routers.ConfigureRoutes(s)

	return http.ListenAndServe(s.config.serverPort, s.Router)
}

func (s *server.ServerApi) configureStore() error {
	db, err := sql.Open(s.config.dbDriver, s.config.dbURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.Store = sqlStorage.New(db)
	return nil
}
