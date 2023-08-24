package server

import (
	"ShopRestAPI/internal/Storage"
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	sessionName = "session"
)

var (
	errNotAuth = errors.New("not auth")
)

type Server struct {
	Router       *http.ServeMux
	Store        Storage.Store
	SessionStore sessions.Store
}

func NewServer(store Storage.Store, sessionStore sessions.Store) *Server {
	s := &Server{
		Router:       http.NewServeMux(),
		Store:        store,
		SessionStore: sessionStore,
	}

	return s
}

func (s *Server) AuthUser(w http.ResponseWriter, r *http.Request) error {
	session, err := s.SessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	id, ok := session.Values["user_id"]
	if !ok {
		return errNotAuth
	}

	_, err = s.Store.User().Find(id.(int))
	if err != nil {
		return err
	}
	return nil
}
