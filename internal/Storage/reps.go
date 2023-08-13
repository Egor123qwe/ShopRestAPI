package Storage

import "ShopRestAPI/internal/Storage/reps"

type Reps struct {
	ProductRep *reps.ProductRep
}

func (s *Store) Product() *reps.ProductRep {
	if s.Reps.ProductRep != nil {
		return s.Reps.ProductRep
	}

	s.Reps.ProductRep = reps.NewProductRep(s.db)
	return s.Reps.ProductRep
}
