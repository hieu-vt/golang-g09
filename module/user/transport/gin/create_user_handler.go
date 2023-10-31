package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/user/biz"
	"g09-to-do-list/module/user/model"
	"g09-to-do-list/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData model.UserCreate

		if err := c.ShouldBind(&userData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)

		biz := biz2.NewBizCreateUser(store)

		if err := biz.CreateUser(c.Request.Context(), &userData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotCreateEntity(model.EntityName, err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
