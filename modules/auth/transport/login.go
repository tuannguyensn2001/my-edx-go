package authtransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authdto "my-edx-go/modules/auth/dto"
	"my-edx-go/modules/auth/repository"

	"my-edx-go/types"
	"net/http"
)

func LoginTransport(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var json authdto.LoginDTO

		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Thất bại",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		user_repository := repository.NewAuthRepository(db)

		ctx.JSON(200, types.HttpResponse{
			Message: "thành công",
			Data:    user_repository.FindById(),
		})
	}
}
