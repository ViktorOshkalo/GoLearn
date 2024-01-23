package repositories

import (
	"database/sql"
	"fmt"
	models "main/models"
	"time"
)

type SkuRepository struct {
	BaseRepository
}

func insertSkuWithDbConn(db IDbExecutable, sku models.Sku) (int64, error) {
	query := "INSERT INTO skus (product_id, amount, price, unit, created) VALUES (?, ?, ?, ?, UTC_TIMESTAMP())"

	result, err := db.Exec(query, sku.ProductId, sku.Amount, sku.Price, sku.Unit)
	if err != nil {
		return 0, err
	}

	skuId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return skuId, nil
}

func (sr SkuRepository) InsertSku(sku models.Sku) (id int64, err error) {
	return ExecuteWithResult[int64](sr.BaseRepository, func(db IDbExecutable) (int64, error) {
		return insertSkuWithDbConn(db, sku)
	})
}

func (sr SkuRepository) GetSkuById(id int64) (*models.Sku, error) {
	return ExecuteWithResult[*models.Sku](sr.BaseRepository, func(db IDbExecutable) (*models.Sku, error) {
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

		err := row.Scan(
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
	})
}
