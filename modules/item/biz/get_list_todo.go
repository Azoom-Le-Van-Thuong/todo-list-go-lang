package biz

import (
	"context"
	"todo-api/common"
	"todo-api/modules/item/model"
)

type GetListTodoStorage interface {
	ListTodo(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,

	) ([]model.Todo, error)
}

type GetListTodoBiz struct {
	store GetListTodoStorage
}

func NewGetListTodoBiz(store GetListTodoStorage) *GetListTodoBiz {
	return &GetListTodoBiz{
		store: store,
	}
}
func (b *GetListTodoBiz) ListTodo(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Todo, error) {
	data, err := b.store.ListTodo(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}
	return data, nil
}
