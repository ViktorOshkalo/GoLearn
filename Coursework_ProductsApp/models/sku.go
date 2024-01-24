package models

import (
	"database/sql"
	"time"
)

type Sku struct {
	Id         int64        `json:"id" db:"id"`
	ProductId  int64        `json:"product_id" db:"product_id"`
	Amount     float32      `json:"amount" db:"amount"`
	Price      float32      `json:"price" db:"price"`
	Unit       string       `json:"unit" db:"unit"`
	Created    time.Time    `json:"created" db:"created"`
	Updated    sql.NullTime `json:"updated" db:"updated"`
	Archived   sql.NullTime `json:"archived" db:"archived"`
	Attributes []Attribute  `json:"attributes"`
}
