package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/model"
	"g09-to-do-list/plugin/pubsub"
	"log"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *model.Like) error
}

type ItemLikeStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeItemBiz struct {
	store UserLikeItemStore
	ps    pubsub.PubSub
}

func NewUserLikeItemBiz(store UserLikeItemStore, ps pubsub.PubSub) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, ps: ps}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikeItem, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
