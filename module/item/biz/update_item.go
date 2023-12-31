package biz

import (
	"context"
	"errors"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/model"
)

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
	FindItem(ctx context.Context, cond map[string]interface{}, itemData *model.TodoItem) error
}

type bizUpdateItem struct {
	store     UpdateItemStorage
	requester common.Requester
}

func NewBizUpdateItem(store UpdateItemStorage, requester common.Requester) *bizUpdateItem {
	return &bizUpdateItem{store: store, requester: requester}
}

func (biz *bizUpdateItem) UpdateItem(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	cond := make(map[string]interface{})

	cond["id"] = id

	var itemData model.TodoItem

	if err := biz.store.FindItem(ctx, cond, &itemData); err != nil {
		return err
	}

	if itemData.Status == common.DELETED {
		return model.ErrItemIsDeleted
	}

	isOwner := biz.requester.GetUserId() == itemData.UserId

	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission(errors.New("No permission"))
	}

	if err := biz.store.UpdateItem(ctx, cond, data); err != nil {
		return err
	}

	return nil
}
