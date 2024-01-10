package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Role struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Task struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"userId"`
}
