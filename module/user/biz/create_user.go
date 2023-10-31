package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/user/model"
)

type storeCreateUser interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type bizCreateUser struct {
	store  storeCreateUser
	hasher Hasher
}

func NewBizCreateUser(store storeCreateUser, hasher Hasher) *bizCreateUser {
	return &bizCreateUser{store: store, hasher: hasher}
}

func (biz *bizCreateUser) CreateUser(ctx context.Context, data *model.UserCreate) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil && err != common.RecordNotFound {
		return err
	}

	if user != nil {
		//if user.Status == 0 {
		//	return error user has been disable
		//}

		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Salt = salt

	data.Password = biz.hasher.Hash(data.Password + data.Salt)

	data.Role = model.RoleUser.String()

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
