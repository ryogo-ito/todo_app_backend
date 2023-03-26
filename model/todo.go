package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

func GetTodos() ([]Todo, error) {
	var todos []Todo

	err := db.Find(&todos).Error

	return todos, err

}

func AddTodo(name string) (*Todo, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	todo := Todo{
		ID:        id,
		Name:      name,
		Completed: false,
	}

	err = db.Create(&todo).Error

	return &todo, err
}

func ChangeCompletedTodo(todoID uuid.UUID) error {
	err := db.Model(&Todo{}).Where("id = ?", todoID).Update("completed", true).Error
	return err
}

func DeleteTodo(todoID uuid.UUID) error {
	err := db.Model(&Todo{}).Where(" id = ?", todoID).Delete(Todo{}).Error
	return err
}
