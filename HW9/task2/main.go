package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

var users map[string]models.User
var teachers map[int]models.Teacher
var school models.School

func initData() {

	user1 := models.User{Id: 1, Username: "user1", Password: "user1"}
	user2 := models.User{Id: 2, Username: "user2", Password: "user2"}
	user3 := models.User{Id: 3, Username: "user3", Password: "user3"}
	user4 := models.User{Id: 4, Username: "user4", Password: "user4"}

	users = map[string]models.User{
		user1.Username: user1,
		user2.Username: user2,
		user3.Username: user3,
		user4.Username: user4,
	}

	teachers = map[int]models.Teacher{
		user1.Id: models.Teacher{Name: "2Pac"},
		user2.Id: models.Teacher{Name: "Jay-Z"},
		user3.Id: models.Teacher{Name: "Snoop Dogg"},
	}

	school = models.School{
		Name: "8 Mile private school",
		Classes: []models.Class{
			{
				Name:    "1A",
				Teacher: teachers[user1.Id],
				Students: []models.Student{
					{Name: "Rabbit", Rates: map[string]int{"Poetry": 12, "Music": 12}},
					{Name: "Wink", Rates: map[string]int{"Poetry": 5, "Music": 6}},
					{Name: "Cheddar", Rates: map[string]int{"Poetry": 9, "Music": 12}},
				},
			},
			{
				Name:    "1B",
				Teacher: teachers[user2.Id],
				Students: []models.Student{
					{Name: "Paul", Rates: map[string]int{"Poetry": 9, "Music": 8}},
					{Name: "Alex", Rates: map[string]int{"Poetry": 10, "Music": 10}},
					{Name: "Papa", Rates: map[string]int{"Poetry": 9, "Music": 11}},
				},
			},
			{
				Name:    "1C",
				Teacher: teachers[user3.Id],
				Students: []models.Student{
					{Name: "Lotto", Rates: map[string]int{"Poetry": 5, "Music": 4}},
					{Name: "Lyckety-Splyt", Rates: map[string]int{"Poetry": 7, "Music": 10}},
					{Name: "Future", Rates: map[string]int{"Poetry": 7, "Music": 9}},
				},
			},
		},
	}

	fmt.Println(school)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")
}

func SchoolHandler(w http.ResponseWriter, r *http.Request) {
	schoolPreview := school.GetPreview()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schoolPreview)
	w.WriteHeader(http.StatusOK)
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, success := r.BasicAuth()
		if !success {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user, found := users[username]
		if !found {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if user.Password != password {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctxWithUser))
	})
}

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "User not provided")
			return
		}

		teacher, ok := teachers[user.Id]
		if !ok {
			fmt.Fprintf(w, "User is not teacher")
			w.WriteHeader(http.StatusOK)
			return
		}

		vars := mux.Vars(r)
		className := vars["className"]
		class, ok := school.GetClassByName(className)
		if !ok {
			fmt.Fprintf(w, "Class not found: %s", className)
			w.WriteHeader(http.StatusOK)
			return
		}

		if teacher.Name != class.Teacher.Name {
			fmt.Fprintf(w, "Teacher are not allowed to see info for provided class: %s", className)
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ClassHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	className := vars["className"]
	class, ok := school.GetClassByName(className)
	if !ok {
		fmt.Fprintf(w, "Class not found: %s", className)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("GO!")
	initData()

	rClass := mux.NewRouter()
	rClass.HandleFunc("/class/{className}", ClassHandler)
	rClass.Use(AuthenticateMiddleware)
	rClass.Use(AuthorizeMiddleware)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/school", SchoolHandler)
	r.PathPrefix("/").Handler(rClass)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error: ", err)
	}
}
