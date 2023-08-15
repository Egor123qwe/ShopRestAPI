package products

import "ShopRestAPI/internal/Storage/models"

func (r *ProductRep) CreateProperty(prop *models.Property, productId int) error {
	colorId := r.findAdditionalId("color", prop.Color)
	sizesId := r.findAdditionalId("sizes", prop.Size)
	photosId := 0
	if err := r.db.QueryRow(
		"INSERT INTO properties (product_id, color_id, photos_id, size_id, amount) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		productId, colorId, photosId, sizesId, prop.Amount,
	).Scan(&prop.Id); err != nil {
		return err
	}
	return nil
}

func (r *ProductRep) EditProperty(id int) error {

	return nil
}

func (r *ProductRep) DeleteProperty(id int) error {
	r.db.QueryRow("DELETE FROM properties WHERE id = $1", id)
	return nil
}

func (r *ProductRep) ProductsSearch(count int, page int, term string) error {

	return nil
}
