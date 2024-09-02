package biz

import (
	"context"
	"todo-api/modules/item/model"
)

type DeleteTodoStorage interface {
	DeleteTodoById(ctx context.Context, id int) error
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
}

type DeleteTodoByIdBiz struct {
	store DeleteTodoStorage
}

func NewDeleteTodoByIdBiz(store DeleteTodoStorage) *DeleteTodoByIdBiz {
	return &DeleteTodoByIdBiz{
		store: store,
	}
}

func (b *DeleteTodoByIdBiz) DeleteTodoById(ctx context.Context, id int) error {
	_, err := b.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if err := b.store.DeleteTodoById(ctx, id); err != nil {
		return err
	}
	return nil
}
