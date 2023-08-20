package products

import (
	"ShopRestAPI/internal/models/products"
)

type ProductRepository interface {
	Create(p *products.Product) error
	Get(id int) (*products.Product, error)
	Delete(id int)
	Edit(p *products.Product) error
	CreateInstance(prop *products.Instance) error
	EditInstance(p *products.Instance) error
	DeleteInstance(id int)
	ProductsSearch(filter *products.Filter) ([]products.Product, error)
	GetPropertyList(table string) ([]string, error)
}
