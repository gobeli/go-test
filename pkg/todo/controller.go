package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobeli/go-test/pkg/shared"
)

type ToDoController struct {
	Service ToDoService
}

// ToDoOverview handles /todo get
func (controller ToDoController) ToDoOverview(w http.ResponseWriter, r *http.Request) {
	todos, _ := controller.Service.GetToDos()
	obj, _ := json.Marshal(todos)
	page := ToDoPage{ToDos: todos, Page: &shared.Page{Title: "Todos", Data: string(obj)}}
	templates := shared.ParseTemplate("todo/index")
	templates.ExecuteTemplate(w, "base.html", page)
}

// ToDoEdit Handles /todo/add get and post
func (controller ToDoController) ToDoEdit(w http.ResponseWriter, r *http.Request) {
	templates := shared.ParseTemplate("todo/add")
	if r.Method == http.MethodGet {
		page := shared.Page{Title: "Edit Todo"}
		templates.ExecuteTemplate(w, "base.html", page)
	} else if r.Method == http.MethodPost {
		text := r.FormValue("Text")
		// if the todo is invalid
		if text == "" {
			page := shared.Page{Title: "Edit Todo", Messages: []shared.Message{shared.Message{Text: "Text invalid", Color: "red"}}}
			templates.ExecuteTemplate(w, "base.html", page)
			return
		}
		err := controller.Service.CreateToDo(ToDo{Text: text, Finished: false})
		fmt.Println(err)
		http.Redirect(w, r, "/todo", 302)
	}
}
