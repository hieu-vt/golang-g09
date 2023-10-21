package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type storageDeleteItem interface {
	DeleteItem(ctx context.Context, id int) error
	FindItem(ctx context.Context, cond map[string]interface{}, itemData *model.TodoItem) error
}

type bizDeleteItem struct {
	store storageDeleteItem
}

func NewBizDeleteItem(store storageDeleteItem) *bizDeleteItem {
	return &bizDeleteItem{store: store}
}

func (biz *bizDeleteItem) DeleteItem(ctx context.Context, id int) error {
	var itemData model.TodoItem
	if err := biz.store.FindItem(ctx, map[string]interface{}{"id": id}, &itemData); err != nil {
		return err
	}

	if itemData.Status == common.DELETED {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.DeleteItem(ctx, id); err == nil {
		return err
	}

	return nil
}
