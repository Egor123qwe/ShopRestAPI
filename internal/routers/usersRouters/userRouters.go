package usersRouters

import (
	"ShopRestAPI/internal/Storage"
	"net/http"
)

type UsersRoutes struct {
	store Storage.Store
}

func ConfigureUsersRoutes(mux *http.ServeMux, store Storage.Store) {
	//r := UsersRoutes{store: store}
}
