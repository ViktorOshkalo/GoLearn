package main

import (
	"fmt"
	"log"
	"main/configuration"
	"main/dbStore"
	m "main/models"
)

func main() {
	fmt.Println("Yoooo G")

	product := m.Product{
		CategoryId:  1,
		Name:        "T-shirt",
		Description: "Super comfortable and warm",
		Skus: []m.Sku{
			{
				Amount: 10,
				Unit:   "unit",
				Price:  200,

				Attributes: []m.Attribute{
					{
						Key:       "Size",
						Value:     "M",
						ValueType: "string",
					},
					{
						Key:       "Color",
						Value:     "Red",
						ValueType: "string",
					},
				},
			},
		},
	}

	dbStore := dbStore.GetNewDbStore(configuration.ConnectionString)

	productId, err := dbStore.Products.InsertProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	productRead, err := dbStore.Products.GetProductById(productId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product read from db: ", productRead)
}