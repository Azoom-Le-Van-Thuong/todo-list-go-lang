package storage

import (
	"context"
	"todo-api/modules/item/model"
)

func (s *Store) UpdateTodo(ctx context.Context, cond map[string]interface{}, data *model.TodoUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
