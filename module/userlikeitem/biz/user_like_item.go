package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/model"
	"log"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type ItemLikeStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeItemBiz struct {
	store     UserLikeItemStore
	itemStore ItemLikeStore
}

func NewUserLikeItemBiz(store UserLikeItemStore, itemStore ItemLikeStore) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	go func() {
		defer common.Recovery()

		if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
