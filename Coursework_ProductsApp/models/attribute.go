package models

type Attribute struct {
	SkuId     int64  `json:"sku_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
}
