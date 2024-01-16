package repositories

import (
	"database/sql"
	"fmt"
	conf "main/configuration"
	models "main/models"
)

func InsertAttribute(attr models.Attribute) error {
	db, err := sql.Open("mysql", conf.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO attributes (`sku_id`, `key`, `value`, `value_type`) VALUES (?, ?, ?, ?)"

	_, err = db.Exec(query, attr.SkuId, attr.Key, attr.Value, attr.ValueType)
	return err
}

func GetAttributesBySkuId(id int64) ([]models.Attribute, error) {
	db, err := sql.Open("mysql", conf.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT
			a.sku_id
			, a.key
			, a.value
			, a.value_type
		FROM attributes a 
		WHERE sku_id = ?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var attributes []models.Attribute

	for rows.Next() {
		var attribute models.Attribute

		err = rows.Scan(
			&attribute.SkuId,
			&attribute.Key,
			&attribute.Value,
			&attribute.ValueType)

		if err != nil {
			return nil, err
		}

		attributes = append(attributes, attribute)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sku not found by id: %d", id)
		} else {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attributes, nil
}
