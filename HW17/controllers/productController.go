package controllers

import (
	"encoding/json"
	"fmt"
	dbStore "main/DbStore"
	"main/models"
	"net/http"
)

type ProductController struct {
	Db dbStore.DbStore
}

func (pc ProductController) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	id, err := pc.Db.Products.InsertProduct(product)
	if err != nil {
		http.Error(w, "unable to insert product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
