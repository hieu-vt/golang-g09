package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
