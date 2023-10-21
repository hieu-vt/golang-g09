package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindItem(ctx context.Context, cond map[string]interface{}, itemData *model.TodoItem) error {
	db := s.db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")

	if err := db.Where(cond).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}

		return common.ErrDB(err)
	}

	return nil
}
