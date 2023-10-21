package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type storageListItem interface {
	ListItem(ctx context.Context, paging *common.Paging, result *[]model.TodoItem) error
}

type bizListItem struct {
	store storageListItem
}

func NewBizListItem(store storageListItem) *bizListItem {
	return &bizListItem{store: store}
}

func (biz *bizListItem) ListItem(ctx context.Context, paging *common.Paging) ([]model.TodoItem, error) {
	var result []model.TodoItem

	if err := biz.store.ListItem(ctx, paging, &result); err != nil {
		return nil, err
	}

	return result, nil
}
