package main

import (
	"encoding/json"
	"fmt"
	dbStore "main/DbStore"
	conf "main/configuration"
	"main/controllers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db dbStore.DbStore
var productController controllers.ProductController

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, success := r.BasicAuth()
		if !success {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != conf.User {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if password != conf.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := db.Products.GetAllProducts()
	if err != nil {
		http.Error(w, "unable to get products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := db.Products.GetProductById(productID)
	if err != nil {
		errMessage := fmt.Sprintf("unable to get product by id: %d", productID)
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func main() {
	fmt.Println("Yoo G")

	db = dbStore.GetNewDbStore(conf.ConnectionString)
	productController = controllers.ProductController{Db: db}

	router := mux.NewRouter()
	router.HandleFunc("/products", GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/products", productController.AddProductHandler).Methods("POST")
	router.HandleFunc("/products/catalog/{id:[0-9]+}", productController.GetProductsByCatalogHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", GetProductHandler).Methods("GET")
	router.Use(AuthenticateMiddleware)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error: ", err)
	}
}

// {
// 	"color": "blue"
// 	"size": "M"
// }
