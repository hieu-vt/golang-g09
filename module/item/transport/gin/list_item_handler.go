package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/model"
	"g09-to-do-list/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(db *gorm.DB) func(c *gin.Context) {
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

		store := storage.NewSQLStore(db)
		biz := biz2.NewBizListItem(store)

		if result, err := biz.ListItem(c.Request.Context(), &queryString.Paging, &queryString.Filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotListEntity(model.TableName, err))

			return
		} else {
			c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
		}
	}
}
