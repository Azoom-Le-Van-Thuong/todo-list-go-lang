package biz

import (
	"context"
	"todo-api/modules/item/model"
)

type createTodoStorage interface {
	CreateTodo(ctx context.Context, item *model.TodoCreation) error
}

type CreateTodoBiz struct {
	store createTodoStorage
}

func NewCreateTodoBiz(store createTodoStorage) *CreateTodoBiz {
	return &CreateTodoBiz{
		store: store,
	}
}

func (b *CreateTodoBiz) CreateNewTodo(ctx context.Context, item *model.TodoCreation) error {
	if err := item.Validate(); err != nil {
		return err
	}
	if err := b.store.CreateTodo(ctx, item); err != nil {
		return err
	}
	return nil
}
