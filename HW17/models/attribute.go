package models

type Attribute struct {
	SkuId     int64  `json:"sku_id" db:"sku_id"`
	Key       string `json:"key" db:"key"`
	Value     string `json:"value" db:"value"`
	ValueType string `json:"value_type" db:"value_type"`
}
