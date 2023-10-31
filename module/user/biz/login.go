package biz

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/component/tokenprovider"
	"g09-to-do-list/module/user/model"
)

type storeFindUser interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type bizLogin struct {
	store         storeFindUser
	hasher        Hasher
	tokenProvider tokenprovider.Provider
	expiry        int
}

func NewBizLogin(
	store storeFindUser,
	hasher Hasher,
	tokenProvider tokenprovider.Provider,
	expiry int) *bizLogin {
	return &bizLogin{
		store:         store,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		expiry:        expiry}
}

func (biz *bizLogin) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil && err == common.RecordNotFound {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if user == nil {
		return nil, common.ErrEntityNotFound(model.EntityName, nil)
	}

	oldPassword := user.Password
	newPassword := biz.hasher.Hash(data.Password + user.Salt)

	if oldPassword != newPassword {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	token, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	//account := usermodel.NewAccount(accessToken, refreshToken)

	return token, nil
}
