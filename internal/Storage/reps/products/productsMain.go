package products

import (
	"ShopRestAPI/internal/Storage/models"
	"database/sql"
	"fmt"
	"log"
)

type ProductRep struct {
	db *sql.DB
}

func NewProductRep(db *sql.DB) *ProductRep {
	return &ProductRep{db: db}
}

func (r *ProductRep) Create(p *models.Product) error {

	typeId := r.findAdditionalId("types", p.Types)
	styleId := r.findAdditionalId("styles", p.Style)
	seasonId := r.findAdditionalId("season", p.Season)
	countryId := r.findAdditionalId("country", p.Country)

	if err := r.db.QueryRow(
		"INSERT INTO products (name, price, description, print, types_id, style_id, "+
			"season_id, country_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		p.Name, p.Price, p.Description, p.Print,
		typeId, styleId, seasonId, countryId,
	).Scan(&p.Id); err != nil {
		return err
	}

	for _, prop := range *p.Properties {
		r.CreateProperty(&prop, p.Id)
	}
	fmt.Println(p.Id)

	return nil
}

func (r *ProductRep) Get(id int) (*models.Product, error) {
	p := &models.Product{}
	p.Properties = &[]models.Property{}

	if err := r.db.QueryRow(
		"SELECT products.id, products.name, products.price, "+
			"products.description, products.print, types.name, "+
			"styles.name, season.name, country.name FROM products "+
			"INNER JOIN types ON products.types_id = types.id "+
			"INNER JOIN styles ON products.style_id = styles.id "+
			"INNER JOIN season ON products.season_id = season.id "+
			"INNER JOIN country ON products.country_id = country.id "+
			"WHERE products.id = $1;",
		id,
	).Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.Print,
		&p.Types, &p.Style, &p.Season, &p.Country); err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		"SELECT properties.id, properties.amount, color.name, sizes.name "+
			"FROM properties "+
			"INNER JOIN color ON properties.size_id = color.id "+
			"INNER JOIN sizes ON properties.color_id = sizes.id "+
			"WHERE properties.product_id = $1;",
		id,
	)
	if err != nil {
		rows.Close()
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		prop := models.Property{}
		err := rows.Scan(&prop.Id, &prop.Amount, &prop.Color, &prop.Size)
		if err != nil {
			return nil, err
		}
		*p.Properties = append(*p.Properties, prop)
	}
	return p, nil

}

func (r *ProductRep) Delete(id int) error {
	r.db.QueryRow("DELETE FROM products WHERE id = $1", id)
	r.db.QueryRow("DELETE FROM properties WHERE product_id = $1", id)

	return nil
}

func (r *ProductRep) Edit(id int) error {

	//Delete product = id and all properties with product_id = id
	return nil
}

func (r *ProductRep) findAdditionalId(table string, name string) int {
	var typeId int
	if err := r.db.QueryRow(
		"SELECT id FROM "+table+" WHERE name = $1",
		name,
	).Scan(&typeId); err != nil {
		fmt.Println(err)
		if err := r.db.QueryRow(
			"INSERT INTO "+table+"(name) VALUES ($1) RETURNING id",
			name,
		).Scan(&typeId); err != nil {
			log.Fatal(err)
		}
	}

	return typeId
}
