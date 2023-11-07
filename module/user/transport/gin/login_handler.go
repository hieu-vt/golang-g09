package gin

import (
	"g09-to-do-list/common"
	biz2 "g09-to-do-list/module/user/biz"
	"g09-to-do-list/module/user/model"
	"g09-to-do-list/module/user/storage"
	tokenprovider2 "g09-to-do-list/plugin/tokenprovider"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func LoginHandler(db *gorm.DB, tokenProvider tokenprovider2.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData model.UserLogin

		if err := c.ShouldBind(&userData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)

		md5Hash := common.NewMd5Hash()

		biz := biz2.NewBizLogin(store, md5Hash, tokenProvider, 30*24*60*60)

		token, err := biz.Login(c.Request.Context(), &userData)

		if err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
