package main

import (
	"encoding/json"
	"fmt"
	conf "main/configuration"
	"main/dbStore"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db dbStore.DbStore

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

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the "id" parameter from the URL
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	// Convert productIDStr to int64
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
	fmt.Println("YooG")

	db = dbStore.GetNewDbStore(conf.ConnectionString)

	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", GetProductHandler).Methods("GET")
	router.Use(AuthenticateMiddleware)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error: ", err)
	}
}
