package controllers

import (
	"encoding/json"
	"fmt"
	dbStore "main/DbStore"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	Db dbStore.DbStore
}

func (pc ProductController) GetProductsByCatalogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catalogIDStr := vars["id"]

	// try read filter data
	var filter map[string]string
	if r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(&filter)
		if err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
	}

	catalogId, err := strconv.ParseInt(catalogIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid catalog id", http.StatusBadRequest)
		return
	}

	var products []models.Product
	if len(filter) > 0 {
		products, err = pc.Db.Products.GetProductsByCatalogIdWithFilter(catalogId, filter)
	} else {
		products, err = pc.Db.Products.GetProductsByCatalogId(catalogId)
	}

	if err != nil {
		errMessage := fmt.Sprintf("unable to get product by catalog id: %d", catalogId)
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
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
