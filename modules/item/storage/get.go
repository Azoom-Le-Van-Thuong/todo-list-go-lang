package storage

import (
	"context"
	"todo-api/modules/item/model"
)

func (s *Store) GetTodo(ctx context.Context, cond map[string]interface{}) (*model.Todo, error) {
	var todo model.Todo
	if err := s.db.Where(cond).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
