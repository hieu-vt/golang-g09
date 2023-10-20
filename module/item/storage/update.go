package storage

import (
	"context"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
