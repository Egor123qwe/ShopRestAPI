package products

import (
	"ShopRestAPI/internal/Storage/repositories/products"
	models "ShopRestAPI/internal/models/products"
	"database/sql"
	"fmt"
	"log"
)

type product struct {
	db *sql.DB
}

func New(db *sql.DB) products.ProductRepository {
	return &product{db: db}
}

func (r *product) Create(p *models.Product) error {

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

	for _, inst := range p.Instances {
		inst.ProductId = p.Id
		if err := r.CreateInstance(&inst); err != nil {
			return err
		}
	}
	fmt.Println(p.Id)

	return nil
}

func (r *product) Get(id int) (*models.Product, error) {
	p := &models.Product{}
	p.Instances = []models.Instance{}

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
			"INNER JOIN color ON properties.color_id = color.id "+
			"INNER JOIN sizes ON properties.size_id = sizes.id "+
			"WHERE properties.product_id = $1;",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		inst := models.Instance{}
		err := rows.Scan(&inst.Id, &inst.Amount, &inst.Color, &inst.Size)
		if err != nil {
			return nil, err
		}
		p.Instances = append(p.Instances, inst)
	}
	return p, nil
}

func (r *product) Delete(id int) {
	r.db.QueryRow("DELETE FROM properties WHERE product_id = $1", id)
	r.db.QueryRow("DELETE FROM products WHERE id = $1", id)
}

func (r *product) Edit(p *models.Product) error {

	typeId := r.findAdditionalId("types", p.Types)
	styleId := r.findAdditionalId("styles", p.Style)
	seasonId := r.findAdditionalId("season", p.Season)
	countryId := r.findAdditionalId("country", p.Country)

	if _, err := r.db.Query(
		"UPDATE products SET name = $1, price = $2, description = $3, print = $4, "+
			"types_id = $5, style_id = $6, season_id = $7, country_id = $8 "+
			"WHERE id = $9",
		p.Name, p.Price, p.Description, p.Print,
		typeId, styleId, seasonId, countryId, p.Id,
	); err != nil {
		return err
	}
	for _, inst := range p.Instances {
		if err := r.EditInstance(&inst); err != nil {
			return err
		}
	}
	return nil
}

func (r *product) findAdditionalId(table string, name string) int {
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
