package subscrible

import (
	"context"
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/storage"
	"g09-to-do-list/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"
)

type HasItemId interface {
	GetItemId() int
}

func IncreaseLikeCountAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Increase like count after user likes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasItemId)

			return storage.NewSQLStore(db).IncreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
