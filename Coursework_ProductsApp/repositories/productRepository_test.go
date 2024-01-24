package repositories

import (
	"main/configuration"
	m "main/models"
	"testing"
)

var repo ProductRepository

func setup() {
	// setup repo
	repo = ProductRepository{
		BaseRepository: BaseRepository{
			ConnectionString: configuration.ConnectionStringTest,
		},
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

// tests
func Test_GetProductsByCatalogIdWithFilter(t *testing.T) {
	// positive
	var id int64 = 1
	filter := map[string]string{
		"Color": "Black",
		"Size":  "M",
	}
	products, err := repo.GetProductsByCatalogIdWithFilter(id, filter)
	if err != nil || products == nil {
		t.Errorf("unable to get products by catalog id and filter")
	}

	for _, p := range products {
		for _, s := range p.Skus {
			found := 0
			for _, a := range s.Attributes {
				if v, exists := filter[a.Key]; exists {
					if v == filter[a.Key] {
						found++
					}
				}
			}
			if found < len(filter) {
				t.Errorf("sku has not enough attributes accordingly to applied filter")
			}
		}
	}

	// negative
	filter = map[string]string{
		"Color": "Yellow",
		"Size":  "XXL",
	}

	products, err = repo.GetProductsByCatalogIdWithFilter(id, filter)
	if err != nil {
		t.Errorf("unable to get products by catalog id and filter")
	}

	if len(products) > 0 {
		t.Errorf("shouldn't be products by this filter")
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

	product, err := repo.GetProductById(id)
	if err != nil {
		t.Errorf("unable to get product")
	}

	if !product.Archived.Valid {
		t.Errorf("archive time is not correct")
	}
}

func Test_UpdateProduct(t *testing.T) {
	var id int64 = 1
	product, err := repo.GetProductById(id)
	if err != nil {
		t.Errorf("unable to get product")
	}
	lastTimeUpdated := product.Updated

	var catalogIdNew int64 = 2
	nameNew := "T-shirt2"
	descriptionNew := "Nice and warm"

	productUpdate := m.Product{
		Id:          id,
		CatalogId:   catalogIdNew,
		Name:        nameNew,
		Description: descriptionNew,
	}

	err = repo.UpdateProduct(productUpdate)
	if err != nil {
		t.Errorf("unable to update product")
	}

	productUpdated, err := repo.GetProductById(id)
	if err != nil {
		t.Errorf("unable to get product")
	}

	newTimeUpdated := productUpdated.Updated

	if !newTimeUpdated.Valid {
		t.Errorf("time updated is not valid")
	}

	if lastTimeUpdated.Valid && newTimeUpdated.Time.Before(lastTimeUpdated.Time) {
		t.Errorf("update time is not correct")
	}

	if productUpdated.CatalogId != int64(catalogIdNew) {
		t.Errorf("catalog id is not updated")
	}

	if productUpdated.Name != nameNew {
		t.Errorf("name is not updated")
	}

	if productUpdated.Description != descriptionNew {
		t.Errorf("description is not updated")
	}
}

func Test_GetAll(t *testing.T) {
	products, err := repo.GetAllProducts()
	if err != nil || products == nil || len(products) == 0 {
		t.Errorf("unable to get all products")
	}
}

func Test_GetById(t *testing.T) {
	var id int64 = 1
	product, err := repo.GetProductById(id)
	if err != nil || product == nil {
		t.Errorf("unable to get product")
	}

	if product.Id != id {
		t.Errorf("wrong product received")
	}
}

func Test_Insert(t *testing.T) {
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
