package ginuserlikeitem

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/biz"
	"g09-to-do-list/module/userlikeitem/model"
	"g09-to-do-list/module/userlikeitem/storage"
	"g09-to-do-list/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func LikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)
		business := biz.NewUserLikeItemBiz(store, ps)
		now := time.Now().UTC()

		if err := business.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    int(id.GetLocalID()),
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
