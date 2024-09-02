package storage

import (
	"context"
	"todo-api/modules/item/model"
)

func (s *Store) CreateTodo(ctx context.Context, item *model.TodoCreation) error {

	if err := s.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}
