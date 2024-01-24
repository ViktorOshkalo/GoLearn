package main

import (
	"fmt"
	dbStore "main/DbStore"
	conf "main/configuration"
	"main/controllers"
	"net/http"

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

func main() {
	fmt.Println("Yoo G")

	db = dbStore.GetNewDbStore(conf.ConnectionString)
	productController = controllers.ProductController{Db: db}

	router := mux.NewRouter()
	router.HandleFunc("/products", productController.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/products", productController.AddProductHandler).Methods("POST")
	router.HandleFunc("/products/catalog/{id:[0-9]+}", productController.GetProductsByCatalogHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", productController.GetProductHandler).Methods("GET")
	router.Use(AuthenticateMiddleware)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error: ", err)
	}
}
