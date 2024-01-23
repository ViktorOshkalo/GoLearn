package controllers

import (
	"bytes"
	"encoding/json"
	dbStore "main/DbStore"
	"main/configuration"
	m "main/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var db dbStore.DbStore = dbStore.GetNewDbStore(configuration.ConnectionString)
var controller ProductController = ProductController{Db: db}

func TestProductController_AddProduct(t *testing.T) {
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

	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fatal("Error marshaling JSON:", err)
	}

	req, err := http.NewRequest("POST", "/insertProduct", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal("Error creating request:", err)
	}

	w := httptest.NewRecorder()
	controller.AddProductHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, received %d", http.StatusOK, w.Code)
	}
}
