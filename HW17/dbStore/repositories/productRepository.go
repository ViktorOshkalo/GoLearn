package repositories

import (
	models "main/models"
)

type ProductRepository struct {
	BaseRepository
}

func (pr ProductRepository) InsertProduct(product models.Product) (int64, error) {
	return ExecuteTransactWithResult[int64](pr.BaseRepository, func(tx IDbExecutable) (int64, error) {
		query := "INSERT INTO products (category_id, name, description, created) VALUES (?, ?, ?, UTC_TIMESTAMP())"

		result, err := tx.Exec(query, product.CategoryId, product.Name, product.Description)
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

func (pr ProductRepository) GetProductById(id int64) (*models.Product, error) {
	return ExecuteWithResult[*models.Product](pr.BaseRepository, func(db IDbExecutable) (*models.Product, error) {
		query := `
		SELECT
			p.id
			, p.category_id
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
		WHERE p.id = ?`

		rows, err := db.Query(query, id)
		if err != nil {
			return nil, err
		}

		// result object
		var product models.Product

		// sku and attributers store
		var skus map[int64]models.Sku = make(map[int64]models.Sku)                       // key - skuId, value - sku
		var attributes map[int64][]models.Attribute = make(map[int64][]models.Attribute) // key - skuId, value - []attribute

		for rows.Next() {
			var sku models.Sku
			var attr models.Attribute

			err = rows.Scan(
				&product.Id,
				&product.CategoryId,
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

			// store skus
			skus[sku.Id] = sku

			// store attributes
			attributes[attr.SkuId] = append(attributes[attr.SkuId], attr)
		}

		// add attributes to sku, add sku to product
		for _, s := range skus {
			s.Attributes = append(s.Attributes, attributes[s.Id]...)
			product.Skus = append(product.Skus, s)
		}

		return &product, nil
	})
}
