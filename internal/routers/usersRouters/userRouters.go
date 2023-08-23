package usersRouters

import (
	"ShopRestAPI/internal/Storage"
	model "ShopRestAPI/internal/models/users"
	"ShopRestAPI/internal/routers/helperRoters"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const (
	sessionName = "session"
)

var (
	errIcorrecEmailorPassword = errors.New("Incorrect email or password")
)

type UsersRoutes struct {
	store        Storage.Store
	sessionStore sessions.Store
}

func ConfigureUsersRoutes(mux *http.ServeMux, store Storage.Store, sessionStore sessions.Store) {
	r := UsersRoutes{
		store:        store,
		sessionStore: sessionStore,
	}
	mux.HandleFunc("/registration", r.CreateRegisterRouter())
	mux.HandleFunc("/auth", r.CreateAuthRouter())
}

func (r *UsersRoutes) CreateRegisterRouter() http.HandlerFunc {
	type response struct {
		Email string `json:"email"`
	}

	return func(w http.ResponseWriter, req *http.Request) {

		var res = &model.Users{}
		if err := json.NewDecoder(req.Body).Decode(res); err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			return
		}

		if err := r.store.User().Create(res); err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			return
		}

		data := response{
			Email: res.Email,
		}
		info, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s\n", info)
	}
}

func (r *UsersRoutes) CreateAuthRouter() http.HandlerFunc {
	type response struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, req *http.Request) {
		res := &response{}
		if err := json.NewDecoder(req.Body).Decode(res); err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			return
		}

		u, err := r.store.User().FindByEmail(res.Email)
		if err != nil || !u.ComparePassword(res.Password) {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, errIcorrecEmailorPassword)
			return
		}

		session, err := r.sessionStore.Get(req, sessionName)
		if err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusInternalServerError, errIcorrecEmailorPassword)
			return
		}

		session.Values["user_id"] = u.Id
		if err := r.sessionStore.Save(req, w, session); err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusInternalServerError, errIcorrecEmailorPassword)
			return
		}

		helperRoters.Respond(w, req, http.StatusOK, nil)
	}
}
