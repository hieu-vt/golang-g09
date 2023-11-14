package subscrible

import (
	"context"
	"g09-to-do-list/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"log"
)

type HasUserId interface {
	GetUserId() int
}

func PushNotificationAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Push notification after user likes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			//db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasUserId)

			log.Println("Push notification to user id:", data.GetUserId())

			return nil
		},
	}
}
