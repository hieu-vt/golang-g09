package storage

import (
	"context"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) FindItem(ctx context.Context, cond map[string]interface{}, itemData *model.TodoItem) error {
	if err := s.db.Where(cond).First(&itemData).Error; err != nil {
		return err
	}

	return nil
}
