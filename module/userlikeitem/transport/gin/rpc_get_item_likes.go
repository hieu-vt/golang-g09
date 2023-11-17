package ginuserlikeitem

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/userlikeitem/storage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetItemLikes(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestData struct {
			Ids []int `json:"ids"`
		}

		var data RequestData

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)

		mapRs, err := store.GetItemLikes(c.Request.Context(), data.Ids)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(mapRs))
	}
}
