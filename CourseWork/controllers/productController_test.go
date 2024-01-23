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

	"github.com/gorilla/mux"
)

var db dbStore.DbStore = dbStore.GetNewDbStore(configuration.ConnectionString)
var controller ProductController = ProductController{Db: db}

func TestProductController_GetProductsByCatalogId(t *testing.T) {
	filter := map[string]string{
		"Color": "Black",
		"Size":  "M",
	}
	filterJSON, err := json.Marshal(filter)
	if err != nil {
		t.Fatal("error converting filter to json:", err)
	}

	testCaseData := [][]byte{
		filterJSON, // with filter
		nil,        // without filter
	}

	for _, data := range testCaseData {
		req, err := http.NewRequest("GET", "/products/catalog/1", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal("error creating request:", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": "1"})

		w := httptest.NewRecorder()
		controller.GetProductsByCatalogHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, received %d", http.StatusOK, w.Code)
		}
	}
}

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
		t.Fatal("error marshaling JSON:", err)
	}

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal("error creating request:", err)
	}

	w := httptest.NewRecorder()
	controller.AddProductHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, received %d", http.StatusOK, w.Code)
	}
}
