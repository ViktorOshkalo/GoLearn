package repositories

import (
	"fmt"
	"log"
	"main/configuration"
	m "main/models"
)

func RunDbTestCase() {

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

	repo := ProductRepository{BaseRepository: BaseRepository{ConnectionString: configuration.ConnectionString}}

	productId, err := repo.InsertProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	productRead, err := repo.GetProductById(productId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product read from db: ", productRead)
}
