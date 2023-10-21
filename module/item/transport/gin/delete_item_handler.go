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

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz2.NewBizDeleteItem(store)

		if err := biz.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotDeleteEntity(model.TableName, err))

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
