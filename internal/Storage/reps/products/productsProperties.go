package products

import (
	"ShopRestAPI/internal/Storage/models"
	"math"
	"strconv"
	"strings"
)

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

func (r *ProductRep) EditProperty(p *models.Property) error {
	colorId := r.findAdditionalId("color", p.Color)
	sizesId := r.findAdditionalId("sizes", p.Size)
	photosId := 0
	if _, err := r.db.Query(
		"UPDATE properties SET color_id = $1, photos_id = $2, size_id = $3, "+
			"amount = $4 "+
			"WHERE id = $5",
		colorId, photosId, sizesId, p.Amount, p.Id,
	); err != nil {
		return err
	}
	return nil
}

func (r *ProductRep) DeleteProperty(id int) error {
	r.db.QueryRow("DELETE FROM properties WHERE id = $1", id)
	return nil
}

func (r *ProductRep) ProductsSearch(count int, page int, filter *models.ProductFilter) (*[]models.Product, error) {
	products := make([]models.Product, 0)

	if filter.MaxPrice == 0 {
		filter.MaxPrice = math.MaxFloat32
	}
	var additionalParams string
	additionalParams += getSqlCheck(filter.Print, "p.print", 6)
	additionalParams += getSqlCheck(filter.Types, "types.name", 7)
	additionalParams += getSqlCheck(filter.Style, "styles.name", 8)
	additionalParams += getSqlCheck(filter.Season, "season.name", 9)
	additionalParams += getSqlCheck(filter.Country, "country.name", 10)
	additionalParams += getSqlCheck(filter.Color, "color.name", 11)
	additionalParams += getSqlCheck(filter.Size, "sizes.name", 12)

	rows, err := r.db.Query(""+
		`SELECT DISTINCT prop.product_id, `+
		"p.name, p.description, p.price, "+
		"p.print, types.name, styles.name, season.name, "+
		"country.name "+
		"FROM properties prop "+
		"INNER JOIN color ON prop.color_id = color.id "+
		"INNER JOIN sizes ON prop.size_id = sizes.id "+
		"INNER JOIN products p ON prop.product_id = p.id "+
		"INNER JOIN types ON p.types_id = types.id "+
		"INNER JOIN styles ON p.style_id = styles.id "+
		"INNER JOIN season ON p.season_id = season.id "+
		"INNER JOIN country ON p.country_id = country.id "+
		"WHERE (p.price > $3 AND price < $4) "+
		"AND (LOWER(p.name) LIKE $5) "+additionalParams+
		"OFFSET $1 LIMIT $2;",
		page*count,
		count,
		filter.MinPrice, filter.MaxPrice,
		"%"+strings.ToLower(filter.Term)+"%",
		getConcatStr(filter.Print), getConcatStr(filter.Types),
		getConcatStr(filter.Style), getConcatStr(filter.Season),
		getConcatStr(filter.Country), getConcatStr(filter.Color),
		getConcatStr(filter.Size),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price,
			&p.Print, &p.Types, &p.Style, &p.Season, &p.Country)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &products, nil
}

func (r *ProductRep) GetPropertyList(table string) (*[]string, error) {
	pList := make([]string, 0)

	rows, err := r.db.Query(
		"SELECT (name) FROM " + table,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s string
	for rows.Next() {
		rows.Scan(&s)
		pList = append(pList, s)
	}

	return &pList, nil
}

func getConcatStr(arr []string) string {
	var res string
	for _, s := range arr {
		res += s
	}
	return res
}

func getSqlCheck(arr []string, param string, i int) string {
	if getConcatStr(arr) == "" {
		return " AND ($" + strconv.Itoa(i) + "=" + "$" + strconv.Itoa(i) + ") "
	}
	return " AND ($" + strconv.Itoa(i) + " LIKE CONCAT('%'," + param + ",'%')) "
}
