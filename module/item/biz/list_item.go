package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type storageListItem interface {
	ListItem(ctx context.Context, paging *common.Paging, filter *model.Filter, result *[]model.TodoItem, moreKeys ...string) error
}

type bizListItem struct {
	store storageListItem
}

func NewBizListItem(store storageListItem) *bizListItem {
	return &bizListItem{store: store}
}

func (biz *bizListItem) ListItem(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	var result []model.TodoItem

	if err := biz.store.ListItem(ctx, paging, filter, &result, "Owner"); err != nil {
		return nil, err
	}

	return result, nil
}
