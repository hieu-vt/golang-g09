package gin

import (
	"g09-to-do-list/common"
	"g09-to-do-list/module/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(*model.User)
		u.SQLModel.Mask(common.DbTypeUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
