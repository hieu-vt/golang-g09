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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewCreateItemBiz(store)

		if err := biz.CreateNewItem(c, &itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))

	}
}
