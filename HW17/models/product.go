package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int64        `json:"id" db:"id"`
	CategoryId  int64        `json:"category_id" db:"category_id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Created     time.Time    `json:"created" db:"created"`
	Updated     sql.NullTime `json:"updated" db:"updated"`
	Archived    sql.NullTime `json:"archived" db:"archived"`
	Skus        []Sku        `json:"skus"`
}

type ProductUpdate struct {
	Id          int64  `json:"id" db:"id"`
	CategoryId  int64  `json:"category_id" db:"category_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
