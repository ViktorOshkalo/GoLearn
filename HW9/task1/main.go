package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	Name        string
	Description string
	Date        time.Time
}

var tasks = []Task{
	{Name: "Cleaning", Description: "Clean up kitchen and bathroom", Date: time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)},
	{Name: "Homework", Description: "Do a homework Math and English", Date: time.Date(2023, time.December, 2, 0, 0, 0, 0, time.UTC)},
	{Name: "Walking", Description: "Walk and breath fresh air 30 mins", Date: time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC)},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	w.WriteHeader(http.StatusOK)
}

func TasksByDateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	dateInput, err := time.Parse(time.DateOnly, date)
	if err != nil {
		fmt.Fprintf(w, "Incorrect date, err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Date ", dateInput)
	var tasksFiltered []Task
	for _, task := range tasks {
		if task.Date == dateInput {
			tasksFiltered = append(tasksFiltered, task)
		}
	}

	if len(tasksFiltered) > 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasksFiltered)
	} else {
		fmt.Fprintf(w, "Tasks not found")
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("GO!")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/tasks", TasksHandler)
	r.HandleFunc("/tasks/{date}", TasksByDateHandler)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
