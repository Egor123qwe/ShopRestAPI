package Storage

import (
	"ShopRestAPI/internal/Storage/repositories/products"
	"ShopRestAPI/internal/Storage/repositories/users"
)

type Store interface {
	Product() products.ProductRepository
	User() users.UserRepository
}
