package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type storeGetItem interface {
	FindItem(ctx context.Context, cond map[string]interface{}, itemData *model.TodoItem) error
}

type bizGetItem struct {
	store storeGetItem
}

func NewBizGetItem(store storeGetItem) *bizGetItem {
	return &bizGetItem{store}
}

func (biz *bizGetItem) GetItem(ctx context.Context, id int) (*model.TodoItem, error) {
	var itemData model.TodoItem

	if err := biz.store.FindItem(ctx, map[string]interface{}{"id": id}, &itemData); err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(model.TableName, err)
		}
		return nil, model.ErrTitleCannotFound
	}

	return &itemData, nil
}
