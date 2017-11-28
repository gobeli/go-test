package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gobeli/go-test/pkg/todo"
)

const port = 8080

func main() {
	// templates := template.Must(template.ParseFiles("templates/index.html"))
	// db, _ := sql.Open("sqlite3", "db/dev.db")
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/myDb?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		migrate(db)
	} else {
		fmt.Println(err)
	}

	todoService := todo.MySqlToDoService{DB: *db}
	todoController := todo.ToDoController{Service: todoService}
	http.HandleFunc("/", todoController.ToDoOverview)
	http.HandleFunc("/todo", todoController.ToDoOverview)
	http.HandleFunc("/todo/add", todoController.ToDoEdit)

	fmt.Println(fmt.Sprintf("Listening on port %d", port))
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&todo.ToDo{})
	fmt.Println("Migration created")
	if err := db.First(&todo.ToDo{}).Error; err != nil {
		todo := todo.ToDo{Finished: false, Text: "My very first Todo"}
		db.Create(&todo)
		fmt.Println("Example created")
	}
}
