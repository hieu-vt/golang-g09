package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

func (s *sqlStore) ListItem(ctx context.Context, paging *common.Paging, filter *model.Filter, moreKeys ...string) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := s.db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")
	if status := filter.Status; status != "" {
		db = db.Where("status = ?", status)
	}
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		result[len(result)-1].Mask()
		paging.NextCursor = result[len(result)-1].FadeId.String()
	}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
