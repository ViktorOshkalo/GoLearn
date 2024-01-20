package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int64        `db:"id"`
	CategoryId  int64        `db:"category_id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	Created     time.Time    `db:"created"`
	Updated     sql.NullTime `db:"updated"`
	Archived    sql.NullTime `db:"archived"`
	Skus        []Sku
}
