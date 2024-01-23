package repositories

import (
	"main/configuration"
	m "main/models"
	"testing"
)

// setup repo
var repo ProductRepository = ProductRepository{
	BaseRepository: BaseRepository{
		ConnectionString: configuration.ConnectionString,
	},
}

// tests
func Test_GetProductsByCatalogIdWithFilter(t *testing.T) {
	var id int64 = 1
	filter := map[string]string{
		"Color": "Black",
		"Size":  "M",
	}
	products, err := repo.GetProductsByCatalogIdWithFilter(id, filter)
	if err != nil || products == nil {
		t.Errorf("unable to get products by catalog id and filter")
	}
}

func Test_GetProductsByCatalogId(t *testing.T) {
	var id int64 = 1
	products, err := repo.GetProductsByCatalogId(id)
	if err != nil || products == nil {
		t.Errorf("unable to get products by catalog id")
	}
}

func Test_ArchiveProduct(t *testing.T) {
	var id int64 = 1
	err := repo.ArchiveProduct(id)
	if err != nil {
		t.Errorf("unable to archive product")
	}
}

func TestProductRepository_UpdateProduct(t *testing.T) {
	var id int64 = 1
	product, err := repo.GetProductById(id)
	if err != nil {
		t.Errorf("unable to get product")
	}

	productUpdate := m.Product{
		Id:          id,
		CatalogId:   2,
		Name:        "T-shirt2",
		Description: "Nice and warm",
	}
	timeUpdatedOrig := product.Updated

	err = repo.UpdateProduct(productUpdate)
	if err != nil {
		t.Errorf("unable to update product")
	}

	productUpdated, err := repo.GetProductById(id)

	timeUpdated := productUpdated.Updated
	if err != nil {
		t.Errorf("unable to get product")
	}

	if !timeUpdated.Valid {
		t.Errorf("time updated is not valid")
	}

	if timeUpdatedOrig.Valid && timeUpdated.Time.Before(timeUpdatedOrig.Time) {
		t.Errorf("update time is not correct")
	}
}

func TestProductRepository_GetAll(t *testing.T) {
	products, err := repo.GetAllProducts()
	if err != nil || products == nil {
		t.Errorf("unable to get all products")
	}
}

func TestProductRepository_GetById(t *testing.T) {
	product, err := repo.GetProductById(1)
	if err != nil || product == nil {
		t.Errorf("unable to get product")
	}
}

func TestProductRepository_Insert(t *testing.T) {
	product := m.Product{
		CatalogId:   1,
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
	if err != nil || productId == 0 {
		t.Errorf("unable to save product")
	}
}
