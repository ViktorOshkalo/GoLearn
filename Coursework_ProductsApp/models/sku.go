package models

import (
	"database/sql"
	"time"
)

type Sku struct {
	Id         int64        `json:"id"`
	ProductId  int64        `json:"product_id"`
	Amount     float32      `json:"amount"`
	Price      float32      `json:"price"`
	Unit       string       `json:"unit"`
	Created    time.Time    `json:"created"`
	Updated    sql.NullTime `json:"updated"`
	Archived   sql.NullTime `json:"archived"`
	Attributes []Attribute  `json:"attributes"`
}
