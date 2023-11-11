package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, itemId int) error {
	var data model.Like

	if err := s.db.Table(data.TableName()).
		Where("user_id = ? and item_id = ?", userId, itemId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
