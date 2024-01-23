package repositories

import (
	"database/sql"
	"fmt"
	models "main/models"
)

type ProductRepository struct {
	BaseRepository
}

func (pr ProductRepository) ArchiveProduct(id int64) error {
	return Execute(pr.BaseRepository, func(db IDbExecutable) error {
		query := `
			UPDATE products
			SET archived = UTC_TIMESTAMP()
			WHERE products.id = ?
		`
		_, err := db.Exec(query, id)
		if err != nil {
			return err
		}
		return nil
	})
}

func (pr ProductRepository) UpdateProduct(pu models.Product) error {
	return Execute(pr.BaseRepository, func(db IDbExecutable) error {
		query := `
			UPDATE products
			SET catalog_id = ?
				, name = ?
				, description = ?
				, updated = UTC_TIMESTAMP()
			WHERE 
				products.id = ?
		`
		_, err := db.Exec(query, pu.CatalogId, pu.Name, pu.Description, pu.Id)
		if err != nil {
			return err
		}
		return nil
	})
}

func (pr ProductRepository) InsertProduct(product models.Product) (int64, error) {
	return ExecuteTransactWithResult[int64](pr.BaseRepository, func(tx IDbExecutable) (int64, error) {
		query := "INSERT INTO products (catalog_id, name, description, created) VALUES (?, ?, ?, UTC_TIMESTAMP())"

		result, err := tx.Exec(query, product.CatalogId, product.Name, product.Description)
		if err != nil {
			return 0, err
		}

		productId, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}

		// insert sku
		for _, sku := range product.Skus {
			sku.ProductId = productId
			skuId, err := insertSkuWithDbConn(tx, sku)
			if err != nil {
				return 0, err
			}

			// insert attributes
			for _, attr := range sku.Attributes {
				attr.SkuId = skuId
				err := insertAttributeWithDbConn(tx, attr)
				if err != nil {
					return 0, err
				}
			}
		}

		return productId, nil
	})
}

func (pr ProductRepository) GetProductsByCatalogId(id int64) ([]models.Product, error) {
	return pr.getProductsWithCondition("WHERE p.catalog_id = ? AND p.archived IS NULL", id)
}

func (pr ProductRepository) GetAllProducts() ([]models.Product, error) {
	return pr.getProductsWithCondition("WHERE p.archived IS NULL")
}

func (pr ProductRepository) GetProductById(id int64) (*models.Product, error) {
	products, err := pr.getProductsWithCondition("WHERE p.id = ?", id)
	if err != nil {
		return nil, err
	}
	if len(products) != 1 {
		return nil, fmt.Errorf("unable to get product bu id: %d", id)
	}
	return &products[0], nil
}

func (pr ProductRepository) getProductsWithCondition(condition string, params ...any) ([]models.Product, error) {
	return ExecuteWithResult[[]models.Product](pr.BaseRepository, func(db IDbExecutable) ([]models.Product, error) {
		query := `
		SELECT
			p.id
			, p.catalog_id
			, p.name
			, p.description
			, p.created
			, p.updated
			, p.archived
			, sku.id
			, sku.product_id
			, sku.amount
			, sku.price
			, sku.unit
			, sku.created
			, sku.updated
			, sku.archived
			, attr.sku_id
			, attr.key
			, attr.value
			, attr.value_type
		FROM products p
			LEFT JOIN skus sku ON
				sku.product_id = p.id
			LEFT JOIN attributes attr ON
				attr.sku_id = sku.id
		` + condition

		rows, err := db.Query(query, params...)
		if err != nil {
			return nil, err
		}

		// helper structs
		type skuNullable struct {
			Id        sql.NullInt64
			ProductId sql.NullInt64
			Amount    sql.NullFloat64
			Price     sql.NullFloat64
			Unit      sql.NullString
			Created   sql.NullTime
			Updated   sql.NullTime
			Archived  sql.NullTime
		}

		type attrNullable struct {
			SkuId     sql.NullInt64
			Key       sql.NullString
			Value     sql.NullString
			ValueType sql.NullString
		}

		// result object
		var products map[int64]models.Product = make(map[int64]models.Product, 0) // key - productId, value = product

		// sku and attributers store
		var skus map[int64]skuNullable = make(map[int64]skuNullable)             // key - skuId, value - sku
		var attributes map[int64][]attrNullable = make(map[int64][]attrNullable) // key - skuId, value - []attribute

		for rows.Next() {
			var product models.Product
			var sku skuNullable
			var attr attrNullable

			err = rows.Scan(
				&product.Id,
				&product.CatalogId,
				&product.Name,
				&product.Description,
				&product.Created,
				&product.Updated,
				&product.Archived,
				&sku.Id,
				&sku.ProductId,
				&sku.Amount,
				&sku.Price,
				&sku.Unit,
				&sku.Created,
				&sku.Updated,
				&sku.Archived,
				&attr.SkuId,
				&attr.Key,
				&attr.Value,
				&attr.ValueType,
			)
			if err != nil {
				return nil, err
			}

			// store product
			products[product.Id] = product

			// store skus
			if sku.Id.Valid {
				skus[sku.Id.Int64] = sku
			}

			// store attributes
			if attr.SkuId.Valid {
				attributes[attr.SkuId.Int64] = append(attributes[attr.SkuId.Int64], attr)
			}
		}

		// build product objects
		var output []models.Product = make([]models.Product, 0)
		for _, p := range products {
			for _, s := range skus {
				if s.ProductId.Valid {
					if s.ProductId.Int64 != p.Id {
						continue
					}

					sku := models.Sku{
						Id:        s.Id.Int64,
						ProductId: s.ProductId.Int64,
						Amount:    float32(s.Amount.Float64),
						Price:     float32(s.Price.Float64),
						Unit:      s.Unit.String,
						Created:   s.Created.Time,
						Updated:   s.Updated,
						Archived:  s.Archived,
					}

					for _, a := range attributes[s.Id.Int64] {
						attr := models.Attribute{
							SkuId:     a.SkuId.Int64,
							Key:       a.Key.String,
							Value:     a.Value.String,
							ValueType: a.ValueType.String,
						}

						sku.Attributes = append(sku.Attributes, attr)
					}

					p.Skus = append(p.Skus, sku)
				}
			}
			output = append(output, p)
		}

		return output, nil
	})
}
