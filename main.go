package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"test/pkg/todo"
)

type Page struct {
	Name string
	DBStatus bool
	ToDo ToDo
	ToDoString string
}

const port = 8080

func main() {
	// templates := template.Must(template.ParseFiles("templates/index.html"))
	// db, _ := sql.Open("sqlite3", "db/dev.db") 
	db, err  := gorm.Open("mysql", "root:password@tcp(db:3306)/myDb?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		migrate(db)
	} else {
		fmt.Println(err)
	}

	todoService := todo.MySqlToDoService{DB: *db}
	todoController := todo.ToDoController{Service: todoService}
	http.HandleFunc("/", todoController.ToDoOverview)
	http.HandleFunc("/todo", todoController.ToDoOverview)
	http.HandleFunc("/todo/edit", todoController.ToDoEdit)

	fmt.Println(fmt.Sprintf("Listening on port %d", port))
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// func(w http.ResponseWriter, r *http.Request) {
// 	p, name := Page{Name: "Goo"}, r.FormValue("name")
// 	if name != "" {
// 		p.Name = name 
// 	}
// 	p.DBStatus = db.HasTable("to_dos")
// 	db.First(&p.ToDo)
// 	obj, _ := json.Marshal(p.ToDo)
// 	p.ToDoString = string(obj)
// 	err := templates.ExecuteTemplate(w, "index.html", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }