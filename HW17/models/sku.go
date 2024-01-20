package models

import (
	"database/sql"
	"time"
)

type Sku struct {
	Id         int64        `db:"id"`
	ProductId  int64        `db:"product_id"`
	Amount     float32      `db:"amount"`
	Price      float32      `db:"price"`
	Unit       string       `db:"unit"`
	Created    time.Time    `db:"created"`
	Updated    sql.NullTime `db:"updated"`
	Archived   sql.NullTime `db:"archived"`
	Attributes []Attribute
}
