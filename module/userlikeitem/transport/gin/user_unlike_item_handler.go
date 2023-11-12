package ginuserlikeitem

import (
	"g09-to-do-list/common"
	storage2 "g09-to-do-list/module/item/storage"
	"g09-to-do-list/module/userlikeitem/biz"
	"g09-to-do-list/module/userlikeitem/storage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UnlikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		itemStore := storage2.NewSQLStore(db)
		business := biz.NewUserUnlikeItemBiz(store, itemStore)

		if err := business.UnlikeItem(c.Request.Context(), requester.GetUserId(), int(id.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
