package routes

import (
	"ShopRestAPI/internal/Storage"
	"net/http"
)

type Routes struct {
	store *Storage.Store
}

func ConfigureRoutes(mux *http.ServeMux, store *Storage.Store) {
	r := Routes{store: store}
	r.ConfigureProductRoutes(mux)
}
