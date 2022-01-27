package authtransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-edx-go/models"
	authdto "my-edx-go/modules/auth/dto"
	authrepository "my-edx-go/modules/auth/repository"
	jwt_provider "my-edx-go/providers/jwt"
	"my-edx-go/types"
	"net/http"
)

type loginResponse struct {
	AccessToken string      `json:"access_token"`
	User        models.User `json:"user"`
}

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

		userRepository := authrepository.NewAuthRepository(db)

		userFound := userRepository.FindById()

		accessToken, _ := jwt_provider.GenerateToken(userFound)

		ctx.JSON(200, types.HttpResponse{
			Message: "thành công",
			Data: loginResponse{
				User:        userRepository.FindById(),
				AccessToken: accessToken,
			},
		})
	}
}
