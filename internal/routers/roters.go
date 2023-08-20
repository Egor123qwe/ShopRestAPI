package routers

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/routers/productRouters"
	"ShopRestAPI/internal/routers/usersRouters"
	"net/http"
)

func ConfigureRoutes(mux *http.ServeMux, store Storage.Store) {
	productRouters.ConfigureProductRoutes(mux, store)
	usersRouters.ConfigureUsersRoutes(mux, store)
}
