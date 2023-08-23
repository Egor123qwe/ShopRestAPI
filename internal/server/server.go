package server

import (
	"ShopRestAPI/internal/Storage"
	"github.com/gorilla/sessions"
	"net/http"
)

type server struct {
	config       *Config
	router       *http.ServeMux
	store        Storage.Store
	sessionStore sessions.Store
}

func (s *Server) newServer() error {

	return http.ListenAndServe(s.Config.ServerPort, s.Router)
}
