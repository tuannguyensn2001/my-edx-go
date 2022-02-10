package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwt_provider "my-edx-go/providers/jwt"
	"my-edx-go/types"
	"net/http"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Authorization not valid")
	}

	return parts[1], nil
}

func Auth(ctx *gin.Context) {
	token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, types.HttpResponse{
			Message: "Dữ liệu không hợp lệ",
			Error:   err,
		})
		return
	}

	user, err := jwt_provider.ValidateToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, types.HttpResponse{
			Message: "Dữ liệu không hợp lệ",
			Error:   err,
		})
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
