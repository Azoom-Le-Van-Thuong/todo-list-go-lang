package storage

import (
	"context"
	"todo-api/common"
	"todo-api/modules/item/model"
)

func (s *Store) ListTodo(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Todo, error) {
	var results []model.Todo

	db := s.db.Where("status <> ?", "DELETED")
	if filter != nil {
		if filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
	}
	if err := db.Table(model.Todo{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Select("*").
		Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
