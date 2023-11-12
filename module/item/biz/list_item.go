package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type repoListItem interface {
	ListItem(
		ctx context.Context,
		paging *common.Paging,
		filter *model.Filter,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type bizListItem struct {
	repo repoListItem
}

func NewBizListItem(repo repoListItem) *bizListItem {
	return &bizListItem{repo: repo}
}

func (biz *bizListItem) ListItem(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	result, err := biz.repo.ListItem(ctx, paging, filter, "Owner")
	if err != nil {
		return nil, err
	}

	return result, nil
}
