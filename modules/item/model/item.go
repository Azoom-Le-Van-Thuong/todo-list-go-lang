package model

import (
	"errors"
	"strings"
	"todo-api/common"
)

var (
	ErrTitleCannotBeEmpty = errors.New("Title cannot be empty")
	ErrTodoIsDeleted      = errors.New("Todo is deleted")
)

type Todo struct {
	common.SQLModel
	Title       string `json:"title,omitempty" gorm:"column:title"`
	Description string `json:"description,omitempty" gorm:"column:description"`
	Status      string `json:"status,omitempty" gorm:"column:status"`
}

func (Todo) TableName() string {
	return "todo"
}

type TodoCreation struct {
	Id          int    `json:"id" gorm:"column:id"`
	Title       string `json:"title,omitempty"  gorm:"default:No Title" gorm:"column:title"`
	Description string `json:"description,omitempty"  gorm:"default:No Description" gorm:"column:description"`
	Status      string `json:"status,omitempty"  gorm:"default:Pending" gorm:"column:status"`
}

func (t *TodoCreation) Validate() error {
	t.Title = strings.TrimSpace(t.Title)
	if t.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

func (TodoCreation) TableName() string {
	return "todo"
}

type TodoUpdate struct {
	Title       string `json:"title,omitempty" gorm:"column:title"`
	Description string `json:"description,omitempty" gor:"column:description"`
	Status      string `json:"status,omitempty" gorm:"column:status"`
}

func (TodoUpdate) TableName() string {
	return "todo"
}
