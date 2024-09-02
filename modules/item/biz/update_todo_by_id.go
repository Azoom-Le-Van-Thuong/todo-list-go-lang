package biz

import (
	"context"
	"todo-api/modules/item/model"
)

type UpdateTodoStorage interface {
	UpdateTodo(ctx context.Context, cond map[string]interface{}, data *model.TodoUpdate) error
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
}

type UpdateTodoByIdBiz struct {
	store UpdateTodoStorage
}

func NewUpdateTodoByIdBiz(store UpdateTodoStorage) *UpdateTodoByIdBiz {
	return &UpdateTodoByIdBiz{
		store: store,
	}
}

func (b *UpdateTodoByIdBiz) UpdateTodoById(ctx context.Context, id int, data *model.TodoUpdate) error {
	todo, err := b.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if todo.Status == "DELETED" || todo.Status == "deleted" {
		return model.ErrTodoIsDeleted
	}
	if err := b.store.UpdateTodo(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
