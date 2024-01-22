package repositories

import (
	"fmt"
	"main/configuration"
	m "models"
	"testing"
)

func TestProductRepositoryInsert(t *testing.T) {

	fmt.Println("Test running")

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
		t.Errorf("Unable to save product")
	}

	if productId == 0 {
		t.Errorf("Unable to save product")
	}
}
