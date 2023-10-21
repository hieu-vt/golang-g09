package upload

import (
	"fmt"
	"g09-to-do-list/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Upload(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrNoPermission(err))
			return
		}

		dst := fmt.Sprintf("static/%d%s", time.Now().UnixNano(), file.Filename)

		c.SaveUploadedFile(file, dst)

		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfill("http://localhost:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
