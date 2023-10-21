package gin

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/model"
	"g09-to-do-list/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewBizGetItem(store)

		if item, err := biz.GetItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotGetEntity(model.TableName, err))

			return
		} else {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse(item))
		}
	}
}
