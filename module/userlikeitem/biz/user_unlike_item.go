package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/model"
	"log"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type ItemUnlikeStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeItemBiz struct {
	store           UserUnlikeItemStore
	itemUnlikeStore ItemUnlikeStore
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStore, itemUnlikeStore ItemUnlikeStore) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store, itemUnlikeStore: itemUnlikeStore}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)

	// Delete if data existed
	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	go func() {
		defer common.Recovery()

		if err := biz.itemUnlikeStore.DecreaseLikeCount(ctx, itemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
