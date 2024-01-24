package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int64        `json:"id"`
	CatalogId   int64        `json:"catalog_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Created     time.Time    `json:"created"`
	Updated     sql.NullTime `json:"updated"`
	Archived    sql.NullTime `json:"archived"`
	Skus        []Sku        `json:"skus"`
}
