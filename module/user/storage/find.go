package storage

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error) {
	db := s.db.Table(model.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user model.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
