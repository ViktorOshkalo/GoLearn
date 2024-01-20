package repositories

import (
	"database/sql"
	"fmt"
	"log"
	conf "main/configuration"
	models "main/models"
	"time"
)

func GetSkuById(id int64) (*models.Sku, error) {
	db, err := sql.Open("mysql", conf.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		SELECT
			s.id
			, s.product_id
			, s.amount
			, s.price
			, s.unit
			, s.created
			, s.updated
			, s.archived
		FROM skus s 
		WHERE id = ?`

	row := db.QueryRow(query, id)

	var sku models.Sku
	var created []uint8

	err = row.Scan(
		&sku.Id,
		&sku.ProductId,
		&sku.Amount,
		&sku.Price,
		&sku.Unit,
		&created,
		&sku.Updated,
		&sku.Archived)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sku not found by id: %d", id)
		} else {
			return nil, err
		}
	}

	sku.Created, err = time.Parse("2006-01-02 15:04:05", string(created))
	if err != nil {
		return nil, err
	}

	return &sku, nil
}

func insertSku(db IDbExecutable, sku models.Sku) (*int64, error) {
	var res int64 = 0
	query := "INSERT INTO skus (product_id, amount, price, unit, created) VALUES (?, ?, ?, ?, UTC_TIMESTAMP())"

	result, err := db.Exec(query, sku.ProductId, sku.Amount, sku.Price, sku.Unit)
	if err != nil {
		return &res, err
	}

	skuId, err := result.LastInsertId()
	if err != nil {
		return &res, err
	}

	res = skuId
	return &res, nil
}

func InsertSku(sku models.Sku) (id *int64, err error) {
	return ExecuteWithResult[int64](func(db IDbExecutable) (id *int64, err error) {
		return insertSku(db, sku)
	})
}
