package authroutes

import (
	"github.com/gin-gonic/gin"
	"my-edx-go/database"
	authtransport "my-edx-go/modules/auth/transport"
)

func RoutesAuth(router *gin.RouterGroup, store *database.Store) {
	router.POST("/login", authtransport.LoginTransport(store.DB))
	router.POST("/register", authtransport.Register(store.DB))
}
