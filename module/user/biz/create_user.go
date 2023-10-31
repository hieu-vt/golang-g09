package biz

import (
	"context"
	"g09-to-do-list/module/user/model"
)

type storeCreateUser interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type bizCreateUser struct {
	store storeCreateUser
}

func NewBizCreateUser(store storeCreateUser) *bizCreateUser {
	return &bizCreateUser{store: store}
}

func (biz *bizCreateUser) CreateUser(ctx context.Context, data *model.UserCreate) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return err
	}

	if user.Email != "" {
		return model.ErrEmailExisted
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
