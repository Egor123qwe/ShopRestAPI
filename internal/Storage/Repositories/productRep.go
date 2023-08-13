package Repositories

import (
	"ShopRestAPI/internal/Storage/models"
	"database/sql"
)

type ProductRep struct {
	db *sql.DB
}

func NewProductRep(db *sql.DB) *ProductRep {
	return &ProductRep{db: db}
}

func (r *ProductRep) Create(p *models.Product) (*models.Product, error) {
	if err := r.db.QueryRow(
		"INSERT INTO products (name, amount, description, type_id, photos_id, price) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING goods_id",
		p.Name, p.Amount, p.Description,
		p.TypeId, p.PhotosId, p.Price,
	).Scan(&p.ID); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRep) Find(id int) (*models.Product, error) {
	p := &models.Product{}
	if err := r.db.QueryRow(
		"SELECT name, amount, description, type_id, photos_id, price FROM products WHERE goods_id = $1",
		id,
	).Scan(
		&p.Name, &p.Amount, &p.Description,
		&p.TypeId, &p.PhotosId, &p.Price,
	); err != nil {
		return nil, err
	}
	return p, nil
}

//brew services restart postgresql@15
