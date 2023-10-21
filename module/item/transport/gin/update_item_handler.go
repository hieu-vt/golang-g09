package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/model"
	"g09-to-do-list/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateItemHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		var updateData model.TodoItemUpdate

		if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz2.NewBizUpdateItem(store)

		if err := biz.UpdateItem(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotUpdateEntity(model.TableName, err))

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
