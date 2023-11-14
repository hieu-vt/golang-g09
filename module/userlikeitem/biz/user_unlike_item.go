package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/model"
	"g09-to-do-list/plugin/pubsub"
	"log"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStore
	ps    pubsub.PubSub
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStore, ps pubsub.PubSub) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store, ps: ps}
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

	if err := biz.ps.Publish(ctx, common.TopicUserUnlikeItem, pubsub.NewMessage(&model.Like{UserId: userId, ItemId: itemId})); err != nil {
		log.Println(err)
	}

	return nil
}
