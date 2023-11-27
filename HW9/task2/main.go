package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users map[string]models.User
var teachers map[int]models.Teacher
var students map[int]models.Student
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
		user1.Id: models.Teacher{Name: "2Pac", UserId: user1.Id},
		user2.Id: models.Teacher{Name: "Jay-Z", UserId: user2.Id},
		user3.Id: models.Teacher{Name: "Snoop Dogg", UserId: user3.Id},
	}

	students := map[int]models.Student{
		1: {Name: "Rabbit", Rates: map[string]int{"Poetry": 12, "Music": 12}, Id: 1},
		2: {Name: "Wink", Rates: map[string]int{"Poetry": 5, "Music": 6}, Id: 2},
		3: {Name: "Cheddar", Rates: map[string]int{"Poetry": 9, "Music": 12}, Id: 3},
		4: {Name: "Paul", Rates: map[string]int{"Poetry": 9, "Music": 8}, Id: 4},
		5: {Name: "Alex", Rates: map[string]int{"Poetry": 10, "Music": 10}, Id: 5},
		6: {Name: "Papa", Rates: map[string]int{"Poetry": 9, "Music": 11}, Id: 6},
		7: {Name: "Lotto", Rates: map[string]int{"Poetry": 5, "Music": 4}, Id: 7},
		8: {Name: "Lyckety-Splyt", Rates: map[string]int{"Poetry": 7, "Music": 10}, Id: 8},
		9: {Name: "Future", Rates: map[string]int{"Poetry": 7, "Music": 9}, Id: 9},
	}

	school = models.School{
		Name: "8 Mile private school",
		Classes: []models.Class{
			{
				Name:    "1A",
				Teacher: teachers[user1.Id],
				Students: []models.Student{
					students[1],
					students[2],
					students[3],
				},
			},
			{
				Name:    "1B",
				Teacher: teachers[user2.Id],
				Students: []models.Student{
					students[4],
					students[5],
					students[6],
				},
			},
			{
				Name:    "1C",
				Teacher: teachers[user3.Id],
				Students: []models.Student{
					students[7],
					students[8],
					students[9],
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
			http.Error(w, "User not provided", http.StatusInternalServerError)
			return
		}

		teacher, ok := teachers[user.Id]
		if !ok {
			http.Error(w, "User is not teacher", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		className := vars["className"]
		class, ok := school.GetClassByName(className)
		if !ok {
			errMessage := fmt.Sprintf("Class not found: %s", className)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		if teacher.Name != class.Teacher.Name {
			errMessage := fmt.Sprintf("Teacher are not allowed to see info for provided class: %s", className)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		ctxWithClass := context.WithValue(r.Context(), "class", class)
		next.ServeHTTP(w, r.WithContext(ctxWithClass))
	})
}

func AuthorizeStudentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "User is not provided", http.StatusInternalServerError)
			return
		}

		teacher, ok := teachers[user.Id]
		if !ok {
			http.Error(w, "User is not teacher", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		studentId, err := strconv.Atoi(vars["id"])
		if err != nil {
			errMessage := fmt.Sprintf("Incorrect student id, err: %s", err)
			http.Error(w, errMessage, http.StatusBadRequest)
			return
		}

		student, found := school.GetTeachersStudentById(teacher, studentId)
		if !found {
			http.Error(w, "Student not belongs to teacher or not found", http.StatusBadRequest)
			return
		}

		ctxWithStudent := context.WithValue(r.Context(), "student", student)
		next.ServeHTTP(w, r.WithContext(ctxWithStudent))
	})
}

func ClassHandler(w http.ResponseWriter, r *http.Request) {
	class, ok := r.Context().Value("class").(models.Class)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value("student").(models.Student)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	fmt.Println("GO!")
	initData()

	rClass := mux.NewRouter()
	rClass.HandleFunc("/class/{className}", ClassHandler)
	rClass.Use(AuthenticateMiddleware)
	rClass.Use(AuthorizeMiddleware)

	rStudent := mux.NewRouter()
	rStudent.HandleFunc("/student/{id}", StudentHandler)
	rStudent.Use(AuthenticateMiddleware)
	rStudent.Use(AuthorizeStudentMiddleware)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/school", SchoolHandler)

	r.PathPrefix("/class").Handler(rClass)
	r.PathPrefix("/student").Handler(rStudent)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error: ", err)
	}
}
