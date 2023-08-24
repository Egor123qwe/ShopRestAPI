package ServerApi

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/Storage/sqlStorage"
	"ShopRestAPI/internal/routers"
	"ShopRestAPI/internal/server"
	"database/sql"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func Start(config *config) error {

	storage, err := configureStore(config.dbDriver, config.dbURL)
	if err != nil {
		log.Fatal(err)
	}
	sessionsStore := sessions.NewCookieStore([]byte(config.sessionKey))

	s := server.NewServer(storage, sessionsStore)
	routers.ConfigureRoutes(s)

	return http.ListenAndServe(config.serverPort, s.Router)
}

func configureStore(dbDriver string, dbURL string) (Storage.Store, error) {
	db, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return sqlStorage.New(db), nil
}
