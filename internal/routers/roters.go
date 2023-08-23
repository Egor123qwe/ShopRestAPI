package routers

import (
	"ShopRestAPI/internal/routers/productRouters"
	"ShopRestAPI/internal/routers/usersRouters"
	"ShopRestAPI/internal/server"
)

func ConfigureRoutes(s *server.ServerApi) {
	productRouters.ConfigureProductRoutes(s.Router, s.Store)
	usersRouters.ConfigureUsersRoutes(s.Router, s.Store, s.SessionStore)
}
