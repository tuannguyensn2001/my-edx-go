package authtransport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authdto "my-edx-go/modules/auth/dto"
	authrepository "my-edx-go/modules/auth/repository"
	"my-edx-go/types"
	"net/http"
)

func Register(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user authdto.Register

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Dữ liệu không hợp lệ",
				Error:   err.Error(),
			})
			return
		}

		repository := authrepository.NewAuthRepository(db)

		_, err := repository.FindByEmail(user.Email)

		if err == nil {
			ctx.JSON(http.StatusBadRequest, types.HttpResponse{
				Message: "Email đã tồn tại",
				Error:   err,
			})
			return
		}

		userCreated, err := repository.Create(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.HttpResponse{
				Message: "Có lỗi xảy ra. Vui lòng thử lại",
				Error:   err,
			})
		}

		ctx.JSON(http.StatusOK, types.HttpResponse{
			Message: "Đăng ký thành công",
			Data:    userCreated,
		})

	}
}
