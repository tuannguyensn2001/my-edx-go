package authtransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-edx-go/models"
	authdto "my-edx-go/modules/auth/dto"
	authrepository "my-edx-go/modules/auth/repository"
	hash_provider "my-edx-go/providers/hash"
	jwt_provider "my-edx-go/providers/jwt"
	"my-edx-go/types"
	"net/http"
)

type loginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

func LoginTransport(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var json authdto.LoginDTO

		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Email hoặc mật khẩu không hợp lệ",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		userRepository := authrepository.NewAuthRepository(db)

		userFound, err := userRepository.FindByEmail(json.Email)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Email hoặc mật khẩu không hợp lệ",
			})
			return
		}

		checkPassword := hash_provider.Compare(json.Password, userFound.Password)

		if !checkPassword {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Email hoặc mật khẩu không hợp lệ",
			})
			return
		}

		accessToken, _ := jwt_provider.GenerateToken(*userFound)

		refreshToken, _ := jwt_provider.GenerateToken(*userFound)

		ctx.JSON(200, types.HttpResponse{
			Message: "Đăng nhập thành công",
			Data: loginResponse{
				User:         userRepository.FindById(),
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		})
	}
}
