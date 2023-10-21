package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/item/biz"
	"g09-to-do-list/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		store := storage.NewSQLStore(db)
		biz := biz2.NewBizListItem(store)

		if result, err := biz.ListItem(c, &paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		} else {
			c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
		}
	}
}
