package main

import (
	"fmt"
	"main/configuration"
	"main/controllers"
	"main/dbstore"
	"net/http"

	"github.com/gorilla/mux"
)

var db dbstore.DbStore
var productController controllers.ProductController

func main() {
	fmt.Println("Yoo G")

	configuration.Setup()

	db = dbstore.GetNewDbStore(configuration.ConnectionString)
	productController = controllers.ProductController{Db: db}

	router := mux.NewRouter()
	router.HandleFunc("/products", productController.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/products", productController.AddProductHandler).Methods("POST")
	router.HandleFunc("/products/catalog/{id:[0-9]+}", productController.GetProductsByCatalogHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", productController.GetProductHandler).Methods("GET")
	router.Use(controllers.AuthenticateMiddleware)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error: ", err)
	}
}
