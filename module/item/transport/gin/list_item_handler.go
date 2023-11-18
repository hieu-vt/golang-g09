package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/model"
	"g09-to-do-list/module/item/repository"
	"g09-to-do-list/module/item/storage"
	"g09-to-do-list/module/item/storage/rpc"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(service goservice.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		queryString.Process()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := service.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		likeStore := rpc.NewItemService(service)
		repo := repository.NewRepoListItem(store, likeStore, requester)
		biz := biz2.NewBizListItem(repo)

		result, err := biz.ListItem(c.Request.Context(), &queryString.Paging, &queryString.Filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotListEntity(model.TableName, err))

			return
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}
