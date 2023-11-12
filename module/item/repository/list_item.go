package repository

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type storageListItem interface {
	ListItem(ctx context.Context,
		paging *common.Paging, filter *model.Filter,
		moreKeys ...string) ([]model.TodoItem, error)
}

type ItemLikeStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type repoListItem struct {
	store     storageListItem
	likeStore ItemLikeStorage
	requester common.Requester
}

func NewRepoListItem(store storageListItem, likeStore ItemLikeStorage, requester common.Requester) *repoListItem {
	return &repoListItem{store: store, likeStore: likeStore, requester: requester}
}

func (repo *repoListItem) ListItem(
	ctx context.Context,
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, repo.requester)

	data, err := repo.store.ListItem(ctxStore, paging, filter, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.TableName, err)
	}

	if len(data) == 0 {
		return data, nil
	}

	ids := make([]int, len(data))

	for i := range ids {
		ids[i] = data[i].Id
	}

	likeUserMap, err := repo.likeStore.GetItemLikes(ctxStore, ids)

	if err != nil {
		return data, nil
	}

	for i := range data {
		data[i].LikeCount = likeUserMap[data[i].Id]
	}

	return data, nil
}
