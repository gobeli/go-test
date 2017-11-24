package todo

import (
	"github.com/jinzhu/gorm"
	"test/pkg/shared"
)

type MySqlToDoService struct {
	DB gorm.DB
}

func (service MySqlToDoService) GetToDos() ([]ToDo, error) {
	var todos []ToDo
	err := service.DB.Find(&todos).Error
	return todos, shared.StandardizeError(err)
}

func (service MySqlToDoService) GetToDo(id uint) { }
func (service MySqlToDoService) CreateToDo(toDo ToDo) (error) {
	err := service.DB.Create(&toDo).Error
	return err
}
func (service MySqlToDoService) RemoveToDo(id uint) { }