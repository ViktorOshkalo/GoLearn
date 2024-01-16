package main

import (
	"fmt"
	"log"
	m "main/models"
	repo "main/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Yooooo!")

	product := m.Product{
		Categoryid:  1,
		Name:        "T-shirt",
		Description: "Super comfortable and warm",
	}

	productId, err := repo.InsertProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	sku := m.Sku{
		ProductId: productId,
		Amount:    10,
		Unit:      "unit",
		Price:     200,
	}

	skuId, err := repo.InsertSku(sku)
	if err != nil {
		log.Fatal(err)
	}

	attributes := []m.Attribute{
		{
			SkuId:     skuId,
			Key:       "Size",
			Value:     "M",
			ValueType: "string",
		},
		{
			SkuId:     skuId,
			Key:       "Color",
			Value:     "Red",
			ValueType: "string",
		},
	}

	for _, attr := range attributes {
		err = repo.InsertAttribute(attr)
		if err != nil {
			log.Fatal(err)
		}
	}

	productRead, err := repo.GetProductById(productId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product read from db: ", productRead)

	skuRead, err := repo.GetSkuById(skuId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sku read from db: ", skuRead)

	attributesRead, err := repo.GetAttributesBySkuId(skuId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Attributes read from db: ", attributesRead)
}
