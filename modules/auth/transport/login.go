package authtransport

import (
	"github.com/gin-gonic/gin"
	"my-edx-go/models"
	"my-edx-go/modules/auth/dto"
	"my-edx-go/types"
	"net/http"
)

func LoginTransport() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var json dto.LoginDTO
		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Thất bại",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		user := models.User{
			Email:    "devpro2001@gmail.com",
			Password: "java2001",
		}

		ctx.JSON(200, types.HttpResponse{
			Message: "thành công",
			Data:    user,
		})
	}
}
