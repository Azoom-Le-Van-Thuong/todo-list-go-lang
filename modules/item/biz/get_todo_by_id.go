package biz

import (
	"context"
	"todo-api/modules/item/model"
)

type GetTodoStorage interface {
	GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error)
}

type GetTodoByIdBiz struct {
	store GetTodoStorage
}

func NewGetTodoByIdBiz(store GetTodoStorage) *GetTodoByIdBiz {
	return &GetTodoByIdBiz{
		store: store,
	}
}

func (b *GetTodoByIdBiz) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	todo, err := b.store.GetTodo(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return todo, nil
}
