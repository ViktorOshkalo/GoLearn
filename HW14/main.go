package main

import (
	"fmt"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Yo G!")

	router := mux.NewRouter()

	uc := controllers.UserController{}
	rc := controllers.RoleController{}
	tc := controllers.TaskController{}

	router.HandleFunc("/users", uc.GetUsers).Methods("GET")
	router.HandleFunc("/roles", rc.GetRoles).Methods("GET")
	router.HandleFunc("/tasks", tc.GetTasks).Methods("GET")
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error: ", err)
	}
}
