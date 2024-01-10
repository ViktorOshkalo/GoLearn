package views

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"main/models"
)

type View struct {
	ContentType string
	Data        []byte
}

func GetUsersView(users []models.User) View {
	// csv
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"ID", "Username", "Email"}
	_ = writer.Write(header)

	for _, user := range users {
		data := []string{fmt.Sprint(user.Id), user.Username, user.Email}
		_ = writer.Write(data)
	}

	writer.Flush()
	return View{"text/plain", buf.Bytes()}
}

func GetRolesView(roles []models.Role) View {
	// json
	responseData, _ := json.Marshal(roles)
	return View{"application/json", responseData}
}

func GetTasksView(tasks []models.Task) View {
	// html
	var buf bytes.Buffer

	buf.WriteString("<!DOCTYPE html><html><head><title>User Data</title></head><body>")
	buf.WriteString("<table border='1'><thead><tr>")
	buf.WriteString("<th>Id</th><th>Title</th><th>UserId</th>")
	buf.WriteString("</tr></thead><tbody>")
	for _, task := range tasks {
		buf.WriteString("<tr>")
		buf.WriteString("<td>")
		buf.WriteString(fmt.Sprint(task.Id))
		buf.WriteString("</td>")
		buf.WriteString("<td>")
		buf.WriteString(task.Title)
		buf.WriteString("</td>")
		buf.WriteString("<td>")
		buf.WriteString(fmt.Sprint(task.UserId))
		buf.WriteString("</td>")
		buf.WriteString("</tr>")
	}
	buf.WriteString("</tbody></table></body></html>")

	return View{"text/html", buf.Bytes()}
}
