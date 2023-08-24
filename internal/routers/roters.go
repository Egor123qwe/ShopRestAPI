package routers

import (
	"ShopRestAPI/internal/routers/productRouters"
	"ShopRestAPI/internal/routers/usersRouters"
	"ShopRestAPI/internal/server"
)

func ConfigureRoutes(s *server.Server) {
	productRouters.ConfigureProductRoutes(s)
	usersRouters.ConfigureUsersRoutes(s.Router, s.Store, s.SessionStore)
}
