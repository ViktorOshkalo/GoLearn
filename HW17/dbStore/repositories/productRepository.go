package repositories

import (
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

func (pr ProductRepository) UpdateProduct(pu models.ProductUpdate) error {
	return Execute(pr.BaseRepository, func(db IDbExecutable) error {
		query := `
			UPDATE products
			SET category_id = ?
				, name = ?
				, description = ?
				, updated = UTC_TIMESTAMP()
			WHERE 
				products.id = ?
		`
		_, err := db.Exec(query, pu.CategoryId, pu.Name, pu.Description, pu.Id)
		if err != nil {
			return err
		}
		return nil
	})
}

func (pr ProductRepository) GetAllProducts() ([]models.Product, error) {
	return ExecuteWithResult[[]models.Product](pr.BaseRepository, func(db IDbExecutable) ([]models.Product, error) {
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
				attr.sku_id = sku.id`

		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}

		// result object
		var products map[int64]models.Product = make(map[int64]models.Product) // key - productId, value = product

		// sku and attributers store
		var skus map[int64]models.Sku = make(map[int64]models.Sku)                       // key - skuId, value - sku
		var attributes map[int64][]models.Attribute = make(map[int64][]models.Attribute) // key - skuId, value - []attribute

		for rows.Next() {
			var product models.Product
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

			// store product
			products[product.Id] = product

			// store skus
			skus[sku.Id] = sku

			// store attributes
			attributes[attr.SkuId] = append(attributes[attr.SkuId], attr)
		}

		// build product objects
		var output []models.Product
		for _, p := range products {
			for _, s := range skus {
				if s.ProductId != p.Id {
					continue
				}
				s.Attributes = append(s.Attributes, attributes[s.Id]...)
				p.Skus = append(p.Skus, s)
			}
			output = append(output, p)
		}

		return output, nil
	})
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
