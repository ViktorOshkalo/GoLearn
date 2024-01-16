package models

type Attribute struct {
	SkuId     int64  `db:"sku_id"`
	Key       string `db:"key"`
	Value     string `db:"value"`
	ValueType string `db:"value_type"`
}
