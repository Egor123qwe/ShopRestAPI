package routers

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/routers/productRouters"
	"net/http"
)

func ConfigureRoutes(mux *http.ServeMux, store Storage.Store) {
	productRouters.ConfigureProductRoutes(mux, store)
}
