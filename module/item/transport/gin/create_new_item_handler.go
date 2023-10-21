package gin

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/model"
	"g09-to-do-list/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData model.TodoItemCreation

		if err := c.ShouldBind(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewCreateItemBiz(store)

		if err := biz.CreateNewItem(c.Request.Context(), &itemData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotCreateEntity(model.TableName, err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))

	}
}
