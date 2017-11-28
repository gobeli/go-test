package todo

import "github.com/gobeli/go-test/pkg/shared"

type ToDo struct {
	ID				uint
	Finished 	bool
	Text 			string
}

type ToDoPage struct {
	*shared.Page
	ToDos []ToDo
}

type ToDoService interface {
	GetToDos() ([]ToDo, error)
	GetToDo(id uint)
	CreateToDo(toDo ToDo) (error)
	RemoveToDo(id uint)
}