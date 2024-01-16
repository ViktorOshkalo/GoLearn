package repositories

import (
	"database/sql"
	"fmt"
	conf "main/configuration"
	models "main/models"
	"time"
)

func InsertProduct(product models.Product) (int64, error) {
	db, err := sql.Open("mysql", conf.ConnectionString)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	query := "INSERT INTO products (category_id, name, description, created) VALUES (?, ?, ?, UTC_TIMESTAMP())"

	result, err := db.Exec(query, product.Categoryid, product.Name, product.Description)
	if err != nil {
		return -1, err
	}

	productId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return productId, nil
}

func GetProductById(id int64) (*models.Product, error) {
	db, err := sql.Open("mysql", conf.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT
			p.id
			, p.category_id
			, p.name
			, p.description
			, p.created
			, p.updated
			, p.archived
		FROM products p 
		WHERE id = ?`

	row := db.QueryRow(query, id)

	var product models.Product
	var created []uint8

	err = row.Scan(
		&product.Id,
		&product.Categoryid,
		&product.Name,
		&product.Description,
		&created,
		&product.Updated,
		&product.Archived)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sku not found by id: %d", id)
		} else {
			return nil, err
		}
	}

	product.Created, err = time.Parse("2006-01-02 15:04:05", string(created))
	if err != nil {
		return nil, err
	}

	return &product, nil
}
