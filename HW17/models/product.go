package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int64        `json:"id" db:"id"`
	CatalogId   int64        `json:"catalog_id" db:"catalog_id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Created     time.Time    `json:"created" db:"created"`
	Updated     sql.NullTime `json:"updated" db:"updated"`
	Archived    sql.NullTime `json:"archived" db:"archived"`
	Skus        []Sku        `json:"skus"`
}
