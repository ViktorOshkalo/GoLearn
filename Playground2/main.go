package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num1, _ := strconv.Atoi(vars["num1"])
	num2, _ := strconv.Atoi(vars["num2"])
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Sum: %d", num1+num2)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num1, _ := strconv.Atoi(vars["num1"])
	num2, _ := strconv.Atoi(vars["num2"])
	if num2 == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Num2 cannot be 0")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Sum: %f", float32(num1)/float32(num2))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/add/{num1}/{num2}", sumHandler)
	r.HandleFunc("/divide/{num1}/{num2}", divideHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error")
	}
}
