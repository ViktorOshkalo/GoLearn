package main

import (
	"fmt"
	"log"
	"main/configuration"
	m "main/models"
	repo "main/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Yooooo!")

	repo.SetConnectionString(configuration.ConnectionString)

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

	productId, err := repo.InsertProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	productRead, err := repo.GetProductById(*productId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product read from db: ", productRead)
}
