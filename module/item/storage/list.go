package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) ListItem(ctx context.Context, paging *common.Paging, filter *model.Filter, result *[]model.TodoItem, moreKeys ...string) error {
	db := s.db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")
	if status := filter.Status; status != "" {
		db = db.Where("status = ?", status)
	}
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
