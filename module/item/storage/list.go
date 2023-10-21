package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) ListItem(ctx context.Context, paging *common.Paging, result *[]model.TodoItem) error {
	db := s.db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return err
	}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return err
	}

	return nil
}
