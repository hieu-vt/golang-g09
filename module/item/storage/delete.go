package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, id int) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": common.DELETED,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
