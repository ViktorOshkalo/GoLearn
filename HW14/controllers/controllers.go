package controllers

import (
	"main/db"
	"main/views"
	"net/http"
)

type BaseController struct{}

func (bc BaseController) WriteViewResponce(v views.View, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", v.ContentType)
	w.Write(v.Data)
}

type UserController struct {
	BaseController
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := db.Users
	view := views.GetUsersView(users)
	uc.WriteViewResponce(view, w, r)
}

type RoleController struct {
	BaseController
}

func (uc RoleController) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles := db.Roles
	view := views.GetRolesView(roles)
	uc.WriteViewResponce(view, w, r)
}

type TaskController struct {
	BaseController
}

func (uc TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := db.Tasks
	view := views.GetTasksView(tasks)
	uc.WriteViewResponce(view, w, r)
}
