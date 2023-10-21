package gin

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/item/biz"
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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		biz := biz.NewBizGetItem(store)

		if item, err := biz.GetItem(c, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		} else {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse(item))
		}
	}
}
