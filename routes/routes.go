package routes

import (
	"github.com/gin-gonic/gin"
	"my-edx-go/database"
	authroutes "my-edx-go/modules/auth/routes"
)

func InitRoute(router *gin.RouterGroup, store *database.Store) func(module string) {
	return func(module string) {
		switch module {
		case "auth":
			authroutes.RoutesAuth(router, store)
		}
	}
}
