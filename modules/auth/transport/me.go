package authtransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-edx-go/types"
	"net/http"
)

func Me(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user")

		ctx.JSON(http.StatusOK, types.HttpResponse{
			Message: "Lấy thông tin người dùng thành công",
			Data:    user,
		})
	}
}
