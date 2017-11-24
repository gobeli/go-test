package todo

import (
	"net/http"
	"encoding/json"
	"test/pkg/shared"
)

type ToDoController struct {
	Service ToDoService
}

func (controller ToDoController) ToDoOverview(w http.ResponseWriter, r *http.Request) {
	todos, _ := controller.Service.GetToDos()
	obj, _ := json.Marshal(todos)
	page := ToDoPage{ToDos: todos, Page: &shared.Page{Title: "Todos", Data: string(obj)} }
	templates := shared.ParseTemplate("todo/index")
	templates.ExecuteTemplate(w, "base.html", page)
} 

func (controller ToDoController) ToDoEdit(w http.ResponseWriter, r *http.Request) {
	templates := shared.ParseTemplate("todo/edit")
	if r.Method == http.MethodGet {
		page := shared.Page{ Title: "Edit Todo" }
		templates.ExecuteTemplate(w, "base.html", page)
	} else if r.Method == http.MethodPost {
		text := r.FormValue("Text")
		// if the todo is invalid
		if text == "" {
			page := shared.Page{ Title: "Edit Todo", Messages: []shared.Message{shared.Message{Text: "Text invalid", Color: "red"} }}
			templates.ExecuteTemplate(w, "base.html", page)
			return
		}

		controller.Service.CreateToDo(ToDo{Text: text, Finished: false})
		http.Redirect(w, r, "/todo", 302)
	}
}

func (controller ToDoController) ToDoEditPost(w http.ResponseWriter, r *http.Request) {
	
}