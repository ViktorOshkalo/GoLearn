package db

import "main/models"

// in-memory db
var Users []models.User = []models.User{
	{Id: 1, Username: "billy123", Email: "billy123@ex.com"},
	{Id: 2, Username: "carl777", Email: "carl777@ex.com"},
}

var Roles []models.Role = []models.Role{
	{Id: 1, Title: "Cleaner", Body: "Cleaning everyting everywhere"},
	{Id: 2, Title: "Washer", Body: "Whasing all you touched"},
}

var Tasks []models.Task = []models.Task{
	{Id: 1, Title: "Clean room", UserId: 1},
	{Id: 2, Title: "Wash clothes", UserId: 2},
}
