package storage

import (
	"context"
	"todo-api/modules/item/model"
)

func (s *Store) DeleteTodoById(ctx context.Context, id int) error {
	if err := s.db.Table(model.Todo{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": "DELETED",
	}).Error; err != nil {
		return err
	}
	return nil

}
