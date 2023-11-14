package subscrible

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/storage"
	"g09-to-do-list/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"
)

func DecreaseLikeCountAfterUserUnlikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Decrease like count after user unlikes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasItemId)

			return storage.NewSQLStore(db).DecreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
