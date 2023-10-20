package biz

import (
	"context"
	"g09-to-do-list/module/item/model"
)

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, id int, data *model.TodoItemUpdate) error
}

type bizUpdateItem struct {
	store UpdateItemStorage
}

func NewBizUpdateItem(store UpdateItemStorage) *bizUpdateItem {
	return &bizUpdateItem{store: store}
}

func (biz *bizUpdateItem) UpdateItem(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	if err := biz.store.UpdateItem(ctx, id, data); err != nil {
		return err
	}

	return nil
}
