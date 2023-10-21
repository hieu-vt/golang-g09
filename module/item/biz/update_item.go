package biz

import (
	"context"
	"g09-to-do-list/module/item/model"
)

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

type bizUpdateItem struct {
	store UpdateItemStorage
}

func NewBizUpdateItem(store UpdateItemStorage) *bizUpdateItem {
	return &bizUpdateItem{store: store}
}

func (biz *bizUpdateItem) UpdateItem(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	cond := make(map[string]interface{})

	cond["id"] = id

	if err := biz.store.UpdateItem(ctx, cond, data); err != nil {
		return err
	}

	return nil
}
