package Storage

import (
	"ShopRestAPI/internal/Storage/repositories/products"
)

type Store interface {
	Product() products.ProductRepository
}
