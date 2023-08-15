package Storage

import (
	"ShopRestAPI/internal/Storage/reps/products"
)

type Reps struct {
	ProductRep *products.ProductRep
}

func (s *Store) Product() *products.ProductRep {
	if s.Reps.ProductRep != nil {
		return s.Reps.ProductRep
	}

	s.Reps.ProductRep = products.NewProductRep(s.db)
	return s.Reps.ProductRep
}
